package main

import "fmt"

type Animal interface {
	sound()
	speed() int
	isWild() bool
	iam() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

// Dog implements methods.
func (d *Dog) sound()  {
	fmt.Println("Bow Bow bow bow...... bow  bow.....!")
}

func (d *Dog) speed() int {
	return 20;
}

func (d *Dog) isWild() bool {
	return true;
}

func (d *Dog) iam() string {
	return d.name;
}

//Cat implements methods.
func (c *Cat) sound() {
	fmt.Println("Meow Meow meow meeoow.... meow..")
}

func (c *Cat) speed() int {
	return 25;
}

func (d *Cat) isWild() bool {
	return true;
}

func (d *Cat) iam() string {
	return d.name;
}

func display(animal Animal)  {
	fmt.Printf("=========== %s ============ \n", animal.iam())
	fmt.Printf("%s sounds as  ", animal.iam())
	animal.sound()
	fmt.Printf("%s runs with speed %d km/hr \n", animal.iam(), animal.speed())
	if animal.isWild() {
		fmt.Printf("%s is wild animal. \n", animal.iam())
	} else {
		fmt.Printf("%s is not wild animal. \n", animal.iam())
	}

}
func main() {
	var dog = new(Dog)
	dog.name = "Dog"

	var cat = new(Cat)
	cat.name = "Cat"

	display(dog)
	display(cat)
}