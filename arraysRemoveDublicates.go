package main

import (
	"fmt"
)


type Users []struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
}

func main(){
var users, allUsers, users2 Users
res1 := []string{}

users = Users{{"1", "1@halykbank.kz"}, {"2", "2@halykbank.kz"},}
allUsers = Users{{"2", "2@halykbank.kz"}, {"3", "3@halykbank.kz"},}
users2 = append(allUsers, users...)

for _, val := range users2{
	res1 = append(res1, val.Email)
}
fmt.Println(res1)

res1 = removeDuplicates(res1)
fmt.Println(res1)
}

func removeDuplicates(elements []string) []string {
    // Use map to record duplicates as we find them.
    encountered := map[string]bool{}
    result := []string{}

    for v := range elements {
        if encountered[elements[v]] == true {
            // Do not add duplicate.
        } else {
            // Record this element as an encountered element.
            encountered[elements[v]] = true
            // Append to result slice.
            result = append(result, elements[v])
        }
    }
    // Return the new slice.
    return result
}