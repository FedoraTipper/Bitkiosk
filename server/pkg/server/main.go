package server

import (
	"github.com/fedoratipper/bitkiosk/server/internal/handlers/middleware"
	log "github.com/fedoratipper/bitkiosk/server/internal/logger"

	"github.com/fedoratipper/bitkiosk/server/internal/orm"

	"github.com/fedoratipper/bitkiosk/server/internal/handlers"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

// Run spins up the server
func Run(orm *orm.ORM) {
	log.Info("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))

	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	// Add auth middleware
	r.Use(middleware.AddContextToGinRequest())
	r.Use(middleware.AddCorsToGinRequest())
	log.Info("Auth middleware injected.")
	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())
	r.POST("/authenticate", handlers.AuthenticationHandler(orm))
	r.POST("/logout", handlers.LogoutHandler())


	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Info("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	log.Info("GraphQL @ " + endpoint + gqlPath)
	// Run the server
	// Inform the user where the server is listening
	log.Info("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatal(r.Run(host + ":" + port))
}
