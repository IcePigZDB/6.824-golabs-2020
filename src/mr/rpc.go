package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

// //
// // example to show how to declare the arguments
// // and reply for an RPC.
// //
// type ExampleArgs struct {
// 	X int
// }

// type ExampleReply struct {
// 	Y int
// }

// Add your RPC definitions here.

// TaskArgs rpc args for worker to request task
type TaskArgs struct {
	WorkerID int // worker's id
}

// TaskReply rpc args for master to return task
type TaskReply struct {
	Task *Task
}

// ReportTaskArgs rpc args for worker to report
type ReportTaskArgs struct {
	Done     bool
	Seq      int       // seq of task
	Phase    TaskPhase // task of phase
	WorkerID int       // worker's id
}

// ReportTaskReply need not return
type ReportTaskReply struct {
}

//RegisterArgs  need not return
type RegisterArgs struct {
}

// RegisterReply rpc args for worker to register worker
type RegisterReply struct {
	WorkerID int
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the master.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func masterSock() string {
	s := "/var/tmp/824-mr-"
	// Getuid return current process id.
	s += strconv.Itoa(os.Getuid())
	return s
}
