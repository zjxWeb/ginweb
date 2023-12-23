package main

// go get -u github.com/gin-gonic/gin
import "github.com/gin-gonic/gin"

func main(){
	//创建一个默认的路由引擎
	r := gin.Default()

	// 注册一个路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	//  启动HTTP服务  默认在0.0.0.0:8000
	r.Run(":8000")
}