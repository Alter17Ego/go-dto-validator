# GoValidator


GoValidator is a Package validator implements value validations for structs and nested structs fields based on tags.


## Tags
----
- ##### [`field` ] value : required
- ##### [`type`] value : number / text
- ##### [`min`] value : (string number)
- ##### [`max`] value : (string number)
- ##### [`startswith`] value : (string)
- ##### [`endswith`] value : (string)
- ##### [`value_of`] value : (string)(can accept 1 or more string value seperated with ,(coma))
- ##### [`format`] value : email /alphabet/alphanumeric


Tag `min` / `max` will automatically validate based on the value type.
String : count the length
Int/numeric : validate the value


### Installation

```
go get github.com/Alter17Ego/go-dto-validator
```
## Usage
---
#### Basic
`````
type User struct {
	FirstName      string       `field:"required"`
	LastName       string       `type:"text"`
	Age            int          `type:"number"`
	Email          string       `format:"email"`
}

func main(){
    MyStruct := User{FirstName:"John",LastName:"S"}
    err := validator.validate(MyStruct)
    if err != nil {
        fmt.Println(err)
    }
}
`````
#### Multiple Validation
`````
type User struct {
	FirstName      string       `field:"required" type:"text" max:"10"`
	LastName       string       `field:"required" type:"text" min:"1"`
	Age            int          `type:"number" max:"1"`
	Email          string       `format:"email"`
}

func main(){
    MyStruct := User{FirstName:"John",LastName:"S"}
    err := validator.validate(MyStruct)
    if err != nil {
        fmt.Println(err)
    }
}
`````
#### Nested Struct validation
`````
type User struct {
	FirstName      string       `field:"required" type:"text" max:"10"`
	LastName       string       `field:"required" type:"text" min:"1"`
	Age            int          `type:"number" max:"1"`
	Email          string       `format:"email"`
}
type Orders struct {
	OrderId             int  `field:"required" max:"10"`
	UserData            User
}

func main(){
    MyStruct := User{FirstName:"John",LastName:"S"}
    MyOrder := Orders{OrderId:9930438,UserData:MyStruct}

    err := validator.validate(MyOrder)
    if err != nil {
        fmt.Println(err)
    }
}
`````
#### Slice Struct Validation
`````
type User struct {
	FirstName      string       `field:"required" type:"text" max:"10"`
	LastName       string       `field:"required" type:"text" min:"1"`
	Age            int          `type:"number" max:"1"`
	Email          string       `format:"email"`
}
type Orders struct {
	OrderId             int  `field:"required" max:"10"`
	UserData            User
}

func main(){
    var MyStruct []User
    MyStruct = append(MyStruct,User{FirstName:"John",LastName:"S"})
    err := validator.validate(Mystruct)
    if err != nil {
        fmt.Println(err)
    }
}
`````
#### Nested Slice Struct Validation
`````
type User struct {
	FirstName      string       `field:"required" type:"text" max:"10"`
	LastName       string       `field:"required" type:"text" min:"1"`
	Age            int          `type:"number" max:"1"`
	Email          string       `format:"email"`
	Orders         []Orders
}
type Orders struct {
	OrderId             int  `field:"required" max:"10"`
}

func main(){
    var myOrders []Orders
    myOrders = append(myOrders,Orders{OrderId:9930438})
    myOrders = append(myOrders,Orders{OrderId:9930432})
    MyStruct := User{FirstName:"John",LastName:"S",Orders:myOrders}
    err := validator.validate(Mystruct)
    if err != nil {
        fmt.Println(err)
    }
}
`````