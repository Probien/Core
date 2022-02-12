package endorsement_domain

import "github.com/gin-gonic/gin"

type EndorsementRepository interface {
	GetById(c *gin.Context) (*Endorsement, error)
	GetAll() (*[]Endorsement, error)
	Create(c *gin.Context) (*Endorsement, error)
}
