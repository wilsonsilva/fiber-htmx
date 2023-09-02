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

## Development

### Without hot reloading

```sh
go run .
```

### With hot reloading

Go Fiber does not have a hot reloading feature. Install [`air`](https://github.com/cosmtrek/air#installation) to run
the app with hot reloading.

```sh
air
```

```
  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ 1.45.0, built with Go 1.21.0

watching .
!exclude tmp
watching views
building...
running...

 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.49.0                   │
 │               http://127.0.0.1:3000               │
 │       (bound on host 0.0.0.0 and port 3000)       │
 │                                                   │
 │ Handlers ............. 6  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 96157 │
 └───────────────────────────────────────────────────┘
```

Then open the browser to http://localhost:3000. You should be able to search for a stock's ticker
upon clicking on a ticker, you should see a list of news related to the ticker.
