package inspector

import (
	"function-golang-http-helloworld/internal/openfaasfn"
	handler "github.com/openfaas-incubator/go-function-sdk"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func MakeRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var input []byte

		if r.Body != nil {
			defer r.Body.Close()

			bodyBytes, bodyErr := ioutil.ReadAll(r.Body)

			if bodyErr != nil {
				log.Printf("Error reading body from request.")
			}

			input = bodyBytes
		}

		req := handler.Request{
			Body:        input,
			Header:      r.Header,
			Method:      r.Method,
			QueryString: r.URL.RawQuery,
		}

		result, resultErr := openfaasfn.Handler(req)

		if result.Header != nil {
			for k, v := range result.Header {
				w.Header()[k] = v
			}
		}

		if resultErr != nil {
			log.Print(resultErr)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if result.StatusCode == 0 {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(result.StatusCode)
			}
		}

		w.Write(result.Body)
	}
}

func ParseIntOrDurationValue(val string, fallback time.Duration) time.Duration {
	if len(val) > 0 {
		parsedVal, parseErr := strconv.Atoi(val)
		if parseErr == nil && parsedVal >= 0 {
			return time.Duration(parsedVal) * time.Second
		}
	}

	duration, durationErr := time.ParseDuration(val)
	if durationErr != nil {
		return fallback
	}
	return duration
}
