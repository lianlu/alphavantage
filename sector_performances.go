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

	"github.com/tradyfinance/marshaler"
)

// SectorPerformance contains the performance of each sector.
//
// See: https://www.alphavantage.co/documentation/#sector-information
type SectorPerformance struct {
	Utilities             marshaler.Percent64
	CommunicationServices marshaler.Percent64 `json:"Communication Services"`
	RealEstate            marshaler.Percent64 `json:"Real Estate"`
	Financials            marshaler.Percent64
	ConsumerDiscretionary marshaler.Percent64 `json:"Consumer Discretionary"`
	ConsumerStaples       marshaler.Percent64 `json:"Consumer Staples"`
	HealthCare            marshaler.Percent64 `json:"Health Care"`
	Industrials           marshaler.Percent64
	Materials             marshaler.Percent64
	Energy                marshaler.Percent64
	InformationTechnology marshaler.Percent64 `json:"Information Technology"`
}

// SectorPerformances contains the performance of each sector over different
// periods of time.
//
// See: https://www.alphavantage.co/documentation/#sector-information
type SectorPerformances struct {
	RealTime   SectorPerformance `json:"Rank A: Real-Time Performance"`
	OneDay     SectorPerformance `json:"Rank B: 1 Day Performance"`
	FiveDay    SectorPerformance `json:"Rank C: 5 Day Performance"`
	OneMonth   SectorPerformance `json:"Rank D: 1 Month Performance"`
	ThreeMonth SectorPerformance `json:"Rank E: 3 Month Performance"`
	YTD        SectorPerformance `json:"Rank F: Year-to-Date (YTD) Performance"`
	OneYear    SectorPerformance `json:"Rank G: 1 Year Performance"`
	ThreeYear  SectorPerformance `json:"Rank H: 3 Year Performance"`
	FiveYear   SectorPerformance `json:"Rank I: 5 Year Performance"`
	TenYear    SectorPerformance `json:"Rank J: 10 Year Performance"`
}

// GetSectorPerformances returns the performance of each sector over different
// periods of time.
//
// See: https://www.alphavantage.co/documentation/#sector-information
func (c *Client) GetSectorPerformances() (sp SectorPerformances, err error) {
	err = c.getJSON("/query", url.Values{"function": []string{"SECTOR"}}, &sp)
	return
}
