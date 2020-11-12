package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type createCheckoutSessionResponse struct {
	SessionID string `json:"id"`
}

func main() {
	server := gin.Default()

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	server.POST("/create-payment-intent", createPaymentIntent)

	server.Static("/", "../../client")

	server.Run(":4242")
}

type orderData struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func createPaymentIntent(c *gin.Context) {
	var order orderData
	c.BindJSON(&order)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(order.Amount),
		Currency: stripe.String(order.Currency),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"publishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
		"clientSecret":   pi.ClientSecret,
	})
}

func confirmPayment(c *gin.Context) {

}
