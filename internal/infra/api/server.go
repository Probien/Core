package api

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/api/router"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine *gin.Engine

	authRouter         router.IAuthRouter
	branchOfficeRouter router.IBranchOfficeRouter
	categoryRouter     router.ICategoryRouter
	customerRouter     router.ICustomerRouter
	employeeRouter     router.IEmployeeRouter
	endorsementRouter  router.IEndorsementRouter
	pawnOrderRouter    router.IPawnOrderRouter
	productRouter      router.IProductRouter
	logRouter          router.ILogRouter
}

func New(
	engine *gin.Engine,
	authRouter router.IAuthRouter,
	branchOfficeRouter router.IBranchOfficeRouter,
	categoryRouter router.ICategoryRouter,
	customerRouter router.ICustomerRouter,
	employeeRouter router.IEmployeeRouter,
	endorsementRouter router.IEndorsementRouter,
	pawnOrderRouter router.IPawnOrderRouter,
	productRouter router.IProductRouter,
	logRouter router.ILogRouter,

) *Server {
	return &Server{
		engine:             engine,
		authRouter:         authRouter,
		branchOfficeRouter: branchOfficeRouter,
		categoryRouter:     categoryRouter,
		customerRouter:     customerRouter,
		employeeRouter:     employeeRouter,
		endorsementRouter:  endorsementRouter,
		pawnOrderRouter:    pawnOrderRouter,
		productRouter:      productRouter,
		logRouter:          logRouter,
	}
}

func (s *Server) BuildServer() {
	//setup your basepath here
	basePathGroup := s.engine.Group("/api/v1")

	//implement interfaces
	s.authRouter.AuthResource(basePathGroup)
	s.branchOfficeRouter.BranchOfficeResource(basePathGroup)
	s.categoryRouter.CategoryResource(basePathGroup)
	s.customerRouter.CustomerResource(basePathGroup)
	s.employeeRouter.EmployeeResource(basePathGroup)
	s.endorsementRouter.EndorsementResource(basePathGroup)
	s.pawnOrderRouter.PawnOrderResource(basePathGroup)
	s.productRouter.ProductResource(basePathGroup)
	s.logRouter.LogResource(basePathGroup)
}

func (s *Server) Run() {
	if err := s.engine.Run(":9000"); err != nil {
		log.Fatalln(err)
	}
}
