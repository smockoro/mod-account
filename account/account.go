package account

import (
	"fmt"

	u "github.com/smockoro/mod-user/user"
)

func AccountHello() {
	fmt.Println("mod Account v1.0.0")
	user := u.User{Name: "Mod account Gopher", Age: 20}
	user.UserHello()
}
