package factories

import "os"

type config struct {
	Key   string
	Value string
}

type Configs []*config

func GetConfigs() *Configs {
	// DB
	DBType := os.Getenv("DB_TYPE")
	DBName := os.Getenv("DB_NAME")
	if DBType == "" {
		DBType = "sqlite3"
		DBName = "development.sqlite3"
	}
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	// PORT
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
	return &Configs{
		{Key: "DBType", Value: DBType},
		{Key: "DBHost", Value: DBHost},
		{Key: "DBName", Value: DBName},
		{Key: "DBUser", Value: DBUser},
		{Key: "DBPass", Value: DBPass},
		{Key: "Port", Value: Port},
	}
}
func (cs *Configs) GetValue(key string) string {
	for _, conf := range *cs {
		if conf.Key == key {
			return conf.Value
		}
	}
	return ""
}

func (cs *Configs) Get(key string) (string, string) {
	for _, conf := range *cs {
		if conf.Key == key {
			return conf.Key, conf.Value
		}
	}
	return "", ""
}
