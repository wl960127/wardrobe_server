package parseconfig



import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

var cfg *ini.File

// PostgresDataSetting .
var PostgresDataSetting = &PostgresDatabase{}

// MysqlDataSetting .
var MysqlDataSetting= &MysqlDatabase{}

// ServerSetting .
var ServerSetting = &Server{}

// AppSetting .
var AppSetting = &App{}

// Server .
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// PostgresDatabase .
type PostgresDatabase struct {
	Type     string
	User     string
	Password string
	Host     string
	HTTPPort string
	Name     string
	SSLMode  string
}

// MysqlDatabase .
type MysqlDatabase struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

// App 其余配置
type App struct {
	JwtSecret string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string
}

func init() {
	var err error

	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
		os.Exit(1)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("postgres", PostgresDataSetting)
	mapTo("mysql", MysqlDataSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
