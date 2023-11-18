package twitterclient

type Interface interface {
	Create(string) error
	Verify() bool
}
