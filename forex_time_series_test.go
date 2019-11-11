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

func TestClient_GetForexTimeSeries(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(
			"timestamp,open,high,low,close"+
			"2019-09-18,1.1063,1.1076,1.1056,1.1065"+
			"2019-09-17,1.1005,1.1075,1.0989,1.1071"
		))
		return &res, nil
	}), "")
	got := []ForexQuote{}
	if err := c.GetForexTimeSeries(
		"EUR", "USD",
		Interval1Day,
		OutputSizeCompact,
		func(q ForexQuote) error {
			got = append(got, q)
			return nil
		},
	); err != nil {
		t.Fatal(err)
	}
	if want := []ForexQuote{
		ForexQuote{
			Timestamp: marshaler.FlexibleTime(time.Date(2019, 9, 18, 0, 0, 0, 0, time.UTC)),
			Open:      1.1063,
			High:      1.1076,
			Low:       1.1056,
			Close:     1.1065,
		},
		ForexQuote{
			Timestamp: marshaler.FlexibleTime(time.Date(2019, 9, 17, 0, 0, 0, 0, time.UTC)),
			Open:      1.1005,
			High:      1.1075,
			Low:       1.0989,
			Close:     1.1071,
		},
	}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
