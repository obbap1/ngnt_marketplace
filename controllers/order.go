package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ngnt_marketplace/helpers"
)

type ExchangeType string

const (
	BuyCoins ExchangeType = "BuyCoins"
	Paxful                = "Paxful"
	Busha                 = "Busha"
)

func (et ExchangeType) IsValid() bool {
	switch et {
	case BuyCoins, Paxful, Busha:
		return true
	}

	return false
}

type NgntOrder struct {
	UserEmail string
	Amount    float64
	Exchange  ExchangeType
	Rate      float64
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order NgntOrder

	fmt.Println("pointer", &order, order)

	decodedBody := json.NewDecoder(r.Body)

	decodedBody.DisallowUnknownFields()

	err := decodedBody.Decode(&order)

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verify that rate is within range

	if isRateValid := helpers.IsRateValid(order.Rate); !isRateValid {
		http.Error(w, "Invalid Rate. Rate should be inbetween 0.9 and 1.1", http.StatusBadRequest)
		return
	}

	// verify exchange is valid
	if isExchangeValid := order.Exchange.IsValid(); !isExchangeValid {
		http.Error(w, "Invalid Exchange. Allowed Exchanges are Buycoins, Paxful and Busha", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", order)
}
