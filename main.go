package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

// host represents data about a record host.
type host struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Architecture string `json:"architecture"`
	OS           string `json:"os"`
	Class        string `json:"class"`
	IPv4         string `json:"ipv4"`
}

// hosts slice to seed record album data.
var hosts = map[string]host{
	"a100012ffffffffffffffffffffff012": {
		ID:           "a100012ffffffffffffffffffffff012",
		Name:         "m900",
		Architecture: "AMD64",
		OS:           "Ubuntu22",
		Class:        "Desktop",
		IPv4:         "10.10.100.12",
	},
}

func main() {
	router := gin.Default()
	router.GET("/hosts", getHosts)
	router.GET("/hosts/:id", getHostByID)
	router.POST("/hosts", registerHost)

	router.Run("localhost:8080")
}

// getHosts responds with the list of all albums as JSON.
func getHosts(c *gin.Context) {
	// keys := make([]string, 0, len(hosts))
	// values := make([]host, 0, len(hosts))

	// for k, v := range hosts {
	// 	keys = append(keys, k)
	// 	values = append(values, v)
	// }
	// c.IndentedJSON(http.StatusOK, values)
	vals := maps.Values(hosts)
	c.IndentedJSON(http.StatusOK, vals)
}

// registerHost adds an album from JSON received in the request body.
func registerHost(c *gin.Context) {
	var newHost host

	// Call BindJSON to bind the received JSON to
	// newHost.
	if err := c.BindJSON(&newHost); err != nil {
		return
	}

	// Add the new host to the map.
	hosts[newHost.ID] = newHost
	c.IndentedJSON(http.StatusCreated, newHost)
}

// getHostByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getHostByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	// for _, a := range hosts {
	if val, ok := hosts[id]; ok {
		c.IndentedJSON(http.StatusOK, val)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// curl http://localhost:8080/hosts \
//     --include \
//     --header "Content-Type: application/json" \
//     --request "POST" \
//     --data '{"id": "a100011ffffffffffffffffffffff011", "name":"qotom", "architecture":"amd64", "os":"ubuntu22", "class":"evio-router","ipv4":"10.10.100.11"}'
