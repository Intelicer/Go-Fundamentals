package structs

import (
	// "fmt"
	"sort"
	"strconv"
)

type Person struct {
	Name        string
	Age         int
	Information map[string]string
}

func NewPerson(name string, age int, information map[string]string) Person {
	person := Person{
		Name:        name,
		Age:         age,
		Information: information,
	}
	person.sortInformation()
	return person
}

// format person information
func (person *Person) PersonFormattedInformation() string {
	personFormatedString := "Person name: " + person.Name + 
												", Age: " + strconv.Itoa(person.Age) + "\nDetailed Info:"
	for key, value := range person.Information {
		personFormatedString += "\n" + key + ": " + value
	}
	return personFormatedString
}


//to change value of age make sure that person is pointer so you can change by referencing to object
func (person *Person) UpdateAge(newAge int){
	person.Age = newAge
}


func (person *Person) AddExtraInformation(infoType string, details string){
	person.Information[infoType] = details
	person.sortInformation()
}

//if the method starts with lower case letter its private
func (person *Person) sortInformation(){

	keys := make([]string, 0, len(person.Information)) //create a slice based on the length of information map.

	for key := range person.Information{
		keys = append(keys, key)
	}

	sort.Strings(keys)

	newInfo := make(map[string]string) 
	for _, key := range keys{
		newInfo[key] = person.Information[key]
	}

	person.Information = newInfo
}