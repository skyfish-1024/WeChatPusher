package main

import (
	"WeChatPusher/db"
	"WeChatPusher/task"
	"fmt"
)

func main() {
	err := db.InitDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	task.Task()
}
