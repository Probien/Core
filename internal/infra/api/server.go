package api

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/api/router"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine *gin.Engine

	branchOfficeRouter router.IBranchOfficeRouter
}

func New(engine *gin.Engine, branchOfficeRouter router.IBranchOfficeRouter) *Server {
	return &Server{
		engine:             engine,
		branchOfficeRouter: branchOfficeRouter,
	}
}

func (s *Server) BuildServer() {
	basePathGroup := s.engine.Group("/api/v1")
	s.branchOfficeRouter.BranchOfficeResource(basePathGroup)
}

func (s *Server) Run() {
	if err := s.engine.Run(":9000"); err != nil {
		log.Fatalln(err)
	}
}
