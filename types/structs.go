package types

import "time"

type FileInfo struct {
    Name    string
    ModTime time.Time
}

type Link struct {
    Href string
    Text string
}

type Tag struct {
    ID     uint   `json:"id" gorm:"primary_key"`
    Name   string `json:"name"`
    Emoji  string `json:"emoji"`
}

