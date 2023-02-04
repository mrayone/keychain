package webserver

func Serve() {

}

// func Router() http.Handler {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		defer r.Body.Close()
// 		w.Header().Add("content-type", "text/plain")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("OK"))
// 	}).Methods(http.MethodGet)

// 	return handlers.CompressHandler(router)
// }
