package main

import (
	"fmt"
	"os"
	"os/user"

	"monkeylanguage.ayrtonguttier.com.br/repl"
)

func main(){

    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! This is the Monkey language!\n", user.Username)
    fmt.Println("Feel free to type in commands")


    repl.Start(os.Stdin, os.Stdout)
}
