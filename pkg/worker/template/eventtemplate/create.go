package eventtemplate

import (
	"bytes"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type kind string

const (
	KindDiscord  kind = "discord"
	KindTwitter  kind = "twitter"
	KindWarpcast kind = "warpcast"
)

var (
	newlin = regexp.MustCompile(`\n+`)
)

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

func (t *Template) Create(eid objectid.ID, kin kind) (string, error) {
	var err error

	var eob *eventstorage.Object
	{
		eob, err = t.searchEvnt(eid)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// It might happen that events get created and at the time the task runs the
	// event got already deleted. In such a case we just stop processing here.
	if eob == nil {
		return "", tracer.Maskf(cancelError, "missing event object")
	}

	var dob *descriptionstorage.Object
	{
		dob, err = t.searchDesc(eid)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// It might happen that events get created without descriptions. In such a
	// case we just stop processing here.
	if dob == nil {
		return "", tracer.Maskf(cancelError, "missing description object")
	}

	var lob labelstorage.Slicer
	{
		lob, err = t.searchLabl(eob)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var tem string
	{
		tem, err = ensureTmpl(dob, eob, lob, kin)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return tem, nil
}

func (t *Template) searchDesc(eid objectid.ID) (*descriptionstorage.Object, error) {
	var err error

	var dob []*descriptionstorage.Object
	{
		dob, err = t.des.SearchEvnt("", []objectid.ID{eid})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(dob) == 0 {
		return nil, nil
	}

	// Sort event descriptions by time in ascending order with second priority.
	// Note that we are interested in the earliest description if no single
	// description has the most likes. The earliest description is then most
	// likely the description the event creator provided during event creation.
	sort.SliceStable(dob, func(i, j int) bool {
		return dob[i].Crtd.Unix() < dob[j].Crtd.Unix()
	})

	// Sort event descriptions by likes in descending order with first priority.
	// Since we want to create an event post with the event description that has
	// the most likes, we have to search the list of descriptions here in reverse
	// order, ensuring the first description in the list ends up having the most
	// likes.
	sort.SliceStable(dob, func(i, j int) bool {
		return dob[i].Mtrc.Data[objectlabel.DescriptionMetricUser] > dob[j].Mtrc.Data[objectlabel.DescriptionMetricUser]
	})

	return dob[0], nil
}

func (t *Template) searchEvnt(eid objectid.ID) (*eventstorage.Object, error) {
	var err error

	var eob []*eventstorage.Object
	{
		eob, err = t.eve.SearchEvnt("", []objectid.ID{eid})
		if eventstorage.IsEventObjectNotFound(err) {
			return nil, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(eob) == 0 {
		return nil, nil
	}

	return eob[0], nil
}

func (t *Template) searchLabl(eob *eventstorage.Object) ([]*labelstorage.Object, error) {
	var err error

	var lid []objectid.ID
	{
		lid = append(lid, eob.Cate...)
		lid = append(lid, eob.Host...)
	}

	var lob labelstorage.Slicer
	{
		lob, err = t.lab.SearchLabl(lid)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return lob, nil
}

func camCas(lis []string) []string {
	var fix []string

	var cas cases.Caser
	{
		cas = cases.Title(language.English)
	}

	for _, x := range lis {
		var spl []string
		{
			spl = strings.Split(x, " ")
		}

		var prt []string
		if len(spl) == 1 {
			prt = spl
		} else {
			for _, y := range spl {
				prt = append(prt, cas.String(y))
			}
		}

		fix = append(fix, strings.Join(prt, ""))
	}

	return fix
}

func createData(dob *descriptionstorage.Object, eob *eventstorage.Object, lob labelstorage.Slicer, kin kind) templateData {
	var key string
	if kin == KindTwitter {
		key = objectlabel.ProfileTwitter
	}
	if kin == KindWarpcast {
		key = objectlabel.ProfileWarpcast
	}

	return templateData{
		Cate: "#" + strings.Join(camCas(lob.Cate().Name()), " #"),
		Desc: dob.Text.Data,
		Host: "@" + strings.Join(camCas(lob.Prfl(key)), " @"),
		Link: eob.Link,
		Time: eob.Time.Format("Mon 02 Jan, 15:04 MST"),
	}
}

func createTmpl(dat templateData) (string, error) {
	var err error

	var nam string
	{
		nam = "eventtemplate.Template/Create"
	}

	var tem *template.Template
	{
		tem, err = template.New(nam).Parse(templateBody)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var byt bytes.Buffer
	{
		err = tem.ExecuteTemplate(&byt, nam, dat)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return byt.String(), nil
}

func ensureTmpl(dob *descriptionstorage.Object, eob *eventstorage.Object, lob labelstorage.Slicer, kin kind) (string, error) {
	var err error

	var dat templateData
	{
		dat = createData(dob, eob, lob, kin)
	}

	var tem string
	for {
		tem, err = createTmpl(dat)
		if err != nil {
			return "", tracer.Mask(err)
		}

		if len(tem) < 280 {
			break
		}

		// In case the template is longer than 280 characters, then remove the last
		// word from the description, if any.
		if dat.Desc != "" {
			dat.Desc = trmRigt(dat.Desc, "...")
		}

		// In case the template is longer than 280 characters, then remove the last
		// category label, if any.
		if dat.Desc == "" {
			dat.Cate = trmRigt(dat.Cate, "")
		}

		// In case the template is longer than 280 characters, then remove the last
		// host label, if any.
		if dat.Cate == "" {
			dat.Host = trmRigt(dat.Host, "")
		}
	}

	// Replace multiple extraneous line breaks with a single one, once we pruned
	// all extraneous content.
	{
		tem = strings.Trim(newlin.ReplaceAllString(tem, "\n\n"), "\n")
	}

	return tem, nil
}

func trmRigt(des string, elp string) string {
	var spl []string
	{
		spl = strings.Split(des, " ")
	}

	if len(spl) <= 2 {
		return ""
	}

	{
		spl = spl[:len(spl)-2]
	}

	if elp != "" {
		spl = append(spl, elp)
	}

	return strings.Join(spl, " ")
}
