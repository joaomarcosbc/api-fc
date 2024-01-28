package main

import "github.com/joaomarcosbc/api-fc/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
