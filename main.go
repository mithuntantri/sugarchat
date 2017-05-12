package main

import (
  "sugarchat/handlers"
  "sugarchat/psql"
  "sugarchat/store"
  "sugarchat/rethink"
  "github.com/gin-gonic/gin"
)

func main(){
  store.Info.Println("Starting Sugar Chat Application")

  store.Info.Println("Connecting to PSQL")
  psql.ConnectPSQL()
  store.Info.Println("Connecting to PSQL")
  rethink.ConnectDB()
  
  store.Success.Println("Sugar Chat Application Started")

  store.Info.Println("Listening to HTTP Requests on Port 3333")
  gin.SetMode(gin.ReleaseMode)
  router := gin.Default()


  v1 := router.Group("v1")
  {
    v1.POST("/register", handlers.RegisterHandler)
    v1.POST("/login", handlers.LoginHandler)

    app := v1.Group("app")
    app.Use(TokenAuthMiddleware())
    {
      app.GET("/profile", handlers.ProfileHandler)
    }
  }

  router.Run(":3333")
}


func TokenAuthMiddleware() gin.HandlerFunc {
  return func (c *gin.Context)  {
    token := c.Request.Header.Get("X-Authorization-Token")

    if token == ""{
      c.JSON(401, gin.H{
        "success": false,
        "message": "API Token Required",
      })
      c.Abort()
    }
    c.Next()
  }
}
