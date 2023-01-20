package config

type AppConfig struct {
	Port        uint16
	MsSqlServer string
	MsSqlDb     string
	MsSqlUser   string
	MsSqlPwd    string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		Port:        8081, // FIXME: порт из ENV
		MsSqlServer: "loldbtest.cd.local",
		MsSqlDb:     "auth",
		MsSqlUser:   "sa",
		MsSqlPwd:    "",
	}
}
