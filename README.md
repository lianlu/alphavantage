# Alpha Vantage

Alpha Vantage is a library for accessing the [Alpha Vantage API](https://www.alphavantage.co/).

## Example

```go
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
```

## Documentation

Documentation is available [here](https://godoc.org/github.com/tradyfinance/alphavantage).

## License

This project is released under the [Apache License, Version 2.0](LICENSE).
