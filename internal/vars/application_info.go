package vars

import (
	"github.com/spf13/viper"
)

type application struct {
	name string
}

var app *application

func newApplication() *application {
	app := new(application)
	name := viper.GetString("application.name")
	if name == "" {
		name = "myApp"
	}

	app.name = name
	return app
}

func init() {
	app = newApplication()
}

func ApplicationName() string {
	return app.getApplicationName()
}

func (app *application) getApplicationName() string {
	return app.name
}
