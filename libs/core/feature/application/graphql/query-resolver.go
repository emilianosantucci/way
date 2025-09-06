package graphql

import (
	"context"
	"fmt"
	"libs/core/graphql/generated"
)

type ApplicationQueryResolver struct {
}

// Applications is the resolver for the applications field.
func (r *ApplicationQueryResolver) Applications(ctx context.Context) ([]*generated.Application, error) {
	panic(fmt.Errorf("not implemented: Applications - applications"))
}
