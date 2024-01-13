package subgraph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/availability/subgraph/generated"
	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/availability/subgraph/model"
)

// UpdateAvailability is the resolver for the updateAvailability field.
func (r *mutationResolver) UpdateAvailability(ctx context.Context, employeeID int, isAvailable bool) (*model.Employee, error) {
	storage.Set(employeeID, isAvailable)
	err := r.NC.Publish(fmt.Sprintf("employeeUpdated.%d", employeeID), []byte(fmt.Sprintf(`{"id":%d,"__typename": "Employee"}`, employeeID)))
	if err != nil {
		return nil, err
	}
	return &model.Employee{ID: employeeID, IsAvailable: isAvailable}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }