package main

import (
	"fmt"
	"os"

	"corsy"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

const VERSION = "v1.0.0"

func main() {

	logger := logrus.New()
	logger.Out = os.Stdout

	c := corsy.DefaultConfig()

	pflag.StringVarP(&c.ListenAddr, "addr", "a", ":8080", "Local address:port to listen on. Default: :8080 ")
	pflag.StringVarP(&c.Remote, "proxy", "p", "", "Remote address to proxy with cors")

	pflag.StringSliceVarP(&c.HeaderBlacklist, "blacklist", "b", c.HeaderBlacklist, "Headers to remove from the request and response")
	pflag.IntVarP(&c.MaxRedirects, "max-redirects", "r", c.MaxRedirects, "Maximum number of redirects to follow")
	pflag.IntVarP(&c.Timeout, "timeout", "t", c.Timeout, "Request timeout")

	help := pflag.BoolP("help", "h", false, "Show this message")
	version := pflag.BoolP("version", "v", false, "Corsy version and information")

	pflag.Parse()

	if *version {
		fmt.Println(VERSION + "")
		os.Exit(0)
	}

	if *help || pflag.NArg() != 0 || c.ListenAddr == "" {
		fmt.Fprintf(os.Stderr, "corsy [help] 📖 \n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		pflag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(0)
	}

	_ = corsy.NewClient(logger, c).Start()

}
