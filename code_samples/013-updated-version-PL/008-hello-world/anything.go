package main

import "fmt"

func main() {
	fmt.Println("Hello Everyone, I'm having fun learning Go programming!!!")
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited...")
}
