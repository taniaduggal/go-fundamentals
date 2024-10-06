package auth

// if i want to keep the logic private and yet want toexport the function, i can dio like this

func extractSession() string {
	return "loggedIn"
}

func GetSession() string {
	return extractSession()
}
