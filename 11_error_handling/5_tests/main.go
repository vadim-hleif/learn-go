package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	//We are going to learn about testing next.
	//For this exercise, take a moment and see how much you can
	//figure out about testing by reading the testing documentation &
	//also Caleb Doxsey’s article on testing. See if you can get a basic example of testing working.
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	marshal, err := json.Marshal(a)

	if err != nil {
		err = fmt.Errorf("custom error: %v", err.Error())
	}

	return marshal, err
}
