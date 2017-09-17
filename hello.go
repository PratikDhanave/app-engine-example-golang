package hello

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

)

func init() {
	http.HandleFunc("/save", handler)
	http.HandleFunc("/retrieve", retrieve)
}




type Store struct {
			Input string
}

func retrieve(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	q := datastore.NewQuery("Store")

	html := ""

	iterator := q.Run(ctx)

	for {
		var entity Store
   	_, err := iterator.Next(&entity)

		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		html += `
			<dt>` + entity.Input + `</dt>		`
}

	w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w,html)
}

func handler(w http.ResponseWriter, r *http.Request) {

  ctx := appengine.NewContext(r)

	param := r.URL.Query().Get("input")

	entity := &Store{}

  entity.Input = param

//	key := datastore.NewKey(ctx, "Store","Input",0, nil)

	key := datastore.NewIncompleteKey(ctx, "Store", nil)

	_, err := datastore.Put(ctx, key, entity)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

fmt.Fprint(w, "Value = ",param,"\tstored in Database")
}
