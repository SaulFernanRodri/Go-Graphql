package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "go-graphql/graph/model"

type Resolver struct {
	Users        []*model.User
	UserCreated  chan *model.User // Canal de eventos para la suscripción
}
