package main

import (
	"github.com/gin-gonic/gin"
        "fmt"
        "os"
        "io"
        "bye_school/services"
)

func main() {
        r := gin.Default()
        r.GET("/ping", func(c *gin.Context) {
                c.JSON(200, gin.H{
                    "message": "pong",
                })
        })
        r.POST("/push", func(c *gin.Context) {
                //通过gin框架构造的http请求头
                msg := c.PostForm("message")
                user_id := c.PostForm("userId")
                title := c.PostForm("title")
                //获取message，userid，title参数，其中message为推送消息的内容部分，
                // userid为使用android客户端的用户id可以在管理员后台看到，
                // 推送服务通过这个这个userid获得用户相关手机的唯一标识符。
                // title为推送消息的标题
                f,_ := os.OpenFile("log.txt",os.O_APPEND|os.O_RDWR,os.ModePerm)
                // 打开日志文件
                if err := services.SjPush(user_id, msg, title); err == nil {
                // 调用推送服务service包SjPush服务
                        c.JSON(200, map[string]string{"type": "ok"})
                // 返回http码200，{"type":"ok"}数据到管理后台
                } else {
                        res := err.Error() + "\n"  
                        f.WriteString(res)
                        // 否则在错误日志中写入错误原因
                        c.JSON(200, map[string]string{"type": "err"})
                        // 返回http码200，{"type":"err"}数据到管理后台
                }
                f.Close()
        })
        r.POST("/up_pic", func(c *gin.Context) {
                file, header, _ := c.Request.FormFile("file")
                filename := c.PostForm("filename")
                article_id := c.PostForm("article_id")
                fmt.Println(file)
                fmt.Println(header.Filename)
                fmt.Println(article_id)
                f,_ := os.Create(filename)
                io.Copy(f,file)
                defer f.Close()   
        })
        r.Run() // listen and server on 0.0.0.0:2333
        }


