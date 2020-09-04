package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println(runtime.Version())
	r := gin.Default()
	log.Println("http://localhost:8080")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog": "www.baidu.com",
		})
	})
	log.Println("http://localhost:8080/users/zhangsan")
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is %s", id)
	})
	//r.GET("/users/list", func(c *gin.Context) {
	//	c.String(200,"hello")
	//
	//})
	log.Println("http://localhost:8080/l/zhangsan")
	r.GET("/l/*list", func(c *gin.Context) {
		c.String(200, c.Param("list"))

	})
	log.Println("http://localhost:8080/query?name=zhangsan")
	r.GET("/query", func(c *gin.Context) {
		c.String(200, c.DefaultQuery("name", "张三"))
	})
	log.Println("http://localhost:8080/queryArray?media=zhangsan&media=lisi")
	r.GET("/queryArray", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})
	log.Println("http://localhost:8080/map?ids[a]=123&ids[b]=456&ids[c]=789")
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, c.QueryMap("ids"))
	})
	r.POST("/", func(c *gin.Context) {
		wechat := c.PostForm("wechat")
		c.String(200, wechat)
	})
	log.Fatal(r.Run(":8080"))

}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
