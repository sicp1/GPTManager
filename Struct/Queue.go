package Struct

type Queue struct {
	Data []interface{}
}
type QueueInterface interface {
	Push(data interface{})
	Pop() interface{}
	Show()
}

func (queue *Queue) Push(data interface{}) {
	queue.Data = append(queue.Data, data)
}
func (queue *Queue) Pop() (interface{}, int) {
	if len(queue.Data) > 1 {
		pop_data := queue.Data[0]
		queue.Data = queue.Data[1:]
		return pop_data, 1
	} else if len(queue.Data) == 1 {
		last_data := queue.Data[0]
		queue.Data = []interface{}{}
		return last_data, 0
	} else {
		return nil, -1
	}
}
func (queue *Queue) Show() {
	for _, item := range queue.Data {
		println(item)
	}
}
