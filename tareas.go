package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) EliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista(){
	for _, tarea := range t.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados(){
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripcion:", tarea.descripcion)
		}
	}
}

func main() {

	t1 := &task{
		nombre:      "Completar mi curso de ruby",
		descripcion: "Completar mi curso de ruby de platzi",
		completado:  false,
	}

	t2 := &task{
		nombre:      "Completar mi curso de python",
		descripcion: "Completar mi curso de python de platzi",
		completado:  false,
	}

	t3 := &task{
		nombre:      "Completar mi curso de Nodejs",
		descripcion: "Completar mi curso de Nodejs de platzi",
		completado:  false,
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarALista(t3)

	/*for i:= 0; i < len(lista.tasks); i++ {
		fmt.Println("Index:", i, "nombre:", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks{
		fmt.Println("Index:", index, "nombre:", tarea.nombre)
	}*/

	/*for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}*/

	lista.tasks[0].marcarCompleta()
	//fmt.Println("Tareas completadas:")
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Nestor"] = lista

	t4 := &task{
		nombre:      "Completar mi curso de go",
		descripcion: "Completar mi curso de go de platzi",
		completado:  false,
	}

	t5 := &task{
		nombre:      "Completar mi curso de java",
		descripcion: "Completar mi curso de java de platzi",
		completado:  false,
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Ricardo"] = lista2

	fmt.Println("tareas de nestor")
	mapaTareas["Nestor"].imprimirLista()
	fmt.Println("tareas de Ricardo")
	mapaTareas["Ricardo"].imprimirLista()

} //termina funcion main

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizaNombre(nombre string) {
	t.nombre = nombre
}
