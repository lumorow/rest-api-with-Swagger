package handlers

import "C"
import (
	operations "day_4/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"os/exec"
	"strconv"
)

type buyCandyImpl struct{}

func NewBuyCandyHandler() operations.BuyCandyHandler {
	return &buyCandyImpl{}
}

func (impl *buyCandyImpl) Handle(params operations.BuyCandyParams) middleware.Responder {
	order := params.Order
	candyCount := int(*order.CandyCount)
	candyType := *order.CandyType
	candyMoney := int(*order.Money)

	candyPrices := map[string]int{
		"CE": 10,
		"AA": 5,
		"NT": 8,
		"DE": 12,
		"YR": 15,
	}

	if candyCount < 0 {
		errorResponse := &operations.BuyCandyBadRequestBody{Error: "Invalid candy count"}
		return operations.NewBuyCandyBadRequest().WithPayload(errorResponse)
	}

	if candyMoney < 0 {
		errorResponse := &operations.BuyCandyBadRequestBody{Error: "Invalid candy money"}
		return operations.NewBuyCandyBadRequest().WithPayload(errorResponse)
	}

	if _, ok := candyPrices[candyType]; !ok {
		errorResponse := &operations.BuyCandyBadRequestBody{Error: "Invalid candy type"}
		return operations.NewBuyCandyBadRequest().WithPayload(errorResponse)
	}

	price := candyPrices[candyType] * candyCount

	if candyMoney < price {
		amount := price - candyMoney
		errorResponse := &operations.BuyCandyPaymentRequiredBody{Error: "You need " + strconv.Itoa(amount) + " more money!"}
		return operations.NewBuyCandyPaymentRequired().WithPayload(errorResponse)
	}
	change := candyMoney - price
	out, err := exec.Command("./exe/cow", "Thank you!").Output()
	if err != nil {
		log.Fatal(err)
	}
	response := &operations.BuyCandyCreatedBody{
		Change: int64(change),
		Thanks: string(out),
	}
	return operations.NewBuyCandyCreated().WithPayload(response)
}
