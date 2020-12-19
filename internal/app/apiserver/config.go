package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	StaticPath  string `toml:"static_path"`
	Prod        bool   `toml:"prod"`
	ConsoleLog  bool   `toml:"console_log"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		LogLevel:   "debug",
		Prod:       true,
		ConsoleLog: false,
	}
}
