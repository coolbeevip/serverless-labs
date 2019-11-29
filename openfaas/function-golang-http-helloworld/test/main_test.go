package test

import (
	"encoding/json"
	"fmt"
	"function-golang-http-helloworld/internal/inspector"
	"function-golang-http-helloworld/internal/openfaasfn"
	handler "github.com/openfaas-incubator/go-function-sdk"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	var msg = openfaasfn.Message{Text: "hello"}
	var request = handler.Request{Host: "127.0.0.1", QueryString: "/", Method: "say", Body: []byte(jsonToString(msg))}
	var expectedResponse, _ = openfaasfn.Handler(request)
	var actual = fmt.Sprintf("Hi,I'm OpenFaaS. I have received your message '%s'", msg.Text)
	var expected = string(expectedResponse.Body[:])
	if actual != expected {
		t.Errorf("actual %s; expected %s", actual, expected)
	}
}

func TestServer(t *testing.T) {
	var msg = openfaasfn.Message{Text: "hello"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", inspector.MakeRequestHandler())
	reader := strings.NewReader(jsonToString(msg))
	r, _ := http.NewRequest(http.MethodPost, "/", reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
	} else {
		var actual = fmt.Sprintf("Hi,I'm OpenFaaS. I have received your message '%s'", msg.Text)
		expected := string(w.Body.Bytes()[:])
		if actual != expected {
			t.Errorf("actual %s; expected %s", actual, expected)
		}
	}
}

func jsonToString(message openfaasfn.Message) string {
	msg, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	return string(msg[:])
}
