package main 

import "fmt"

type animal interface {
   typeOfDog() 	string
   breathe() 	string
   walk()  	string
   name()   	string
   born() 	int
}

type dog struct {
     age int
}


func (boris dog) breathe() string {
    return "can breathes"
}

func (boris dog) walk() string {
	return "can walk"
}

func (boris dog) name()  string {
	return "the name is Boris"
}
func (boris dog) typeOfDog() string {
	return "This is one dog"
}
func (boris dog) born() int {
	return 2018
}

func main(){
	var animal_interface animal
	animal_interface = dog{age: 5}
	animal_interface.breathe()
	animal_interface.walk()
	animal_interface.name()
	animal_interface.typeOfDog()
	animal_interface.born()
	fmt.Println("interfaces")
	fmt.Println(animal_interface.typeOfDog() , animal_interface.walk() , animal_interface.breathe() , animal_interface.name(), "Was born:" ,animal_interface.born())
}