package main

import "github.com/jcasanella/chat_app/config"

func main() {
	// msg := sayHello("Jordi")
	// fmt.Println(msg)
	config.Init("local")
	config.GetConfig()
}

// func sayHello(name string) string {
// 	return fmt.Sprintf("Hello %s, you are my new buddy", name)
// }
