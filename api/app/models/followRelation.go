package models

// "a follower (follower_id) is following this user (followed_id)"
type FollowRelation struct {
	Id         uint `json:"id"`
	FollowedId uint `json:"followed_id"`
	FollowerId uint `json:"follower_id"`
}
