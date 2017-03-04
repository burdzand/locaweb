package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Statuses represents the result of a Tweet search.
type Statuses struct {
	Tweets []Tweet `json:"statuses"`
}

// SearchTweetParams are the parameters for GetTweets
type SearchTweetParams struct {
	Query           string `url:"q,omitempty"`
	Geocode         string `url:"geocode,omitempty"`
	Lang            string `url:"lang,omitempty"`
	Locale          string `url:"locale,omitempty"`
	ResultType      string `url:"result_type,omitempty"`
	Count           int    `url:"count,omitempty"`
	SinceID         int64  `url:"since_id,omitempty"`
	MaxID           int64  `url:"max_id,omitempty"`
	Until           string `url:"until,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"`
}

// GetTweets returns a collection of Tweets matching a search query.
func GetTweets() (*Statuses, error) {
	statuses := new(Statuses)
	req, err := http.NewRequest("GET", "http://tweeps.locaweb.com.br/tweeps", nil)

	client := &http.Client{}
	req.Header.Add("Username", "andersondborges@gmail.com")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &statuses)
	if err != nil {
		fmt.Println("Erro Json")
	}
	//	body, err := ioutil.ReadAll(resp.Body)

	return statuses, nil
}
