package main 

import ( 
	"net/http"
	"fmt"
	"io"
	// "mime/multipart"
	// "reflect"
)

func pingHandler(w http.ResponseWriter, r *http.Request) { 	
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Pong!")
} // helloHandler

func greetHandler(w http.ResponseWriter, r*http.Request) { 
	if r.Method != "POST" { 
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return 
	} // if 

	err := r.ParseForm()
	if err != nil { 
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	} // if 

	file := r.FormValue("name")
	if err != nil { 
		http.Error(w, "Couldn't retrieve file", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello,", file)
} // greetHandler

// func multipartHandler(w http.ResponseWriter, r *http.Request) { 
// 	if r.Method != "POST" { 
// 		http.Error("Only POST allowed", http.StatusMethodNotAllowed)
// 	}

// 	mimeReader, err := r.MultipartReader()
	
// 	if err != nil { 
// 		http.Error("Not a multipart form", http.StatusBadRequest)
// 	} // if 

// 	var []byte maxFileSize = 10 << 20
// 	if mimeReader.ReadForm() > (10 << 20) { 
// 		http.Error("File size ")
// 	}
// 	// need to read contents of multipart form and check if it exceeds maximum memory size

// 	// ReadForm returns a pointer to a Form -> *Form
// 	// Form has 
// 	// Reader is an iterator over all parts in a MIME body 
// 		/* 
// 		trying to figure out how to use Reader to check file size (if possible), 
// 		and deny request if total file size (including metadata) exceeds 10 << 20 (10 MB)
// 		*/

// 	/* okay apparently the header variable for each part is available 
// 		and each individual part can be read using Read()

// 		each part needs to have its information read into a buffer
// 		using a channel



// 		NOTE: 
// 		it looks like the streaming approach using Read() will be more
// 		difficult as it will require reading parts of the file incrementally
// 	*/
// 	for { 
// 		currPart, err := mimeReader.NextPart()
		
// 		if err != nil { 
// 			fmt.Fprintln("No more parts to process")
// 		}

// 		inspectPartContents(currPart)
		
// 	} // for 

// } // multipartHandler


// /* Used for testing and inspecting contents of parts of a stream */ 
// func inspectPartContents(p *Part) { 
// 	partStruct := reflect.TypeOf(p)

// 	for i:=0; i < partStruct.NumField(); i++) { 
// 		currField = partStruct.Field(i)
// 		fmt.Println("Field Name:", field.Name, "Field Type", field.Type)
// 	} // for 


// }

func multipartHandler(w http.ResponseWriter, r *http.Request) { 
	if r.Method != "POST" { 
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
	} // if 

	// populate Multipart Form to retrieve file and file header 
	err := r.ParseMultipartForm(10 << 20)
	
	if err != nil { 
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	} // if 

	// If file form fields are supplied, iterate through map and process
	if len(r.MultipartForm.File) != 0 {
		if r.MultipartForm.File["files"] != nil { 
			for _, header := range r.MultipartForm.File["files"] {
				fileReader, err := header.Open(); if err != nil { fmt.Println(err)}
				contents, err := io.ReadAll(fileReader); if err != nil { fmt.Println(err)}
				_ = fileReader
				_ = contents 
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
	
	
} // multipartHandler

func main() { 
	http.HandleFunc("/ping", pingHandler)

	http.HandleFunc("/greet", greetHandler)

	http.HandleFunc("/upload", multipartHandler)

	http.ListenAndServe(":8080", nil)
	
} // main 