package views

import (
	"net/http"
	"log"
	"html/template"
	//"io/ioutil"
	"strconv"
	"path"
	"time"
)

//Properties of page variable
type PageVariables struct {
	Date string
	Time string
	Difference string
}



//details of the function indexHandler
func IndexHandler(w http.ResponseWriter, r *http.Request){
	//to get the form element name "dob" on templates link
	now, err :=  time.Parse("2006-01-02", r.FormValue("dob"))
	
	//to get the form element name "age" on templates link 
	age :=  r.FormValue("age")
	//converting the string to interger for calculation
	ageNew, err := strconv.Atoi(age)

	//Add the "age" input from above to get the birthday
	day := now.AddDate(+ageNew, 0, 0)

	//implementing the pageVariable properties from above
	HomePageVars := PageVariables {
		Date: time.Now().Format("Monday-January-2006"),
		Time: time.Now().Format("15:04:05"),
		Difference: day.Format("Monday"),
	}


	//passing the template to templates/index.html i.e rendering the templates
	pathJoin := path.Join("templates", "index.html") 
	tmpl, err := template.ParseFiles(pathJoin) 
	if err != nil {    
		http.Error(w, err.Error(), http.StatusInternalServerError)    
	 	return  
	 }
 
	if err := tmpl.Execute(w, HomePageVars); err != nil {    
		http.Error(w, err.Error(), http.StatusInternalServerError)  
	} 
}
