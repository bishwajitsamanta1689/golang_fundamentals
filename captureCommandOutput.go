package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main(){
	cmd, err:= exec.Command("ls","-lath").CombinedOutput()
	if err != nil {
		log.Fatalln("Error is", err)
	}
	fmt.Println("Command Output", string(cmd))
}
