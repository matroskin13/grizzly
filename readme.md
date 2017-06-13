# Grizzly codegen

[![Build Status](https://travis-ci.org/matroskin13/grizzly.svg?branch=master)](https://travis-ci.org/matroskin13/grizzly)
[![codecov](https://codecov.io/gh/matroskin13/grizzly/branch/master/graph/badge.svg)](https://codecov.io/gh/matroskin13/grizzly)

## Generation of collections

```bash
$ go get github.com/matroskin13/grizzly

$ grizzly create user id:int name:string age:int
or
$ $GOPATH/bin/grizzly create users id:int name:string age:int

```

## Use of collections after generation

```go

package test

import (
    "fmt"
    "test/collections"
)

func main() {
    users := collections.NewUserCollection([]*collections.User{
        {Id: 1, Name: "John", Age: 20},
        {Id: 2, Name: "Tom", Age: 22},
        {Id: 3, Name: "Billy", Age: 20},
        {Id: 4, Name: "Mister X", Age: 30},
    })

    city := collections.NewCitiesCollection([]*collections.Cities{
        {CityId: 1},
    }).Find(func (city *collections.Cities) bool {
        return true
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

    uniqAges := users.UniqByAge()
    sortedAges := users.SortByAge("asc").MapToInt(func (user *collections.User) int {
        return user.Age
    })

    fmt.Println("tom", Tom)
    fmt.Println("young users ids", youngUsersIds)
    fmt.Println("first city", city)
    fmt.Println("uniq ages", uniqAges)
    fmt.Println("sorted ages", sortedAges)
}
```

## Generate from config

Create a file grizzly.json in your root directory

```json
{
  "collections": [
    {
      "name": "user",
      "types": {
        "id": "int",
        "name": "string",
        "age": "int"
      }
    },
    {
      "name": "city",
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

You can also specify the required methods:

```json
{
  "name": "User",
  "types": {
    "id": "int",
    "name": "string",
    "age": "int"
  },
  "methods": ["find", "filter"]
}
```

List of default methods:  "find", "filter", "maps", "array", "get", "uniq", "sort"

## Methods of collection

The following methods will be available for the collection:

```json
{
  "collections": [
    {
      "name": "Users",
      "types": {
        "id": "int",
        "name": "string"
      }
    }
  ]
}
```

```go
type SearchCallbackUsers func(item *Users) bool

type Users struct {
    Id   int
    Name string
}

type UsersCollection struct {
    Items []*Users
}

func NewUsersCollection(items []*Users) *UsersCollection

func (c *UsersCollection) Filter(callback SearchCallbackUsers) *UsersCollection

func (c *UsersCollection) Find(callback SearchCallbackUsers) *Users

func (c *UsersCollection) Get(index int) (model *Users)

func (c *UsersCollection) MapToInt(callback func(item *Users) int) []int

func (c *UsersCollection) MapToString(callback func(item *Users) string) []string

func (c *UsersCollection) Pop() *Users

func (c *UsersCollection) Push(item *Users) *UsersCollection

func (c *UsersCollection) Shift() *Users

func (c *UsersCollection) SortById(mode string) *UsersCollection

func (c *UsersCollection) SortByName(mode string) *UsersCollection

func (c *UsersCollection) UniqById() *UsersCollection

func (c *UsersCollection) UniqByName() *UsersCollection

func (c *UsersCollection) Unshift(item *Users) *UsersCollection
```