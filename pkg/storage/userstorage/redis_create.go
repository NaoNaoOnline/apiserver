package userstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp *Object) (*Object, error) {
	var err error

	{
		err = inp.Verify()
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out *Object
	{
		out, err = r.SearchSubj(inp.Subj[0])
		if IsSubjectClaimMapping(err) {
			// The user does not appear to exist. So first, create the mapping between
			// external subject claim and internal user ID.
			{
				inp.Crtd = time.Now().UTC()
				inp.User = objectid.New(inp.Crtd)
			}

			{
				err = r.red.Simple().Create().Element(useCla(inp.Subj[0]), inp.User.String())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				err = r.red.Simple().Create().Element(useNam(inp.Name), inp.User.String())
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
				err = r.red.Simple().Create().Element(useObj(inp.User), jsn)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			return inp, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		} else if out.Imag != inp.Imag || out.Name != inp.Name {
			// The user exists and we update it due to changes in profile picture
			// and/or username. The user's mapping between name and ID is updated by
			// first creating the new reference and then deleting the old one.
			{
				err = r.red.Simple().Create().Element(useNam(inp.Name), out.User.String())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				_, err = r.red.Simple().Delete().Multi(useNam(out.Name))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				out.Imag = inp.Imag
				out.Name = inp.Name
			}

			{
				err = r.red.Simple().Create().Element(useObj(out.User), musStr(out))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}
	}

	return out, nil
}
