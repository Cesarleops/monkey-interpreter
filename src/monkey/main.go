package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)


func main(){

	//An interpreter for the Monkey programming language

	//The main function is the entry point of the interpreter

	user, err := user.Current()
	
	if err != nil {
		panic(err)
	}


	fmt.Printf("Hello %s! This is the Monkey programming language!\n" ,user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
