package mr

import (
	"fmt"
	"log"
)

// TaskPhase is the enumeration of task phase.
type TaskPhase int

// Maphase
const (
	MapPhase    TaskPhase = 0
	ReducePhase TaskPhase = 1
)

// Debug flag
const Debug = false

// DPrintf is the debug output
func DPrintf(format string, v ...interface{}) {
	if Debug {
		log.Printf(format+"\n", v...)
	}
}

// Task struct
type Task struct {
	FileName string    // task's file
	NReduce  int       // number of reduce tasks
	NMaps    int       // number of map tasks
	Seq      int       // task seq
	Phase    TaskPhase // current task's phase
	Alive    bool      // worker should exit when alive is false
}

func reduceName(mapIdx, reduceIdx int) string {
	return fmt.Sprintf("mr-%d-%d", mapIdx, reduceIdx)
}

func mergeName(reduceIdx int) string {
	return fmt.Sprintf("mr-out-%d", reduceIdx)
}
