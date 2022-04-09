package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Printf("Interval Runner\n")

	var wait time.Duration
	flag.DurationVar(&wait, "initial-wait", 60*time.Second, "Initial wait time that can be parsed by `time.ParseDuration`")

	var interval time.Duration
	flag.DurationVar(&interval, "interval", 60*time.Second, "Interval time that can be parsed by `time.ParseDuration`")

	flag.Parse()
	args := flag.Args()

	fmt.Printf("initial wait = %v\n", wait)
	fmt.Printf("interval     = %v\n", interval)
	fmt.Printf("task         = %v\n", args)

	time.Sleep(wait)
	for {
		out, err := exec.Command(args[0], args[1:]...).Output()
		fmt.Printf("%v", string(out))
		if err != nil {
			fmt.Printf("[error] %v", err)
		}

		time.Sleep(interval)
	}
}
