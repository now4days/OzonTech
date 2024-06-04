package resolvers

import (
	"context"
	"example.com/graphql/internal/db"
	"example.com/graphql/internal/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Posts(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	db.DB.Preload("Comments").Find(&posts)
	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post
	if err := db.DB.Preload("Comments").First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
