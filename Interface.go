package main

import io "fmt"

type Student struct{
  name string
}

type Speaker interface{
  SayHi()
}

func (s Student)SayHi(){
  io.Println("hi, my name is %s.", s.name)
}

func main(){
  s := Student{"lee"}
  var speaker Speaker
  speaker = s
  speaker.SayHi()
}
