package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/tariqc80/oui-challenge/cmd/gqlgen/graph/model"
	"github.com/tariqc80/oui-challenge/pkg/provider"
)

type Resolver struct {
	Provider *provider.Set
	Sets     []*model.Set
}

func (r *Resolver) Close() {
	r.Provider.Close()
}
