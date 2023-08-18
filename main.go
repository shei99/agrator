package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	var (
		batriumInfluxConnectionUrl = viper.GetString("INFLUX_CONNECTION_URL")
		batriumInfluxToken         = viper.GetString("INFLUX_TOKEN")
		batriumInfluxOrganisation  = viper.GetString("INFLUX_ORG")
		batriumInfluxBucket        = viper.GetString("INFLUX_BUCKET")
	)

	router := gin.Default()
	// incoming.NewBatriumRouter(router, *batriumController)
	router.Run(":9090")
}
