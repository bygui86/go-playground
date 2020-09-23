package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	// https://api.pro.coinbase.com/products/BTC-USDC/candles?granularity=60&start=2020-03-19T11:02:26Z&end=2020-03-19T11:03:26Z
	// https://api.pro.coinbase.com/products/BTC-USDC/candles?granularity=60&start=2020-03-19T11:00:01Z&end=2020-03-19T11:01:01Z
	// https://api.pro.coinbase.com/products/BTC-USDC/candles?granularity=60&start=2020-03-19T18:09:01Z&end=2020-03-19T18:10:01Z
	url            = "https://api.pro.coinbase.com/products/BTC-USDC/candles"
	urlQueryFormat = "?granularity=%d&start=%s&end=%s"
)

func main() {

	granularity := 60
	end := time.Now().Add(-3 * time.Minute)
	start := end.Add(-1 * time.Minute)

	path := url + buildUrlValues(granularity, &start, &end)

	request, reqErr := http.NewRequest(http.MethodGet, path, nil)
	if reqErr != nil {
		log.Fatal("Request error:", reqErr)
	}

	// requestUrlQuery := request.URL.Query()
	// requestUrlQuery.Add("granularity", strconv.Itoa(granularity))
	// requestUrlQuery.Add("start", start.String())
	// requestUrlQuery.Add("end", end.String())
	// request.URL.RawQuery = requestUrlQuery.Encode()

	fmt.Printf("Request: %s\n", request.URL.String())

	response, resErr := http.DefaultClient.Do(request)
	if resErr != nil {
		log.Fatal("Response error:", resErr)
	}

	// fmt.Printf("Response headers: %+v\n", response.Header)

	// response data
	// var responseBody [][]interface{} // OK
	// // var responseBody [][]float64 // OK
	// decodeErr := encoding.NewDecoder(response.Body).Decode(&responseBody)
	// if decodeErr != nil {
	// 	log.Fatal(decodeErr)
	// }
	// fmt.Printf("Response body: %+v\n", responseBody)
	//
	// resCloseErr := response.Body.Close()
	// if resCloseErr != nil {
	// 	log.Fatal(resCloseErr)
	// }

	// response data string
	data, _ := ioutil.ReadAll(response.Body)
	resCloseErr := response.Body.Close()
	if resCloseErr != nil {
		log.Fatal(resCloseErr)
	}
	fmt.Printf("%s\n", data)
}

// INFO: here is made the translation of any kind of time to UTC
func buildUrlValues(granularity int, start, end *time.Time) string {
	return fmt.Sprintf(urlQueryFormat,
		granularity,
		strings.Split(start.UTC().Format(time.RFC3339), "+")[0],
		strings.Split(end.UTC().Format(time.RFC3339), "+")[0])
}
