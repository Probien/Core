package provider

import (
	"github.com/JairDavid/Probien-Backend/internal/app"
	"github.com/JairDavid/Probien-Backend/internal/infra/adapter/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/api"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/handler"
	"github.com/JairDavid/Probien-Backend/internal/infra/api/router"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/JairDavid/Probien-Backend/internal/infra/resource/postgres"
	"github.com/JairDavid/Probien-Backend/internal/infra/resource/redis"
	"github.com/gin-gonic/gin"
)

type Container struct{}

func New() *Container {
	return &Container{}
}

func (c *Container) Build() *api.Server {
	engine := gin.New()

	//clients
	postgresClient := postgres.NewPostgresConnection("postgres://postgres:root@localhost:5432/probien?sslmode=disable")
	redisClient := redis.New("redis-10391.c279.us-central1-1.gce.cloud.redislabs.com", "10391")

	//components
	authenticator := component.NewAuthenticator()

	//DI branchOffices
	branchOfficeRepo := adapter.NewBranchOfficeRepositoryImp(postgresClient.GetConnection())
	branchOfficeApp := application.NewBranchOfficeApp(branchOfficeRepo)
	branchOfficeHandler := handler.NewBranchOfficeHandler(branchOfficeApp)
	branchOfficeRouter := router.NewBranchOfficeRouter(authenticator, redisClient, branchOfficeHandler)

	//API server instance
	server := api.New(engine, branchOfficeRouter)
	server.BuildServer()
	return server
}
