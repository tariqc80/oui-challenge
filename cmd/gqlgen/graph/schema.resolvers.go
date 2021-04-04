package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/generated"
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/model"
)

func (r *mutationResolver) CreateSet(ctx context.Context, input model.SetInput) (*model.Set, error) {
	var set *model.Set
	newid, err := r.Provider.Create(input.Members)

	if err != nil {
		log.Print(err)
	} else {
		set, err = r.Provider.Get(newid)
	}

	return set, err
}

func (r *queryResolver) Sets(ctx context.Context) ([]*model.Set, error) {
	sets, err := r.Provider.GetCollection()

	if err != nil {
		log.Print(err)
	}

	return sets, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
