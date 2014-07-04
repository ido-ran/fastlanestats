package goplus

import (
        "net/http"
		"time"
		"fmt"
	//	"io/ioutil"

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
		http.HandleFunc("/test", test)
		http.HandleFunc("/testfetch", testfetch)
}

func test(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

	pp := PricePoint {
		PointInTime: time.Now(),
		Value: 4.43,
	}
	
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "pricepoint", nil), &pp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var pp2 PricePoint
    if err = datastore.Get(c, key, &pp2); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Stored and retrieved the PricePoint at %q value %q", pp2.PointInTime, pp)
}

func testfetch(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get("https://www.fastlane.co.il/Mobile.aspx")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	doc, e := goquery.NewDocumentFromResponse(resp);
	if (e != nil) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	doc.Find("span#lblPrice").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        price := s.Text()
        fmt.Fprintf(w, "price is %s\n", price)
    })

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	
	// n := bytes.Index(body, []byte{0})
	// s := string(body[:n])
   // fmt.Fprintf(w, "HTTP GET returned status %v", s)
}