//1- The main package that executes the program
package main

//2- import of all the required packages
import (
	"net/http"
	"log"
	"views"
)


/*3- The main function that handles server connection and relates with other public 
functions. */

func main () {

	//The index directory and the function that works on the dirctory
	http.HandleFunc("/index/", views.IndexHandler)

	//Parse the css, js etc for the templates
	http.Handle("/static/", http.FileServer(http.Dir("public")))


	/*It request connection to server through port :8080 and responded with 
	"I don listen" on succesful connection */
	log.Println("I don listen")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


