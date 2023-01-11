package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	// "github.com/stretchr/testify/require"
)

func TestHealhCheck(t *testing.T) {
	url := "/healthCheck"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Rquest error occured", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	// Process the response

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println("Helow")

}
