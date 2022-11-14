package main

import (
	"fmt"
	"log"

	"github.com/404th/goplay_gin/config"
	"github.com/404th/goplay_gin/handler"
	"github.com/404th/goplay_gin/storage/postgres"

	"github.com/gin-gonic/gin"

	"github.com/404th/goplay_gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Book Store
// @version         1.0
// @description     BookStore
// @termsOfService  http://swagger.io/terms/
// @contact.name   	404th
// @contact.url    	http://t.me/myevenway
// @contact.email  	umarov.doniyor.2002@gmail.com
// @license.name   	Apache 2.0
// @license.url    	http://www.apache.org/licenses/LICENSE-2.0.html
// @host      	   	http://localhost:7676
// @BasePath  		/api
// @securityDefinitions.basic  BasicAuth
func main() {
	// Load .env
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("err: cannot load config: %v", err)
	}

	// swagger
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Book Store API"
	docs.SwaggerInfo.Description = "Book Store"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = cfg.PROJECT_HOST
	docs.SwaggerInfo.BasePath = fmt.Sprintf(":%s/api", cfg.PROJECT_PORT)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	str := fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.POSTGRES_PORT, cfg.POSTGRES_HOST, cfg.POSTGRES_USER, cfg.POSTGRES_DB, cfg.POSTGRES_PASSWORD, cfg.PGSSLMODE,
	)

	// connect db
	strg, err := postgres.NewPostgresRepo(str)
	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}
	defer strg.DBClose()

	h := handler.NewHandler(strg, cfg)

	r := gin.New()
	r.Use(gin.Logger())

	r.POST("/api/author", h.CreateAuthor)
	r.GET("/api/author/:id", h.GetAuthorByID)
	r.GET("/api/author", h.GetAllAuthor)
	r.PUT("/api/author/:id", h.UpdateAuthor)
	r.DELETE("/api/author/:id", h.DeleteAuthor)

	// TODO - CRUD Book

	// CreateBook
	// GetBookByID
	// GetAllBook
	// UpdateBook
	// DeleteBook

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(fmt.Sprintf(":%s", cfg.PROJECT_PORT))
}
