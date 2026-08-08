package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "用法: roman <to|from> <数字或罗马串>")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "to":
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "请输入整数")
			os.Exit(1)
		}
		r := ToRoman(n)
		if r == "" {
			fmt.Fprintln(os.Stderr, "只支持 1~3999")
			os.Exit(1)
		}
		fmt.Println(r)
	case "from":
		n, err := FromRoman(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(n)
	default:
		fmt.Fprintln(os.Stderr, "只支持 to / from")
		os.Exit(2)
	}
}
