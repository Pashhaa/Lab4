package engine

import (
	"fmt"
)

type printCommand struct {
	Arg string
}

func (print *printCommand) Execute(loop handler) {
	fmt.Println(print.Arg)
}

type PrintcCommand struct {
	Arg1 string
	Arg2 int
}

func (pcc *PrintcCommand) Execute(Handler handler) {
	res := pcc.Arg1
	for i := 1; i < pcc.Arg2; i++ {
		res = res + pcc.Arg1
	}
	Handler.Post(&printCommand{Arg: res})
}

type stopCommand struct{}

func (s stopCommand) Execute(h handler) {
	h.(*EventLoop).stop = true
}
