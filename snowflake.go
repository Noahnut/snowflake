package snowflake

import (
	"context"
	"sync/atomic"
	"time"
)

type Snowflake struct {
	machineID    int
	datacenterID int
	serialNum    int32
	cancel       context.CancelFunc
}

func NewSnowflakeUniqueGenerator(machineID int, datacenterID int) *Snowflake {
	s := Snowflake{
		machineID:    machineID,
		datacenterID: datacenterID,
	}

	return &s
}

func (s *Snowflake) serialNumberReseter(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	go func() {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			atomic.AddInt32(&s.serialNum, 0)
		}
	}()
}

func (s *Snowflake) StartSnowflake() {
	atomic.AddInt32(&s.serialNum, 0)
	var ctx context.Context
	ctx, s.cancel = context.WithCancel(context.Background())
	s.serialNumberReseter(ctx)
}

func (s *Snowflake) StopSnowflake() {
	s.cancel()
}

func (s *Snowflake) GetUniqueID() uint64 {
	var id uint64
	snumber := atomic.LoadInt32(&s.serialNum)
	atomic.AddInt32(&s.serialNum, 1)
	id = (id<<41 | uint64(time.Now().Unix()))
	id = (id<<5 | uint64(s.datacenterID))
	id = (id<<5 | uint64(s.machineID))
	id = (id<<12 | uint64(snumber))
	return id
}
