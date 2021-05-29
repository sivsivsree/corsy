package corsy

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewClient(t *testing.T) {

	tests := []struct {
		name   string
		config *Config
		log    *logrus.Logger
		err    error
	}{
		{
			name: "no remote URL provided",
			log:  logrus.New(),
			config: &Config{
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
			},
			err: fmt.Errorf("no remote URL provided"),
		},

		{
			name: "invalid url",
			log:  logrus.New(),
			config: &Config{
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
				Remote: "invalid url",
			},
			err: fmt.Errorf("invalid remote URL %v provided", "invalid url"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			c := NewClient(tt.log, tt.config)
			go func() {
				err := c.Start()
				if err.Error() != tt.err.Error() {
					t.Errorf("failed the error: (%v/%v) %v", err, tt.err, err.Error() != tt.err.Error())
				}
			}()
			c.Stop()
		})
	}
}
