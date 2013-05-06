package main

import (
  "errors"
  "fmt"
  "html/template"
  "io/ioutil"
  "net/http"
  "regexp"
)

// Here are some simple tasks you might want to tackle on your own:

// Store templates in tmpl/ and page data in data/.
// Add a handler to make the web root redirect to /view/FrontPage.
// Spruce up the page templates by making them valid HTML and adding some CSS rules.
// Implement inter-page linking by converting instances of [PageName] to 
// <a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)

const lenPath = len("/view/")
const filePath = "data/"

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

type Page struct {
  Title string
  Body  []byte
}

func (p *Page) save() error {
  filename := filePath + p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := filePath + title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, req *http.Request, title string) {
  // title := req.URL.Path[lenPath:]
  // title, err := getTitle(w, req)
  // if err != nil {
  //   return
  // }

  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w, req, "/edit/" + title, http.StatusFound)
    return 
  }
  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, req *http.Request, title string) {
  // title := req.URL.Path[lenPath:]
  // title, err := getTitle(w, req)
  // if err != nil {
  //   return
  // }
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, req *http.Request, title string) {
  // title := req.URL.Path[lenPath:]
  // title, err := getTitle(w, req)
  // if err != nil {
  //   return
  // }
  body := req.FormValue("body")
  p := &Page{Title: title, Body: []byte(body)}
  err := p.save()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return 
  }
  http.Redirect(w, req, "/view/" + title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  err := templates.ExecuteTemplate(w, tmpl + ".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func getTitle(w http.ResponseWriter, req *http.Request) (title string, err error) {
  title = req.URL.Path[lenPath:]
  if !titleValidator.MatchString(title) {
    http.NotFound(w, req)
    err = errors.New("Invalid Page Title")
  }
  return
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  return func(w http.ResponseWriter, req *http.Request) {
    title := req.URL.Path[lenPath:]
    if !titleValidator.MatchString(title) {
      http.NotFound(w, req)
      return
    }

    fn(w, req, title)
  }
}

func main() {
  fmt.Println("Moogo starting...")
  http.HandleFunc("/view/", makeHandler(viewHandler))
  http.HandleFunc("/edit/", makeHandler(editHandler))
  http.HandleFunc("/save/", makeHandler(saveHandler))
  http.ListenAndServe(":8210", nil)
}