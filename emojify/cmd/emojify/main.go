package main

import (
	"encoding/json"
	"github.com/hamologist/emojify-go/pkg/transformer"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hamologist/emojify-go/pkg/model"
)

func defaultError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	}, err
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload := model.EmojifyPayload{}
	err := json.Unmarshal([]byte(event.Body), &payload)
	if err != nil {
		return defaultError(err)
	}

	emojifyResponse, err := transformer.EmojifyTransformer(payload)
	if err != nil {
		return defaultError(err)
	}

	response, err := json.Marshal(emojifyResponse)
	if err != nil {
		return defaultError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(handler)
}
