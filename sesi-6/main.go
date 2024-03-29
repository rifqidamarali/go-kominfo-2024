package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	// fmt.Println("hi")
	NetHttp()
}

func NetHttp(){
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request)  {
		// get all users
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users)
			return
		}	
		// post user
		if r.Method == http.MethodPost{
			user := User{}
			// only bind username and email
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			user.ID = uint(len(users) + 1)
			users = append(users, user)

			w.WriteHeader(http.StatusAccepted)
			return
		}
	})

	http.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pathValue, _ := strconv.Atoi(r.PathValue("id"))
			for _, user := range users {
				if user.ID == uint(pathValue) {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(user)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		}

		if r.Method == http.MethodDelete {
			// var newUsers []User
			pathValue, _ := strconv.Atoi(r.PathValue("id"))
			for i, user := range users {
				if user.ID == uint(pathValue){
					users = append(users[:i], users[i+1:]...)
                	w.WriteHeader(http.StatusNoContent)
                	return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		}

		if r.Method == http.MethodPut {
			var updatedUser User
        err := json.NewDecoder(r.Body).Decode(&updatedUser)
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
		pathValue, _ := strconv.Atoi(r.PathValue("id"))
        for i, user := range users {
            if user.ID == uint(pathValue) {
                // Update user details
                users[i] = updatedUser
				users[i].ID = uint(pathValue) 
                w.WriteHeader(http.StatusOK)
                json.NewEncoder(w).Encode(users[i])
                return
            }
        }
        // If user not found, return 404 Not Found
        w.WriteHeader(http.StatusNotFound)



		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}