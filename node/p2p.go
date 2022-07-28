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
// Input is the ip of the node being added, and whether this is on the mainnet. Set true if this node is on the mainnet.
// Output is a bool. Returns true if the peer was added, false if not.
func (n *Node) AddNode(nodeIp string, mainnet bool) bool {

	// Split the ip and port sent from the http get response
	host, _, splitErr := net.SplitHostPort(nodeIp)

	if splitErr != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Node port could not be split. Error:"+splitErr.Error()))
		return false
	}

	if mainnet {
		// If node is on the mainnet

		nodeIp = net.JoinHostPort(host, "8181")
		fmt.Println(color.Colorize(color.Green, "[NODE]: Node added to the mainnet"))

	} else {
		// If the node is on the testnet

		nodeIp = net.JoinHostPort(host, "8180")
		fmt.Println(color.Colorize(color.Green, "[NODE]: Node added to the testnet"))
	}

	resp, httpErr := http.Get(nodeIp + "/status")

	if httpErr != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Non-Node attempted to join the network. Error: "+httpErr.Error()))
		return false
	}

	if resp.StatusCode != http.StatusOK {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Node became unresponsive. No longer adding them to peers"))
		return false
	}

	n.Peers = append(n.Peers, nodeIp)

	return true
}

// This function removes a peer from the list.
// Input is the ip of the node.
// Returns true if they were removed, false if they were not on the list of peers.
func (n *Node) RemoveNode(nodeIp string) bool {

	// Search for the node
	for index := 0; index < len(n.Peers); index += 1 {

		// Removes the node
		if n.Peers[index] == nodeIp {

			n.Peers = append(n.Peers[:index], n.Peers[index+1:]...)
			return true
		}
	}

	return false
}

// This function senda data to a node of choice.
// Inputs are the nodes ip, the path the func will transmit to, ex "/tx", and the data.
// Returns a bool, true if successful, false if not successful.
func (n *Node) SendData(nodeIp string, path string, data *bytes.Buffer) bool {

	resp, httpErr := http.Post(nodeIp+path, "data/json", data)

	if httpErr != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Error: "+httpErr.Error()))

		if n.RemoveNode(nodeIp) {
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
func (n *Node) SendDataToAll(path string, data *bytes.Buffer) {

	// Loops through all of the known peers.
	for index := 0; index < len(n.Peers); index += 1 {

		resp, httpErr := http.Post(n.Peers[index]+path, "data/json", data)

		if httpErr != nil {

			fmt.Println(color.Colorize(color.Red, "[NODE]: Error: "+httpErr.Error()))

			if n.RemoveNode(n.Peers[index]) {
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
