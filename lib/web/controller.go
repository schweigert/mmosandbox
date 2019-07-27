package web

import "github.com/gin-gonic/gin"

// Controller interface
type Controller interface {
	Routes(*gin.Engine)
}
