package openfaasfn

import (
	"encoding/json"
	"fmt"
	handler "github.com/openfaas-incubator/go-function-sdk"
)

func Handler(req handler.Request) (handler.Response, error) {
	var err error
	var messageJSON Message
	jsonErr := json.Unmarshal(req.Body, &messageJSON)
	if jsonErr != nil {
		panic(jsonErr)
	}
	message := fmt.Sprintf("Hi,I'm OpenFaaS. I have received your message '%s'", messageJSON.Text)

	return handler.Response{
		Body: []byte(message),
	}, err
}
