package main

import (
	"fmt"
	"time"
)

// Iterator を返すだけ
type Aggregator interface {
	iterator() Iterator
}

type Iterator interface {
	hasNext() bool
	next() interface{}
}

type Task struct {
	title    string
	deadline time.Time
}

type Tasks struct {
	tasks []*Task
}

func (ts *Tasks) iterator() Iterator {
	return &TasksIterator{tasks: ts, index: 0}
}

func (ts *Tasks) length() int {
	return len(ts.tasks)
}

func (ts *Tasks) appendTask(t *Task) {
	ts.tasks = append(ts.tasks, t)
}

type TasksIterator struct {
	tasks *Tasks
	index int
}

func (ti *TasksIterator) hasNext() bool {
	return ti.index < ti.tasks.length()
}

func (ti *TasksIterator) next() interface{} {
	t := ti.tasks.tasks[ti.index]
	ti.index++
	return t
}

func main() {
	ts := Tasks{}
	ts.appendTask(&Task{title: "design"})
	ts.appendTask(&Task{title: "write code"})
	ts.appendTask(&Task{title: "test"})
	it := ts.iterator()
	for it.hasNext() {
		// Iterator インターフェイスを汎用的にするために、next の返り値を interface{} しているため
		// type assertion する必要がある
		t := it.next().(Task)
		fmt.Println(t.title)
	}
}
