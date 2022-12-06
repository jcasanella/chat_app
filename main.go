package main

import (
	"fmt"
)

func main() {
	msg := sayHello("Jordi")
	fmt.Println(msg)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s, you are my new buddy", name)
}
