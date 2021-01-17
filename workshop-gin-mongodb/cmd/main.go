package main

import (
	"demo"
	_ "demo/docs"
	"fmt"
	"time"
	_ "time/tzdata"
)

func main() {
	fmt.Print("Local time zone ")
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now())
	demo.StartServer()
}
