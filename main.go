package main

import "fmt"

type User struct {
	name string
	age int32
	email string
	password int
	score []int
}

func (u User) getHighScore() int {
	if len(u.score) == 0 {
		return 0 // return 0, if slice is empty
	}

	q := u.score[0] // assume the first element is the smallest

	for _, sc := range u.score {
		if sc < q { // you can write > and it will give the largest number
			q = sc // update the value of the smallest number
		}
	}

	return q
}

func (u *User) getName(name1 string) {
u.name = name1
}


func (u *User) isElder() bool {
	a := u.age 
	isTrue := false 
	
	if a >= 18 {
		isTrue = true
	} else if a < 18 {
		isTrue = false
	}
	return isTrue
}

func change(u *User) {
	u.age = 18
	u.name = "Ivan"
	u.email = "IvanPichkov12@gmai.com"
}

func main() {
	user := User{"Maksim", 17, "MaksimKidrov99@gmai.com", 1284, []int{12,444,4343,11,10100,9}}

	change(&user) // pointer to user < User{"Maksim", 17, "MaksimKidrov99@gmai.com", 1284} >

	user.getName("Daniel")

	fmt.Printf("Users options: %+v\n", user) // Users options: {name:Daniel age:17 email:IvanPichkov12@gmai.com password:1284}
	fmt.Printf("Name is %v Emailaddress is %v\n", user.name, user.email) // Name is Daniel Email is IvanPichkov12@gmai.com

	fmt.Println(user.isElder()) // true, because in change() age = 18

	fmt.Println(user.getHighScore()) // 9

}