package main

import (
	"go-graphql/graph"
	"go-graphql/internal/order"
	"go-graphql/internal/user"
	"go-graphql/pkg/config"
	"go-graphql/pkg/database"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize and migrate the database
	db := database.InitDB()
	database.Migrate(db)

	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize User components
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userResolver := user.NewUserResolver(userService)

	// Initialize Order components
	orderRepo := order.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepo)
	orderResolver := order.NewOrderResolver(orderService)

	resolver := graph.NewResolver(userResolver, orderResolver)

	// Create GraphQL server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Add HTTP and WebSocket transports
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Allow all origins; customize as needed
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second, // Ping interval for WebSocket keep-alive
	})

	http.Handle("/query", srv)

	// Start the server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}