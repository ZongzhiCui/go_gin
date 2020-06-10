package common

import (
	"time"

	"ZongzhiCui/go_gin/config"

	"go.uber.org/zap/zapcore"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func OutputJson(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
		"time": time.Now(),
	})
}

// 判断文件夹是否存在
//func PathExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}

//basePath是固定目录路径 basePath = storage/log/
//func CreateDateDir(basePath string) (dirPath, dataString string) {
//	folderName := time.Now().Format("20060102")
//	folderPath := filepath.Join(basePath, folderName)
//	if exist, _ := PathExists(folderPath); !exist {
//		// 必须分成两步
//		// 先创建文件夹
//		_ = os.MkdirAll(folderPath, os.ModePerm)
//		// 再修改权限
//		_ = os.Chmod(folderPath, os.ModePerm)
//	}
//	return folderPath, folderName
//}

//返回zap的logger实例
func ZapLogger() (logger *zap.Logger) {
	//logger, _ = zap.NewProduction(zap.ErrorOutput(zapcore.AddSync(gin.DefaultWriter)))
	//core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(gin.DefaultWriter), zapcore.DebugLevel)

	cfg := config.ZapConf()

	//自定义日志级别：自定义Info级别
	logLevel := config.ZapLogLevel
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	//声明一个mutiCore，输出到多个终端
	core := zapcore.NewTee(
		// 每个终端都是由ioCore来实现
		zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(gin.DefaultWriter), infoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(gin.DefaultErrorWriter), warnLevel),
	)
	//原文链接：https://blog.csdn.net/lfhlzh/article/details/106194463
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	return
}
