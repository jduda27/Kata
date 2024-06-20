package qkata1

import (
	"fmt"
	"testing"
)

func TestAddToQueue(t *testing.T) {

	fmt.Println("Test add to queue")
	queue := Queue{}
	queue.EnterQueue("000001")
	queue.EnterQueue("000002")
	queue.EnterQueue("000003")

	Show(queue)

}

func TestRemoveNextInQueue(t *testing.T) {

	fmt.Println("Test remove next in queue")
	queue := Queue{}
	queue.EnterQueue("000001")
	queue.EnterQueue("000002")
	queue.EnterQueue("000003")
	queue.EnterQueue("000004")

	Show(queue)
	queue.RemoveNextInQueue()
	Show(queue)
	queue.RemoveNextInQueue()
	Show(queue)
	queue.RemoveNextInQueue()
	Show(queue)
	queue.RemoveNextInQueue()
	Show(queue)
	queue.RemoveNextInQueue()
}

func TestExitQueue(t *testing.T) {

	var err error
	fmt.Println("Test remove from queue")
	queue := Queue{}
	queue.EnterQueue("000001")
	queue.EnterQueue("000002")
	queue.EnterQueue("000003")
	queue.EnterQueue("000004")

	Show(queue)
	err = queue.ExitQueue("000002")
	if err != nil {
		fmt.Println(err)
	}
	Show(queue)
	err = queue.ExitQueue("000001")
	if err != nil {
		fmt.Println(err)
	}
	Show(queue)
	err = queue.ExitQueue("000004")
	if err != nil {
		fmt.Println(err)
	}
	Show(queue)
	err = queue.ExitQueue("000005")
	if err != nil {
		fmt.Println(err)
	}

}
