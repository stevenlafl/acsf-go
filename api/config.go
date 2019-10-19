// @todo don't hardcode this
package api

const (
	USE_PROXY  = false
	PROXY_ADDR = "127.0.0.1:10000"
	API_USER   = ""
	API_PASS   = ""
)

func GetAPIUrls() map[string]string {
	var APIUrls map[string]string = make(map[string]string)
	APIUrls["dev"] = "https://www.dev-example.acsitefactory.com"
	APIUrls["test"] = "https://www.test-example.acsitefactory.com"
	APIUrls["live"] = "https://www.example.acsitefactory.com"

	return APIUrls
}
