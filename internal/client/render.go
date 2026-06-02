package client

import "fmt"

func (c *Client) Render() {
	c.UI.Mu.RLock()
	defer c.UI.Mu.RUnlock()
	fmt.Print("\033[H\033[2J") //clears terminal(idk)

	fmt.Println("=================")
	fmt.Println("|    kaoluma    |")
	fmt.Println("=================")

	for _, log := range c.UI.Logs {
		fmt.Println(log)
	}

	if !c.UI.Online {
		fmt.Println("________________")
		fmt.Println("| Disconnected |")
	} else {
		code := c.JoinCode
		fmt.Println("________________________")
		fmt.Println("| Connected | ", code, " |")
	}

	fmt.Print(">>>")

}
