package main

import (
  "fmt"
)

var CmdNew = &Command {
  UsageLine:  "new [path]",
  Short:      "create a skeleton Moogo application",
  Long: `
  这里是对这个命令的详细解释
  `,
}

func init() {
  CmdNew.Run = newApp
}

func newApp(args []string) {
  if len(args) == 0 {
    errorf("No import path given.\nRun 'moogo help new' for usage.\n")
  }
  fmt.Println(args)
}