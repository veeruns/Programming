package cms
import (
	"net/http"
	"time"
	)
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET" :
		Tmpl.ExecuteTemplate(w,"new", nil)
	case "POST" :
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		if( contentType == "page") {
				Tmpl.ExecuteTemplate(w,"page", &Page{
					Title: title,
					 Content : content,
					})
				return
			}
		if( contentType == "post" ){
				
			Tmpl.ExecuteTemplate(w,"post", &Post{
					Title: title,
					 Content : content,
					})
				return
			}
		
		

	}
}

	func ServeIndex (w http.ResponseWriter, r *http.Request) {
		p:= &Page{
			Title: "Go Projects CMS",
			Content : " Welcome to some stuff",
			Posts : []*Post{
				&Post {
					Title: "H1",
					Content : "Tet 1",
					DatePublished: time.Now()

				},
				&Post {
					Title: "H2",
					Content: "Test 2",
					DatePublished: time.Now().Add(-time.Hour),
					Comments: []*Comment{
						&Comments {
						 Author: "T",
						 Comment: "Some Comment",
						 DatePublished: time.Now.Add(-time.Hour /2 )
						}
					}
				}
			}
		}
		Tmpl.ExecuteTemplate(w, "page", p)
	}