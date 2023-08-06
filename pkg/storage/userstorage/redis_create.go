package userstorage

import (
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp *Object) (*Object, error) {
	var err error

	var out *Object
	{
		out, err = r.Search(inp.Subj[0], "")
		if IsNotFound(err) {
			// The user does not appear to exist. So first, create the mapping between
			// external subject claim and internal user ID.
			{
				inp.Crtd = time.Now().UTC()
				inp.User = scoreid.New(inp.Crtd)
			}

			{
				err = r.red.Simple().Create().Element(fmt.Sprintf(keyfmt.UserClaim, inp.Subj[0]), inp.User.String())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// Second, create the mapping between internal user ID and internal user
			// object.
			var jsn string
			{
				jsn = musStr(inp)
			}

			{
				err = r.red.Simple().Create().Element(fmt.Sprintf(keyfmt.UserObject, inp.User), jsn)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		} else if err != nil {
			return nil, tracer.Mask(err)
		} else if out.Imag != inp.Imag || out.Name != inp.Name {
			// The user exists and we update it due to changes in profile picture
			// and/or username.
			{
				out.Imag = inp.Imag
				out.Name = inp.Name
			}

			var jsn string
			{
				jsn = musStr(out)
			}

			{
				err = r.red.Simple().Create().Element(fmt.Sprintf(keyfmt.UserObject, out.User), jsn)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}
	}

	return out, nil
}
