package bambooAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetBambooAPIURL() string {
	const bambooAPIKey string = ""
	const bambooAPIURL = "https://" + bambooAPIKey + ":x@api.bamboohr.com/"

	return bambooAPIURL
}

// Make a GET request to a URL.
func GetAPIData(url string, acceptHeader string) []byte {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("accept", acceptHeader)
	response, err := client.Do(request)
	//response, err := http.Get(url)
	if err != nil {
		fmt.Print("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Print("Errored reading the response body %s\n", err2)
		os.Exit(1)
	}

	return body
}
