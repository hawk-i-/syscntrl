package rest

import (
	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"errors"
)

type Context struct {
	config *viper.Viper
	router *mux.Router
}

var context Context
var initialized bool

func InitializePackage (c Context) (err error) {
	if initialized {
		err = errors.New("package already initialized")
		return
	}

	context = c
	initialized = true
	return
}

func IsInitialized () (bool) {
	return initialized
}

func loadRouter () (err error) {
	err = errors.New("not implemented yet")
	return
}