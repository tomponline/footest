package main

import "fmt"
import "github.com/tomponline/footest/submod/helloworld"
import "github.com/tomponline/footest/somepkg"

func main() {
	fmt.Println(helloworld.Message+somepkg.Message)
}
