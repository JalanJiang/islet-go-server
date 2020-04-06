package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/JalanJiang/islet-go-server/graph/generated"
	"github.com/JalanJiang/islet-go-server/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	// panic(fmt.Errorf("not implemented"))
	return &model.Book{}, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
