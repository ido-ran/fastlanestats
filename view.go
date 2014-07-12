package fastlanestat

import (
        "net/http"
		"html/template"
//		"fmt"

        // Import appengine urlfetch package, that is needed to make http call to the api
        "appengine"
    	"appengine/datastore"
)

type ViewContext struct {
	PricePoints []PricePoint
}

func viewStatsHandler(w http.ResponseWriter, r *http.Request) {
   	c := appengine.NewContext(r)

	// The Query type and its methods are used to construct a query.
	q := datastore.NewQuery("PricePoint").
//	        Filter("PointInTime >=", "2014-07-12 11:42:37").
//	        Filter("PointInTime <=", "2014-07-13 11:42:37").
	        Order("PointInTime")

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	var pricePoints []PricePoint
	//_, err := 
	q.GetAll(c, &pricePoints)
	// handle error
	// ...

	viewContext := ViewContext{ PricePoints: pricePoints }
	t, _ := template.ParseFiles("templates/simple.htmltemplate")
	t.Execute(w, viewContext)
}