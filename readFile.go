package main

import (
	"log"
	"os"
)

func readArgument() string {
	args := os.Args
	if len(args) <= 1 {
		log.Fatalln("Too few arguments, Enter required argument.")
	}
	if len(args) >= 3 {
		log.Fatalln("Too much arguments, Only one argument needed.")
	}
	return args[1]
}
