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

type DBConfig struct {
	DBHost     	string
	DBPort     	int
	DBName     	string
	DBUser     	string
	DBPwd		interface{}
}

type RpcConfig struct {
	Address string
	Port    string
}

type Const struct {
	//todo 默认头像ID
	UserAvatar	string
	PetAvatar   string
	//todo 宠物默认训练项
	Name1       string
	Name2       string
	Name3       string
	Voice1      string
	Voice2      string
	Voice3      string
	PetDuration	int32
	//todo 文件服务地址
	FileServer  string
	//todo 通知服务地址
	NoticeServer  string
	//todo mysql异常字符串
	NotFound    string
}
