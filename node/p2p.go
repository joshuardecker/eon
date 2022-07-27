package node

import (
	"bytes"
	"fmt"
	"net"
	"net/http"

	"github.com/TwiN/go-color"
)

// This function adds a node to your copy of peers.
// If they are unresponsive, they will not be added.
// Input is the ip of the node being added.
// Output is a bool. Returns true if the peer was added, false if not.
func (s *Server) AddNode(nodeIp string) bool {

	resp, httpErr := http.Get(nodeIp + "/status")

	if httpErr != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Error: "+httpErr.Error()))
		return false
	}

	if resp.StatusCode != http.StatusOK {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Node became unresponsive. No longer adding them to peers"))
		return false
	}

	s.Peers = append(s.Peers, nodeIp)

	return true
}

// This function removes a peer from the list.
// Input is the ip of the node.
// Returns true if they were removed, false if they were not on the list of peers.
func (s *Server) RemoveNode(nodeIp string) bool {

	// Search for the node
	for index := 0; index < len(s.Peers); index += 1 {

		// Removes the node
		if s.Peers[index] == nodeIp {

			s.Peers = append(s.Peers[:index], s.Peers[index+1:]...)
			return true
		}
	}

	return false
}

// This function senda data to a node of choice.
// Inputs are the nodes ip, the path the func will transmit to, ex "/tx", and the data.
// Returns a bool, true if successful, false if not successful.
func (s *Server) SendData(nodeIp string, path string, data *bytes.Buffer) bool {

	resp, httpErr := http.Post(nodeIp+path, "data/json", data)

	if httpErr != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Error: "+httpErr.Error()))

		if s.RemoveNode(nodeIp) {
			// If the node was known/saved

			fmt.Println(color.Colorize(color.Red, "[NODE]: Non-responsive node removed"))

		} else {
			// If the node was not known/saved

			fmt.Println(color.Colorize(color.Red, "[NODE]: Tried to contact a non-recognized peer"))
		}

		return false
	}

	// If the http response given was good
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {

		fmt.Println(color.Colorize(color.Green, "[NODE]: Successfully talked with peer"))
		return true
	}

	fmt.Println(color.Colorize(color.Red, "[NODE]: Potensial Error: Peer responded with a negative http status code"))
	return false
}

// This function sends data to all known peers.
// Inputs are the path being transmitted to, ex "/tx", and the data being sent.
// Returns nothing.
func (s *Server) SendDataToAll(path string, data *bytes.Buffer) {

	// Loops through all of the known peers.
	for index := 0; index < len(s.Peers); index += 1 {

		resp, httpErr := http.Post(s.Peers[index]+path, "data/json", data)

		if httpErr != nil {

			fmt.Println(color.Colorize(color.Red, "[NODE]: Error: "+httpErr.Error()))

			if s.RemoveNode(s.Peers[index]) {
				// If the node was known/saved

				fmt.Println(color.Colorize(color.Red, "[NODE]: Non-responsive node removed"))

			} else {
				// If the node was not known/saved

				fmt.Println(color.Colorize(color.Red, "[NODE]: Tried to contact a non-recognized peer"))
			}
		}

		// If the http response given was good
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {

			fmt.Println(color.Colorize(color.Green, "[NODE]: Successfully talked with peer"))

		} else {

			fmt.Println(color.Colorize(color.Red, "[NODE]: Potensial Error: Peer responded with a negative http status code"))
		}
	}
}

// This function gets the outbound ip of the machine running this software.
// This will be used as the ip the node is known for.
// Returns the string ip of this node.
func GetOutboundIP() string {

	connection, conErr := net.Dial("udp", "8.8.8.8:80")

	if conErr != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Error: "+conErr.Error()))
	}

	// Close the connection when the function ends
	defer connection.Close()

	localIp := connection.LocalAddr().(*net.UDPAddr)

	return localIp.IP.String()
}
