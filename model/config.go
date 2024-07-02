package model

type AppConfig struct {
	Port     int
	Database Database
}

type Database struct {
	Driver      string
	Credential  string
	MaxOpenConn int
	MaxIdleConn int
	MaxIdleTime int
	MaxLifeTime int
}
