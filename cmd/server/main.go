package main

import "fmt"

// Run - is responsible for the instantation and startup of the app
func Run() error {
	fmt.Println("starting up the app")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		// custom error handling logic
		fmt.Println(err)
	}
}
