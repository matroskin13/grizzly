# Grizzly

[![Build Status](https://travis-ci.org/matroskin13/grizzly.svg?branch=master)](https://travis-ci.org/matroskin13/grizzly)
[![codecov](https://codecov.io/gh/matroskin13/grizzly/branch/master/graph/badge.svg)](https://codecov.io/gh/matroskin13/grizzly)

Grizzly allows you to use collections in GO without generics. With Grizzly you can use the methods Map, Filter, Find, etc.

## Usage with generation

Install Grizzly:

```bash
$ go get github.com/matroskin13/grizzly
```

And update your working file:

```go
//go:generate grizzly generate main.go

package main

import (
    "fmt"
)

//grizzly:generate
type User struct {
    Id   int
    Name string
    Age  int
}

func main() {
    users := NewUserCollection([]*User{
        {Id: 1, Name: "John", Age: 20},
        {Id: 2, Name: "Tom", Age: 22},
        {Id: 3, Name: "Billy", Age: 20},
        {Id: 4, Name: "Mister X", Age: 30},
    })

    youngUsers := users.Filter(func (user *User) bool {
        return user.Age < 30
    })

    youngUsersIds := youngUsers.MapToInt(func (user *User) int {
        return user.Id
    })

    fmt.Println("young users ids", youngUsersIds)
}
```

And run go generate:

```bash
$ go generate
```

## Generate from config

Create file grizzly.json in your root directory

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

And run Grizzly

```bash
$ grizzly update
```

Now you can use the collection code:

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

List of default methods:  "find", "filter", "maps", "array", "get", "uniq", "sort", "qsort"

## Methods of collection

The following methods will be available for the collection:

grizzly.json

```json
{
  "collections": [
    {
      "name": "user",
      "types": {
        "id": "int",
        "name": "string"
      }
    }
  ]
}
```

by go generate

```go
//grizzly:generate
type User struct {
	Id   int
	Name string
}
```

```go
func NewUserCollection(items []*User) *UserCollection
func NewEmptyUserCollection() *UserCollection

func (c *UserCollection) Len() int

func (c *UserCollection) ForEach(callback func(item *User))

func (c *UserCollection) Filter(callback SearchCallbackUser) *UserCollection

func (c *UserCollection) Find(callback SearchCallbackUser) *User

func (c *UserCollection) Get(index int) (model *User)

func (c *UserCollection) MapToInt(callback func(item *User) int) []int

func (c *UserCollection) MapToString(callback func(item *User) string) []string

func (c *UserCollection) Pop() *User

func (c *UserCollection) Push(item *User) *UserCollection

func (c *UserCollection) Shift() *User

func (c *UserCollection) SortById(mode string) *UserCollection

func (c *UserCollection) SortByName(mode string) *UserCollection

func (c *UserCollection) Sort(comparator func(a *Model, b *Model) int) *UserCollection

func (c *UserCollection) UniqById() *UserCollection

func (c *UserCollection) UniqByName() *UserCollection

func (c *UserCollection) Unshift(item *User) *UserCollection

func (c *UserCollection) ForEach(callback func(item *User))
```
