package twitterclient

func (c *Client) Verify() bool {
	return c.cli != nil && c.cli.IsReady()
}
