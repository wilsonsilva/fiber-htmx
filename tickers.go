package main

import (
	"context"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var limit = 5

func initPolygonClient() *polygon.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	return polygon.New(apiKey)
}

func SearchTicker(ticker string) []models.Ticker {
	c := initPolygonClient()

	// To avoid rate limits, limit results to 5 and manually fetch next page.
	res := c.ListTickers(context.Background(), &models.ListTickersParams{Search: &ticker, Limit: &limit})

	var tickers []models.Ticker
	for i := 0; i < limit && res.Next(); i++ {
		tickers = append(tickers, res.Item())
	}

	return tickers
}

func SearchTickerNews(ticker string) []models.TickerNews {
	c := initPolygonClient()

	// To avoid rate limits, limit results to 5 and manually fetch next page.
	res := c.ListTickerNews(context.Background(), &models.ListTickerNewsParams{TickerEQ: &ticker, Limit: &limit})

	var news []models.TickerNews
	for i := 0; i < limit && res.Next(); i++ {
		news = append(news, res.Item())
	}

	return news
}
