package main

import (
	"fmt"
	"function-golang-helloworld/internal/openfaasfn"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Unable to read standard input: %s", err.Error())
	}
	fmt.Println(openfaasfn.Handler(input))
}
