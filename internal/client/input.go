package client

import (
	"bufio"
	"kaoluma/internal/core/protocol"
	"os"
	"strings"
)

func (c *Client) InputLoop() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		c.handleCommand(line)
	}
}

func (c *Client) handleCommand(line string) {
	parts := strings.Fields(line)
	c.Log(">" + line)
	switch parts[0] {

	case "help":
		c.Log("Command list:")

	case "list":
		c.SendQueue <- protocol.Packet{Type: protocol.ListClientsRequest}

	default:
		c.Log("unknown command:" + parts[0])
	}
}
