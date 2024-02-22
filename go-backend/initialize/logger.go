package initialize

import (
	"github.com/fatih/color"
	"go-backend/global"
	"go-backend/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitLogger() {
	// 创建根目录
	CreateRootDir()
	// 设置日志等级
	SetLogLevel()

	logg := zap.New(getZapCore(), options...)

	// 创建logger实例
	zap.ReplaceGlobals(logg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	global.GlobLogger = logg // 注册到全局变量中
	color.Blue("Logger inits successfully\n")
}

func CreateRootDir() {
	if ok, _ := utils.PathExists(global.GlobConfig.Logger.RootDir); !ok {
		_ = os.Mkdir(global.GlobConfig.Logger.RootDir, os.ModePerm)
	}

}

func SetLogLevel() {
	switch global.GlobConfig.Logger.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	if global.GlobConfig.Logger.ShowLine {
		options = append(options, zap.AddCaller())
	}
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder
	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.GlobConfig.App.Env + "." + l.String())
	}

	// 设置编码器
	if global.GlobConfig.Logger.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.GlobConfig.Logger.RootDir + "/" + global.GlobConfig.Logger.Filename,
		MaxSize:    global.GlobConfig.Logger.MaxSize,
		MaxBackups: global.GlobConfig.Logger.MaxBackups,
		MaxAge:     global.GlobConfig.Logger.MaxAge,
		Compress:   global.GlobConfig.Logger.Compress,
	}
	return zapcore.AddSync(file)
}
