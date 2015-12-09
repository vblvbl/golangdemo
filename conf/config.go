package config

//Set the configuration values on this file

// Get gets you the value for the key
func Get(t string) string {
	switch t {
	case "WEBSERVER_PORT":
		return "8899"
	case "":
		return ""
	case "SESSION_NAME":
		return "master-session"
	default:
		return ""
	}
}
