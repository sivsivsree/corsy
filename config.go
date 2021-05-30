package corsy

// Config is used to store corsy inital configurations
type Config struct {
	HopHeaders      []string
	HeaderBlacklist []string
	MaxRedirects    int
	Timeout         int
	ListenAddr      string

	Remote string
}
