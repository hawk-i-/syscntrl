package main

import (
	"github.com/hawk-i-/syscntrl/core"
	"github.com/spf13/viper"
)

type AppContext struct {
	config *viper.Viper
}

var Context AppContext

func init() {

}

func main() {
	config, err := getConfig("main")

	if err != nil {
		panic(err)
	}

	Context = AppContext{
		config: config,
	}

	err = core.InitializePackage(core.Context{
		Config:     nil,
		DBProvider: getDBGenerator(Context.config.Sub("db")),
	})

	if err != nil {
		panic(err)
	}
}
