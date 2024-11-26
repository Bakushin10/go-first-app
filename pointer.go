package main

import "fmt"


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

  
  s := "hello, world"

  s = "hello"
  // strchange(s)
  fmt.Println(s)


}
