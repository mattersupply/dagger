package dagger

import "github.com/spf13/viper"

func NewConfig(appName string, env string) (*viper.Viper, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/" + appName)
	viper.AddConfigPath("$HOME/." + appName)
	viper.AddConfigPath("./config/")

	err := viper.MergeInConfig()
	if err != nil {
		return nil, err
	}

	viper.SetConfigName("config." + env)
	err = viper.MergeInConfig()
	if err != nil {
		return nil, err
	}

	viper.SetEnvPrefix(appName)
	// AutomaticEnv allows us to override any config variable
	// via environment variable. This might be a little dangerous
	// in some instances. Use viper.BindEnv("some_attribute") to bind
	// individual attributes.
	viper.AutomaticEnv()

	return viper.GetViper(), nil
}
