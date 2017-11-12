package main

import (
	"encoding/json"
	"fmt"
	 "net/http"
	 "github.com/gorilla/mux"
)
//Test 

type Person struct{
	ID string
	FirstName string
	LastName string
	Address *Address
}

type Address struct{
	City string
	State string
}

func createPersonHandller(w http.ResponseWriter, r *http.Request){
json.NewEncoder(w).Encode(people)
}

func getPersonHandller(w http.ResponseWriter, req *http.Request){

	params:=mux.Vars(req)
for _, item:=range people{

	json.NewEncoder(w).Encode(item)

	fmt.Println(params["Id"])

	if item.Id==params["Id"]{
		json.NewEncoder(w).Encode(item)
		return
	}else{
		fmt.Println("Person is not Exist")
	}
}

}

var people []Person

func main(){
fmt.Println("Rest Api Start...")
people=append(people,Person{Id:"1",FirstName:"Test",LastName:"Test",Address:&Address{City:"istanbul",State:"Turkey"}})
r:=mux.NewRouter()
r.HandleFunc("/",createPersonHandller).Methods("GET") 
r.HandleFunc("/get",getPersonHandller).Methods("GET") 
http.Handle("/",r) 
http.ListenAndServe(":8080",r)	
}




 