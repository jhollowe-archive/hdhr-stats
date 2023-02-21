package main

import "github.com/spf13/viper"

type Config struct {
	General struct {
		TunerHost string
	}
	Channels struct {
		InitialScan bool
		Ignore      []string
	}
}

func initConfig() {

	// General settings
	viper.SetDefault("general.tuner_host", "hdhomerun.local")

	// Channel settings

	viper.AutomaticEnv()
}
