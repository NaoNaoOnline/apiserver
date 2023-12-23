package discordclient

type Interface interface {
	Create(string) error
	Verify() bool
}
