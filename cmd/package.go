package main

import (
  "fmt"
)

var CmdPackage = &Command {
  UsageLine:  "package [import path]",
  Short:      "package a Moogo application (e.g. for deployment)",
  Long:`
  这里是对这个命令的详细解释
  `,
}


func init() {
  CmdPackage.Run = packageApp
}

func packageApp(args []string) {
  if len(args) == 0 {
    errorf("No import path given.\nRun 'moogo help new' for usage.\n")
  }
  fmt.Println(args)
}