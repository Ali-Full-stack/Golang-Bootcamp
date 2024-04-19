package main

import (
	"fmt"
	p "paralelism/paralel"
)

func main(){
	// Goroutine
	// go p.Say("hello")
	// p.Say("world")

	// // Channels
	// s :=[]int{7,2,8,-9,4,0}

	// c :=make(chan int)

	// go p.Sum(s[:len(s)/2],c)
	// go p.Sum(s[len(s)/2:],c)
	// x,y := <-c,<-c // receive from chanel

	// fmt.Println(x,y,x+y)

	ch :=make(chan int)

	go p.Sum2(1,2,ch)
	fmt.Println(<-ch)
}