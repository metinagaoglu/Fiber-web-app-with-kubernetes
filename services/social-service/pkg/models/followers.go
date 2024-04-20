package models

type Follower struct {
	Followee int64  `json:"followee"`
	Follower int64  `json:"followers"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
