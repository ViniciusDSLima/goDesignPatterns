package main

import "fmt"

type stop interface {
	execute()
	setNext(stop)
}

type openFlap struct {
	next stop
}

type pressStartButton struct {
	next stop
}

type enterPassword struct {
	next stop
}

func (r *openFlap) execute() {
	fmt.Println("openFlap")
	r.next.execute()
}

func (r *openFlap) setNext(next stop) {
	r.next = next
}

func (r *pressStartButton) execute() {
	fmt.Println("pressStartButton")
	r.next.execute()
}

func (r *pressStartButton) setNext(next stop) {
	r.next = next
}

func (e *enterPassword) execute() {
	fmt.Println("enterPassword")

}

func (e *enterPassword) setNext(next stop) {
	e.next = next
}

//func main() {
//	enterPassword := &enterPassword{}
//
//	pressStartButton := &pressStartButton{}
//	pressStartButton.setNext(enterPassword)
//
//	openFlap := &openFlap{}
//	openFlap.setNext(pressStartButton)
//
//	openFlap.execute()
//}
