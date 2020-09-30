package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// This binds the method ServeHTTP with userController class, this is how we create functions on an object
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Conversion into a Slice data type
	w.Write([]byte("Hello from the User Controller"))
}

// This is constructor function. There is no special meaning of this function, it's based on convention in Go.
// All the constructor functions are named as new____ to designate it as constructor function.
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}