package utils

import (
	"sync/atomic"
	"time"
)

const (
	// IDStartTimeEnvName         = "ID_START_TIME"
	defaultEpochStr        = "2020-05-26T00:00:00.000Z"
	defaultEpochNano int64 = 1590422400000000000
)

//IDInstance holds the configurations of id generating
type IDInstance struct {
	timeBits      uint
	timeMask      int64
	seqBits       uint
	seqMask       int64
	machineIDBits uint
	machineID     int64
	machineIDMask int64
	lastID        int64
}

// NewCommonIDInstance create common id-generating instance
func NewCommonIDInstance(machineID int64) *IDInstance {
	return NewIDInstance(41, 12, 10, machineID)
}

// NewIDInstance create new instance
func NewIDInstance(timeBits, seqBits, machineBits uint, machineID int64) *IDInstance {
	machineIDMask := ^(int64(-1) << machineBits) // the maximum value of machineID
	return &IDInstance{
		timeBits:      timeBits,
		seqBits:       seqBits,
		machineIDBits: machineBits,
		timeMask:      ^(int64(-1) << timeBits),
		seqMask:       ^(int64(-1) << seqBits),
		machineIDMask: machineIDMask,
		machineID:     machineID & machineIDMask,
	}
}

// NextID generates id
func (i *IDInstance) NextID() int64 {
	for {
		localLastID := atomic.LoadInt64(&i.lastID)
		seq := i.GetSeqFromID(localLastID)
		lastIDTime := i.GetTimeFromID(localLastID)
		now := i.getCurrentTimestamp()
		if now > lastIDTime {
			seq = 0
		} else if seq >= i.seqMask {
			time.Sleep(time.Duration(0xFFFFF - (time.Now().UnixNano() & 0xFFFFF)))
			continue
		} else {
			seq++
		}

		newID := now<<(i.machineIDBits+i.seqBits) + seq<<i.machineIDBits + i.machineID
		if atomic.CompareAndSwapInt64(&i.lastID, localLastID, newID) {
			return newID
		}
		time.Sleep(time.Duration(20))
	}
}

//GetSeqFromID get seq number from id
func (i *IDInstance) GetSeqFromID(id int64) int64 {
	return (id >> i.machineIDBits) & i.seqMask
}

func (i *IDInstance) getCurrentTimestamp() int64 {
	return (time.Now().UnixNano() - defaultEpochNano) >> 20 & i.timeMask
}

// GetTimeFromID get timestamp from id
func (i *IDInstance) GetTimeFromID(id int64) int64 {
	return id >> (i.machineIDBits + i.seqBits)
}
