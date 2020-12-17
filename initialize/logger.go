package initialize

import (
	"ff/g"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitLogger() {
	hook := lumberjack.Logger{
		Filename:   g.VP.GetString("log.filename"),   // 日志文件路径
		MaxSize:    g.VP.GetInt("log.max-file-size"), // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: g.VP.GetInt("log.max-backups"),   // 日志文件最多保存多少个备份
		MaxAge:     g.VP.GetInt("log.max-age"),       // 文件最多保存多少天
		Compress:   g.VP.GetBool("log.compress"),     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(zap.InfoLevel)) // 日志级别

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()

	// 开启文件及行号
	development := zap.Development()

	// 构建日志
	g.Logger = zap.New(core, caller, development)
	fmt.Println("==>成功加载日志配置！")
}
