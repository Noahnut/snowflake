package snowflake

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestOneSnowFlake(t *testing.T) {
	s := NewSnowflakeUniqueGenerator(1, 1)
	s.StartSnowflake()
	defer s.StopSnowflake()

	for i := 0; i < 70; i++ {
		ID := s.GetUniqueID()
		fmt.Println(strconv.FormatUint(ID, 2))
		time.Sleep(time.Microsecond)
	}

}

func TestManySnowFlake(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {

		go func(i int) {
			s := NewSnowflakeUniqueGenerator(i, i)
			s.StartSnowflake()
			defer s.StopSnowflake()

			for j := 0; j < 30; j++ {
				ID := s.GetUniqueID()
				fmt.Println(i, i, strconv.FormatUint(ID, 2))
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
