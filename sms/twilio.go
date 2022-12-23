package sms

import (
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	twil "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
)

const (
	accountSid = "ACe5a8d3716f4ef6d0c2447d69587c538b"
)

func Send(body string) error {
	// read token from env
	authToken, ok := os.LookupEnv("TWILIO_AUTH_TOKEN")
	if !ok {
		log.Fatal("error on auth token secret lookup")
	}

	fmt.Println(authToken)

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twil.CreateMessageParams{}
	params.SetTo("+447805444181")
	params.SetFrom("+1 445 545 2603")
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}
