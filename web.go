package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct {}

func (escritorWeb) Write(p []byte) (int, error){
	fmt.Println(string(p))
	return len(p), nil
}

func main()  {
	respuesta, err :=http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}

	e := escritorWeb{}
	//copy toma el flujo que viene del body y lo pasa al editor para sacarlo lo mejor posible
	io.Copy(e, respuesta.Body)

}

func Var_dump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("Var_dump: %#v", expression))
}