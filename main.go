package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"api/db"
	"api/handlers"
)

func main() {
	if err := db.Open(); err != nil {
		panic("Could not open a connection to the database")
	}
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}

	r.Use(cors.New(config))
	AddGinHandlers(r)
	r.Run(":8080")
}

func AddGinHandlers(r *gin.Engine) {
	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", handlers.Login)
			public.POST("/register", handlers.Register)
		}
	}
}