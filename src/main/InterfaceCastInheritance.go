package main

import "fmt"

type ATester interface {
	test()
}

type BTester interface {
	test()
}

type Subject struct {
	name string
}

func (subject *Subject) test() {
	fmt.Println(subject)
}

func main() {
	interfaceDemo()
	fmt.Println("-----------")
	interfaceDemo2()
	fmt.Println("-----------")
	interfaceDemo3()
	fmt.Println("-----------")
}

func interfaceDemo() {
	var testerA ATester = &Subject{"Test"}
	var testerB BTester = testerA
	testerA.test()
	testerB.test()
}

type SuperTester interface {
	stest()
}

type ParentTester interface {
	ptest()
}

type ChildTester interface {
	SuperTester
	ParentTester
	ctest()
}

type Topic struct {
	name string
}

func (topic *Topic) ptest() {
	fmt.Printf("ptest %s\n", topic)
}

func (topic *Topic) ctest() {
	fmt.Printf("ctest %s\n", topic)
}

func (topic *Topic) stest() {
	fmt.Printf("stest %s\n", topic)
}

func interfaceDemo2() {
	var tester ChildTester = &Topic{"Test"}
	tester.ptest()
	tester.ctest()
	tester.stest()
}

type Tester interface {
	stest()
	ptest()
	ctest()
}

func interfaceDemo3() {
	var ctester ChildTester = &Topic{"Test"}
	var tester Tester = ctester
	tester.stest()
	tester.ptest()
	tester.ctest()
}
