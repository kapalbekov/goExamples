package main 
import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)	

type Claims struct {
	User *User `json:"user"`
	*jwt.StandardClaims
}

type User struct {
	Username   string    `json:"username"`
	Mail       string    `json:"mail"`
	Name       string    `json:"name"`
	Department string    `json:"department"`
	Manager    string    `json:"manager"`
	JobTitle   string    `json:"jobTitle"`
	City       string    `json:"city"`
	Groups     []*Group  `json:"groups"`
	Modules    []*Module `json:"modules"`
}

type Group struct {
	Name  string `json:"name"`
	Type  string `db:"type"json:"type"`
	Level int    `db:"level"json:"level"`
}

type Module struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func main(){
	//ExampleNewWithClaims_standardClaims()
	ExampleParseWithClaims_customClaimsType()
}

func ExampleNewWithClaims_standardClaims() {
	mySigningKey := []byte("alfa-jwt-key")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	//Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwMCwiaXNzIjoidGVzdCJ9.yGmjOUKu7jaYMVbdTyTzZrg80AWW4RaGZuU_rJ0YgR0 <nil>
}

func ExampleParseWithClaims_customClaimsType() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwfQ.iOLmmFcotY27Py8LMlTY1Ctjaev2X1TWI2V5GJzBIvY"

	type MyCustomClaims struct {
		Foo string `json:"BAR"`
		jwt.StandardClaims
	}

	// sample token is expired.  override time so it parses as valid
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("alfa-jwt-key"), nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
		} else {
			fmt.Println(err)
		}
	})

	// Output: bar 15000
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}