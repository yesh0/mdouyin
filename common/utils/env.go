package utils

import (
	"log"
	"os"
)

type EnvVars struct {
	Host      string
	Cassandra string
	Rdbms     string
	Etcd      string
}

const (
	ENV_MDOUYIN_HOST      = "ENV_MDOUYIN_HOST"
	ENV_MDOUYIN_CASSANDRA = "ENV_MDOUYIN_CASSANDRA"
	ENV_MDOUYIN_RDBMS     = "ENV_MDOUYIN_RDBMS"
	ENV_MDOUYIN_ETCD      = "ENV_MDOUYIN_ETCD"
)

func retrieveEnvOrPanic(name string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	} else {
		log.Fatalln("expecting environmental variable ", name)
		return ""
	}
}

var Env EnvVars

func InitEnvVars() {
	Env = EnvVars{
		Host:      retrieveEnvOrPanic(ENV_MDOUYIN_HOST),
		Cassandra: retrieveEnvOrPanic(ENV_MDOUYIN_CASSANDRA),
		Rdbms:     retrieveEnvOrPanic(ENV_MDOUYIN_RDBMS),
		Etcd:      retrieveEnvOrPanic(ENV_MDOUYIN_ETCD),
	}
}
