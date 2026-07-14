package puppy

import (
	"fmt"

	"github.com/zssok/dog"
)

func Bark() string {
	return "汪!!!!!!!!!"
}

func Barks() string {
	return "汪汪汪汪汪汪汪汪!!!!!!!!!"
}

func BigBark() string {

	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {

	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("From11")
}
