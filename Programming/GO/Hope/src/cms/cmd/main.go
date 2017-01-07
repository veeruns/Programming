package main

import (
 "cms"
 "os"
 )
func main() {
	p := &cms.Page {
			Title : "Hello World!",
			Content: "This is some weird webpage",
	}
	cms.Tmpl.ExecuteTemplate(os.Stdout, "index",p)
}
