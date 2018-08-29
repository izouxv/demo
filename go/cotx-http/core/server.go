package core

type ServiceConfig struct {
	Name      string
	Host      string
	Port      int
	MaxIdle   int
	MaxActive int
	Enable    bool
	Interval  int
	Password  string
}
