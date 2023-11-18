package twitterclient

func (c *Client) Verify() bool {
	return c.cli.IsReady()
}
