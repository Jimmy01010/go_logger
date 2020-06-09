package app


type Config struct{
	// Log to syslog. Defaults to writing to stdout/stderr.
	Syslog bool `json:"syslog"`
}


func NewConfig()(*Config){
	return &Config{}
}