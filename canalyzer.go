package main

import (
	"fmt"
	"os"
)

var k uint32
var asm *os.File

func printCall(n *Node) {
	if n.c == nil {
		return
	}
	if len(n.c) != 2 {
		if len(n.c) == 4 {
			printCall(n.c[1])
			printCall(n.c[3])
		}
		return
	}
	fmt.Fprintln(asm, "    pushq %rbp")
	fmt.Fprintln(asm, "    movq %rsp, %rbp")
	fmt.Fprintln(asm, "    subq $32, %rsp")
	fmt.Fprintf(asm, "    leaq msg%d(%%rip), %%rcx\n", k)
	fmt.Fprintln(asm, "    call puts")
	fmt.Fprintln(asm, "    addq $32, %rsp")
	fmt.Fprintln(asm, "    popq %rbp")
	k++
	printCall(n.c[1])
	if len(n.c) == 4 {
		printCall(n.c[3])
	}
}

func printMsg(n *Node) {
	if n.c == nil {
		return
	}
	if len(n.c) != 2 {
		if len(n.c) == 4 {
			printMsg(n.c[1])
			printMsg(n.c[3])
		}
		return
	}
	fmt.Fprintf(asm, "msg%d:\n", k)
	fmt.Fprintf(asm, "    .asciz \"%s\"\n", n.c[0].c[0].v.get()[1:])
	k++
	printMsg(n.c[1])
}

func analyzer(n *Node, s string) error {
	var err error
	asm, err = os.Create(s + ".s")
	if err != nil {
		return err
	}
	defer asm.Close()
	fmt.Fprintln(asm, ".global main")
	fmt.Fprintln(asm, ".extern puts")
	fmt.Fprintln(asm, "main:")
	if len(n.c) > 0 {
		k = uint32(0)
		printCall(n.c[0])
	}
	fmt.Fprintln(asm, "    ret")
	if len(n.c) > 0 {
		k = uint32(0)
		printMsg(n.c[0])
	}
	return nil
}
