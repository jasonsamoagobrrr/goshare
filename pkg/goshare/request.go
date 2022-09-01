package goshare

type Request struct {
	Method   string
	Host     string
	Url      string
	Headers  map[string]string
	Data     []byte
	Input    map[string][]byte
	Position int
	Raw      string
}

func NewRequest(conf *Config) Request {
	var req Request
	req.Method = conf.Method
	req.Url = conf.Url
	return req
}

func BaseRequest(conf *Config) Request {
	req := NewRequest(conf)
	req.Data = []byte(conf.Data)
	return req
}
