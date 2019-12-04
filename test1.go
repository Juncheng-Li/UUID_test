package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
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

	//Testing
	singleRoutineTest(totalNum)
	multiRoutineTest(totalNum, 1)
	multiRoutineTest(totalNum, 21)
	multiRoutineTest(totalNum, 22)
	multiRoutineTest(totalNum, 23)
	multiRoutineTest(totalNum, 3)
	multiRoutineTest(totalNum, 4)
	multiRoutineTest(totalNum, 5)

}

/**
单协程uuid性能测试
*/
func singleRoutineTest(totalNum int) {
	start := time.Now()

	for i := 0; i < totalNum; i++ {
		_, _ = uuid.NewUUID()
	}

	//Finish
	fmt.Println("singleRoutineTest finish.")
	timeCost := time.Since(start)
	fmt.Println("timeCost: ", timeCost)
	fmt.Println("uuid/s: ", int(float64(totalNum)/timeCost.Seconds()))
	fmt.Println()
}

/**
多协程uuid性能测试
*/
func multiRoutineTest(totalNum int, ver int) {
	runtime.GOMAXPROCS(8)
	start := time.Now()

	//Start running
	switch ver {
	case 1:
		//版本1
		fmt.Println("Version 1 start...")
		waitGroup.Add(100)
		for i := 0; i < 100; i++ {
			go uuid1Test(totalNum / 100)
		}
		waitGroup.Wait()
		fmt.Println("version: 1")
	case 21:
		//版本2 - Group
		fmt.Println("Version 2 - Group start...")
		waitGroup.Add(100)
		for i := 0; i < 100; i++ {
			go uuid21Test(totalNum / 100)
		}
		waitGroup.Wait()
		fmt.Println("version: 3")
	case 22:
		//版本2 - Person
		fmt.Println("Version 2 - Person start...")
		waitGroup.Add(100)
		for i := 0; i < 100; i++ {
			go uuid22Test(totalNum / 100)
		}
		waitGroup.Wait()
		fmt.Println("version: 3")
	case 23:
		//版本2 - DCE security
		fmt.Println("Version 2 - DCE security start...")
		waitGroup.Add(100)
		for i := 0; i < 100; i++ {
			go uuid23Test(totalNum / 100)
		}
		waitGroup.Wait()
		fmt.Println("version: 3")
	case 3:
		//版本3
		fmt.Println("Version 3 start...")
		waitGroup.Add(100)
		for i := 0; i < 100; i++ {
			go uuid3Test(totalNum / 100)
		}
		waitGroup.Wait()
		fmt.Println("version: 3")
	case 4:
		//版本4
		fmt.Println("Version 4 start...")
		waitGroup.Add(100)
		for i := 0; i < 100; i++ {
			go uuid4Test(totalNum / 100)
		}
		waitGroup.Wait()
		fmt.Println("version: 4")
	case 5:
		//版本5
		fmt.Println("Version 5 start...")
		waitGroup.Add(100)
		for i := 0; i < 100; i++ {
			go uuid5Test(totalNum / 100)
		}
		waitGroup.Wait()
		fmt.Println("version: 5")
	}

	//Finish
	timeCost := time.Since(start)
	fmt.Println("timeCost: ", timeCost)
	fmt.Println("uuid/s: ", int(float64(totalNum)/timeCost.Seconds()))
	fmt.Println()
}

/**
不同版本的uuid
*/
//版本1 - NodeID（本机MAC地址，clock sequence，当前时间）
func uuid1Test(num int) {
	for i := 0; i < num; i++ {
		_, _ = uuid.NewUUID()
	}
	waitGroup.Done()
	//fmt.Println("finished")
}

//版本2.1
func uuid21Test(num int) {
	for i := 0; i < num; i++ {
		_, _ = uuid.NewDCEGroup()
	}
	waitGroup.Done()
	//fmt.Println("finished")
}

//版本2.2
func uuid22Test(num int) {
	for i := 0; i < num; i++ {
		_, _ = uuid.NewDCEPerson()
	}
	waitGroup.Done()
	//fmt.Println("finished")
}

//版本2.3
func uuid23Test(num int) {
	for i := 0; i < num; i++ {
		//Domain可以是 Person，Group 或者 Organization
		_, _ = uuid.NewDCESecurity(uuid.Domain(0), uint32(os.Getuid()))
	}
	waitGroup.Done()
	//fmt.Println("finished")
}

//版本3
func uuid3Test(num int) {
	for i := 0; i < num; i++ {
		uuid.NewMD5(uuid.NameSpaceDNS, []byte("random string"))
	}
	waitGroup.Done()
	//fmt.Println("finished")
}

//版本4
func uuid4Test(num int) {
	for i := 0; i < num; i++ {
		_, _ = uuid.NewRandom()
	}
	waitGroup.Done()
	//fmt.Println("finished")
}

//版本5
func uuid5Test(num int) {
	for i := 0; i < num; i++ {
		uuid.NewSHA1(uuid.NameSpaceDNS, []byte("random string"))
	}
	waitGroup.Done()
	//fmt.Println("finished")
}
