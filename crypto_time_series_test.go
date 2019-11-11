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
	"time"

	"github.com/tradyfinance/httpext"
	"github.com/tradyfinance/marshaler"
)

func TestClient_GetCryptoTimeSeries(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(
			"timestamp,open (CNY),high (CNY),low (CNY),close (CNY),open (USD),high (USD),low (USD),close (USD),volume,market cap (USD)"+
			"2019-09-18,72248.58941200,72456.80759600,72211.71153200,72422.83739500,10187.48000000,10216.84000000,10182.28000000,10212.05000000,466.36058500,466.36058500"+
			"2019-09-17,72689.70559200,72869.27250000,71826.83411900,72251.00065800,10249.68000000,10275.00000000,10128.01000000,10187.82000000,22914.32456300,22914.32456300"
		))
		return &res, nil
	}), "")
	got := []CryptoQuote{}
	if err := c.GetCryptoTimeSeries(
		"BTC", "CNY",
		Interval1Day,
		func(q CryptoQuote) error {
			got = append(got, q)
			return nil
		},
	); err != nil {
		t.Fatal(err)
	}
	if want := []CryptoQuote{
		CryptoQuote{
			Timestamp: marshaler.FlexibleTime(time.Date(2019, 9, 18, 0, 0, 0, 0, time.UTC)),
			Open:      10187.48000000,
			High:      10216.84000000,
			Low:       10182.28000000,
			Close:     10212.05000000,
			Volume:    466.36058500,
			MarketCap: 466.36058500,
		},
		CryptoQuote{
			Timestamp: marshaler.FlexibleTime(time.Date(2019, 9, 17, 0, 0, 0, 0, time.UTC)),
			Open:      10249.68000000,
			High:      10275.00000000,
			Low:       10128.01000000,
			Close:     10187.82000000,
			Volume:    22914.32456300,
			MarketCap: 22914.32456300,
		},
	}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
