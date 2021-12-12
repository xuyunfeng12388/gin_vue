package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"io/ioutil"

	"github.com/spf13/viper"
)

const (
	configType = "yaml"
)

func Init(output io.Writer, configFile string) error {
	if output == nil {
		output = ioutil.Discard
	}

	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	viper.SetConfigType(configType) // or viper.SetConfigType("YAML")
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		return err
	}
	// 监控配置文件改变
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		_, _ = fmt.Fprintf(output, "Config file changed %s \n", e.Name)
	})
	return nil
}

func MustInit(output io.Writer, conf string) { // MustInit if fail panic
	if err := Init(output, conf); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}


