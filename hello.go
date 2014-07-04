package goplus

import (
        "net/http"
		"time"
		"fmt"
		"strconv"

        // Import appengine urlfetch package, that is needed to make http call to the api
        "appengine"
        "appengine/urlfetch"
    	"appengine/datastore"

		"github.com/PuerkitoBio/goquery"
)

type PricePoint struct {
    PointInTime time.Time
    Value  float32
}

// gopherFallback is the official gopher URL (in case we don't find any in the Google+ stream)
const gopherFallback = "http://golang.org/doc/gopher/gophercolor.png"

// init is called before the application start
func init() {
        // Register a handler for /gopher URLs.
		http.HandleFunc("/fetchprice", fetchprice)
}

func fetchprice(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get("https://www.fastlane.co.il/Mobile.aspx")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	doc, err := goquery.NewDocumentFromResponse(resp);
	if (err != nil) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	doc.Find("span#lblPrice").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        priceText := s.Text()
        fmt.Fprintf(w, "price is %s\n", priceText)

		price, err := strconv.ParseFloat(priceText, 32)
		if (err != nil) {
	        http.Error(w, err.Error(), http.StatusInternalServerError)
	        return		
		}
		
		pp := PricePoint {
			PointInTime: time.Now(),
			Value: float32(price),
		}
		
		_, err = datastore.Put(c, datastore.NewIncompleteKey(c, "PricePoint", nil), &pp)
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        	return
    	}

    })

}