package config

// Url:scheme://host:port/path?query#fragment
type Url struct {
	Scheme string
	Host   string
	Port   string
	Path   string
	Query  string
}
