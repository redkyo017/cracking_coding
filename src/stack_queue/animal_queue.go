package stack_queue

import "log"

type Animal interface {
	setOrder(order int)
	getOrder() int
}
type Dog struct {
	order int
	name  string
}

func (d *Dog) setOrder(order int) {
	d.order = order
}
func (d *Dog) getOrder() int {
	return d.order
}

type Cat struct {
	order int
	name  string
}

func (c *Cat) setOrder(order int) {
	c.order = order
}
func (c *Cat) getOrder() int {
	return c.order
}

type AnimalShelter struct {
	Cats  []*Animal
	Dogs  []*Animal
	order int
}

func (as *AnimalShelter) Enqueue(animal Animal) {
	animal.setOrder(as.order)
	switch animal.(type) {
	case *Dog:
		as.Dogs = append(as.Dogs, &animal)
	case *Cat:
		as.Cats = append(as.Cats, &animal)
	}
	as.order++
}

func (as *AnimalShelter) DequeueDog() *Animal {
	if len(as.Dogs) == 0 {
		return nil
	}
	dog := as.Dogs[0]
	as.Dogs = as.Dogs[1:]
	return dog
}
func (as *AnimalShelter) DequeueCat() *Animal {
	if len(as.Cats) == 0 {
		return nil
	}
	cat := as.Cats[0]
	as.Cats = as.Cats[1:]
	return cat
}

func (as *AnimalShelter) DequeueAny() *Animal {
	if len(as.Dogs) == 0 {
		return as.DequeueCat()
	}
	if len(as.Cats) == 0 {
		return as.DequeueDog()
	}

	cat := *as.Cats[0]
	dog := *as.Dogs[0]
	if cat.getOrder() > dog.getOrder() {
		return as.DequeueDog()
	}
	return as.DequeueCat()
}

func ImplementAnimalShelter() {
	as := &AnimalShelter{}
	as.order = 0
	dog1 := &Dog{name: "dog_1"}
	dog2 := &Dog{name: "dog_2"}
	dog3 := &Dog{name: "dog_3"}
	cat1 := &Cat{name: "cat_1"}
	cat2 := &Cat{name: "cat_2"}
	cat3 := &Cat{name: "cat_3"}
	cat4 := &Cat{name: "cat_4"}
	as.Enqueue(dog1)
	as.Enqueue(cat1)
	as.Enqueue(cat2)
	as.Enqueue(cat3)
	as.Enqueue(dog2)
	as.Enqueue(cat4)
	as.Enqueue(dog3)

	c := *as.DequeueCat()
	d := *as.DequeueDog()
	d = *as.DequeueDog()
	as.DequeueAny()
	a := *as.DequeueAny()
	log.Println("con cho", d)
	log.Println("con meo", c)
	log.Println("con co be be", a)
}
