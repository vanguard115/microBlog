package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//ConfigFile path
var ConfigFile = "./data/mapping.json"

//RouterMap router map
var RouterMap map[string]ArtMap

//ArtNodes will carry json data
type ArtNodes struct {
	ArtNodes []ArtMap `json:"articles_map"`
}

func init() {

}

//ArtMap = Article Mapping
type ArtMap struct {
	ArtTitle    string `json:"article_title"`
	ArtName     string `json:"article_name"`
	ArtHtml     string `json:"html_file"`
	ArtTemplate string `json:"layout_file"`
}

func main() {

	fmt.Println("microBlog Started")

	// //load confiuration
	RouterMap = ConfigLoader(ConfigFile)
	fmt.Println(RouterMap)

	fmt.Println("\t.config files loaded")

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

	fmt.Println("\t.Webserver listening")

}

func serverTemplate(w http.ResponseWriter, r *http.Request) {

	urlPath := filepath.Clean(r.URL.Path)
	urlExt := filepath.Ext(urlPath)

	fp := "./data/pages/welcome-page.html"
	lp := "./data/templates/welcome-page.html"

	//lp := filepath.Join("./data/templates", "layout.html")
	//fp := filepath.Join("./data/pages", urlPath)

	//get details from  configfile
	//fmt.Println("Debug :: filepath ", urlPath)

	//fmt.Println("Debug :: map data ", RouterMap)

	//if called url has no extension call router
	if urlExt == "" && urlPath != "" && urlPath != "/" {
		urlPath := strings.Replace(urlPath, "/", "", -1)
		r, exist := RouterMap[urlPath]
		fmt.Println("Debug :: exist ", exist)
		if exist == true {
			fp = "./data/pages/" + r.ArtHtml
			lp = "./data/templates/" + r.ArtTemplate
			fmt.Printf("Processing article named %v whith files html %v layout %v", r.ArtName, r.ArtHtml, r.ArtTemplate)
		} else {
			//fmt.Println("Debug :: Inside Error ", urlPath)
			fp = "./data/pages/error.html"
			lp = "./data/templates/error.html"
		}
	}

	fmt.Println("Debug :: ##### ", fp, lp)

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
