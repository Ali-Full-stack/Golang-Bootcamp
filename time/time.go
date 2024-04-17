package time

// display the current time with date/time/timezone
now :=time.Now()
fmt.Println(now)

// Unix time 1970 year january 1  till now
fmt.Println(now.Unix)

// display time in different mode
fmt.Println(now.Day())
fmt.Println(now.Month())
fmt.Println(now.Year())
fmt.Println(now.Hour())
fmt.Println(now.Minute())
fmt.Println(now.Second())

// to sleep the program for 2 second
time.Sleep(2 * time.Second)