package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer > 0; timer-- {
		fmt.Println("You have", timer,
			"seconds to reach minimum safe distance.")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("BOOM!")
}
