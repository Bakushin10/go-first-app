package main

import "fmt"

	
type person struct {
    name string
    age  int
}

func zeroval(ival int){
  ival = 0
}

func zeroptr(iptr *int){
  *iptr = 0
}


func strchange(s string){
  s = "hello"
}

func main(){

  i := 1
  zeroval(i)
  zeroptr(&i)
  fmt.Println(i)

  
  // s = "hello"
  // strchange(s)
  // fmt.Println(si)

  s := person{ name: "Shin", age : 50}

  sp := &s

  s.age = 62

  fmt.Println(sp)
  fmt.Println(s)

}
