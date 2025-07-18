package main

import (
	"fmt"
	"strings"
)

type Task interface{
	getName() string
	getTimeEstimate() int
	printTask(indent int)
}

type SimpleTask struct{
	name string
	time int
}

func NewSimpleTask(name string, time int) *SimpleTask{
	return &SimpleTask{name: name,time: time}
}

func(t*SimpleTask) getName() string {return t.name}
func(t*SimpleTask) getTimeEstimate() int {return t.time}
func(t*SimpleTask) printTask(indent int) {
	fmt.Printf("%s- task %q (%d h)\n",strings.Repeat(" ", indent), t.name, t.time)
}

type TaskGroup struct {
	name string
	tasks []Task
}

func NewTaskGroup(name string) *TaskGroup{
	return &TaskGroup{name: name, tasks: make([]Task,0)}
}

func (g*TaskGroup)getTimeEstimate() int {
	sum := 0
	for _,t:=range g.tasks{
		sum += t.getTimeEstimate()
	}
	return sum
}

func (g*TaskGroup)addTask(t Task) {
	g.tasks = append(g.tasks, t)
}

func (g*TaskGroup)getName() string {return g.name}

func (g*TaskGroup)printTask(indent int) {
	fmt.Printf("%s* %s: total time %dh\n", strings.Repeat(" ", indent), g.getName(), g.getTimeEstimate())
	for _,t:=range g.tasks{
		t.printTask(indent + 1)
	}
}

func main() {
	fmt.Println("=== COMPOSITE")
	uiDesign := NewSimpleTask("Design UI", 5)
	backendDesign := NewSimpleTask("Backend Design", 7)
	featureA := NewSimpleTask("feature A", 10)
	featureB := NewSimpleTask("feature B", 8)

	designPhase := NewTaskGroup("Design Phase")

	designPhase.addTask(uiDesign)
	designPhase.addTask(backendDesign)

	devPhase := NewTaskGroup("Development Phase")
	devPhase.addTask(featureA)
	devPhase.addTask(featureB)

	project := NewTaskGroup("Project Alpha")
	project.addTask(designPhase)
	project.addTask(devPhase)

	fmt.Printf("Total project time: %d hours\n", project.getTimeEstimate())
	project.printTask(0)
}
