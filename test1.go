package main

import (
	"fmt"
	"github.com/google/uuid"
	"runtime"
	"sync"
	"time"
)

var (
	NameSpaceDNS  = uuid.Must(uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceURL  = uuid.Must(uuid.Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceOID  = uuid.Must(uuid.Parse("6ba7b812-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceX500 = uuid.Must(uuid.Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8"))
	Nil           uuid.UUID // empty UUID, all zeros
)
var waitGroup sync.WaitGroup

func main() {
	var totalNum int = 2 * 1e8
	//Start
	fmt.Println("UUID test running...")

	//Running
	start := time.Now()
	singleRoutineTest(totalNum)
	//multiRoutineTest(totalNum)
	timeCost := time.Since(start)

	//Finish
	fmt.Println("version: Random UUID")
	fmt.Println("timeCost: ", timeCost)
	fmt.Println("uuid/s: ", int(float64(totalNum)/timeCost.Seconds()))
	fmt.Println("UUID test finished.")

	fmt.Println()
	//Running1
	start1 := time.Now()
	//singleRoutineTest(totalNum)
	multiRoutineTest(totalNum)
	timeCost1 := time.Since(start1)

	//Finish
	fmt.Println("version: Random UUID")
	fmt.Println("timeCost: ", timeCost1)
	fmt.Println("uuid/s: ", int(float64(totalNum)/timeCost1.Seconds()))
	fmt.Println("UUID test finished.")

}

/**
单协程uuid性能测试
*/
func singleRoutineTest(totalNum int) {
	for i := 0; i < totalNum; i++ {
		uuid.New()
	}
	fmt.Println("singleRoutineTest finish.")
}

/**
多协程uuid性能测试
*/
func multiRoutineTest(totalNum int) {
	runtime.GOMAXPROCS(8)
	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go uuidTest(totalNum / 100)
	}
	waitGroup.Wait()
	fmt.Println("multiRoutineTest finish.")
}

func uuidTest(num int) {
	for i := 0; i < num; i++ {
		uuid.New()
	}
	waitGroup.Done()
	//fmt.Println("finished")
}
