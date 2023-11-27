# HTMX + Go Fiber

A proof of concept HTMX app with a Go Fiber backend. Displays a list of tickers from polygon.io
and a list of news related to the ticker.

<p align="center">
  <img src=".github/assets/home-page.png?raw=true" alt="Home page" width="25%" height="auto"/>
</p>

## Requirements

- Go `1.21+`
- An API key from [Polygon.io](https://polygon.io/dashboard/api-keys)

## Installation

Copy `.env.example` to `.env` and add your Polygon.io API key.

Fetch the dependencies
```sh
go get
```

## Development

Run the following commands to build the Tailwind CSS classes:

```sh
npm install
npm run dev
```

Or if you prefer to use [bun](https://bun.sh/):

```sh
bun install
bun run dev
```

There are two options to start the web server:
- without hot reloading
- with hot reloading

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
/_/--\ |_| |_| \_ v1.49.0, built with Go go1.21.4

watching .
watching bin
!exclude node_modules
watching public
watching src
!exclude tmp
watching views
building...
running...

 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.51.0                   │
 │               http://127.0.0.1:3000               │
 │       (bound on host 0.0.0.0 and port 3000)       │
 │                                                   │
 │ Handlers ............. 7  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 76782 │
 └───────────────────────────────────────────────────┘
```

Then open the browser to http://localhost:3000. You should be able to search for a stock's ticker
upon clicking on a ticker, you should see a list of news related to the ticker.
