package main
import (
  "github.com/sirupsen/logrus"
  
  "github.com/gin-gonic/gin"
  "errors"
  )

func configureLogger() *logrus.Logger {
	logger := logrus.New()
	customFormatter := &logrus.TextFormatter{}
	customFormatter.TimestampFormat = "02.01.2006 15:04:05"//"02.01.2006 15:04:05"//"dd.mm.yyyy HH24:MI:SS"
	customFormatter.FullTimestamp = true
	logger.SetFormatter(customFormatter)
	return logger
}

func main(){
test := configureLogger()
//log := logrus.New()
// hooks, config,...

r := gin.New()
//r.Use(ginlogrus.Logger(log), gin.Recovery())

// pingpong
r.GET("/ping", func(c *gin.Context) {
	
	test.Errorf("begin, cant parse page param to int, requestId:%s, username:%s, error:%v",
					100,
					"Malik",
					errors.New("Test Error"))
	c.Data(200, "text/plain", []byte("pong"))
	test.Errorf("end, cant parse page param to int, requestId:%s, username:%s, error:%v",
					100,
					"Malik",
					errors.New("Test Error"))
})

r.Run("127.0.0.1:8080")

}