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
	"encoding/json"
	"net/url"

	"github.com/tradyfinance/marshaler"
)

// An ExchangeRate is the exchange rate for a currency pair.
type ExchangeRate struct {
	FromCurrencyCode string             `json:"1. From_Currency Code"`
	FromCurrencyName string             `json:"2. From_Currency Name"`
	ToCurrencyCode   string             `json:"3. To_Currency Code"`
	ToCurrencyName   string             `json:"4. To_Currency Name"`
	ExchangeRate     float64            `json:"5. Exchange Rate,string"`
	LastRefreshed    marshaler.DateTime `json:"6. Last Refreshed"`
	TimeZone         string             `json:"7. Time Zone"`
	BidPrice         float64            `json:"8. Bid Price,string"`
	AskPrice         float64            `json:"9. Ask Price,string"`
}

// MarshalJSON implements the json.Marshaler interface.
func (er ExchangeRate) MarshalJSON() ([]byte, error) {
	var v struct {
		RealtimeCurrencyExchangeRate struct {
			FromCurrencyCode string             `json:"1. From_Currency Code"`
			FromCurrencyName string             `json:"2. From_Currency Name"`
			ToCurrencyCode   string             `json:"3. To_Currency Code"`
			ToCurrencyName   string             `json:"4. To_Currency Name"`
			ExchangeRate     float64            `json:"5. Exchange Rate,string"`
			LastRefreshed    marshaler.DateTime `json:"6. Last Refreshed"`
			TimeZone         string             `json:"7. Time Zone"`
			BidPrice         float64            `json:"8. Bid Price,string"`
			AskPrice         float64            `json:"9. Ask Price,string"`
		} `json:"Realtime Currency Exchange Rate"`
	}
	v.RealtimeCurrencyExchangeRate.FromCurrencyCode = er.FromCurrencyCode
	v.RealtimeCurrencyExchangeRate.FromCurrencyName = er.FromCurrencyName
	v.RealtimeCurrencyExchangeRate.ToCurrencyCode = er.ToCurrencyCode
	v.RealtimeCurrencyExchangeRate.ToCurrencyName = er.ToCurrencyName
	v.RealtimeCurrencyExchangeRate.ExchangeRate = er.ExchangeRate
	v.RealtimeCurrencyExchangeRate.LastRefreshed = er.LastRefreshed
	v.RealtimeCurrencyExchangeRate.TimeZone = er.TimeZone
	v.RealtimeCurrencyExchangeRate.BidPrice = er.BidPrice
	v.RealtimeCurrencyExchangeRate.AskPrice = er.AskPrice
	return json.Marshal(&v)
}

// UnmarshalJSON implements the json.Marshaler interface.
func (er *ExchangeRate) UnmarshalJSON(b []byte) error {
	var v struct {
		RealtimeCurrencyExchangeRate struct {
			FromCurrencyCode string             `json:"1. From_Currency Code"`
			FromCurrencyName string             `json:"2. From_Currency Name"`
			ToCurrencyCode   string             `json:"3. To_Currency Code"`
			ToCurrencyName   string             `json:"4. To_Currency Name"`
			ExchangeRate     float64            `json:"5. Exchange Rate,string"`
			LastRefreshed    marshaler.DateTime `json:"6. Last Refreshed"`
			TimeZone         string             `json:"7. Time Zone"`
			BidPrice         float64            `json:"8. Bid Price,string"`
			AskPrice         float64            `json:"9. Ask Price,string"`
		} `json:"Realtime Currency Exchange Rate"`
	}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	er.FromCurrencyCode = v.RealtimeCurrencyExchangeRate.FromCurrencyCode
	er.FromCurrencyName = v.RealtimeCurrencyExchangeRate.FromCurrencyName
	er.ToCurrencyCode = v.RealtimeCurrencyExchangeRate.ToCurrencyCode
	er.ToCurrencyName = v.RealtimeCurrencyExchangeRate.ToCurrencyName
	er.ExchangeRate = v.RealtimeCurrencyExchangeRate.ExchangeRate
	er.LastRefreshed = v.RealtimeCurrencyExchangeRate.LastRefreshed
	er.TimeZone = v.RealtimeCurrencyExchangeRate.TimeZone
	er.BidPrice = v.RealtimeCurrencyExchangeRate.BidPrice
	er.AskPrice = v.RealtimeCurrencyExchangeRate.AskPrice
	return nil
}

// GetExchangeRate returns the exchange rate for a currency pair.
func (c *Client) GetExchangeRate(from, to string) (er ExchangeRate, err error) {
	err = c.getJSON("/query", url.Values{
		"function":      []string{"CURRENCY_EXCHANGE_RATE"},
		"from_currency": []string{from},
		"to_currency":   []string{to},
	}, &er)
	return
}
