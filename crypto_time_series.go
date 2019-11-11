// Copyright 2019 Miles Barr <milesbarr2@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alphavantage

import (
	"errors"
	"net/url"

	"github.com/tradyfinance/csvext"
	"github.com/tradyfinance/marshaler"
)

// A CryptoQuote is a quote for a cryptocurrency.
//
// See: https://www.alphavantage.co/documentation/#digital-currency
type CryptoQuote struct {
	Timestamp marshaler.FlexibleTime `csv:"timestamp"`
	Open      float64                `csv:"open (USD)"`
	High      float64                `csv:"high (USD)"`
	Low       float64                `csv:"low (USD)"`
	Close     float64                `csv:"close (USD)"`
	Volume    float64                `csv:"volume"`
	MarketCap float64                `csv:"market cap (USD)"`
}

// GetCryptoTimeSeries gets cryptocurrency time series data, calling f for each
// quote.
//
// See: https://www.alphavantage.co/documentation/#digital-currency
func (c *Client) GetCryptoTimeSeries(symbol, market string, interval Interval, f func(CryptoQuote) error) error {
	query := url.Values{
		"symbol": []string{symbol},
		"market": []string{market},
	}
	switch interval {
	case Interval1Min:
		fallthrough
	case Interval5Min:
		fallthrough
	case Interval15Min:
		fallthrough
	case Interval30Min:
		fallthrough
	case Interval60Min:
		return errors.New("only daily, weekly, and monthly intervals are supported for cryptocurrencies on Alpha Vantage")
	case Interval1Day:
		fallthrough
	case Interval1Week:
		fallthrough
	case Interval1Month:
		query.Set("function", "DIGITAL_CURRENCY_"+string(interval))
	}
	return c.getCSV("/query", query, func(header, record []string) error {
		var q CryptoQuote
		if err := csvext.UnmarshalRecord(header, record, &q); err != nil {
			return err
		}
		return f(q)
	})
}
