package initialize

import (
	"fmt"
	"github.com/fatih/color"
	"go-backend/global"
	"go-backend/models"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func InitDB() {
	switch global.GlobConfig.Database.Driver {
	case "mysql":
		initMySqlGorm()
	default:
		initMySqlGorm()
	}

}

func CloseDB() {
	db := global.GlobDB
	if db != nil {
		sqlDB, _ := db.DB()
		err := sqlDB.Close()
		if err != nil {
			global.GlobLogger.Error("Failed to connect mysql, err:", zap.Any("err", err))
			panic("Failed to close mysql, err:" + err.Error())
		}
	}
}

func initMySqlGorm() {
	dbCfg := global.GlobConfig.Database
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		dbCfg.UserName,
		dbCfg.Password,
		dbCfg.Host,
		strconv.Itoa(dbCfg.Port),
		dbCfg.Database,
		dbCfg.Charset)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   getGormLogger(), // 使用自定义 Logger
	})
	if err != nil {
		global.GlobLogger.Error("Failed to connect mysql, err:", zap.Any("err", err))
		panic("Failed to connect mysql, err:" + err.Error())
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbCfg.MaxIdleCons)
		sqlDB.SetMaxOpenConns(dbCfg.MaxOpenCons)
		initMySqlTables(db)
		global.GlobDB = db
		color.Blue("DB inits successfully\n")

	}

}

func getGormLogWriter() logger.Writer {
	var writer io.Writer
	// 是否启用日志文件
	if global.GlobConfig.Database.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.GlobConfig.Logger.RootDir + "/" + global.GlobConfig.Database.LogFilename,
			MaxSize:    global.GlobConfig.Logger.MaxSize,
			MaxBackups: global.GlobConfig.Logger.MaxBackups,
			MaxAge:     global.GlobConfig.Logger.MaxAge,
			Compress:   global.GlobConfig.Logger.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.GlobConfig.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  logMode,                // 日志级别
		IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  true,
	})
}

func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
		models.Media{},
	)
	if err != nil {
		global.GlobLogger.Error("migrate table failed", zap.Any("err", err))
		panic("Migrate table failed")
	}
}
