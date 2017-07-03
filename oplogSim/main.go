package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetOutput(os.Stdout)

	for i := 0; i < 12; i++ {
		fmt.Println("This is a printl", i)
		log.Println("This is a LOG", i)
		time.Sleep(1 * time.Second)
		if i == 5 {
			panic("FUCK!")
		}
	}

}
