package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {

	http.Handle("/static/assets/images/", http.StripPrefix("/static/assets/", http.FileServer(http.Dir(path.Join(".", "/static/assets/")))))
	http.Handle("/static/assets/css/", http.StripPrefix("/static/assets/css/", http.FileServer(http.Dir(path.Join(".", "/static/assets/css/")))))
	http.Handle("/static/assets/js/", http.StripPrefix("/static/assets/js/", http.FileServer(http.Dir(path.Join(".", "/static/assets/js/")))))

	http.HandleFunc("/", serverTemplate)

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func serverTemplate(w http.ResponseWriter, r *http.Request) {

	urlPath := filepath.Clean(r.URL.Path)
	urlExt := filepath.Ext(urlPath)

	lp := filepath.Join("./data/templates", "layout.html")
	fp := filepath.Join("./data/pages", urlPath)

	if urlPath == "/" {
		fp = "./data/pages/welcome-page.html"
		lp = "./data/templates/welcome-page.html"
	}

	if urlExt != ".html" {
		fmt.Println("Debug :: UrlPath", urlPath)
		fmt.Println("Debug :: File extension", filepath.Base(urlPath))

	}

	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {

			http.NotFound(w, r)
			return
		}

	}

	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, _ := template.ParseFiles(lp, fp)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}
