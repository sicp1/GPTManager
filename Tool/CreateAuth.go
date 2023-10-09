package Tool

func CreateAuth(api_key string) string {
	return "Bearer " + api_key
}
