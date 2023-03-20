package _0230320

import (
	"fmt"
	"testing"
)

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}
		this.in2out()
	}

	first := this.outStack[0]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return first
}
func (this *CQueue) in2out() {
	inLength := len(this.inStack)
	if inLength > 0 {
		this.outStack = append(this.outStack, this.inStack[inLength-1])
		this.inStack = this.inStack[:inLength-1]
	}

}
func TestSlice(t *testing.T) {
	q:=Constructor()
	fmt.Println("1-->",q.DeleteHead())

	q.AppendTail(5)
	q.AppendTail(2)


	fmt.Println("2-->",q.DeleteHead())
	fmt.Println("3-->",q.DeleteHead())



}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
