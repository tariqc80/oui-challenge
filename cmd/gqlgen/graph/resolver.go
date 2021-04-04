package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/model"
	"github.com/tariqc80/oui-challenge/pkg/provider"
)

// Resolver struct stores the provider connections and instances of data models
type Resolver struct {
	Db    *provider.Pg
	Cache *provider.Redis
	Sets  []*model.Set
}

// Close closes any external connections
func (r *Resolver) Close() {
	r.Db.Close()
	//r.Cache.Close()
}
