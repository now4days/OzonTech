package services

import "example.com/graphql/internal/db"

func GetPosts() ([]db.Post, error) {
	//
}

func CreatePost(title string, content string, allowComments bool) (db.Post, error) {
	//
}
