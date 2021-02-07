package main

import (
	"fmt"
	"log"
	"omdbapi/config"
	"omdbapi/middleware"
	"omdbapi/routes"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	startApp()
}

func startApp() {

	defaultMiddleware := middleware.DefaultMiddleware{}
	router := gin.Default()
	zap, _ := config.NewLogger()
	var err error
	config.DB, err = gorm.Open("mysql", config.DBURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()

	router.Use(config.Logger(3*time.Second, zap))
	router.Use(defaultMiddleware.CORSMiddleware())
	routes.Route(router)

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)
	router.Run(serverString)
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}

	return ""
}
