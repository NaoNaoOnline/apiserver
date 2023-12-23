package discordclient

func (c *Client) Verify() bool {
	return c.tkn != ""
}
