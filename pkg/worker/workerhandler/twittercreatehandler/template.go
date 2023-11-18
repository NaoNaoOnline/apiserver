package twittercreatehandler

type templateData struct {
	Cate string
	Desc string
	Host string
	Link string
	Time string
}

const templateBody = `Online event added to NaoNao. {{ .Host }} welcomes you to chat about {{ .Cate }}!

{{ .Desc }}

Join {{ .Time }}.

{{ .Link }}
`
