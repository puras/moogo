package main

import (
  "fmt"
)

var CmdRun = &Command {
  UsageLine:  "run [import path] [run mode] [port]",
  Short:      "run a Moogo application",
  Long: `
  这里是对这个命令的详细解释
  `,
}


func init() {
  CmdRun.Run = runApp
}

func runApp(args []string) {
  if len(args) == 0 {
    errorf("No import path given.\nRun 'moogo help new' for usage.\n")
  }
  fmt.Println(args)
}