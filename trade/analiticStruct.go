package trade

type analystQueue struct {
	Count int
	Queue [120]float64
}

// Добавить элемент в очередь
func (q *analystQueue) Add(element float64) {
	if q.Count <= 119 {
		q.Queue[q.Count] = element
		q.Count++
		return
	}
	sl := q.Queue[1:]
	i := 0
	for _, v := range sl {
		q.Queue[i] = v
		i++
	}
	q.Queue[q.Count-1] = element

}
func NewQueue() *analystQueue {
	return &analystQueue{}
}

//true == rise
func (q *analystQueue) GetSolving() bool {
	max, min := q.Queue[0], q.Queue[0]
	lastElem := q.Queue[119]
	for _, v := range q.Queue {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	difrentMax := max - lastElem
	difrentMin := lastElem - min
	if difrentMax > difrentMin {
		return true
	}
	return false

}
