package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	routes "github.com/giann02/finalBackEndGo/cmd/server/router"
	"github.com/giann02/finalBackEndGo/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/giann02/finalBackEndGo/docs"
	_ "github.com/go-sql-driver/mysql"
)

const (
	puerto = "8080"
)

// @title           Gestion odontologica API
// @version         1.0
// @description     Final Back End.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	// Recover from panic.
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Load the environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database.
	db := connectDB()

	// Create a new Gin engine.
	router := gin.New()
	router.Use(gin.Recovery())
	// Add the logger middleware.
	router.Use(middleware.Logger())

	// Add the swagger handler.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the application.
	runApp(db, router)

	// Close the connection.
	defer db.Close()

}

func runApp(db *sql.DB, engine *gin.Engine) {
	// Run the application.
	router := routes.NewRouter(engine, db)
	// Map all routes.
	router.MapRoutes()
	if err := engine.Run(fmt.Sprintf(":%s", puerto)); err != nil {
		panic(err)
	}

}

// connectDB connects to the database.
func connectDB() *sql.DB {
	
	dbUsername := os.Getenv("DBusername")
	dbPassword := os.Getenv("DBpassword")
	dbHost := os.Getenv("DBhost")
	dbPort := os.Getenv("DBport")
	dbName := os.Getenv("DBname")

	// string de conexion
	// "username:password@tcp(host:puerto)/base_datos"
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}