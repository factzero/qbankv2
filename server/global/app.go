package global

import (
	"gozero/config"

	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Config
}

var App = new(Application)
