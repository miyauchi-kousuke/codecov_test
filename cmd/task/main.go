package main

import (
	"flag"
	"fmt"
)

type args []string

func (a *args) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func (a *args) String() string {
	return "Task IDs to be excuted"
}

var (
	taskIDs = args{}
	debug   = flag.Bool("debug", false, "debug log level flag")
)

func main() {
	flag.Var(&taskIDs, "task", "task id")
	flag.Parse()

	fmt.Println(taskIDs)
}

