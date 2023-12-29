package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"server/helpers"
	"server/models"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
)

func generateStock(id int, name string, ticker string, value int, timestamp string) models.Stock {
	return models.Stock{
		ID:        id,
		Name:      name,
		Ticker:    ticker,
		Value:     value,
		Timestamp: timestamp,
	}
}

func fluctuateStockPrice(ctx context.Context, priceCh chan<- int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	//tick every second
	ticker := time.NewTicker(time.Second)

	// var limit int
	// limit = 10

outerloop:
	for {
		// if limit == 0 {
		// 	break outerloop
		// }
		select {
		case <-ctx.Done():
			break outerloop
		case <-ticker.C:
			p := r.Intn(100)
			priceCh <- p
		}
		// limit--
	}

	ticker.Stop()

	close(priceCh)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	log.Info().Msgf("Name: %s", name)

	stocks := []models.Stock{
		generateStock(
			1,
			name,
			strings.ToUpper(name),
			rand.Intn(1000),
			time.Now().UTC().Format(time.RFC3339),
		),
	}

	helpers.WriteJSON(w, http.StatusOK, stocks)
}

func GetStocksStream(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	log.Info().Msgf("Name: %s", name)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")


	stock := generateStock(
		1,
		name,
		strings.ToUpper(name),
		rand.Intn(1000),
		time.Now().UTC().Format(time.RFC3339),
	)

	priceCh := make(chan int)

	go fluctuateStockPrice(r.Context(), priceCh)

	for price := range priceCh {
		stock.Value = price
		stock.Timestamp = time.Now().UTC().Format(time.RFC3339)

		marshalledStock, err := json.Marshal(stock)

		if err != nil {
			log.Error().Err(err).Msg("Error marshalling stock")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//send the stock price to the client
		w.Write([]byte("event: single-stock-stream\n"))
		w.Write([]byte(string(fmt.Sprintf("data: %s\n\n", marshalledStock))))
		flusher.Flush()
	}

	helpers.WriteJSON(w, http.StatusOK, struct{
		Message string `json:"message"`
	}{
		Message: "Stream ended",
	})
}
