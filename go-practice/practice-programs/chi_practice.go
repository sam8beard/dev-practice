package main 

import ( 
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() { 
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/ping", pingHandler) 
	router.Post("/upload", uploadHandler)

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	} // server
	fmt.Println("Listening on port", server.Addr)
	err := server.ListenAndServe(); if err != nil {fmt.Println("Failed to listen to server", err)}
} // main

func pingHandler(w http.ResponseWriter, r *http.Request) { 
	w.Write([]byte("Pong!"))
} // pingHandler

func uploadHandler(w http.ResponseWriter, r *http.Request) { 
	// populate Multipart Form to retrieve file and file header 
	err := r.ParseMultipartForm(10 << 20)
	
	if err != nil { 
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	} // if 

	// If file form fields are supplied, iterate through map and process
	if len(r.MultipartForm.File) != 0 {
		if r.MultipartForm.File["files"] != nil { 
			for _, header := range r.MultipartForm.File["files"] {
				fmt.Fprintf(w, "Size of %s: %v bytes\n", header.Filename, header.Size)
			} // for 
		} else { 
			fmt.Fprintln(w, "For files, please use the field name, 'files'.")
		}
	} // if 
	
	// If non-file form fields are supplied, iterate through map and process
	if len(r.MultipartForm.Value) != 0 { 
		if r.MultipartForm.Value["urls"] != nil { 
			for _, value := range r.MultipartForm.Value["urls"] { 
				fmt.Fprintln(w, value)
			} // for	
		} else { 
			fmt.Fprintln(w, "For urls, please use the field name, 'urls'.")
		} // if 
	} // if 
} // uploadHandler