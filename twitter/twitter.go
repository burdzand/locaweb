package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

// Statuses represents the result of a Tweet search.
type Statuses struct {
	Tweets []Tweet `json:"statuses"`
}

// FilterLocaweb filter all tweets that have Locaweb
func (statuses *Statuses) FilterLocaweb() []Tweet {
	var tweets []Tweet
	for _, tweet := range statuses.Tweets {
		for _, usermention := range tweet.Entities.UserMentions {
			if usermention.ID == 42 {
				tweets = append(tweets, tweet)
			}
		}
	}
	return tweets
}

// GetTweets returns a collection of Tweets matching a search query.
func GetTweets() ([]Tweet, error) {
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

	return statuses.FilterLocaweb(), nil
}

// GetTweetsOrderUser get tweets order by users
func GetTweetsOrderUser() []UserTweets {
	//usertweets := make(UserTweets)
	tweets, _ := GetTweets()
	var tusers []UserTweets
	for _, tweet := range tweets {
		if index, ok := getUserTweetIndex(tusers, tweet); ok {
			tusers[index].AddTweet(tweet)
		} else {
			tuser := UserTweets{ScreenName: tweet.User.ScreenName, FollowersCount: tweet.User.FriendsCount, ID: tweet.User.ID}
			tuser.AddTweet(tweet)
			tusers = append(tusers, tuser)
		}
	}
	for _, tuser := range tusers {
		sort.Sort(SortByRetweets(tuser.Tweets))
	}
	return tusers
}

func getUserTweetIndex(tusers []UserTweets, tweet Tweet) (int, bool) {
	for index, usertweet := range tusers {
		if usertweet.ScreenName == tweet.User.ScreenName {
			return index, true
		}
	}
	return 0, false
}
