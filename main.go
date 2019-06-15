package main

import (
	"github.com/Dev-ManavSethi/blogzzz/controllers"
	"net/http"
)

func main() {

	multiplexer := http.NewServeMux()

	//home
	//multiplexer.HandleFunc("/", home)

	//session related functions
	multiplexer.HandleFunc("/login", controllers.login)
	// multiplexer.HandleFunc("/signup", signup)
	// multiplexer.HandleFunc("/logout", logout)

	// //user related functions
	// multiplexer.HandleFunc("/user", userProfile)
	// multiplexer.HandleFunc("/editUser", editUser)
	// multiplexer.HandleFunc("/deleteUser", deleteUser)

	// //blog related fucntions
	// multiplexer.HandleFunc("/createBlog", createBlog)
	// multiplexer.HandleFunc("/editBlog", editBlog)
	// multiplexer.HandleFunc("/deleteBlog", deleteBlog)

}
