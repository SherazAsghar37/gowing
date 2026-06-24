package config

type Route struct {
	Location  string     `yaml:"location"`
	Strategy  string     `yaml:"strategy"`
	ProxyPass string     `yaml:"proxy_pass"`
	Upstreams []Upstream `yaml:"upstreams"`
}

type Upstream struct {
	URL string `yaml:"url"`
}
