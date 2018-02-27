package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Talker interface {
	CanTalk() bool
}

type Animal struct {
	NoLegS int    `json:"-"`
	NoLegs int    `json:"noLegs"`
	Name   string `json:"name"`
}

func (a Animal) CanTalk() bool {
	return false
}

func (a Animal) String() string {
	return fmt.Sprintf("Animal{NoLegs=%d, Name=%s, NoLegS=%d}", a.NoLegs, a.Name, a.NoLegS)
}

func main() {
	var creature Talker
	creature = Animal{
		NoLegs: 4,
		Name:   "Pig",
	}
	fmt.Println(creature)

	//reading the file
	fileContent, err := ioutil.ReadFile("main/animals.json")
	if err != nil {
		fmt.Println("Unable to open the file")
		panic(err)
	}

	fmt.Println("The content of animals.json is:")
	fmt.Println(string(fileContent))

	//de-serialization
	var animals []Animal
	err = json.Unmarshal(fileContent, &animals)
	if err != nil {
		fmt.Println("Unable to de-serialize the animals")
		panic(err)
	}

	//check the values de-serialized
	fmt.Println("The animals are:")
	fmt.Println(animals)

	if serializedAnimals, err := json.Marshal(animals); err != nil {
		fmt.Println("Unable to serialize the animals")
		panic(err)
	} else {
		fmt.Println(string(serializedAnimals))
	}
}
