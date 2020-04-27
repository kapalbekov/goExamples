package main

import (
	"github.com/satori/go.uuid"
	"fmt"
	)


func IsValidUUID(u string) bool {
  _, err := uuid.FromString(u)
  return err == nil
}

func main (){
	fmt.Println("strin is UUID ", IsValidUUID("437665f1-2a01-49ab-85e1-8329b3b1c31"))
}