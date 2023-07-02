package main

import (
	_ "github.com/claranceliberi/ampersand/cmd/docs"
	"github.com/claranceliberi/ampersand/pkg/agents"
	"github.com/claranceliberi/ampersand/pkg/batteries"
	"github.com/claranceliberi/ampersand/pkg/battery_telematics"
	"github.com/claranceliberi/ampersand/pkg/common/db"
	"github.com/claranceliberi/ampersand/pkg/drivers"
	"github.com/claranceliberi/ampersand/pkg/stations"
	"github.com/claranceliberi/ampersand/pkg/swaps"
	"github.com/claranceliberi/ampersand/pkg/vehicles"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Ampersand API
// @version         1.0
// @description     This is the API for Ampersand, a car battery swapping service.
// @contact.name   Clarance Liberi Ntwari
// @contact.url    https://twitter.com/claranceliberi
// @contact.email  ntwaricliberi@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /api/v1
func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.GetString("PORT")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPass := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbUrl := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"

	router := gin.Default()

	// enable cors
	router.Use(cors.Default())

	db.Init(dbUrl)

	h := db.Init(dbUrl)

	r := router.Group("/api/v1")

	// Register routes
	vehicles.RegisterRoutes(r, h)
	swaps.RegisterRoutes(r, h)
	batteries.RegisterRoutes(r, h)
	battery_telematics.RegisterRoutes(r, h)
	agents.RegisterRoutes(r, h)
	stations.RegisterRoutes(r, h)
	drivers.RegisterRoutes(r, h)

	// swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
}
