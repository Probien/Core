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
}

func New(engine *gin.Engine, authRouter router.IAuthRouter, branchOfficeRouter router.IBranchOfficeRouter) *Server {
	return &Server{
		engine:             engine,
		authRouter:         authRouter,
		branchOfficeRouter: branchOfficeRouter,
	}
}

func (s *Server) BuildServer() {
	//setup your basepath here
	basePathGroup := s.engine.Group("/api/v1")
	s.branchOfficeRouter.BranchOfficeResource(basePathGroup)
	s.authRouter.AuthResource(basePathGroup)
}

func (s *Server) Run() {
	if err := s.engine.Run(":9000"); err != nil {
		log.Fatalln(err)
	}
}
