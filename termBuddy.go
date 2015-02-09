package main 

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)


type Buddy struct {
	Name string
	Level, Xp, Age, Happiness int
}

type Dog struct {
	Buddy
	Breed string
	Type string
}

type Cat struct {
	Buddy
	Type string
}

func (self Dog) setDefaults(name, breed string) Dog{
	self.Name = name
	self.Breed = breed
	self.Level = 1
	self.Xp = 0
	self.Age = 0
	self.Happiness = 10
	self.Type = "Dog"
	return self
}

func (self Cat) setDefaults(name string) Cat{
	self.Name = name
	self.Level = 1
	self.Xp = 0
	self.Age = 0
	self.Happiness = 10
	self.Type = "Cat"
	return self
}

func startGame () {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to TermBuddy,\nThe game where you can raise your own buddy right in the terminal!")
	fmt.Println("First off you need to make a buddy!")
	fmt.Print("Would you like a dog or cat?  >>")
	buddyChoice, err := reader.ReadString('\n')
	if err!=nil {
		fmt.Println(err)
	}
	buddyChoice = strings.TrimSpace(buddyChoice)
	switch buddyChoice {
		case "dog": createDog()
		case "cat": createCat()
		default: fmt.Println("Please choose one of the options above")
	}
}

func createDog() {
	fmt.Println("\n\nYou chose a Dog! Now its time to choose a name and breed!")
	fmt.Print("\n\nWhat would you like to name your dog? >>")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	name = strings.TrimSpace(name)
	fmt.Printf("\n\nWhat breed would you like %s to be? >>", name)
	breed, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	dog := Dog{}
	dog = dog.setDefaults(name, breed)
	fmt.Printf("\n\nNew Dog\nName: %s\nBreed: %sAge: %d\nLevel: %d\nHappiness: %d", dog.Name, dog.Breed, dog.Age, dog.Level, dog.Happiness)
}

func createCat() {
		fmt.Println("\n\nYou chose a Cat! Now its time to choose a name!")
	fmt.Print("\n\nWhat would you like to name your cat? >>")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	name = strings.TrimSpace(name)
	cat := Cat{}
	cat = cat.setDefaults(name)
	fmt.Printf("\n\nNew Cat\nName: %s\nAge: %d\nLevel: %d\nHappiness: %d\n", cat.Name, cat.Age, cat.Level, cat.Happiness)
}

func usersTurn () {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your turn! Choose on of the following. (levelup, play, wait)")
	fmt.Println("Enter text: ")
	text, err := reader.ReadString('\n')
	if err!=nil {
		fmt.Println(err)
	}
	switch strings.TrimSpace(text) {
		case "levelup": fmt.Println("Your buddy went up a level!")
		case "play": fmt.Println("Your buddy is now happier because you played with him!")
		case "wait": fmt.Println("Waiting one turn")
		default: fmt.Println(text + " is not a proper response")
	}
}

func main() {
	startGame()
	usersTurn()
}