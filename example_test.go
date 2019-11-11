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

package alphavantage_test

import (
	"fmt"
	"log"
	"os"

	"github.com/tradyfinance/alphavantage"
)

func ExampleClient_GetCryptoTimeSeries() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	if err := c.GetCryptoTimeSeries(
		"BTC", "CNY",
		alphavantage.Interval1Day,
		func(q alphavantage.CryptoQuote) error {
			fmt.Printf("%+v\n", q)
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleGetDigitalCurrencies() {
	if err := alphavantage.GetDigitalCurrencies(func(c alphavantage.Currency) error {
		fmt.Printf("%+v\n", c)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_GetDigitalCurrencies() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	if err := c.GetDigitalCurrencies(func(c alphavantage.Currency) error {
		fmt.Printf("%+v\n", c)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleGetPhysicalCurrencies() {
	if err := alphavantage.GetPhysicalCurrencies(func(c alphavantage.Currency) error {
		fmt.Printf("%+v\n", c)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_GetPhysicalCurrencies() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	if err := c.GetPhysicalCurrencies(func(c alphavantage.Currency) error {
		fmt.Printf("%+v\n", c)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_GetExchangeRate() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	er, err := c.GetExchangeRate("USD", "JPY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", er)
}

func ExampleClient_GetForexTimeSeries() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	if err := c.GetForexTimeSeries(
		"EUR", "USD",
		alphavantage.Interval1Day,
		alphavantage.OutputSizeCompact,
		func(q alphavantage.ForexQuote) error {
			fmt.Printf("%+v\n", q)
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_GetLatestStockQuote() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	q, err := c.GetLatestStockQuote("MSFT")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", q)
}

func ExampleClient_Search() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	if err := c.Search("BA", func(r alphavantage.SearchResult) error {
		fmt.Printf("%+v\n", r)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_GetSectorPerformances() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	sp, err := c.GetSectorPerformances()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", sp)
}

func ExampleClient_GetStockTimeSeries() {
	c := alphavantage.NewClient(nil, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	if err := c.GetStockTimeSeries(
		"MSFT",
		alphavantage.Interval1Day,
		alphavantage.OutputSizeCompact,
		func(q alphavantage.StockQuote) error {
			fmt.Printf("%+v\n", q)
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}
}
