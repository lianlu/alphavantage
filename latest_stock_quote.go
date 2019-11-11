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
	"io"
	"net/url"

	"github.com/tradyfinance/csvext"
	"github.com/tradyfinance/marshaler"
)

// A LatestStockQuote is the latest quote for a stock.
//
// See: https://www.alphavantage.co/documentation/#latestprice
type LatestStockQuote struct {
	Symbol        string              `csv:"symbol"`
	Open          float64             `csv:"open"`
	High          float64             `csv:"high"`
	Low           float64             `csv:"low"`
	Price         float64             `csv:"price"`
	Volume        int64               `csv:"volume"`
	LatestDay     marshaler.Date      `csv:"latestDay"`
	PreviousClose float64             `csv:"previousClose"`
	Change        float64             `csv:"change"`
	ChangePercent marshaler.Percent64 `csv:"changePercent"`
}

// GetLatestStockQuote returns the latest quote for a stock.
//
// See: https://www.alphavantage.co/documentation/#latestprice
func (c *Client) GetLatestStockQuote(symbol string) (q LatestStockQuote, err error) {
	err = c.getCSV("/query", url.Values{
		"function": []string{"GLOBAL_QUOTE"},
		"symbol":   []string{symbol},
	}, func(header, record []string) error {
		if err := csvext.UnmarshalRecord(header, record, &q); err != nil {
			return err
		}
		return io.EOF
	})
	if err == io.EOF {
		err = nil
	}
	return
}
