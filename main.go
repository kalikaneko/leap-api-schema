package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/qri-io/jsonschema"
)

func main() {

	argLength := len(os.Args[1:])
	if argLength != 1 {
		fmt.Println("Usage: leap-api-schema <EIPURL>")
		return
	}

	eipURL := os.Args[1]

	ctx := context.Background()
	schemaData, err := ioutil.ReadFile("schema.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	response, err := http.Get(eipURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	eipData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	errs, err := rs.ValidateBytes(ctx, eipData)
	if err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	} else {
		fmt.Println("OK:", eipURL)
	}
}
