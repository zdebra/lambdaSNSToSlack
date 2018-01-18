package main

import (
	"errors"
	"log"

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
)

const slackWebhookKey = "WEBHOOK_URL"

var (
	ErrInvalidStatusCode    = errors.New("invalid status code")
	ErrSlackWebhookNotFound = errors.New("slack webhook not found in env variables")
)

func Handler(request SNSMessage) error {
	log.Printf("processing message from SNS: %v\n", request)
	if err := request.Validate(); err != nil {
		return err
	}
	slackURL, found := os.LookupEnv(slackWebhookKey)
	if !found {
		return ErrSlackWebhookNotFound
	}
	input := request.Records[0].SNS
	payload := SlackPayload{
		Text: fmt.Sprintf("%s: %s", input.Subject, input.Message),
	}
	payloadJSON, _ := json.Marshal(payload)
	resp, err := http.Post(slackURL, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ErrInvalidStatusCode
	}
	log.Printf("message with id %s send", input.MessageID)
	return nil
}

func main() {
	lambda.Start(Handler)
}
