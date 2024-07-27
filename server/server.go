package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DefaultApiPrefix = "/api/boilerplate/v1"
)

type (
	ApiServer struct {
		config *ApiServerConfig

		engine *gin.Engine
		db     *gorm.DB

		handler *Handler
	}

	ApiServerConfig struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
)

func DefaultApiServerConfig() (config *ApiServerConfig) {
	bindHost := "127.0.0.1"
	config = &ApiServerConfig{
		Host: bindHost,
		Port: "8000",
	}
	return
}
func NewServer(config *ApiServerConfig) (apiServer *ApiServer, err error) {
	if config == nil {
		config = DefaultApiServerConfig()
	}
	apiServer = &ApiServer{
		config: config,
		engine: gin.Default(),
	}
	if apiServer.db, err = NewDb(nil); err != nil {
		return
	}
	if apiServer.handler, err = NewHandler(apiServer.db); err != nil {
		return
	}
	if err = apiServer.Setup(); err != nil {
		return
	}
	return
}

func (apiServer *ApiServer) Setup() (err error) {
	apiServer.engine.GET("/", apiServer.handler.rootHandler)

	authorized := apiServer.engine.Group(DefaultApiPrefix)
	// TODO: AuthMiddleware not implemented
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/job_templates", apiServer.handler.authRootHandler)
	}

	return
}

func (apiServer *ApiServer) Run() (err error) {
	if err = apiServer.engine.Run(fmt.Sprintf(
		"%s:%s",
		apiServer.config.Host,
		apiServer.config.Port,
	)); err != nil {
		return
	}
	return
}
