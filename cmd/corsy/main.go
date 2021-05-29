package main

import (
	"fmt"
	"os"

	"corsy"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {

	logger := logrus.New()
	logger.Out = os.Stdout
	//logger.Formatter = &logrus.JSONFormatter{}

	c := corsy.Config{
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

	pflag.StringVarP(&c.ListenAddr, "addr", "a", c.ListenAddr, "address:port to listen on :8080 ")
	pflag.StringVarP(&c.Remote, "proxy", "p", c.Remote, "remote address to proxy with cors")

	pflag.StringSliceVarP(&c.HeaderBlacklist, "blacklist", "b", c.HeaderBlacklist, "Headers to remove from the request and response")
	pflag.IntVarP(&c.MaxRedirects, "max-redirects", "r", c.MaxRedirects, "Maximum number of redirects to follow")
	pflag.IntVarP(&c.Timeout, "timeout", "t", c.Timeout, "Request timeout")

	help := pflag.BoolP("help", "h", false, "Show this message")

	pflag.Parse()

	if *help || pflag.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "Usage: corsy [OPTIONS]\n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		pflag.PrintDefaults()
		os.Exit(1)
	}

	_ = corsy.NewClient(logger, &c).Start()

}
