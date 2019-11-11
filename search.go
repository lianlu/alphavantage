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
)

// A SearchResult is a search result from Alpha Vantage.
//
// See: https://www.alphavantage.co/documentation/#symbolsearch
type SearchResult struct {
	Symbol      string  `csv:"symbol"`
	Name        string  `csv:"name"`
	Type        string  `csv:"type"`
	Region      string  `csv:"region"`
	MarketOpen  string  `csv:"marketOpen"`
	MarketClose string  `csv:"marketClose"`
	TimeZone    string  `csv:"timezone"`
	Currency    string  `csv:"currency"`
	MatchScore  float64 `csv:"matchScore"`
}

// Search searches Alpha Vantage by keyword, calling f for each result.
//
// See: https://www.alphavantage.co/documentation/#symbolsearch
func (c *Client) Search(keywords string, f func(SearchResult) error) error {
	return c.getCSV("/query", url.Values{
		"function": []string{"SYMBOL_SEARCH"},
		"keywords": []string{keywords},
	}, func(header, record []string) error {
		var result SearchResult
		if err := csvext.UnmarshalRecord(header, record, &result); err != nil {
			return err
		}
		return f(result)
	})
}
