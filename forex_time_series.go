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
	"net/url"

	"github.com/tradyfinance/csvext"
	"github.com/tradyfinance/marshaler"
)

// A ForexQuote is a quote for a currency pair.
//
// See: https://www.alphavantage.co/documentation/#fx
type ForexQuote struct {
	Timestamp marshaler.FlexibleTime `csv:"timestamp"`
	Open      float64                `csv:"open"`
	High      float64                `csv:"high"`
	Low       float64                `csv:"low"`
	Close     float64                `csv:"close"`
}

// GetForexTimeSeries gets forex time series data, calling f for each quote.
//
// See: https://www.alphavantage.co/documentation/#fx
func (c *Client) GetForexTimeSeries(from, to string, interval Interval, outputSize OutputSize, f func(ForexQuote) error) error {
	query := url.Values{
		"from_symbol": []string{from},
		"to_symbol":   []string{to},
		"outputsize":  []string{string(outputSize)},
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
		query.Set("function", "FX_INTRADAY")
		query.Set("interval", string(interval))
	case Interval1Day:
		fallthrough
	case Interval1Week:
		fallthrough
	case Interval1Month:
		query.Set("function", "FX_"+string(interval))
	}
	return c.getCSV("/query", query, func(header, record []string) error {
		var q ForexQuote
		if err := csvext.UnmarshalRecord(header, record, &q); err != nil {
			return err
		}
		return f(q)
	})
}
