package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"go-graphql/internal/order"
	"go-graphql/internal/user"
)

type Resolver struct {
	UserResolver  *user.UserResolver
	OrderResolver *order.OrderResolver
}

func NewResolver(userResolver *user.UserResolver, orderResolver *order.OrderResolver) *Resolver {
	return &Resolver{
		UserResolver:  userResolver,
		OrderResolver: orderResolver,
	}
}
