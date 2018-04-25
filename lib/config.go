package lib

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func setDefaults(viper *viper.Viper) {
	// viper.SetDefault()
}

// Config reads a configuration file and returns the result.
func Config(configFile string) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigName("envbootstrap")
	v.AddConfigPath(".")
	if configFile != "" {
		cleaned := filepath.Clean(configFile)
		dir, file := filepath.Split(cleaned)
		v.AddConfigPath(dir)
		v.SetConfigName(strings.TrimRight(file, filepath.Ext(cleaned)))
	}
	err = v.ReadInConfig()
	return
}
