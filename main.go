package main
import ( //import required libraries 
    "encoding/json"
    "io"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

type Post struct {
    ID string `json:"id"`
  }

var posts []Post


// method to get all records from the API
func getPosts(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Getting posts... \n")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

//method to delete a record from the API
func deletePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range posts {
    if item.ID == params["id"] {
      posts = append(posts[:index], posts[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(posts)
}


func createPost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range posts {
      if item.ID == params["id"] {
        posts = append(posts[:index])
        break
      }
    }
    json.NewEncoder(w).Encode(&posts)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range posts {
    if item.ID == params["id"] {
      posts = append(posts[:index], posts[index+1:]...)
      var post Post
      _ = json.NewDecoder(r.Body).Decode(&post)
      post.ID = params["id"]
      posts = append(posts, post)
      json.NewEncoder(w).Encode(&post)
      return
    }
  }
  json.NewEncoder(w).Encode(posts)
}


// Main program
func main() {
  //The port our server will be listening on
  port := 8000 

  //Gorilla mux router
  r := mux.NewRouter()

  posts = append(posts, Post{ID: "1"})
  posts = append(posts, Post{ID: "2"})
  posts = append(posts, Post{ID: "3"})

  r.HandleFunc("/api/decrypt/", getPosts).Methods("GET")
  r.HandleFunc("/api/decrypt/{id}", createPost).Methods("POST")
  r.HandleFunc("/api/decrypt/{id}", updatePost).Methods("PUT")
  r.HandleFunc("/api/decrypt/{id}", deletePost).Methods("DELETE")

  log.Printf("GO REST server running on http://localhost:%d/", port)
  err := http.ListenAndServe(":" + strconv.Itoa(port), r)
  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }

}
