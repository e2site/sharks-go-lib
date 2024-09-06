package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oleiade/reflections"
	"os"
	"reflect"
)

func LoadConfig(cfg interface{}) error {
	godotenv.Load(".env")

	var err error = nil

	var listParam = make(map[string]string)

	t := reflect.TypeOf(cfg)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			filed := t.Field(i)
			tag := filed.Tag.Get("config")
			if tag != "" {
				listParam[filed.Name] = tag
			}
		}
	}

	for key, val := range listParam {
		osEnvVal := os.Getenv(val)
		if osEnvVal == "" {
			err := errors.New(fmt.Sprintf("Can't find %s env", val))
			return err
		}
		err = reflections.SetField(cfg, key, osEnvVal)
		if err != nil {
			return err
		}
	}

	return err
}
