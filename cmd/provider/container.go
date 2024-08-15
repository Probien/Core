package provider

import (
	"github.com/JairDavid/Probien-Backend/internal/app"
	postgresAdapter "github.com/JairDavid/Probien-Backend/internal/infra/adapter/postgres"
	redisAdapter "github.com/JairDavid/Probien-Backend/internal/infra/adapter/redis"
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
	postgresClient := postgres.NewPostgresConnection("postgres://postgres:root@probien-database/probien?sslmode=disable")
	redisClient := redis.New("redis-session:6379", "")

	//components
	authenticator := component.NewAuthenticator()
	cookieManager := redisAdapter.NewSessionRepositoryImp(redisClient.GetConnection())

	//DI authentication
	sessionRepo := redisAdapter.NewSessionRepositoryImp(redisClient.GetConnection())
	authRepo := postgresAdapter.NewAuthRepositoryImp(postgresClient.GetConnection())
	authApp := application.NewAuthApp(authRepo, sessionRepo)
	authHandler := handler.NewAuthHandler(authApp)
	authRouter := router.NewAuthRouter(authHandler)

	//DI branch offices
	branchOfficeRepo := postgresAdapter.NewBranchOfficeRepositoryImp(postgresClient.GetConnection())
	branchOfficeApp := application.NewBranchOfficeApp(branchOfficeRepo)
	branchOfficeHandler := handler.NewBranchOfficeHandler(branchOfficeApp)
	branchOfficeRouter := router.NewBranchOfficeRouter(authenticator, cookieManager, branchOfficeHandler)

	//DI categories
	categoryRepo := postgresAdapter.NewCategoryRepositoryImpl(postgresClient.GetConnection())
	categoryApp := application.NewCategoryApp(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryApp)
	categoryRouter := router.NewCategoryHandler(authenticator, cookieManager, categoryHandler)

	//DI customers
	customerRepo := postgresAdapter.NewCustomerRepositoryImpl(postgresClient.GetConnection())
	customerApp := application.NewCustomerApp(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerApp)
	customerRouter := router.NewCustomerRouter(authenticator, cookieManager, customerHandler)

	//DI employees
	employeeRepo := postgresAdapter.NewEmployeeRepositoryImpl(postgresClient.GetConnection())
	employeeApp := application.NewEmployeeApp(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeApp)
	employeeRouter := router.NewEmployeeRouter(authenticator, cookieManager, employeeHandler)

	//DI endorsements
	endorsementRepo := postgresAdapter.NewEndorsementRepositoryImpl(postgresClient.GetConnection())
	endorsementApp := application.NewEndorsementApp(endorsementRepo)
	endorsementHandler := handler.NewEndorsementHandler(endorsementApp)
	endorsementRouter := router.NewEndorsementRouter(authenticator, cookieManager, endorsementHandler)

	//DI pawn orders
	pawnOrderRepo := postgresAdapter.NewPawnOrderRepositoryImpl(postgresClient.GetConnection())
	pawnOrderApp := application.NewPawnOrderApp(pawnOrderRepo)
	pawnOrderHandler := handler.NewPawnOrderHandler(pawnOrderApp)
	pawnOrderRouter := router.NewPawnOrderRouter(authenticator, cookieManager, pawnOrderHandler)

	//DI products
	productRepo := postgresAdapter.NewProductRepositoryImpl(postgresClient.GetConnection())
	productApp := application.NewProductApp(productRepo)
	productHandler := handler.NewProductHandler(productApp)
	productRouter := router.NewProductRouter(authenticator, cookieManager, productHandler)

	//DI logs
	logRepo := postgresAdapter.NewLogsRepositoryImp(postgresClient.GetConnection())
	logApp := application.NewLogApp(logRepo)
	logHandler := handler.NewLogHandler(logApp)
	logRouter := router.NewLogRouter(authenticator, cookieManager, logHandler)

	//API server instance
	server := api.New(
		engine,
		authRouter,
		branchOfficeRouter,
		categoryRouter,
		customerRouter,
		employeeRouter,
		endorsementRouter,
		pawnOrderRouter,
		productRouter,
		logRouter,
	)
	server.BuildServer()
	return server
}
