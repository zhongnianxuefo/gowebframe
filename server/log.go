package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// 初始化日志
func getInitLogger(filePath string) (logger *zap.SugaredLogger, err error) {

	encoder := getEncoder()
	//两个判断日志等级的interface
	//WarnLevel以下属于info
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
	//WarnLevel及以上属于warn
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter, err := getLogWriter(filePath + "/info.log")
	if err != nil {
		return
	}
	warnWriter, err := getLogWriter(filePath + "/warn.log")
	if err != nil {
		return
	}

	//创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, os.Stdout, infoLevel),
		zapcore.NewCore(encoder, infoWriter, infoLevel),
		zapcore.NewCore(encoder, warnWriter, warnLevel),
	)
	log := zap.New(core, zap.AddCaller())
	logger = log.Sugar()
	return
}

// 日志编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志写入接口
func getLogWriter(filePath string) (writer zapcore.WriteSyncer, err error) {
	//now := time.Now()
	//fileName := filePath + "_" + now.Format("20060102") + ".log"
	//fmt.Println(fileName)
	lumberJackLogger := &lumberjack.Logger{

		Filename:   filePath,
		MaxSize:    100, // 单个文件最大100M
		MaxBackups: 60,  // 多于 60 个日志文件后，清理较旧的日志
		MaxAge:     1,   // 一天一切割
		Compress:   false,
	}
	writer = zapcore.AddSync(lumberJackLogger)
	return
}
