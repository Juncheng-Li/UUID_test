package main

import (
	"github.com/google/uuid"
)

const (
	Person = uuid.Domain(0)
	Group  = uuid.Domain(1)
	Org    = uuid.Domain(2)
)

func main() {

	//fmt.Println(os.Getuid())
	//fmt.Println(Group)
	//fmt.Println(Group)
	//fmt.Println(uuid.NewDCESecurity(Group, uint32(os.Getuid())))
	//fmt.Println(uuid.NewDCESecurity(Group, uint32(os.Getuid())))
	//fmt.Println(uuid.NewDCESecurity(Group, uint32(os.Getuid())))
	//
	//fmt.Println()
	//fmt.Println([]byte("random"))
	//fmt.Println("uuid.NameSpace:", uuid.NameSpaceDNS)
	//fmt.Println("uuid.NameSpace:", uuid.NameSpaceURL)
	//fmt.Println("uuid.NameSpace:", uuid.NameSpaceDNS)
	//
	//fmt.Println(uuid.NewMD5(uuid.NameSpaceDNS, []byte("random2")))
	//fmt.Println(uuid.NewMD5(uuid.NameSpaceDNS, []byte("random2")))
	//fmt.Println(uuid.NewMD5(uuid.NameSpaceDNS, []byte("random3")))
	//fmt.Println(uuid.NewMD5(uuid.NameSpaceDNS, []byte("random3")))
	//fmt.Println(uuid.NewSHA1(uuid.NameSpaceDNS, []byte("random3")))

	//fmt.Println(NameSpaceDNS)
	//fmt.Println(NameSpaceDNS.NodeID())
	//fmt.Println("ClockSequence: ", NameSpaceDNS.ClockSequence())
	//fmt.Println(NameSpaceDNS.Domain())
	//fmt.Println(NameSpaceDNS.ID())
	//fmt.Println(NameSpaceDNS.MarshalBinary())
	//fmt.Println(NameSpaceDNS.MarshalText())
	//fmt.Println(NameSpaceURL.MarshalText())
	//fmt.Println(len(NameSpaceDNS.String()))
	//fmt.Println(NameSpaceDNS.Time())
	//
	//fmt.Println()
	//fmt.Println(uuid.UUID{})
	//fmt.Println("nil: ", Nil)
	//fmt.Println(uuid.NewUUID())
	//fmt.Println(uuid.NewUUID())
	//fmt.Println(uuid.NewUUID())
	//fmt.Println(uuid.NewUUID())
	//
	//fmt.Println()
	//
	//var tmp, _ = uuid.NewUUID()
	//
	//start := time.Now()
	//for i := 0; i < 1e8; i++ {
	//	fmt.Println(tmp)
	//	fmt.Println(i)
	//	previousTmp := tmp
	//	tmp, _ = uuid.NewUUID()
	//
	//	if previousTmp.String() == tmp.String() {
	//		break
	//	}
	//}
	//timeCost := time.Since(start)
	//fmt.Println(timeCost)
	//
	//fmt.Println(uuid.New())
	//fmt.Println(uuid.New())
	//fmt.Println(uuid.New())
	//fmt.Println(uuid.NewUUID())
	//fmt.Println(uuid.NewUUID())
	//fmt.Println(uuid.NewUUID())
}
