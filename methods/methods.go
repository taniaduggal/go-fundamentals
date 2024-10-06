package main

// alternative of classes
import "fmt"

//there is no iheritance and super parents like things here

func main() {
	fmt.Println("structs in golang")
	tania := User{"Tania", "tsnisduggsl@gamil.com", true, 19}
	fmt.Println(tania)
	fmt.Printf("tania details are: %+v\n", tania) //+v provides more details
	fmt.Printf("Name is %v and email is %v\n", tania.name, tania.Email)
	tania.GetStatus()
	tania.NewMail()
}

//it should be in uppercase that we can import it easily
type User struct {
	name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user activa: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of user is: ", u.Email)
}

//whenever you passes these structucts, it actualyy passes on a copy of object or of  that func, thus pointer came. if you really want to pass the real object, should you pass the reference of that object orpointer to that
