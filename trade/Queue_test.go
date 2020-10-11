package trade

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAnalystQueue_Add(t *testing.T) {
	q:= NewQueue()
	for i:=0; i<190;i++ {
		rand.Seed(int64(i))
		e:= rand.Float64()
		q.Add(e)
	}
	fmt.Println(q.Queue)
	fmt.Println(q.GetSolving())

}
