package config

import (
	"github.com/fairusatoir/golang-gin-rest/pkg/constants"
	"github.com/fairusatoir/golang-gin-rest/pkg/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Name           string `mapstructure:"app_name"`
	Version        string `mapstructure:"app_version"`
	Environment    string `mapstructure:"APP_ENVIRONMENT"`
	Debug          bool   `mapstructure:"app_debug"`
	DM_Driver      string `mapstructure:"ds_datamaster_driver"`
	DM_URL         string `mapstructure:"ds_datamaster_url"`
	Q_PRODUCT_ALL  string `mapstructure:"Q_PRODUCT_ALL"`
	Q_PRODUCT_FIND string `mapstructure:"Q_PRODUCT_FIND"`
}

func Load() (*Config, error) {
	vp := viper.New()

	vp.SetConfigType("env")
	vp.SetConfigName("")

	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, constants.ErrLoadConfig.Error())
	}

	Config := Config{}
	err = vp.Unmarshal(&Config)
	if err != nil {
		return nil, errors.Wrap(err, constants.ErrUnmarshalConfig.Error())
	}

	return &Config, nil
}

func InitConfig() {
	cfg, err := Load()
	if err != nil {
		log.Fatal(err.Error(), logrus.Fields{log.Category: log.Config})
	}
	log.InfoF(constants.InfoEnvConfig, logrus.Fields{log.Category: log.Config}, cfg.Environment)

}
