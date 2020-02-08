package envconf

import (
	"fmt"
	"os"
	"strconv"
)

type Env struct {
	Name         string
	DefaultValue string
}

func (env Env) Get() string {
	if value, exists := os.LookupEnv(env.Name); exists {
		return value
	}
	return env.DefaultValue
}

func (env Env) GetInt() int {
	stringValue := env.Get()
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse int from '%s': '%s'", stringValue, err))
	}
	return intValue
}
