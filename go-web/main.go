package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
	"log"
)

var client *redis.Client
var templates *template.Template


func main(){
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	templates = template.Must(template.ParseGlob("templates/*.html"))
    // http.Handle("/", http.FileServer(http.Dir("/tmp")))
	r:= mux.NewRouter()
	r.HandleFunc("/",indexHandler).Methods("GET") 
	http.Handle("/",r) 	
	http.ListenAndServe(":8080",nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
			renderTemplate(w, "templates/index.html", nil)	
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}
