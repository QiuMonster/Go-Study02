package main

//导入自己的包 调用自己的方法
import "golang/golang2/service"

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

func main() {
	// go_defer()
	// go_init()
	// new_type()
	// go_struct()
	// intface_intface()
	// go_ocp()
	// go_extend()
	// go_constructor()
	service.TestService() //调用自己包中的方法

	//使用扩展包 使用
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// go_go()
	// go_go2()
	// go_chan()
	// go_witegroup()
	// go_runtime()
	// go_mutex()
	// go_for_chan()
	// go_select()
	// go_timer()
	// go_ticker()
	// go_atomic()
	go_atomic_detail()
}
