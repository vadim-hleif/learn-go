package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (arr byAge) Len() int {
	return len(arr)
}

func (arr byAge) Less(i, j int) bool {
	return arr[i].Age < arr[j].Age
}

func (arr byAge) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type byLast []user

func (arr byLast) Len() int {
	return len(arr)
}

func (arr byLast) Less(i, j int) bool {
	return arr[i].Last < arr[j].Last
}

func (arr byLast) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	//Starting with this code, sort the []user by
	//age
	//last

	//Also sort each []string “Sayings” for each user
	//print everything in a way that is pleasant
	for _, user := range users {
		sort.Strings(user.Sayings)
	}

	sort.Sort(byAge(users))
	fmt.Println(users)

	sort.Sort(byLast(users))
	fmt.Println(users)
}
