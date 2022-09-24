package functional

import (
	"fmt"
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	start := time.Now()

	p1 := ProcessAsync(heavyTask1, "first-task")
	p2 := ProcessAsync(heavyTask2, 10)
	p3 := ProcessAsync(heavyTask3, "this-gonna-return-error")

	fmt.Println(time.Since(start))
	if time.Since(start) > time.Second {
		t.Errorf("the processes must be started asyncronously")
	}

	res1 := p1.WaitForResult().GetResult()
	res2 := p2.WaitForResult().GetResult()
	res3 := p3.WaitForResult().GetError()

	fmt.Printf("%v, %T \n", *res1, res1)

	fmt.Printf("%v, %T  \n", *res2, res2)

	if *res1 != "output-of-first-task" {
		t.Errorf("wrong return for task 1")
	}

	if *res2 != 15 {
		t.Errorf("wrong return for task 2")
	}

	if res3.Error() != "error!" {
		t.Errorf("wrong return for task 3")
	}

	if time.Since(start) < time.Duration(time.Duration.Seconds(2)) {
		t.Errorf("The processes must be performed in 5 seconds (max duration of the three tasks)")
	}
}

func TestFutureIdempotence(t *testing.T) {

	f1 := NewFuture(heavyTask1, "first-task")
	f1.Process().Process().Process()
	f1.Process()

	f1.WaitForResult()
	f1.WaitForResult()

	res1 := f1.WaitForResult().GetResult()

	if *res1 != "output-of-first-task" {
		t.Errorf("wrong return for task 1")
	}

}

func heavyTask1(input string) (*string, error) {
	time.Sleep(5 * time.Second)
	str := "output-of-" + input
	return &str, nil
}

func heavyTask2(input int) (*int, error) {
	time.Sleep(2 * time.Second)
	out := 5 + input
	return &out, nil
}

func heavyTask3(input string) (*int, error) {
	time.Sleep(1 * time.Second)
	return nil, fmt.Errorf("error!")
}
