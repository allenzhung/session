package main

import (
	"fmt"
	"log"
	"time"

	"session/sessions"
	sqlitecookie "session/sqlite" //package namme sqlitecookie

	"github.com/gin-gonic/gin"
)

func main() {
	mm()

	return

	tt()

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	store := sqlitecookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		val, _ := c.Cookie("mysession")
		log.Println("Cookie:%s", val)

		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}

func tt() {
	currentTime := time.Now()                                            //获取当前时间，类型是Go的时间类型Time
	t1 := currentTime.Year()                                             //年
	t2 := currentTime.Month()                                            //月
	t3 := currentTime.Day()                                              //日
	t4 := currentTime.Hour()                                             //小时
	t5 := currentTime.Minute()                                           //分钟
	t6 := currentTime.Second()                                           //秒
	t7 := currentTime.Nanosecond()                                       //纳秒
	currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local) //获取当前时间，返回当前时间Time
	fmt.Println("t6=", t6)
	fmt.Println("t7=", t7)
	fmt.Println(currentTime)                                            //打印结果：2017-04-11 12:52:52.794351777 +0800 CST<br>    fmt.Println(t1,t2,t3,t4,t5,t6)     //打印结果：2017 April 11 12 52 52
	fmt.Println(currentTime.Format("2006-01-02 15:04:05.99999999 MST")) //打印结果：2017-04-11 12:52:52.794351777 +0800 CST<br>    fmt.Println(t1,t2,t3,t4,t5,t6)     //打印结果：2017 April 11 12 52 52
	fmt.Println(currentTimeData)
	fmt.Println(currentTimeData.Format("2006-01-02 15:04:05.99999999 MST")) //打印结果：2017-04-11 12:52:52.794351777 +0800 CST<br>    fmt.Println(t1,t2,t3,t4,t5,t6)     //打印结果：2017 April 11 12 52 52

	fmt.Println("-------------------------------")
	now := time.Now()

	local1, err1 := time.LoadLocation("") //等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local") //服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles")
	if err3 != nil {
		fmt.Println(err3)
	}
	local4, err4 := time.LoadLocation("Asia/Chongqing")
	if err4 != nil {
		fmt.Println(err4)
	}

	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))
	fmt.Println(now.In(local4).Format(time.RFC3339))
	fmt.Println(now.In(local1).Format("2006-01-02T15:04:05.000000000Z07:00"))
	//output:
	// t6= 39
	// t7= 484106300
	// 2019-08-07 10:59:39.4841063 +0800 CST m=+0.006995901
	// 2019-08-07 10:59:39.4841063 CST
	// 2019-08-07 10:59:39.4841063 +0800 CST
	// 2019-08-07 10:59:39.4841063 CST
	// -------------------------------
	// 2019-08-07 02:59:39.5430709 +0000 UTC
	// 2019-08-07 10:59:39.5430709 +0800 CST
	// 2019-08-06 19:59:39.5430709 -0700 PDT
	// 2019-08-07T10:59:39+08:00
	// 2019-08-07T02:59:39.543070900Z

	fmt.Println("-------------------------------")
	timeFormat1 := "2006-01-02T15:04:05.000000000Z07:00"
	ts1 := now.In(local4).Format(timeFormat1) //"Asia/Chongqing"
	time1, _ := time.Parse(timeFormat1, ts1)
	fmt.Println("ts1  =", ts1)
	fmt.Println("time1=", time1)
	fmt.Println("loc1 =", time1.In(local1).Format("2006-01-02T15:04:05.999Z"))

	timeFormat2 := "2006-01-02T15:04:05.000000000"
	ts2 := now.In(local1).Format(timeFormat2)                  //UTC
	time2, _ := time.ParseInLocation(timeFormat2, ts2, local4) //"Asia/Chongqing"
	fmt.Println("ts2  =", ts2)
	fmt.Println("time2=", time2)
	fmt.Println("loc1 =", time2.In(local1)) //"Asia/Chongqing"
}
