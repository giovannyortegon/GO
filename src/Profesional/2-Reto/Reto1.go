package main

import "fmt"

// lista de tarea
type TaskList struct {
	tasks[] * Task
}

func (tl * TaskList) appendTask (t * Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl * TaskList) removeTask (index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index + 1:]...)
}

type Task struct {
	name string
	desc string
	completed bool
}

func (t * Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescription: %s\nConectado: %t\n",
			   t.name, t.desc, t.completed)
}

func main() {
	t1 := Task {"Curso Go","APrendizaje Go desde Principante a Master", false}
	t1.toPrint()

	t2 := Task {
		name: "Curso Python",
		desc:	"Desarrollo de Scripts for Cybersecurity",
		completed: false,
	}
	t2.toPrint()

	lista := TaskList{}
	lista.appendTask(&t1)
	lista.appendTask(&t2)

	lista.removeTask(1)

	for idx, task := range lista.tasks {
		fmt.Println(idx, task.name)
	}

}
