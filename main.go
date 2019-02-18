package main

import (
	"github.com/aidansteele/toyrobot/pkg/toyrobot"
	"os"
)

func main() {
	toyrobot.RunStandardSimulation(os.Stdin, os.Stdout)
}
