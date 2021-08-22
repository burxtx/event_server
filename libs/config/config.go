package config

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string, bean interface{}) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath("./deploy/")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error on parsing configuration file: %s", err.Error())
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed:", e.Name)
	})
	beanValue := reflect.ValueOf(bean)
	if beanValue.Kind() != reflect.Ptr {
		return fmt.Errorf("need a pointer to a value")
	}
	err = v.Unmarshal(bean)
	if err != nil {
		return fmt.Errorf("Unable to decode into struct, %s", err.Error())
	}
	return nil
}

func relativePath(basedir string, path *string) {
	p := *path
	if p != "" && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}
