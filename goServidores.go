package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	//timer
	inicio := time.Now()
	canal := make(chan string)

	//lista de servidores que la aplicación esta revisando
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for  {
		if i >2{
			break
		}
		//recorrer slice
		for _, servidor := range servidores{
			go revisarServidor(servidor, canal)
			//leer un canal
			//fmt.Println(<-canal)
		}
		time.Sleep(1*time.Second)
		fmt.Println(<-canal)
		i++
	}

	//calcula el tiempo de ejecución
	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución %s\n", tiempoPaso)

}//termina main

func revisarServidor(servidor string, canal chan string)  {
	_, err := http.Get(servidor)
	if err != nil{
		canal <- servidor + " No se encuentra disponible"
	}else {
		canal <- servidor + " Esta funcionando"
	}
}
