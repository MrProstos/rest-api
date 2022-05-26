package main

func main() {
	/*router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/registration/", server2.Registration).Methods("GET")

	//router.HandleFunc("/", middleware(http.HandlerFunc(server.Auth))).Methods("GET")
	router.HandleFunc("/auth/", server2.Auth).Methods("GET")
	//	router.HandleFunc("/db/",server.Middleware(http.HandlerFunc(server.)))
	router.HandleFunc("/db/showclietns/{phone_num}", server2.ShowClients).Methods("GET")
	router.HandleFunc("/db/addclient/", server2.AddClient).Methods("POST")
	router.HandleFunc("/db/updateclient/", server2.UpdateClient).Methods("PUT")
	router.HandleFunc("/db/delclient/", server2.DelClient).Methods("DELETE")

	router.HandleFunc("/db/showorder/{client_id}", server2.ShowOrder).Methods("GET")
	router.HandleFunc("/db/addorder/", server2.AddOrder).Methods("POST")
	router.HandleFunc("/db/updateorder/", server2.UpdateOrder).Methods("PUT")
	router.HandleFunc("/db/delorder/", server2.DelOrder).Methods("DELETE")
	err := http.ListenAndServe(":2000", router)
	if err != nil {
		log.Fatal(err)
	}
	*/
}
