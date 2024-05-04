package util

import (
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, config interface{}) error {
	return LoadYamlFile(path, "config", config)
}

func LoadYamlFile(path, name string, result interface{}) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	if err := viper.Unmarshal(&result, viper.DecodeHook(springbootLikeDecodeHook)); err != nil {
		return err
	}

	return nil

}

// ============================================================================================
// Permet de récupérer les valeurs comme springboot
// Exemple :
//
// server:
//
//	property: ${AN_ENV_VARIABLE:default-value}
//
// ============================================================================================
func springbootLikeDecodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String {
		stringData := data.(string)

		// Gestion du format ${ENV_VAR:default-value}
		re := regexp.MustCompile(`^\${([^:]+):([^}]+)}$`)
		match := re.FindStringSubmatch(stringData)
		if len(match) == 3 {
			varName := match[1]
			defaultValue := match[2]
			envValue, found := os.LookupEnv(varName)
			if !found {
				envValue = defaultValue
			}
			return envValue, nil
		}

		// Gestion du format ${ENV_VAR}
		if strings.HasPrefix(stringData, "${") && strings.HasSuffix(stringData, "}") {
			envVarValue := os.Getenv(strings.TrimPrefix(strings.TrimSuffix(stringData, "}"), "${"))
			if len(envVarValue) > 0 {
				return envVarValue, nil
			}
		}
	}
	return data, nil
}

// ============================================================================================
