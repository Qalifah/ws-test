package main

import (
	"github.com/brianvoe/gofakeit/v6"
)

type Tweet struct {
	Id        string `fake:"{uuid}" json:"id"`
	Text      string `fake:"{sentence:10}" json:"text"`
	Author    string `fake:"{username}" json:"author"`
	CreatedAt int    `fake:"{second}" json:"createdAt"`
}

func genTweet() (*Tweet, error) {
	var tweet Tweet
	if err := gofakeit.Struct(&tweet); err != nil {
		return nil, err
	}
	return &tweet, nil
}
