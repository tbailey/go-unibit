# go-unibit


__Simple client written in go that wraps [unibit.ai](https://unibit.ai/) API's__

## Installation

`go get github.com/tbailey/go-unibit`

## Usage

First sign up for your api token here [unibit.ai](https://unibit.ai/)

Follow this basic example:
```go
c := client.NewClient("your_token_here")

// Request Parameters
request := client.Request{
    Tickers: []string{"AAPL","WORK"},
    StartDate: "2019-08-16",
    EndDate: "2019-08-20",
}

// API calls
result, err := c.GetRealtimeStockPrice(request)
result, err := c.GetHistoricalStockPrice(request)
result, err := c.GetCompanyFinancials(request)
result, err := c.GetCompanyProfile(request)
result, err := c.GetCompanyFinancialSummary(request)
result, err := c.GetCompanyOwnershipStructure(request)
result, err := c.GetCompanyInsiderTransactions(request)
result, err := c.GetSecFilingLink(request)
result, err := c.GetStockNews(request)
result, err := c.GetCorporateSplits(request)
result, err := c.GetCorporateDividends(request)
```
