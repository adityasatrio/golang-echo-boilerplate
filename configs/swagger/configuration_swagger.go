package swagger

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/cmd/docs"
)

func InitSwagger() {
	docs.SwaggerInfo.Title = "Micro-go-template Service"
	docs.SwaggerInfo.Description = "Please welcome a holy high-speed and high-performance Echo service!"
	docs.SwaggerInfo.Version = viper.GetString("application.version")
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = "/" + viper.GetString("application.name")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	log.Infof("initialized swagger components : %s/swagger/index.html", docs.SwaggerInfo.Host+docs.SwaggerInfo.BasePath)
}
