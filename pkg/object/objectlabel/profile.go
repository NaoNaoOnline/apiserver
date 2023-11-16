package objectlabel

const (
	ProfileTwitter  string = "Twitter"
	ProfileWarpcast string = "Warpcast"
)

func SearchPrfl() []string {
	return []string{
		ProfileTwitter,
		ProfileWarpcast,
	}
}
