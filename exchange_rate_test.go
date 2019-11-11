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

	"github.com/tradyfinance/marshaler"

	"github.com/tradyfinance/httpext"
)

func TestClient_GetExchangeRate(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(`{
			"Realtime Currency Exchange Rate": {
				"1. From_Currency Code": "USD",
				"2. From_Currency Name": "United States Dollar",
				"3. To_Currency Code": "JPY",
				"4. To_Currency Name": "Japanese Yen",
				"5. Exchange Rate": "108.24000000",
				"6. Last Refreshed": "2019-09-18 03:22:34",
				"7. Time Zone": "UTC",
				"8. Bid Price": "108.24000000",
				"9. Ask Price": "108.26000000"
			}
		}`))
		return &res, nil
	}), "")
	got, err := c.GetExchangeRate("USD", "JPY")
	if err != nil {
		t.Fatal(err)
	}
	if want := (ExchangeRate{
		FromCurrencyCode: "USD",
		FromCurrencyName: "United States Dollar",
		ToCurrencyCode:   "JPY",
		ToCurrencyName:   "Japanese Yen",
		ExchangeRate:     108.24000000,
		LastRefreshed:    marshaler.DateTime(time.Date(2019, 9, 18, 3, 22, 34, 0, time.UTC)),
		TimeZone:         "UTC",
		BidPrice:         108.24000000,
		AskPrice:         108.26000000,
	}); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
