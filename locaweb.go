package main

import (
	"fmt"
	"github.com/andersondborges/locaweb/twitter"
)

func main() {
	fmt.Println("Start Project")
	/*
		client := &http.Client{}
		resp, err := client.Get("http://tweeps.locaweb.com.br/tweeps")
		fmt.Println(err)
		fmt.Println(resp)

		req, err := http.NewRequest("GET", "http://tweeps.locaweb.com.br/tweeps", nil)
		req.Header.Add("Username", "andersondborges@gmail.com")
		resp, err = client.Do(req)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		fmt.Println(err)
		fmt.Println(body)
	*/

	statuses, _ := twitter.GetTweets()
	fmt.Println(statuses)

}
