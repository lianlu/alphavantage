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
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/tradyfinance/httpext"
)

func TestClient_Search(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(
			"symbol,name,type,region,marketOpen,marketClose,timezone,currency,matchScore"+
			"BA,The Boeing Company,Equity,United States,09:30,16:00,UTC-04,USD,1.0000"+
			"BAC,Bank of America Corporation,Equity,United States,09:30,16:00,UTC-04,USD,0.8000"
		))
		return &res, nil
	}), "")
	got := []SearchResult{}
	if err := c.Search("BA", func(r SearchResult) error {
		got = append(got, r)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if want := []SearchResult{
		SearchResult{
			Symbol:      "BA",
			Name:        "The Boeing Company",
			Type:        "Equity",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  1.0000,
		},
		SearchResult{
			Symbol:      "BAC",
			Name:        "Bank of America Corporation",
			Type:        "Equity",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  0.8000,
		},
	}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
