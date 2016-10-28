package main

import(
  io "fmt"
)

type Person struct{
  name, sex string
}

func main(){
  p:Person{name:"leeseaung", sex:"M"}
  p.print()
}

func (p Person)print(){
  p.name = "seaung"
  p.sex = "m"
  io.Println(p.name, p.sex)
}
