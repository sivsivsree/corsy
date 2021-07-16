package corsy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Log interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Client struct {
	*http.Client

	srv    *http.Server
	config *Config
	log    Log
}

func NewClient(logger Log, config *Config) *Client {

	c := &Client{
		config: config,
		log:    logger,
	}

	c.Client = &http.Client{
		Timeout: time.Second * time.Duration(c.config.Timeout),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= c.config.MaxRedirects {
				return fmt.Errorf("stopped after %d redirects", c.config.MaxRedirects)
			}
			return nil
		},
	}

	return c
}

func (c *Client) handleCORS(w http.ResponseWriter, req *http.Request) {

	if req.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		return
	}
	p := strings.TrimLeft(req.URL.Path, "/")

	if q := req.URL.RawQuery; q != "" {
		p += "?" + q
	}

	u, err := url.Parse(c.config.Remote + "/" + p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		c.log.Errorf("error: %v", err)
		return
	}

	if u.Scheme == "" {
		u.Scheme = "http"
	}

	nreq, err := http.NewRequest(req.Method, u.String(), req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		c.log.Errorf("error: %v", err)
		return
	}

	for key := range req.Header {
		for _, ignore := range append(c.config.HopHeaders, c.config.HeaderBlacklist...) {
			if key == ignore {
				continue
			}
		}
		nreq.Header.Set(key, req.Header.Get(key))
	}

	c.log.Infof("Request Proxied -> / %s", p)
	resp, err := c.Do(nreq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		c.log.Errorf("Error: %v", err)
		return
	}

	defer resp.Body.Close()

	expose := []string{"X-Request-URL"}
	for key := range resp.Header {
		for _, ignore := range append(c.config.HopHeaders, c.config.HeaderBlacklist...) {
			if key == ignore {
				continue
			}
		}
		w.Header().Set(key, resp.Header.Get(key))
		expose = append(expose, key)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", strings.Join(expose, ","))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("X-Request-URL", u.String())
	w.Header().Set("X-Final-URL", resp.Request.URL.String())

	w.WriteHeader(resp.StatusCode)

	if _, err = io.Copy(w, resp.Body); err != nil {
		c.log.Warnf("error: %s", err)
		return
	}
}

func (c *Client) Start() error {
	if c.config.Remote == "" {
		return fmt.Errorf("no remote URL provided")
	}

	if _, err := url.ParseRequestURI(c.config.Remote); err != nil {
		return fmt.Errorf("invalid remote URL %v provided", c.config.Remote)
	}

	c.log.Infof("CORSY started on %v proxing %v with CORS", c.config.ListenAddr, c.config.Remote)
	c.srv = &http.Server{Addr: c.config.ListenAddr, Handler: http.HandlerFunc(c.handleCORS)}

	if err := c.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("listening failed")
	}

	return nil
}

func (c *Client) Stop() {
	if c.srv != nil {
		_ = c.srv.Close()
	}
}
