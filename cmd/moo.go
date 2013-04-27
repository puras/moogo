package main

import (
  "flag"
  "fmt"
  "io"
  "os"
  "strings"
  "text/template"
)

// 定义命令结构体
type Command struct {
  Run                     func(args []string) // 命令的执行方法
  UsageLine, Short, Long  string              // 命令、短说明、长说明
}

//
type LoggedError struct {
  error
}

// 获取命令的名称
func (cmd *Command) Name() (name string) {
  name = cmd.UsageLine
  // 获取命令中的空格，以空格分隔开命令与其参数
  i := strings.Index(name, " ")
  if i >= 0 {
    name = name[:i]
  }
  return
}

// 如何从new.go文件中引入相应的内容
// 将其进行编译后便可导入


// 定义命令列表
var commands = []*Command {
  CmdNew,
  CmdRun,
  CmdPackage,
}

func main() {
  fmt.Fprintf(os.Stdout, header)
  flag.Usage = usage
  flag.Parse()
  args := flag.Args()

  if len(args) < 1 || args[0] == "help" {
    if len(args) > 1 {
      for _, cmd := range commands {
        if cmd.Name() == args[1] {
          tmpl(os.Stdout, helpTemplate, cmd)
          return 
        }
      }
    }
    usage()
  }

  defer func() {
    if err := recover(); err != nil {
      if _, ok := err.(LoggedError); !ok {
        panic(err)
      }
      os.Exit(1)
    }
  }()

  for _, cmd := range commands {
    if cmd.Name() == args[0] {
      fmt.Println(args)
      cmd.Run(args[1:])
      return
    }
  }
  errorf("unkonwn command %q\nRun 'moogo help' for usage.\n", args[0])
}

func errorf(format string, args ...interface{}) {
  // HasSuffix 判断字符串format是否是以\n结尾
  if !strings.HasSuffix(format, "\n") {
    format += "\n"
  }
  fmt.Fprintf(os.Stderr, format, args...)
  panic(LoggedError{})
}

const header = `~
~ MooGo! http://puras.github.com/moogo
~
`

const usageTemplate = `usage: moogo command [arguments]

The commands are:
{{range .}}
  {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "moogo help [command] for more information."
`

var helpTemplate = `usage: moogo {{.UsageLine}}
{{.Long}}
`

func usage() {
  tmpl(os.Stderr, usageTemplate, commands)
  os.Exit(2)
}

func tmpl(w io.Writer, text string, data interface{}) {
  t := template.New("top")
  template.Must(t.Parse(text))
  if err := t.Execute(w, data); err != nil {
    panic(err)
  }
}