package twittercreatehandler

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
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

var (
	newlin = regexp.MustCompile(`\n+`)
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eid objectid.ID
	{
		eid = objectid.ID(tas.Meta.Get(objectlabel.EvntObject))
	}

	var eob *eventstorage.Object
	{
		eob, err = h.searchEvnt(eid, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var lob labelstorage.Slicer
	{
		lob, err = h.searchLabl(eob, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// It may very well be that none of the event's labels has any associate
	// profile for an external platform. In that case we are done, because we have
	// nobody to tweet at.
	if len(lob.Prfl(objectlabel.ProfileTwitter)) == 0 {
		return nil
	}

	var dob *descriptionstorage.Object
	{
		dob, err = h.searchDesc(eid, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var twt string
	{
		twt, err = ensureTmpl(dob, eob, lob)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.twi.Create(twt)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *SystemHandler) searchDesc(inp objectid.ID, bud *budget.Budget) (*descriptionstorage.Object, error) {
	var err error

	var des []*descriptionstorage.Object
	{
		des, err = h.des.SearchEvnt("", []objectid.ID{inp})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Sort event descriptions by time with second priority. Note that we are
	// interested in the earliest description if no single description has the
	// most likes. The earliest description is then most likely the description
	// the event creator provided during event creation.
	sort.SliceStable(des, func(i, j int) bool {
		return des[i].Crtd.Unix() < des[j].Crtd.Unix()
	})

	// Sort event descriptions by likes with first priority. Since we want to
	// create an event post with the event description that has the most likes, we
	// have to search the list of descriptions here in reverse order, ensuring the
	// first description in the list ends up having the most likes.
	sort.SliceStable(des, func(i, j int) bool {
		return des[i].Like.Data > des[j].Like.Data
	})

	return des[0], nil
}

func (h *SystemHandler) searchEvnt(inp objectid.ID, bud *budget.Budget) (*eventstorage.Object, error) {
	var err error

	var eve []*eventstorage.Object
	{
		eve, err = h.eve.SearchEvnt("", []objectid.ID{inp})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return eve[0], nil
}

func (h *SystemHandler) searchLabl(inp *eventstorage.Object, bud *budget.Budget) ([]*labelstorage.Object, error) {
	var err error

	var lid []objectid.ID
	{
		lid = append(lid, inp.Cate...)
		lid = append(lid, inp.Host...)
	}

	var lob labelstorage.Slicer
	{
		lob, err = h.lab.SearchLabl(lid)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return lob, nil
}

func camCas(lis []string) []string {
	var fix []string

	for _, x := range lis {
		fix = append(fix, strings.ReplaceAll(x, " ", ""))
	}

	return fix
}

func createData(dob *descriptionstorage.Object, eob *eventstorage.Object, lob labelstorage.Slicer) templateData {
	return templateData{
		Cate: "#" + strings.Join(camCas(lob.Cate().Name()), " #"),
		Desc: dob.Text.Data,
		Host: "@" + strings.Join(lob.Prfl(objectlabel.ProfileTwitter), " @"),
		Link: eob.Link,
		Time: eob.Time.Format("Mon 02 Jan, 15:04 MST"),
	}
}

func createTmpl(dat templateData) (string, error) {
	var err error

	var nam string
	{
		nam = "twittercreatehandler.Handler/Ensure"
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

func ensureTmpl(dob *descriptionstorage.Object, eob *eventstorage.Object, lob labelstorage.Slicer) (string, error) {
	var err error

	var dat templateData
	{
		dat = createData(dob, eob, lob)
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
