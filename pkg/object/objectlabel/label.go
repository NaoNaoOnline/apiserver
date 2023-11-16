package objectlabel

const (
	LabelDiscord  string = "Discord"
	LabelGoogle   string = "Google"
	LabelTwitter  string = "Twitter"
	LabelTwitch   string = "Twitch"
	LabelUnlonely string = "Unlonely"
	LabelYouTube  string = "YouTube"
	LabelZoom     string = "Zoom"
)

func SearchLabel() []string {
	return []string{
		LabelDiscord,
		LabelGoogle,
		LabelTwitter,
		LabelTwitch,
		LabelUnlonely,
		LabelYouTube,
		LabelZoom,
	}
}
