package config

import "time"

// MySQL信息
type mysqlConfig struct {
	MasterHost                string        `yaml:"masterHost"`
	SlaverHost                []string        `yaml:"slaverHost"`
	Port                      string        `yaml:"port"`
	User                      string        `yaml:"user"`
	Password                  string        `yaml:"password"`
	Database                  string        `yaml:"database"`
	Charset                   string        `yaml:"charset"`   //要支持完整的UTF-8编码,需设置成: utf8mb4
	ParseTime                 bool          `yaml:"parseTime"` //解析time.Time类型
	TimeZone                  string        `yaml:"timeZone"`  // 时区,若设置 Asia/Shanghai,需写成: Asia%2fShanghai
	Gorm                      gormConfig    `yaml:"gorm"`
	SlowSql                   time.Duration `yaml:"slowSql"`                   //慢SQL
	LogLevel                  string        `yaml:"logLevel"`                  // 日志记录级别
	IgnoreRecordNotFoundError bool          `yaml:"ignoreRecordNotFoundError"` // 是否忽略ErrRecordNotFound(未查到记录错误)
	Enable                    bool          `yaml:"enable"`
}

// gorm 配置信息
type gormConfig struct {
	SkipDefaultTx   bool   `yaml:"skipDefaultTx"`                            //是否跳过默认事务
	PreparedStmt    bool   `yaml:"prepareStmt"`                              // 设置SQL缓存
	CloseForeignKey bool   `yaml:"disableForeignKeyConstraintWhenMigrating"` // 禁用外键约束
	TablePrefix     string `yaml:"tablePrefix"`                              // 表前缀
	SingularTable   bool   `yaml:"singularTable"`                            //是否使用单数表名(默认复数)，启用后，User结构体表将是user
}
