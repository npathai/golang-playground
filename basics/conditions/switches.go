package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{
		Method: "GET",
	}

	switch r.Method {
	case "GET":
		println("Get request")
	case "POST":
		println("Post request")
	case "PATCH":
		println("Patch request")
	case "DELETE":
		println("Delete request")
	}

	// With fallthrough
	switch r.Method {
	case "GET":
		println("Get request")
		fallthrough
	case "POST":
		println("Post request")
	case "PATCH":
		println("Patch request")
	case "DELETE":
		println("Delete request")
	}

	// With default case
	switch r.Method {
	case "POST":
		println("Post request")
	case "PATCH":
		println("Patch request")
	case "DELETE":
		println("Delete request")
	default:
		println("Get request")
	}
}
