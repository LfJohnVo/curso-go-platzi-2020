package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct {}

type pez struct {}

type pajaro struct {}

func (perro) mover() string{
	return "Soy un perro y camino"
}

func (pez) mover() string {
	return "soy un pez y nado"
}

func (pajaro) mover() string{
	return "Soy un pajaro y vuelo"
}

func moverAnimal(a animal){
	fmt.Println(a.mover())
}

func main(){
	p := perro{}
	moverAnimal(p)

	pez := pez{}
	moverAnimal(pez)

	pa := pajaro{}
	moverAnimal(pa)

}//termina main


