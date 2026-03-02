package config

type AppConf struct {
	Mysql
	Redis
	RabbitMQ
}

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

type RabbitMQ struct {
	Username string
	Password string
	Host     string
	Port     int
}
