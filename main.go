/*
Generation of collection code from a config or cli
Example:
  	package main

	import (
	    "fmt"
	    "test/collections"
	)

	func main() {
	    users := collections.NewUserCollection([]*collections.User{
		{Id: 1, Name: "John", Age: 20},
		{Id: 2, Name: "Tom", Age: 22},
		{Id: 3, Name: "Billy", Age: 19},
		{Id: 4, Name: "Mister X", Age: 30},
	    })

	    youngUsers := users.Filter(func (user *collections.User) bool {
		return user.Age < 30
	    })

	    Tom := youngUsers.Find(func (user *collections.User) bool {
		return user.Name == "Tom"
	    })

	    youngUsersIds := youngUsers.MapToInt(func (user *collections.User) int {
		return user.Id
	    })

	    fmt.Println(Tom, youngUsersIds)
	}
*/

package main

import (
	"github.com/matroskin13/grizzly/cmd"
)

func main() {
	cmd.Init()
}
