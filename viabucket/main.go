package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/404th/viabucket/image"
	"github.com/joho/godotenv"
)

func main() {
	// initializing .env files
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while loading .env: %s", err.Error())
	}

	// loading...
	http.HandleFunc("/image", uploadImage)

	http.ListenAndServe(":9798", nil)
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// tempFile, err := ioutil.TempFile("imgs", "upload-*.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := image.ImageUpload(fileBytes)
	if err != nil {
		fmt.Println("error while uploading file: %w", err)
		return
	}

	fmt.Println(resp)
	fmt.Fprintf(w, resp.Location)
}
