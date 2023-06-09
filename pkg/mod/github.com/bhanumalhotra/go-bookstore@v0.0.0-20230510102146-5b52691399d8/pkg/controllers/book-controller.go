// Declare a package named "controllers"
package controllers

// Import necessary packages
import(
    "encoding/json" // package to encode and decode data in JSON format
    "fmt" // package to provide formatted I/O functions
    "github.com/gorilla/mux" // package to provide HTTP router and dispatcher
    "net/http" // package to provide HTTP client and server implementations
    "strconv" // package to provide functions to convert strings to numeric types
    "github.com/bhanumalhotra/go-bookstore/pkg/utils" // custom package that provides utility functions
    "github.com/bhanumalhotra/go-bookstore/pkg/models" // custom package that defines data models
)
// Declare a variable named "NewBook" of type "models.Book"
var NewBook models.Book

// Function to handle GET requests for the "/book" endpoint
func GetBook(w http.ResponseWriter, r *http.Request){
	// Call the "GetAllBook" function to get a list of all books from the database
	newBooks := models.GetAllBook()

	// Encode the list of books as JSON using the "json.Marshal" function
	// "json.Marshal" takes a Go data structure (in this case, a slice of books) and returns its JSON representation as a byte slice
	res, err := json.Marshal(newBooks)
	if err != nil {
		// If there is an error encoding the JSON, send a 500 Internal Server Error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the "Content-Type" header to "application/json" to indicate that the response is in JSON format
	// MIME type is used to identify the type of data being sent over the Internet
	w.Header().Set("Content-Type", "application/json")

	// Send a 200 OK response to the client
	w.WriteHeader(http.StatusOK)

	// Write the encoded JSON to the response writer to send it to the client
	// The "Write" function writes a byte slice to the response writer, sending it to the client
	w.Write(res)
}


// Function to handle GET requests for the "/book/{bookId}" endpoint
func GetBookById(w http.ResponseWriter, r *http.Request){
	// Extract the "bookId" parameter from the request URL using the Gorilla Mux library
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Parse the bookId string as an integer using the "strconv.ParseInt" function
	// The third argument (0) indicates that the base should be determined automatically
	// The first argument of strconv.ParseInt() is the string that we want to convert to an integer. 
	//The second argument is the base system of the input string. For example, if the input string is in binary format, 
	// we need to set the base to 2. Similarly, if the input string is in hexadecimal format, we need to set the base to 16.
	// The third argument of strconv.ParseInt() is the bit size of the resulting integer. This indicates how many bits are required to represent the integer. 
	// A value of 0 means that the function should determine the bit size automatically based on the input string.So, when we use strconv.ParseInt(bookId, 0, 0),
	// 00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000099 it means that we want to convert the bookId string to an integer and we want Go to determine the base system automatically and the bit size based on the input string.
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If there is an error parsing the bookId, log the error message to the console and return a 400 Bad Request response to the client
		fmt.Println("error while parsing bookId:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the "GetBookById" function to get the details of the book with the given ID from the database
	// The "&" character before "NewBook" is a pointer reference to the variable "NewBook"
	// The "GetBookById" function will modify the value of "NewBook" using this pointer reference
	bookDetails, err := models.GetBookById(&NewBook, ID)
	if err != nil {
		// If there is an error getting the book details, log the error message to the console and return a 500 Internal Server Error response to the client
		fmt.Println("error while getting book details:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encode the book details as JSON using the "json.Marshal" function
	// "json.Marshal" takes a Go data structure (in this case, a book) and returns its JSON representation as a byte slice
	res, err := json.Marshal(bookDetails)
	if err != nil {
		// If there is an error encoding the JSON, log the error message to the console and return a 500 Internal Server Error response to the client
		fmt.Println("error while encoding JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
		

	// Set the "Content-Type" header to "application/json" to indicate that the response is in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Send a 200 OK response to the client
	w.WriteHeader(http.StatusOK)

	// Write the encoded JSON to the response writer to send it to the client
	// The "Write" function writes a byte slice to the response writer, sending it to the client
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
    CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    // Get the bookId parameter from the request URL using mux.Vars
    vars := mux.Vars(r)
    bookId := vars["bookId"]

    // Parse the bookId parameter into an integer
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("Error while parsing bookId")
    }

    // Call the DeleteBook method on the models package with the parsed ID
    deletedBook := models.DeleteBook(ID)

    // Marshal the deletedBook into a JSON string
    res, _ := json.Marshal(deletedBook)

    // Set the response header to indicate that the response body is JSON
    w.Header().Set("Content-Type", "pkglication/json")

    // Write the JSON string to the response with a 200 OK status code
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Create a pointer to a Book struct to hold the updated book details
	var updateBook = &models.Book{}
	// Parse the request body into the updateBook struct
	utils.ParseBody(r, updateBook)
	// Get the book ID from the URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// Convert the book ID from string to integer
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If there is an error while parsing the book ID, print an error message
		fmt.Println("error while parsing")
	}
	// Get the book details from the database
	bookDetails, db := models.GetBookById(ID)
	// Update the book details if the corresponding fields in the updateBook struct are not empty
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	// Save the updated book details to the database
	db.Save(&bookDetails)
	// Convert the updated book details to JSON format
	res, _ := json.Marshal(bookDetails)
	// Set the content type of the response to JSON
	w.Header().Set("Content-Type", "pkglication/json")
	// Set the status code of the response to 200 OK
	w.WriteHeader(http.StatusOK)
	// Write the JSON data to the response body
	w.Write(res)
}
