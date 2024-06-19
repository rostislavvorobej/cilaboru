	type Operation struct {
		Name       string
		Method     string
		Path       string
		Handler    http.HandlerFunc
		Middleware []Middleware
	}  
