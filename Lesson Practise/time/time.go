package time

import (
	"math/rand"
	"time"
)

// display the current time with date/time/timezone
// now :=time.Now()
// fmt.Println(now)

// // Unix time 1970 year january 1  till now
// fmt.Println(now.Unix)

// // display time in different mode
// fmt.Println(now.Day())
// fmt.Println(now.Month())
// fmt.Println(now.Year())
// fmt.Println(now.Hour())
// fmt.Println(now.Minute())
// fmt.Println(now.Second())

// // to sleep the program for 2 second
// time.Sleep(2 * time.Second)

// // to find out difference
// time.Since(now)

// to createrandom date
func RandomDate() (t time.Time) {
	year := rand.Intn(2024)
	month := rand.Intn(12)
	day := rand.Intn(29)
	hour := rand.Intn(24)
	minute := rand.Intn(60)
	second := rand.Intn(60)
	location, _ := time.LoadLocation("Asia/Tashkent")

	t = time.Date(year,time.Month(month),day,hour,minute,second,0,location)
	return t

}
