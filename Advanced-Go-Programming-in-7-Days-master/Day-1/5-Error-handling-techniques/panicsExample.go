package main

import (
	"reflect"
	"sync"
	"log"
	"fmt"
)

func main() {

	//not all defers will work
	//will not compile
	//defer append([]string{"1", "2", "3"}, "4")

	//be careful when assigning and comparing interfaces
	var v interface{} = nil
	var arr []int = nil
	v = arr

	//will panic as []int is an uncomparable type
	fmt.Println(reflect.TypeOf(v))
	//fmt.Println(v == v)

	defer println("Will print string")
	//panic("Panic!")
	defer println("Will not print string")

	var l = &sync.Mutex{}

	defer func() {
		if v := recover(); v != nil {
			log.Println("Recovered from: ", v)
			log.Printf("Lock is: %v\n", l) // {1, 0} indicates that its locked
		}
	}()

	callLockUnsafe(l)
	callLockSafe(l)


	// Panic suppression. FIFO applies here
	defer panic(1)
	defer panic(2)
	panic(3)
}

func callLockSafe(l *sync.Mutex) {
	l.Lock()
	defer l.Unlock()
	panic("Panic1!!!")
}

func callLockUnsafe(l *sync.Mutex) {
	l.Lock()
	panic("Panic2!!!")
	// will not unlock
	l.Unlock()
}
