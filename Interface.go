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

type Tearch struct{
  id uint32
}

type Manager struct{
  Tearch
  manager_name string
}

type TearchInfo interface{
  slary()
}

type ManagerInfo interface{
  slary()
  managerInfo()
}

func (t Tearch)slary(){
  io.Println(t)
}

func (m Manager) managerInfo(){
  io.Println(m)
}

func main(){
  s := Student{"lee"}
  var speaker Speaker
  speaker = s
  speaker.SayHi()
  var mn ManagerInfo = Manager{Tearch{123}, "seaung"}
  mn.managerInfo()
  mn.slary()
  var te Tearch = mn
  te.managerInfo()
  te.slary()
}
