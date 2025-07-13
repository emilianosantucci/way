package configuration

type Configuration struct {
	Profile string                `mapstructure:"PROFILE"`
	DB      DatabaseConfiguration `mapstructure:",squash"`
	Web     WebConfiguration      `mapstructure:",squash"`
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
