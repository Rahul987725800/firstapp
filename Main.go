package main

import "fmt"
import "github.com/Pallinder/go-randomdata"
import "github.com/Rahul987725800/firstapp/maths"

func main() {
	fmt.Println("Hello go")
	fmt.Println(maths.Add(1, 2))
	fmt.Println(randomdata.SillyName())
}