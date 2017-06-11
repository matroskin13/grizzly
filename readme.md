## Grizzly codegen

### Generation of collections

```bash
$ go get github.com/matroskin13/grizzly

$ grizzly create user id:int name:string age:int
or
$ $GOPATH/bin/grizzly create users id:int name:string age:int

```

### Use of collections after generation

```go

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

```

# Generate from config

Create a file grizzly.json in your root directory

```json
{
  "collections": [
    {
      "name": "User",
      "types": {
        "id": "int",
        "name": "string",
        "age": "string"
      }
    },
    {
      "name": "Cities",
      "types": {
        "cityId": "int"
      }
    }
  ]
}
```

And run the grizzly

```bash
$ grizzly update
```