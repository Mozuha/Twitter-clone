package models

type Like struct {
	Id      uint `json:"id"`
	UserId  uint `json:"user_id"`
	TweetId uint `json:"tweet_id"`
}
