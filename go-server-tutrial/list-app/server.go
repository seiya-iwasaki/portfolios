package main

import (
	"fmt"
    "bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

type BookList struct {
    Books []string
}

func New(books []string) *BookList {
    return &BookList{Books: books}
}


func fileRead(fileName string) []string {
    var booklist []string
    file, err := os.Open(fileName)
    if os.IsNotExist(err) {
        return nil
    }
    defer file.Close()
    scaner := bufio.NewScanner(file)
    for scaner.Scan() {
        booklist = append(booklist, scaner.Text())
    }
    return booklist
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    bookList := fileRead("reading.txt")
    html, err := template.ParseFiles("view.html")
    if err != nil {
        log.Fatal(err)
    }
    getBooks := New(bookList)//*BookList{Books: books}
    if err := html.Execute(w, getBooks); err != nil {
        log.Fatal(err)
    }
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    formValue := r.FormValue("value")
    file, err := os.OpenFile("reading.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }
    _, err = fmt.Fprintln(file, formValue)
    if err != nil {
        log.Fatal(err)
    }
    http.Redirect(w, r, "/view", http.StatusFound)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    hello := []byte("Hello World!")
    _, err := w.Write(hello)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/view", viewHandler)
    http.HandleFunc("/view/create", createHandler)
    fmt.Println("Server start up...")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

