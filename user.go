package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Birthday() {
	*u = User{Name: u.Name, Age: u.Age + 1}
	fmt.Println(u.Name, "turns", u.Age)
}
func main() {
	u := User{Name: "Pietro", Age: 30}
	fmt.Println(u.Name, "is now", u.Age)
	u.Birthday()
	fmt.Println(u.Name, "is now", u.Age)
}
