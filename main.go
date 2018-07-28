package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)


func headers(w http.ResponseWriter, r *http.Request){
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w,h)
}

func bodyRecv(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte, len)
	_ , err := r.Body.Read(body)
	if err != nil{
		log.Println(err)
	}
	fmt.Println(string(body))
	fmt.Fprintf(w , string(body))
}

func handleForm(w http.ResponseWriter, r *http.Request){
	//r.ParseForm()
	//fmt.Fprintln(w, r.Form)
	//fmt.Fprintln(w, r.PostForm)

	//name := r.FormValue("name")
	//pass := r.FormValue("pass")
	name := r.PostFormValue("name")
	pass := r.PostFormValue("pass")
	fmt.Println(name ,pass)
	fmt.Fprintln(w, name+pass)
}

func  fileRecv(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("upload")
	if err == nil{
		data ,err := ioutil.ReadAll(file)
		if err == nil{
			fmt.Fprintln(w, string(data))
			return
		}
	}
	//r.ParseMultipartForm(1024)
	//fileHeader := r.MultipartForm.File["upload"][0]
	//file , err := fileHeader.Open()
	//if err == nil{
	//	data ,err := ioutil.ReadAll(file)
	//	if err == nil{
	//		fmt.Fprintln(w, string(data))
	//		return
	//	}
	//}
	fmt.Fprintln(w, err.Error())
}

func  jsonRecv(w http.ResponseWriter, r *http.Request){
	//data := make([]byte, 1024)
	data , err := ioutil.ReadAll(r.Body)
	//_, err := r.Body.Read(data)
	if err !=nil{
		fmt.Fprintln(w, err.Error())
		return
	}

	s := struct{
		Name string
		Pass string
		Age  int `json:"age,string"`
	}{}

	err = json.Unmarshal(data, &s)
	if err != nil{
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Println(s)
	fmt.Fprintln(w, string(data))
}

func writeHtml(w http.ResponseWriter, r *http.Request){
	data , err := ioutil.ReadFile("a.html")
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(data)
}

func writeJson(w http.ResponseWriter, r *http.Request){
	s := struct{
		Name string	`json:"name"`
		Pass string `json:"pass"`
		Age  int `json:"age,string"`
	}{
		Name:"chenchao",
		Pass:"123",
		Age:22,
	}
	js , err := json.Marshal(&s)
	if err != nil{
		fmt.Fprintln(w,err.Error())
		return
	}
	//w.Header().Set("Content-Type","application/json")
	w.Write(js)
}

func status(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("no such service"))
	/*
	Todo: xxxx
	*/
}

func redirect(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "/redirect")
	w.WriteHeader(http.StatusFound)
}

func main() {
	http.HandleFunc("/header", headers)
	http.HandleFunc("/body", bodyRecv)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/upload", fileRecv)
	http.HandleFunc("/json", jsonRecv)

	http.HandleFunc("/write", writeHtml)
	http.HandleFunc("/writeJson", writeJson)
	http.HandleFunc("/status", status)

	http.HandleFunc("/redirect", redirect)

	log.Fatal(http.ListenAndServe(":7878", nil))
}
