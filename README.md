# HTMX + Go Fiber

A proof of concept HTMX app with a Go Fiber backend. Displays a list of tickers from polygon.io
and a list of news related to the ticker.

## Requirements

- Golang `1.21+`
- An API key from [Polygon.io](https://polygon.io/dashboard/api-keys)

## Installation

Copy `.env.example` to `.env` and add your Polygon.io API key.

Fetch the dependencies
```sh
go get
```

## Execution

Run the server
```sh
go run .
```

Then open the browser to http://localhost:3000. You should be able to search for a stock's ticker
upon clicking on a ticker, you should see a list of news related to the ticker.
