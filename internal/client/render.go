package client

import "fmt"

func (c *Client) Render() {
	fmt.Print("\033[H\033[2J") //clears terminal(idk)

	fmt.Println("=================")
	fmt.Println("|    Kaoluma    |")
	fmt.Println("=================")
	fmt.Println()

	for _, log := range c.UI.Logs {
		fmt.Println(log)
	}

	online := "Disconnected"
	if c.UI.Online {
		online = "Connected"
	}
	visible := "Hidden"
	if c.UI.Visible {
		visible = "Visible"
	}

	fmt.Println("_________________________")
	fmt.Println(online, " | ", visible)
	fmt.Print(">>>")

}

func (c *Client) Log(msg string) {
	c.UI.Logs = append(c.UI.Logs, msg)
	c.Render()
}
