package engine

import (
    "fmt"
)

type PrintCommand struct {
    Arg string
}

func (print PrintCommand) Execute(loop Handler) {
    fmt.Println(print.Arg)
}

type PrintcCommand struct {
    Arg1 string
    Arg2 int
}

func (pccPrintcCommand) Execute(handler Handler) {
    res := pcc.Arg1
    for i := 0; i < pcc.Arg2; i++ {
        res = res + pcc.Arg1
    }
    handler.Post(&PrintCommand{Arg: res})
}

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
    h.(*EventLoop).stop = true
}