package runner

import (
	"github.com/imayberoot/ggshare/pkg/goshare"
	"net/http"
	"net/url"
	"time"
)

type BasicRunner struct {
	config *goshare.Config
	client *http.Client
}

func NewBasicRunner(conf *goshare.Config, replay bool) *interface{} {
	var basicrunner BasicRunner
	proxyURL := http.ProxyFromEnvironment
	customProxy := ""

	if replay {
		customProxy = conf.ReplayProxyURL
	} else {
		customProxy = conf.ProxyURL
	}
	if len(customProxy) > 0 {
		pu, err := url.Parse(customProxy)
		if err == nil {
			proxyURL = http.ProxyURL(pu)
		}
	}
	basicrunner.config = conf
	basicrunner.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       time.Duration(time.Duration(conf.Timeout) * time.Second),
		Transport: &http.Transport{
			Proxy:               proxyURL,
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 500,
			MaxConnsPerHost:     500,
		}}

	return &basicrunner
}
