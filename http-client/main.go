/*
for example use some endpoint from https://www.football-data.org/
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiBaseURL = "https://api.football-data.org/v2"

type cli struct {
	httpClient http.Client
}

func main() {
	c := cli{}

	// path - all competitions in the TIER_ONE plan
	path := "/competitions?plan=TIER_ONE"

	fullUrl := fmt.Sprintf("%s%s", apiBaseURL, path)
	resp, err := c.httpClient.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Api only allows a certain amount of free requests per day, this prints
	// remaining count
	fmt.Println(resp.Header.Get("X-Requests-Available"))
	defer resp.Body.Close()

	// read response body
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// print out response body as a string...
	fmt.Println(string(resBody))

	// TODO: decode json body to structs, add some methods maybe, build a cli etc
}
