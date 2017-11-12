package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
)

var client *redis.Client
var templates *template.Template

//ı will continue..
func main(){
	templates = template.Must(template.ParseGlob("templates/*.html"))
 client=redis.NewClient(&redis.Options{
	 Addr:"localhost:8080",
 })	
	
	r:= mux.NewRouter()
	r.HandleFunc("/",indexHandler).Methods("GET") 
	http.Handle("/",r) 	
	http.ListenAndServe(":8080",nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	comments, err:=client.LRange("comments",0,10).Result()
	 if err!=nil{
	 	return
	 }
	templates.ExecuteTemplate(w,"index.html",comments)
}
