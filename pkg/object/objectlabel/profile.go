package objectlabel

const (
	ProfileTwitch   string = "Twitch"
	ProfileTwitter  string = "Twitter"
	ProfileWarpcast string = "Warpcast"
)

func SearchPrfl() []string {
	return []string{
		ProfileTwitch,
		ProfileTwitter,
		ProfileWarpcast,
	}
}
