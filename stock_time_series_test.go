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

func TestClient_GetStockTimeSeries(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(
			"timestamp,open,high,low,close,volume\n"+
			"2019-09-17,136.9600,137.5200,136.4250,137.3900,13611682\n"+
			"2019-09-16,135.8300,136.7000,135.6600,136.3300,16013000"
		))
		return &res, nil
	}), "")
	got := []StockQuote{}
	if err := c.GetStockTimeSeries("MSFT", Interval1Day, OutputSizeCompact, func(q StockQuote) error {
		got = append(got, q)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if want := []StockQuote{
		StockQuote{
			Timestamp: marshaler.FlexibleTime(time.Date(2019, 9, 17, 0, 0, 0, 0, time.UTC)),
			Open:      136.9600,
			High:      137.5200,
			Low:       136.4250,
			Close:     137.3900,
			Volume:    13611682,
		},
		StockQuote{
			Timestamp: marshaler.FlexibleTime(time.Date(2019, 9, 16, 0, 0, 0, 0, time.UTC)),
			Open:      135.8300,
			High:      136.7000,
			Low:       135.6600,
			Close:     136.3300,
			Volume:    16013000,
		},
	}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestClient_GetStockTimeSeriesAdjusted(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(
			"timestamp,open,high,low,close,adjusted_close,volume,dividend_amount,split_coefficient\n"+
			"2019-09-17,136.9600,137.5200,136.4250,137.3900,137.3900,13585841,0.0000,1.0000\n"+
			"2019-09-16,135.8300,136.7000,135.6600,136.3300,136.3300,16013000,0.0000,1.0000"
		))
		return &res, nil
	}), "")
	got := []StockQuoteAdjusted{}
	if err := c.GetStockTimeSeriesAdjusted(
		"MSFT",
		Interval1Day,
		OutputSizeCompact,
		func(q StockQuoteAdjusted) error {
			got = append(got, q)
			return nil
		},
	); err != nil {
		t.Fatal(err)
	}
	if want := []StockQuoteAdjusted{
		StockQuoteAdjusted{
			Timestamp:        marshaler.FlexibleTime(time.Date(2019, 9, 17, 0, 0, 0, 0, time.UTC)),
			Open:             136.9600,
			High:             137.5200,
			Low:              136.4250,
			Close:            137.3900,
			AdjustedClose:    137.3900,
			Volume:           13585841,
			DividendAmount:   0.0000,
			SplitCoefficient: 1.0000,
		},
		StockQuoteAdjusted{
			Timestamp:        marshaler.FlexibleTime(time.Date(2019, 9, 16, 0, 0, 0, 0, time.UTC)),
			Open:             135.8300,
			High:             136.7000,
			Low:              135.6600,
			Close:            136.3300,
			AdjustedClose:    136.3300,
			Volume:           16013000,
			DividendAmount:   0.0000,
			SplitCoefficient: 1.0000,
		},
	}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
