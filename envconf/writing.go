package envconf

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
	"path/filepath"
	"strings"
)

var ConfigFileDelimiter = Env{"ENVCONF_CONFIG_NAME_VALUE_DELIMITER", "="}
const EnvDelimiter string = "="

var log = logging.MustGetLogger("envconf")

func Apply(prefix string, file string) {
	content := ""
	for _, env := range os.Environ() {
		if !strings.HasPrefix(env, prefix) {
			continue
		}
		envName := strings.SplitN(env, EnvDelimiter, 2)[0]
		configName := strings.Replace(envName, prefix, "", 1)
		configValue := os.Getenv(envName)
		content += configName + ConfigFileDelimiter.Get() + configValue + "\n"
	}
	writeConfig(content, file)
}

func writeConfig(config string, file string) {
	log.Debugf("Writing the following config to %s:\n%s", file, config)
	err := os.MkdirAll(filepath.Dir(file), os.ModePerm)
	handleError(err)
	f, err := os.Create(file)
	handleError(err)
	defer f.Close()
	f.WriteString(config)
}

func handleError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Unable to write config: '%s'", err))
	}
}
