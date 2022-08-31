//go:build ignore
// +build ignore

/*
 * @Author: tj
 * @Date: 2022-08-30 17:55:41
 * @LastEditors: tj
 * @LastEditTime: 2022-08-31 17:43:36
 * @FilePath: \fsm\examples\simple.go
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
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event(context.Background(), "open", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event(context.Background(), "close", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())
}
