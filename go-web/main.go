package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
)

var client *redis.Client
var templates *template.Template

func main(){
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r:= mux.NewRouter()
	r.HandleFunc("/",indexHandler).Methods("GET") 
	http.Handle("/",r) 	
	http.ListenAndServe(":8080",nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w,"index.html",nil)
}