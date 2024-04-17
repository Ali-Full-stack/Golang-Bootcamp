package time

// display the current time with date/time/timezone
now :=time.Now()
fmt.Println(now)

// Unix time 1970 year january 1  till now
fmt.Println(now.Unix)