// This package provides unique id in distribute system
// the algorithm is inspired by Twitter's famous snowflake
// its link is: https://github.com/twitter/snowflake/releases/tag/snowflake-2010
//

//  0                   42	            52             64
// +-------------------+---------------+--------------+
// | timestamp(ms)(42) | worker id(10) | sequence(12) |
// +-------------------+---------------+--------------+

// Copyright (C) 2016 by zheng-ji.info

package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	CEpoch_V0      = 1474543688381
	Version        = 0  // Protocol version, for future update
	CWorkerIdBits  = 10 // Num of WorkerId Bits
	CSenquenceBits = 12 // Num of Sequence Bits
	CTimeStampBits = 40 // Num of Sequence Bits
	CSequenceMask  = 0xfff
	CMaxWorker     = 0x3ff

	CWorkerIdShift  = 12
	CTimeStampShift = 22
	CVersionShift   = 62
)

func ParseId(id int64) (t time.Time, ts int64, workerId int64, seq int64) {
	seq = id & CSequenceMask
	workerId = (id >> CWorkerIdShift) & CMaxWorker
	ts = id>>CTimeStampShift + CEpoch_V0
	t = time.Unix(ts/1000, (ts%1000)*1000000)
	return
}

// IdWorker Struct
type IdWorker struct {
	workerId      int64
	lastTimeStamp int64
	sequence      int64
	lock          *sync.Mutex
}

// NewIdWorker Func: Generate NewIdWorker with Given workerid
func NewIdWorker(workerid uint64) (iw *IdWorker, err error) {
	iw = new(IdWorker)

	if workerid > CMaxWorker || workerid < 0 {
		return nil, errors.New("worker not fit")
	}
	iw.workerId = int64(workerid)
	iw.lastTimeStamp = -1
	iw.sequence = 0
	iw.lock = new(sync.Mutex)
	return iw, nil
}

// return in ms
func (iw *IdWorker) timeGen() int64 {
	return time.Now().UnixNano() / 1000000
}

func (iw *IdWorker) timeReGen(last int64) int64 {
	ts := iw.timeGen()
	for {
		if ts <= last {
			ts = iw.timeGen()
		} else {
			break
		}
	}
	return ts
}

// NewId Func: Generate next id
func (iw *IdWorker) NextId() (int64, error) {
	iw.lock.Lock()
	defer iw.lock.Unlock()
	ts := iw.timeGen()
	if ts == iw.lastTimeStamp {
		iw.sequence = (iw.sequence + 1) & CSequenceMask
		if iw.sequence == 0 {
			ts = iw.timeReGen(ts)
		}
	} else {
		iw.sequence = 0
	}

	if ts < iw.lastTimeStamp {
		return 0, errors.New("Clock moved backwards, Refuse gen id")
	}
	iw.lastTimeStamp = ts
	ts -= CEpoch_V0
	id := ts<<CTimeStampShift | iw.workerId<<CWorkerIdShift | iw.sequence
	return id, nil
}
