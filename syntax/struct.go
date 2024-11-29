package main

import "fmt"

type Person struct{
  name string
  age int
  address  Address
}

type Address struct{
  zip int
  city string
  state string
}

func (person Person) checkLegalAge() bool {
  return person.age > 21
}

func (person Person) isTexas() bool {
  return person.address.state == "TX"
}

func getPeople() []Person{
  person1 := Person{
        name: "shin", // Explicit field assignment (recommended for clarity)
        age:  24,
        address: Address{
            zip:   10000,
            city:  "Dallas",
            state: "TX",
        },
    }

  person2 := Person{
        name: "jose", // Explicit field assignment (recommended for clarity)
        age:  24,
        address: Address{
            zip:   20000,
            city:  "Chicago",
            state: "IL",
        },
    }

  return []Person{person1, person2}

}

func main(){

  people := getPeople() 

  for _, p := range people {
   if p.checkLegalAge() && p.isTexas(){
      fmt.Println("state", p.address.state, "Legal", p.age)
    }else{
      fmt.Println("not a legal age", p.age)
    } 
  }

}
