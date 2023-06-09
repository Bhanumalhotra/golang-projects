package testpackage

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// keep the first letter capitol so that func is exported
// if dont want to export keep first letter small

func MyFunction(step int) {
	if step == 1 {
		fmt.Println("Step 1")
	} else if step == 2 {
		fmt.Print("Step 2")
	} else {
		fmt.Println("Step not supported")
	}
}

// understanding error
func main() {
	//http.HandleFunc("/create-payment-intnet", handleCreatePayment)
	//http.HandleFunc("/health")
	//var err error = http.ListenAndServe("localhost:4242", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	var err error = returnsError("supersecretpassword")
	if err != nil {
		fmt.Println(err)
	}
}
func returnsError(password string) error {
	var secretPassword = "wrongpassword"

	if password == secretPassword {
		return nil
	} else {
		return errors.New("invalid password")
	}
}

// first class functions are supported
func main() {
	//http.HandleFunc("/create-payment-intnet", handleCreatePayment)
	//http.HandleFunc("/health")
	//var err error = http.ListenAndServe("localhost:4242", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	functionOne(anotherFunction)

}

func functionOne(functionTwo func()) {
	functionTwo()
}

func anotherFunction() {
	fmt.Println("another function was called")
}

// learning about slices //go is typesafe language and not dynamic like python
//dynamic is recipie for disaster

func main() {
	var names []string = []string{"James", "Bill"}
	fmt.Println(names[1])
	var numbers []int = []int{30, 50, 60}
	fmt.Println(numbers[2])
}

// response can have multiple values
func returnMultiple() (string, int, bool) {
	return "string", 1, true
}

//after this we will learn new way of writing variables using :=
package main

import (
"fmt"
"log"
"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePayment)
	http.HandleFunc("/health", handleHealth)
	log.Println("listening on localhost:4242....")

	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleCreatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED!")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	var response []byte = []byte("Server is up and running")
	writer.Write(response)
}



{

someString , someInt, someBool := returnsMultiple()

fmt.Println(someString)
fmt.Println(someInt)
fmt.Println(someBool)

}

func retursMultiple() (string, int, bool){
	return "string", 1, true
}


//declaringAndAssigningAValue
var variableName string
var variableName string = "something"
or
variableName := "something"