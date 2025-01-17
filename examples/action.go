//go:build ignore
// +build ignore

/*
 * @Author: tj
 * @Date: 2022-08-30 17:58:56
 * @LastEditors: tj
 * @LastEditTime: 2022-08-31 17:45:24
 * @FilePath: \fsm\examples\action.go
 */
package main

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "open", Src: []string{"closed"}, Dst: "open", ActionFunc: fsm.Callbacks{
				"action1": func(_ context.Context, e *fsm.Event) {
					fmt.Println("action1: " + e.FSM.Current())
				},
				"action2": func(_ context.Context, e *fsm.Event) {
					fmt.Println("action2: " + e.FSM.Current())
				},
			}},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event(context.Background(), "open", "action1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fsm.Current())

	err = fsm.Event(context.Background(), "open", "action2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event(context.Background(), "open", "action3")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())
}
