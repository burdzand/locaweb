package twitter

import "fmt"

// SortByFollows  sort tweets by Number of Followers
type SortByFollows []Tweet

// SortByRetweets sort tweets by Number of Retweets
type SortByRetweets []Tweet

// SortByFavorite sort tweets by Number of Likes
type SortByFavorite []Tweet

func (s SortByFollows) Len() int      { return len(s) }
func (s SortByFollows) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByFollows) Less(i, j int) bool {
	fmt.Println(s[i].User.FollowersCount < s[j].User.FollowersCount)
	return s[i].User.FollowersCount > s[j].User.FollowersCount
}

func (s SortByRetweets) Len() int      { return len(s) }
func (s SortByRetweets) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByRetweets) Less(i, j int) bool {
	return s[i].RetweetCount > s[j].RetweetCount
}

func (s SortByFavorite) Len() int      { return len(s) }
func (s SortByFavorite) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByFavorite) Less(i, j int) bool {
	return s[i].FavoriteCount > s[j].FavoriteCount
}
