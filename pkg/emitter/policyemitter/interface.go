package policyemitter

type Interface interface {
	Buffer() error
	Scrape() error
}
