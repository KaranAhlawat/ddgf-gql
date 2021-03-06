// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Advice struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Tags    []*Tag `json:"tags"`
}

type Page struct {
	ID      string    `json:"id"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
}

type Tag struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}
