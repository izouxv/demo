package core

type ServiceConfig struct {
	Name      string
	Host      string
	Port      int
	User      string
	Password  string
	MaxIdle   int
	MaxActive int
	Enable    bool
	Interval  int
}
