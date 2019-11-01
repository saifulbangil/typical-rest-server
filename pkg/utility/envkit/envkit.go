package envkit

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

const configKey = "CONFIG"

// CliLoadEnvFile is cli version of LoadEnvFile
func CliLoadEnvFile(ctx *cli.Context) (err error) {
	return LoadEnvFile()
}

// LoadEnvFile to load environment from .env file
func LoadEnvFile() (err error) {
	configSource := os.Getenv(configKey)
	var configs []string
	var envMap map[string]string
	if configSource == "" {
		envMap, _ = godotenv.Read()
	} else {
		configs = strings.Split(configSource, ",")
		envMap, err = godotenv.Read(configs...)
		if err != nil {
			return
		}
	}
	var builder strings.Builder
	if len(envMap) > 0 {
		builder.WriteString(fmt.Sprintf("Read the environment %s\n", configSource))
		for key, value := range envMap {
			err = os.Setenv(key, value)
			builder.WriteString(" +" + key)
		}
		fmt.Println(builder.String())
	}
	return
}
