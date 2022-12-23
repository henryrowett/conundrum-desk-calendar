package sms

import (
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	twil "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
	"strings"
)

func Send(body string) error {
	// read token from env
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	numbers := os.Getenv("NUMBERS")

	numbersArray := strings.Split(numbers, ",")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	for _, n := range numbersArray {
		params := &twil.CreateMessageParams{}
		params.SetTo(n)
		params.SetFrom("+1 445 545 2603")
		params.SetBody(body)

		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			response, _ := json.Marshal(*resp)
			fmt.Println("Response: " + string(response))
		}
	}

	return nil
}
