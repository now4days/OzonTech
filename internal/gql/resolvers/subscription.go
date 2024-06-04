package resolvers

import (
	"context"
	"example.com/graphql/internal/models"
)

type subscriptionResolver struct{ *Resolver }

var subscribers = map[string]chan *models.Comment{}

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postId string) (<-chan *models.Comment, error) {
	ch := make(chan *models.Comment, 1)
	subscribers[postId] = ch
	go func() {
		<-ctx.Done()
		delete(subscribers, postId)
	}()
	return ch, nil
}
