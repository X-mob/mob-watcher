/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/X-mob/mob-watcher/cmd"
	"github.com/X-mob/mob-watcher/task"
)

func main() {
	task.ScanNewMob()
	cmd.Execute()
}
