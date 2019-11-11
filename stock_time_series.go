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

// A StockQuote is a quote for a stock.
//
// See: https://www.alphavantage.co/documentation/#time-series-data
type StockQuote struct {
	Timestamp marshaler.FlexibleTime `csv:"timestamp"`
	Open      float64                `csv:"open"`
	High      float64                `csv:"high"`
	Low       float64                `csv:"low"`
	Close     float64                `csv:"close"`
	Volume    marshaler.RobustInt64  `csv:"volume"`
}

// A StockQuoteAdjusted is an adjusted quote for a stock.
//
// See: https://www.alphavantage.co/documentation/#time-series-data
type StockQuoteAdjusted struct {
	Timestamp        marshaler.FlexibleTime `csv:"timestamp"`
	Open             float64                `csv:"open"`
	High             float64                `csv:"high"`
	Low              float64                `csv:"low"`
	Close            float64                `csv:"close"`
	AdjustedClose    float64                `csv:"adjusted_close"`
	Volume           marshaler.RobustInt64  `csv:"volume"`
	DividendAmount   float64                `csv:"dividend_amount"`
	SplitCoefficient float64                `csv:"split_coefficient"`
}

// GetStockTimeSeries gets stock time series data, calling f for each quote.
//
// See: https://www.alphavantage.co/documentation/#time-series-data
func (c *Client) GetStockTimeSeries(symbol string, interval Interval, outputSize OutputSize, f func(StockQuote) error) error {
	query := url.Values{
		"symbol":     []string{symbol},
		"outputsize": []string{string(outputSize)},
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
		query.Set("function", "TIME_SERIES_INTRADAY")
		query.Set("interval", string(interval))
	case Interval1Day:
		fallthrough
	case Interval1Week:
		fallthrough
	case Interval1Month:
		query.Set("function", "TIME_SERIES_"+string(interval))
	}
	return c.getCSV("/query", query, func(header, record []string) error {
		var q StockQuote
		if err := csvext.UnmarshalRecord(header, record, &q); err != nil {
			return err
		}
		return f(q)
	})
}

// GetStockTimeSeriesAdjusted gets adjusted stock time series data, calling f
// for each quote.
//
// See: https://www.alphavantage.co/documentation/#time-series-data
func (c *Client) GetStockTimeSeriesAdjusted(symbol string, interval Interval, outputSize OutputSize, f func(StockQuoteAdjusted) error) error {
	query := url.Values{
		"symbol":     []string{symbol},
		"outputsize": []string{string(outputSize)},
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
		query.Set("function", "TIME_SERIES_INTRADAY")
		query.Set("interval", string(interval))
	case Interval1Day:
		fallthrough
	case Interval1Week:
		fallthrough
	case Interval1Month:
		query.Set("function", "TIME_SERIES_"+string(interval)+"_ADJUSTED")
	}
	return c.getCSV("/query", query, func(header, record []string) error {
		var q StockQuoteAdjusted
		if err := csvext.UnmarshalRecord(header, record, &q); err != nil {
			return err
		}
		return f(q)
	})
}
