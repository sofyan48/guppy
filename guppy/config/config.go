package config

// Configs ...
type Configs struct {
	Urls        []string
	DialTimeOut int
}

// NewConfig ...
func NewConfig() *Configs {
	return &Configs{}
}
