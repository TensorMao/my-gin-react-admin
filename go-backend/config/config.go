package config

type ServerConfig struct {
	// gorm
	App      AppConfig    `mapstructure:"app" json:"app" yaml:"app"`
	Database DBConfig     `mapstructure:"database" json:"database" yaml:"database"`
	Logger   LoggerConfig `mapstructure:"log" json:"log" yaml:"log"`
	Redis    RedisConfig  `mapstructure:"redis"`
	JWT      JWTConfig    `mapstructure:"jwt"`
}
