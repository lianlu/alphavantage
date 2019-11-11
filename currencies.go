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

import "github.com/tradyfinance/csvext"

// A Currency is a digital or physical currency.
type Currency struct {
	Code string `csv:"currency code"` // Currency code.
	Name string `csv:"currency name"` // Currency name.
}

// GetDigitalCurrencies gets a list of digital currencies, calling f for each
// currency.
//
// GetDigitalCurrencies is a wrapper around DefaultClient.GetDigitalCurrencies.
func GetDigitalCurrencies(f func(Currency) error) error {
	return DefaultClient.GetDigitalCurrencies(f)
}

// GetDigitalCurrencies gets a list of digital currencies, calling f for each
// currency.
func (c *Client) GetDigitalCurrencies(f func(Currency) error) error {
	return c.getCSV("/digital_currency_list/", nil, func(header, record []string) error {
		var c Currency
		if err := csvext.UnmarshalRecord(header, record, &c); err != nil {
			return err
		}
		return f(c)
	})
}

// GetPhysicalCurrencies gets a list of physical currencies, calling f for each
// currency.
//
// GetPhysicalCurrencies is a wrapper around
// DefaultClient.GetPhysicalCurrencies.
func GetPhysicalCurrencies(f func(Currency) error) error {
	return DefaultClient.GetPhysicalCurrencies(f)
}

// GetPhysicalCurrencies gets a list of physical currencies, calling f for each
// currency.
func (c *Client) GetPhysicalCurrencies(f func(Currency) error) error {
	return c.getCSV("/physical_currency_list/", nil, func(header, record []string) error {
		var c Currency
		if err := csvext.UnmarshalRecord(header, record, &c); err != nil {
			return err
		}
		return f(c)
	})
}
