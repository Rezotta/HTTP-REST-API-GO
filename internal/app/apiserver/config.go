package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
}


//new Config returns a new instance of Config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}