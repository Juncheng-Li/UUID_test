package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	fmt.Println(NameSpaceDNS)
	fmt.Println(NameSpaceDNS.NodeID())
	fmt.Println("ClockSequence: ", NameSpaceDNS.ClockSequence())
	fmt.Println(NameSpaceDNS.Domain())
	fmt.Println(NameSpaceDNS.ID())
	fmt.Println(NameSpaceDNS.MarshalBinary())
	fmt.Println(NameSpaceDNS.MarshalText())
	fmt.Println(NameSpaceURL.MarshalText())
	fmt.Println(len(NameSpaceDNS.String()))
	fmt.Println(NameSpaceDNS.Time())

	fmt.Println()
	fmt.Println(uuid.UUID{})
	fmt.Println("nil: ", Nil)
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())

	fmt.Println()

	var tmp, _ = uuid.NewUUID()

	start := time.Now()
	for i := 0; i < 1e8; i++ {
		fmt.Println(tmp)
		fmt.Println(i)
		previousTmp := tmp
		tmp, _ = uuid.NewUUID()

		if previousTmp.String() == tmp.String() {
			break
		}
	}
	timeCost := time.Since(start)
	fmt.Println(timeCost)

	fmt.Println(uuid.New())
	fmt.Println(uuid.New())
	fmt.Println(uuid.New())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
}
