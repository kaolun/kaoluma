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

	case "connect":
		if !c.UI.Online {
			err := c.Connect()
			if err != nil {
				c.Log("ERROR:" + err.Error())
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
			}
			c.UI.Online = false
			c.Log("Connection closed")
		} else {
			c.Log("Not connected")
		}

	default:
		c.Log("unknown command:" + parts[0])
	}
}

func (c *Client) Log(msg string) {
	c.UI.Mu.Lock()
	c.UI.Logs = append(c.UI.Logs, msg)
	c.UI.Mu.Unlock()
	c.Render()
}
