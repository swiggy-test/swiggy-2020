package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime.Format(time.RFC850))
}
