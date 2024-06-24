package types

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Creation string `json:"creation"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FileInfo struct {
	Name    string
	ModTime time.Time
}

/*
type Link struct {
    Href string
    Text string
}
*/

type Tag struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}
