package main

import (
  "github.com/gin-gonic/gin"
  "recipebook/api"
)

func main() {
  router := gin.Default()

  router.POST("/", gin.WrapH(api.RecipeGraphqlAPI()))

  router.LoadHTMLGlob("view/*")
  router.GET("/play", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
      "title": "Recipe Book",
    })
  })
  router.Run(":3001")
}
