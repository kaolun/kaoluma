package client

import (
	"bufio"
	"fmt"
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
	line = strings.ToLower(line)
	parts := strings.Fields(line)
	c.Log(">" + line)
	switch parts[0] {
	case "help":
		c.Log("Command list:")
		//TODO add commands
		c.Log("bro what are you doing yk this works")
		c.Log("bublubbluhbluhbblubblbubb")
	case "clear":
		c.UI.Mu.Lock()
		c.UI.Logs = nil
		c.UI.Mu.Unlock()
		c.Render()
	case "exit", "close":
		fmt.Print("\033[H\033[2J")
		os.Exit(130)

	case "server":
		if len(parts) < 2 {
			c.Log("'help' to see usage")
		}
		switch parts[1] {
		case "connect":
			if !c.UI.Online {
				err := c.Connect()
				if err != nil {
					c.Log("ERROR:" + err.Error())
					return
				}
				c.Render()
			} else {
				c.Log("Already connected")
			}
		case "disconnect":
			if c.UI.Online {
				close(c.Done)
				err := c.Conn.Close()
				if err != nil {
					c.Log("ERROR:" + err.Error())
					return
				}
				c.UI.Online = false
				c.Log("Connection closed")
			} else {
				c.Log("Not connected")
			}
		case "code":
			if c.UI.Online {
				c.RequestJoinCode()
			} else {
				c.Log("Not connected")
			}
		default:
			c.Log("Unknown arg:" + parts[1])
		}
	case "peer":
		if len(parts) < 2 {
			c.Log("'help' to see usage")
			return
		}
		if !c.UI.Online {
			c.Log("Error: not connected")
			return
		}
		switch parts[1] {
		case "connect":
			if c.PeerID != 0 {
				c.Log("Error: Already connected to peer")
				return
			}
			if len(parts) < 3 {
				c.Log("Usage: Peer Connect Joincode")
				return
			}
			err := c.RequestPeerConnection(strings.ToUpper(parts[2]))
			if err != nil {
				c.Log("ERROR:" + err.Error())
			}
		case "disconnect":
			if c.PeerID == 0 {
				c.Log("Error: Not connected to peer")
			}
			err := c.RequestPeerDisconnect()
			if err != nil {
				c.Log("ERROR:" + err.Error())
			}

		default:
			c.Log("Unknown arg:" + parts[1])
		}

	default:
		c.Log("Unknown command:" + parts[0])
	}
}

func (c *Client) Log(msg string) {
	c.UI.Mu.Lock()
	c.UI.Logs = append(c.UI.Logs, msg)
	c.UI.Mu.Unlock()
	c.Render()
}
