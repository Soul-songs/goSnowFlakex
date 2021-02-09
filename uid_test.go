package snowflake

import (
	"fmt"
	"testing"
)

func TestTimeGen(t *testing.T) {
	iw, _ := NewIdWorker(2)
	ts := iw.timeGen()
	fmt.Println(ts - CEpoch_V0)
}

func TestSnowFlake(t *testing.T) {
	fmt.Println("start generate")
	iw, _ := NewIdWorker(2)
	for i := 0; i < 4095; i++ {
		id, err := iw.NextId()
		if i > 4090 {

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(id)
				fmt.Println(ParseId(id))
			}
		}
	}
	for i := 0; i < 4095; i++ {
		id, err := iw.NextId()
		if i < 10 {

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(id)
				fmt.Println(ParseId(id))
			}
		}
	}
	fmt.Println("end generate")
}
