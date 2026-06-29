package vars

import (
	"github.com/spf13/viper"
)

type application struct {
	name string
}

var app *application

func newApplication() *application {
	a := new(application)
	name := viper.GetString("application.name")
	if name == "" {
		name = "myApp"
	}

	a.name = name
	return a
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
