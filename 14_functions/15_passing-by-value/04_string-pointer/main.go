package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println("main:", &name) // 0x82023c080

	changeMe(&name)

	fmt.Println("main:", &name) //0x82023c080
	fmt.Println("main:", name)  //Rocky
}

func changeMe(z *string) {
	fmt.Println("changeMe:", z)  // 0x82023c080
	fmt.Println("changeMe:", *z) // Todd
	*z = "Rocky"
	fmt.Println("changeMe:", z)  // 0x82023c080
	fmt.Println("changeMe:", *z) // Rocky
}
