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

func TestClient_GetLatestStockQuote(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(
			"symbol,open,high,low,price,volume,latestDay,previousClose,change,changePercent"+
			"MSFT,136.9600,137.5200,136.4250,137.3900,13611682,2019-09-17,136.3300,1.0600,0.7775%"
		))
		return &res, nil
	}), "")
	got, err := c.GetLatestStockQuote("MSFT")
	if err != nil {
		t.Fatal(err)
	}
	if want := (LatestStockQuote{
		Symbol:        "MSFT",
		Open:          136.9600,
		High:          137.5200,
		Low:           136.4250,
		Price:         137.3900,
		Volume:        13611682,
		LatestDay:     marshaler.Date(time.Date(2019, 9, 17, 0, 0, 0, 0, time.UTC)),
		PreviousClose: 136.3300,
		Change:        1.0600,
		ChangePercent: 0.007775,
	}); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
