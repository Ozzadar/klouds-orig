// CREATING A BASIC WEB APPLICATION IN GO

package main

import (
	"net/http"
	"github.com/russross/blackfriday"
	//"encoding/json"
	"gopkg.in/unrolled/render.v1"
	"html/template"
	"path"
)


type Book struct {
	Title string	`json:"title"`
	Author string	`json:"author"`
}

func main() {

	/* SERVE WHOLE DRIVE TO THE OUTSIDE
	http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir("/")))
	*/

	/* RANDOM THING
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe("0.0.0.0:8080", nil) */

	r:= render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now"))

	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))

		})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello":"json"})

		})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		book := Book {"Touching Children is Bad: A real life story", "Michael Jackson"}

		r.HTML(w, http.StatusOK, "index", book)

		})



	//http.HandleFunc("/", ShowBooks)
	http.ListenAndServe("0.0.0.0:8080", mux)

}

// HTML IMPLEMENTATION 
func ShowBooks( w http.ResponseWriter, r *http.Request) {
	book := Book {"Touching Children is Bad: A real life story", "Michael Jackson"}

	fp := path.Join ("templates", "index.html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


/*JSON IMPLEMENTATION
func ShowBooks( w http.ResponseWriter, r *http.Request) {
	book := Book{"The Untold Story: How I Fucked Up", "Daniel Sont"}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
*/
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}