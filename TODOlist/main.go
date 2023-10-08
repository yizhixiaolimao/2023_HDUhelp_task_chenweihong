package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

type TODO struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

var todos []TODO

func main() {
	gin.SetMode(gin.TestMode)
	var user string
	var password int
	fmt.Println("请输入用户名")
	fmt.Scanln(&user)
	fmt.Println("请输入密码")
	fmt.Scanln(&password)

	if user == "hdu" && password == 123456 {
		fmt.Println("登录成功")
		r := gin.Default()

		// 添加 TODO
		r.POST("/todo", func(c *gin.Context) {
			var todo TODO
			c.BindJSON(&todo)
			todos = append(todos, todo)
			fmt.Println(todos)

			c.JSON(200, gin.H{"status": "ok"})
		})
		// 删除 TODO
		r.DELETE("/todo/:index", func(c *gin.Context) {
			index, _ := strconv.Atoi(c.Param("index"))
			todos = append(todos[:index], todos[index+1:]...)
			c.JSON(200, gin.H{"status": "ok"})
		})
		// 修改 TODO
		r.PUT("/todo/:index", func(c *gin.Context) {
			index, _ := strconv.Atoi(c.Param("index"))
			var todo TODO
			c.BindJSON(&todo)
			todos[index] = todo
			c.JSON(200, gin.H{"status": "ok"})
		})

		// 获取 TODO
		r.GET("/todo", func(c *gin.Context) {
			c.JSON(200, todos)
		})
		// 查询 TODO
		r.GET("/todo/:index", func(c *gin.Context) {
			index, _ := strconv.Atoi(c.Param("index"))
			c.JSON(200, todos[index])
		})
		var baocuntodo []TODO
		for index1 := range todos {
			baocuntodo = append(baocuntodo, todos[index1])

		}
		r.Run(":8080")

	} else {
		fmt.Println("你没有权限")
	}
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	jsontodo, _ := json.Marshal(todos)
	n, err := f.Write(jsontodo)
	if err == nil && n < len(jsontodo) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
