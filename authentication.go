package main

import io "fmt"

type Authentication struct{
  username, password string
}

func main(){
  a:Authentication{}
  a.Author()
}

func (a Authentication)Author string{
  var i = 0
  var ms string
  io.Println("please enter your username and password:")
  io.Scanln(&a.username, &a.password)
  for ; i < 3; i++{
    if a.username == 'seaung' && a.password == '123456'{
      break
    }
  }
  
  return ms = "success!"
}
