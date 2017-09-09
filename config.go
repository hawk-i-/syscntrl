package main

import "github.com/spf13/viper"

func getConfig(file string) (config *viper.Viper, err error) {
	viper.SetConfigName(file + ".config")
	viper.AddConfigPath("./config")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	config = viper.GetViper()
	return
}

func GetSubConfig(file string, key string) (config *viper.Viper, err error) {
	parentConfig, err := getConfig(file)
	 if err != nil {
	 	return
	 }
	 config = parentConfig.Sub(key)
	 return
}
