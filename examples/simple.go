//go:build ignore
// +build ignore

/*
 * @Author: tj
 * @Date: 2022-08-30 17:55:41
 * @LastEditors: tj
 * @LastEditTime: 2022-08-30 17:56:39
 * @FilePath: \newfsm\examples\simple.go
 */

package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event("open", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event("close", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())
}
