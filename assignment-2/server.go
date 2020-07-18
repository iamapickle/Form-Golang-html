package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func form(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(res, nil)
}

func uploadFile(res http.ResponseWriter, req *http.Request) string {
	req.ParseMultipartForm(10 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Here: ", handler.Filename)
	fmt.Println("File-size: ", handler.Size)
	fmt.Println("MIME Header: ", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "file-*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer tempFile.Close()
	name := tempFile.Name()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	tempFile.Write(fileBytes)
	fmt.Println("successfully done")
	return name
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("main.html")
	if err != nil {
		log.Fatal(err)
	}
	reqData := req.Form
	fmt.Println(reqData)
	reqName := req.FormValue("PName")
	reqArea := req.FormValue("Area")

	reqFileName := uploadFile(res, req)
	fmt.Println(reqName, reqArea, reqFileName)
	// req.ParseMultipartForm(1024)
	// fileHeader := req.MultipartForm.File["file"][0]
	// // fmt.Println(fileHeader)
	// file, err := fileHeader.Open()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	data := `<div style="margin: 0 auto">
		<div><b>Name:</b>` + reqName + `</div><br><br>
		<div><b>State:</b>` + reqArea + ` </div><br><br>
		<img id="output" width=100% alt="" src="./temp-images/` + reqFileName[12:] + `"/>
	</div>`
	// fmt.Println(data)

	t.Execute(res, data)
}

func main() {
	mux := mux.NewRouter()
	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}
	mux.HandleFunc("/", form)
	mux.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("./"))))
	mux.PathPrefix("/temp-images/").Handler(http.StripPrefix("/temp-images", http.FileServer(http.Dir("./temp-images/"))))
	mux.HandleFunc("/show", mainPage)
	server.ListenAndServe()
}
