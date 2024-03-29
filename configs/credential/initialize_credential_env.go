package credential

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	fileCredentialNameDefault = "secret"
	pathCredentialNameDefault = "."
	fileCredentialType        = "env"
)

func InitCredentialEnv(e *echo.Echo) {

	// -credentials-path="ABC" -credentials-name="XYZ"
	var credentialConfigPath string
	flag.StringVar(&credentialConfigPath, "credentials-path", pathCredentialNameDefault, "your credential credentials path config, default /credential")

	var credentialConfigName string
	flag.StringVar(&credentialConfigName, "credentials-name", fileCredentialNameDefault, "your credential credentials file config, default /credential/secret.env")
	flag.Parse()

	credential := GetCredential()
	credential.AddConfigPath(credentialConfigPath)
	credential.SetConfigName(credentialConfigName)
	credential.SetConfigType(fileCredentialType)

	log.Debugf("credential file : " + credential.ConfigFileUsed())
	err := credential.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
		panic(e)
	}

	//set default variable for undefined on credential/secret.env
	initDefaultCredential()

	credential.WatchConfig()
	log.Infof("initialized WatchConfig(): success : credential")
	credential.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed:", e.Name)
	})

	log.Infof("initialized configs viper: success : credential")

}

func initDefaultCredential() {

	credential := GetCredential()

	//comment when config is not use passwords
	credential.SetDefault("db.configs.username", "root")
	credential.SetDefault("db.configs.password", "password")

	credential.SetDefault("db.configs.host", "127.0.0.1")
	credential.SetDefault("db.configs.port", "3306")
	credential.SetDefault("db.configs.database", "echo_sample")

	//comment when config is not use password
	credential.SetDefault("cache.configs.redis.username", "root")
	credential.SetDefault("cache.configs.redis.password", "password")

	credential.SetDefault("cache.configs.redis.db", 0)
	credential.SetDefault("cache.configs.redis.poolSize", 10)
	credential.SetDefault("cache.configs.redis.isTls", false)
	credential.SetDefault("cache.configs.redis.insecureSkipVerify", true)

	credential.SetDefault("rabbitmq.configs.username", "guest")
	credential.SetDefault("rabbitmq.configs.password", "guest")
	credential.SetDefault("rabbitmq.configs.host", "127.0.0.1")
	credential.SetDefault("rabbitmq.configs.port", "5672")
}
