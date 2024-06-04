package resolvers

import (
	"context"
	"example.com/graphql/internal/db"
	"example.com/graphql/internal/models"
	"fmt"
	"math/rand"
	"strconv"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, allowComments bool) (*models.Post, error) {
	post := models.Post{
		ID:            strconv.Itoa(rand.Int()),
		Title:         title,
		Content:       content,
		AllowComments: allowComments,
	}
	db.DB.Create(&post)
	return &post, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, postId string, content string, parentId *string) (*models.Comment, error) {
	var post models.Post
	if err := db.DB.First(&post, "id = ?", postId).Error; err != nil {
		return nil, err
	}
	if !post.AllowComments {
		return nil, fmt.Errorf("comments are not allowed for this post")
	}
	comment := models.Comment{
		ID:       strconv.Itoa(rand.Int()),
		Content:  content,
		PostID:   postId,
		ParentID: parentId,
	}
	db.DB.Create(&comment)

	// Notify subscribers
	if ch, ok := subscribers[postId]; ok {
		ch <- &comment
	}

	return &comment, nil
}
