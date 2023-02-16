package utils

import (
	"log"
	"os"
)

type EnvVars struct {
	Base      string
	Cassandra string
	Rdbms     string
	Etcd      string
	Secret    string
}

const (
	ENV_MDOUYIN_BASE      = "ENV_MDOUYIN_BASE"
	ENV_MDOUYIN_CASSANDRA = "ENV_MDOUYIN_CASSANDRA"
	ENV_MDOUYIN_RDBMS     = "ENV_MDOUYIN_RDBMS"
	ENV_MDOUYIN_ETCD      = "ENV_MDOUYIN_ETCD"
	ENV_MDOUYIN_SECRET    = "ENV_MDOUYIN_SECRET"
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
		Base:      retrieveEnvOrPanic(ENV_MDOUYIN_BASE),
		Cassandra: retrieveEnvOrPanic(ENV_MDOUYIN_CASSANDRA),
		Rdbms:     retrieveEnvOrPanic(ENV_MDOUYIN_RDBMS),
		Etcd:      retrieveEnvOrPanic(ENV_MDOUYIN_ETCD),
		Secret:    retrieveEnvOrPanic(ENV_MDOUYIN_SECRET),
	}
}
