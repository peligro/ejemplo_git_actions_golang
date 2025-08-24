package main

//go mod init nombre_proyecto
//snap install go --classic
//go get github.com/pilu/fresh
//go run github.com/pilu/fresh
//go run main.go
import (
	"proyecto-golang/modelos"
	"proyecto-golang/rutas"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//cors
	router.Use(corsMiddleware())

	//migrar la bd
	modelos.Migraciones()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"estado": "ok", "mensaje": "Hola mundo desde Golang con Gin Framework con GORM ORM"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"estado": "ok", "mensaje": "Todo bien"})
	})

	//custom error 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"estado": "error", "message": "Recurso no disponible"})
	})
	//rutas

	router.GET("/categorias", rutas.Categoria_get)
	router.GET("/categorias/:id", rutas.Categoria_get_con_parametro)
	router.POST("/categorias", rutas.Categoria_post)
	router.PUT("/categorias/:id", rutas.Categoria_put)
	router.DELETE("/categorias/:id", rutas.Categoria_delete)

	

	//variables globales
	errorVariables := godotenv.Load()
	if errorVariables != nil {

		panic(errorVariables)

	}

	//inicio servidor
	router.Run(":" + os.Getenv("PORT"))
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

 