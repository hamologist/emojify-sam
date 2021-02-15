package main

import (
	"crypto/ed25519"
	"emojify/pkg/model/discord"
	discordtransformer "emojify/pkg/transformer/discord"
	"encoding/hex"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

func verify(signature string, hash []byte, publicKey string) bool {
	decodedSignature, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	decodedPublicKey, err := hex.DecodeString(publicKey)
	if err != nil {
		return false
	}

	return ed25519.Verify(decodedPublicKey, hash, decodedSignature)
}

func defaultError() events.APIGatewayProxyResponse {
	return defaultResponse("Failed to process request")
}

func defaultResponse(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body: body,
	}
}

func discordHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload := discord.Interaction{}
	err := json.Unmarshal([]byte(event.Body), &payload)
	if err != nil {
		return defaultError(), err
	}

	headers := make(map[string]string)
	for k,v := range event.Headers {
		headers[strings.ToLower(k)] = v
	}

	signature := headers["x-signature-ed25519"]
	timestamp := headers["x-signature-timestamp"]

	if !verify(signature, []byte(timestamp+event.Body), os.Getenv("DISCORD_PUBLIC_KEY")) {
		return events.APIGatewayProxyResponse{
			StatusCode:        401,
			Body:              "bad signature",
		}, nil
	}

	if payload.Type == 1 {
		return defaultResponse("{\"type\": 1}"), nil
	}

	discordResponse, err := discordtransformer.InteractionToResponse(&payload)
	if err != nil {
		return defaultError(), nil
	}

	response, err := json.Marshal(discordResponse)
	if err != nil {
		return defaultError(), nil
	}

	return defaultResponse(string(response)), nil
}

func main() {
	lambda.Start(discordHandler)
}
