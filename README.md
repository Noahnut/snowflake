# Distributed global unique ID generator (Snowflake)

Use the golang to implement Snowflake global unique ID generator from twitter. The unit64 type number
split this 64 bits number to different part.
1 bit is keeping for future use
41 bits for the timestamp, use the unix time in this project
5 bits for the data center ID 
5 bits for the machine ID 
12 bits for the machine serial Number which will reset in 1 millisecond.
Use the above feature to make sure each ID generater from the different machine could be unique.

## Install 
```shell
go get github.com/Noahnut/snowflake
```

## Example Usage
```golang
package main

import (
	"fmt"

	"github.com/Noahnut/snowflake"
)

func main() {
    // Create the new snow flake
	s := NewSnowflakeUniqueGenerator(1, 1)

    //Start the Snowflake
	s.StartSnowflake()
	defer s.StopSnowflake()

	for i := 0; i < 70; i++ {

        //get the ID from the snowflake
		ID := s.GetUniqueID()
		fmt.Println(strconv.FormatUint(ID, 2))
		time.Sleep(time.Microsecond)
	}
}
```