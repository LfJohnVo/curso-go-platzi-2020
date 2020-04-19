package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	//timer
	inicio := time.Now()

	//lista de servidores que la aplicación esta revisando
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	//recorrer slice
	for _, servidor := range servidores{
		revisarServidor(servidor)
	}

	//calcula el tiempo de ejecución
	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución %s\n", tiempoPaso)

}//termina main

func revisarServidor(servidor string)  {
	_, err := http.Get(servidor)
	if err != nil{
		fmt.Println(servidor, "no esta disponible :(")
	}else {
		fmt.Println(servidor, "Operando normalmente :)")
	}
}
