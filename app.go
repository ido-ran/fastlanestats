package fastlanestat

import (
        "net/http"

)

// init is called before the application start
func init() {
        // Register a handler for /gopher URLs.
		http.HandleFunc("/fetchprice", fetchpriceHandler)
		http.HandleFunc("/", viewStatsHandler)
}