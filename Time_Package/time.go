package main

import (
	"fmt"
	"time"
)

func main (){
	formatTime := time.Now()
	currentTime := formatTime.Format("2006/01/02, 15:20:40")
	fmt.Println(currentTime)
}