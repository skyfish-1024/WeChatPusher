package task

import (
	"github.com/robfig/cron"
	"log"
)

func Task() {
	i := 0
	c := cron.New()
	spec := "*/10 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		count := 0
		err := Push()
		for err != nil && count < 3 {
			count++
			err = Push()
		}
		if err != nil {
			//发送失败信息给我
			pushToMe("失败")
		} else {
			//发送成功信息给我
			pushToMe("成功")
		}
		log.Println("running:", i)
	})
	c.Start()

	select {}
}
