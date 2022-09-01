package goshare

import "context"

type Config struct {
	Cancel              context.CancelFunc `json:"-"`
	Colors              bool               `json:"colors"`
	CommandKeywords     []string           `json:"-"`
	CommandLine         string             `json:"cmdline"`
	ConfigFile          string             `json:"configfile"`
	Context             context.Context    `json:"-"`
	Data                string             `json:"postdata"`
	Delay               optRange           `json:"delay"`
	Json                bool               `json:"json"`
	MaxTime             int                `json:"maxtime"`
	MaxTimeJob          int                `json:"maxtime_job"`
	Method              string             `json:"method"`
	OutputDirectory     string             `json:"outputdirectory"`
	OutputFile          string             `json:"outputfile"`
	OutputFormat        string             `json:"outputformat"`
	OutputSkipEmptyFile bool               `json:"OutputSkipEmptyFile"`
	ProgressFrequency   int                `json:"-"`
	ProxyURL            string             `json:"proxyurl"`
	Quiet               bool               `json:"quiet"`
	Rate                int64              `json:"rate"`
	Recursion           bool               `json:"recursion"`
	RecursionDepth      int                `json:"recursion_depth"`
	RecursionStrategy   string             `json:"recursion_strategy"`
	ReplayProxyURL      string             `json:"replayproxyurl"`
	SNI                 string             `json:"sni"`
	StopOnErrors        bool               `json:"stop_errors"`
	Threads             int                `json:"threads"`
	Timeout             int                `json:"timeout"`
	Url                 string             `json:"url"`
	Verbose             bool               `json:"verbose"`
}

type InputProviderConfig struct {
	Name     string `json:"name"`
	Keyword  string `json:"keyword"`
	Value    string `json:"value"`
	Template string `json:"template"` // the templating string used for sniper mode (usually "§")
}

func NewConfig(ctx context.Context, cancel context.CancelFunc) Config {
	var conf Config
	conf.CommandKeywords = make([]string, 0)
	conf.Context = ctx
	conf.Cancel = cancel
	conf.Data = ""
	conf.Delay = optRange{0, 0, false, false}
	conf.Json = false
	conf.MaxTime = 0
	conf.MaxTimeJob = 0
	conf.Method = "POST"
	conf.ProgressFrequency = 125
	conf.ProxyURL = ""
	conf.Quiet = false
	conf.Rate = 0
	conf.Recursion = false
	conf.RecursionDepth = 0
	conf.RecursionStrategy = "default"
	conf.SNI = ""
	conf.StopOnErrors = false
	conf.Timeout = 10
	conf.Url = ""
	conf.Verbose = false
	return conf
}

func (c *Config) SetContext(ctx context.Context, cancel context.CancelFunc) {
	c.Context = ctx
	c.Cancel = cancel
}
