package httpGo
import (
	"net/http"
	"github.com/gorilla/mux"
 	//"regexp/syntax"
	"regexp"
	"fmt"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

/*
Router using gorilla/mux
*/
func MyNewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}


	router.HandleFunc("/SNPutObj", SNPutObj)
	//router.HandleFunc("/SNodeListenNoP2P", SNodeListenNoP2P)

	router.HandleFunc("/SNPutObjP2PRequest", SNPutObjP2PRequest)
	router.HandleFunc("/ReturnObjProxy", ReturnObjProxy)
	router.HandleFunc("/prepSN", prepSN)
	router.HandleFunc("/putObj", PutObjAPI)	// Todo: account/container/object
	router.HandleFunc("/putAcc", PutAccAPI)
	router.HandleFunc("/SNPutAcc", SNPutAcc)
	router.HandleFunc("/SNPutAccP2PRequest", SNPutAccP2PRequest)
	router.HandleFunc("/SNPutCont", SNPutCont)
	router.HandleFunc(`/{rest:[a-zA-Z0-9/\-\/]+}`, route)



	return router
}

var triple = regexp.MustCompile(`[a-zA-Z0-9]+/[a-zA-Z0-9]+/[a-zA-Z0-9]`)  // Has digit(s)
var double = regexp.MustCompile(`[a-zA-Z0-9]+/[a-zA-Z0-9]`)  // Has digit(s)
var simple = regexp.MustCompile(`[a-zA-Z0-9]`) // Contains "abc"
//has 2, 3 or 4 slashes
func route(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	switch {
	case triple.MatchString("/"+r.URL.Path):
		tripleF(w, r)
	case double.MatchString("/"+r.URL.Path):
		doubleF(w, r)
	case simple.MatchString("/"+r.URL.Path):
		simpleF(w, r)
	default:
		w.Write([]byte("Unknown Pattern"))
	}
}
func tripleF(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		PutObjAPI(w,r)
	}
	fmt.Println("triple")
}
func doubleF(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		PutContAPI(w,r)
	}
	fmt.Println("double")
}
func simpleF(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		PutAccAPI(w,r)
	}
	fmt.Println("simple")

}
var routes = Routes{

	Route{
		"GetNodes",
		"GET",
		"/GetNodes",
		GetNodes,
	},
	Route{
		"/GetNodesForKey",
		"GET",
		"/GetNodesForKey",
		GetNodesForKey,
	},
	Route{
		"/SNPutObjGetChunks",
		"POST",
		"/SNPutObjGetChunks",
		SNPutObjGetChunks,
	},



}
