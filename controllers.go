package main
import (
	"net/http"
	"fmt"
	"log"
	"io"
	"os/exec"
	 
    "github.com/julienschmidt/httprouter"
)

func testRedis() {
	println(rdxGet("ac"))
	println(rdxSet("ac", "{Name: 'ac', Age: 25}"))
}

func init() {
	go func() {
	cmd := exec.Command("redis-server")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	}()
	go testRedis()
	
}


func Buy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", nil)
				tmpl.ExecuteTemplate(w, "buy.html", "Buy")
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("POST")
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func Sell(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", nil)
				tmpl.ExecuteTemplate(w, "sell.html", "Sell")
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("POST")
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func Deliver(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", nil)
				tmpl.ExecuteTemplate(w, "deliver.html", "Deliver")
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("POST")
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}