package environment

type Configuration struct {
	Profile    string                `mapstructure:"PROFILE"`
	DB         DatabaseConfiguration `mapstructure:",squash"`
	Web        WebConfiguration      `mapstructure:",squash"`
	Messaging  Messaging             `mapstructure:",squash"`
	Pagination Pagination            `mapstructure:",squash"`
}

type DatabaseConfiguration struct {
	Uri  string `mapstructure:"DATABASE_URI"`
	Host string `mapstructure:"DATABASE_HOST"`
	Port int    `mapstructure:"DATABASE_PORT"`
	User string `mapstructure:"DATABASE_USER"`
	Pass string `mapstructure:"DATABASE_PASSWORD"`
	Name string `mapstructure:"DATABASE_NAME"`
}

type WebConfiguration struct {
	Host string `mapstructure:"WEB_HOST"`
	Port int    `mapstructure:"WEB_PORT"`
}

type Messaging struct {
	ServerReadyTimeout int `mapstructure:"MESSAGING_SERVER_READY_TIMEOUT"`
}

type Pagination struct {
	MaxSize int `mapstructure:"PAGINATION_MAX_SIZE"`
}
