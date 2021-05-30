package corsy

// Config is used to store corsy initial configurations
type Config struct {
	HopHeaders      []string
	HeaderBlacklist []string
	MaxRedirects    int
	Timeout         int
	ListenAddr      string

	Remote string
}

func DefaultConfig() *Config {
	return &Config{
		MaxRedirects: 10,
		Timeout:      15,
		ListenAddr:   ":8001",
		HopHeaders: []string{
			"Connection",
			"Keep-Alive",
			"Public",
			"Proxy-Authenticate",
			"Transfer",
			"Upgrade",
		},
	}
}
