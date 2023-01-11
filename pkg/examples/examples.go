package examples

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExternalCallToAnotherAPI() {
	type PostResponse struct {
		UserId int    `json:"userId"`
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("Failed to http.Get>>> %s\n", err.Error())
	}

	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to ioutil.ReadAll>>> %s\n", err.Error())
	}

	var postResponse PostResponse

	err = json.Unmarshal(jsonData, &postResponse)
	if err != nil {
		fmt.Printf("Failed to Unmarshal>>> %s\n", err.Error())
	}

	fmt.Printf("Title>>> %s\nId>>> %d\nUserId>>> %d\nBody>>> %s", postResponse.Title, postResponse.Id, postResponse.UserId, postResponse.Body)

	defer resp.Body.Close()

	fmt.Printf("jsonData>>> %s\n", jsonData)

}
