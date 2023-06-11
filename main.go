package main
import (
	"html/template"
	"net/http"
  	"log"
	"fmt"
	 
    "github.com/julienschmidt/httprouter"
)

const PORT = "localhost:4000"
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
testRedis()
	HandleRoutes()
	return
}

func HandleRoutes() {

	router := httprouter.New()
	router.GET("/", Index)
	/*
	router.GET("/new", NewPhotoGet)
	router.POST("/upload", NewPhotoPost)
	router.GET("/post/:postid", ShowPost)
	router.GET("/delete/:postid", DeletePost)
	router.GET("/user/:name", UserProfile)	
*/	

	router.GET("/register", Register)
	router.POST("/register", Register)
	router.GET("/login", loginHandler)
	router.POST("/login", loginHandler)
	router.GET("/logout", logoutHandler)


	router.NotFound = http.FileServer(http.Dir(""))
	router.ServeFiles("/img/*filepath", http.Dir("uploads"))
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	
	log.Println("Starting erver on ", PORT)
	err := http.ListenAndServe(PORT, router)
//err := http.ListenAndServe(GetPort(), router)
 	if err != nil {
		log.Fatal("error starting http server : ", router)
 	}

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		if isLoggedIn(r) {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nav.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		} else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}