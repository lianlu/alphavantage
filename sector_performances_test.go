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

func TestClient_GetSectorPerformances(t *testing.T) {
	c := NewClient(httpext.WithTransportFunc(nil, func(req *http.Request) (*http.Response, error) {
		var res http.Response
		res.StatusCode = http.StatusOK
		res.Body = ioutil.NopCloser(strings.NewReader(`{
			"Meta Data": {
				"Information": "US Sector Performance (realtime & historical)",
				"Last Refreshed": "04:20 PM ET 09/17/2019"
			},
			"Rank A: Real-Time Performance": {
				"Real Estate": "1.40%",
				"Utilities": "0.89%",
				"Materials": "0.71%",
				"Consumer Discretionary": "0.60%",
				"Consumer Staples": "0.55%",
				"Information Technology": "0.35%",
				"Communication Services": "0.29%",
				"Health Care": "0.14%",
				"Financials": "0.09%",
				"Industrials": "-0.04%",
				"Energy": "-1.52%"
			}
		}`))
		return &res, nil
	}), "")
	got, err := c.GetSectorPerformances()
	if err != nil {
		t.Fatal(err)
	}
	if want := (SectorPerformances{
		RealTime: SectorPerformance{
			RealEstate:            0.013999999999999999,
			Utilities:             0.0089,
			Materials:             0.0070999999999999995,
			ConsumerDiscretionary: 0.0060,
			ConsumerStaples:       0.0055000000000000005,
			InformationTechnology: 0.0034999999999999996,
			CommunicationServices: 0.0029,
			HealthCare:            0.0014000000000000002,
			Financials:            0.0009,
			Industrials:           -0.0004,
			Energy:                -0.0152,
		},
	}); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
