package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t)
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")

	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

	c_time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(c_time)

	t3, _ := time.Parse("2006-01-02 15:04:05", c_time)

	newTime := t3.AddDate(0, 0, 7)
	fmt.Println(newTime)
	//t3 = t3.In(time.UTC)
	timestamp := t3.Unix()
	fmt.Println(timestamp)

	t4 := time.Unix(timestamp, 0)
	p(t4)

	formattedTime := t4.Format("2006-01-02 15:04:05")

	fmt.Println(formattedTime)

	// t3, err := time.Parse("2006-01-02 15:04:05", c_time)
	// if err != nil {
	// 	fmt.Println("解析时间出错:", err)
	// 	return
	// }
	// t3 = t3.In(time.UTC) // 明确指定时区为UTC
	// timestamp = t3.Unix()
	// fmt.Println(timestamp)
}
