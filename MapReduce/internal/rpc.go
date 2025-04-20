package mr

import (
	"os"
	"strconv"
)

type GetTaskArgs struct {
}

type GetTaskReply struct {
	Name    string   // file name from the data is read
	Number  int      // task number for each phase which we uses for referring to intermediate  and output file.
	NReduce int      // partitation number
	Type    TaskType // map or reduce
}

type TaskType string

type MarkTaskAsDoneArgs struct {
	Name string
	Type TaskType
}

type MarkTaskAsDoneReply struct {
}

var (
	mType TaskType = "map"
	rType TaskType = "reduce"
)

func coordinatorSock() string {
	s := "/var/temp/mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
