package userstorage

import (
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
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

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var out *Object
	{
		out, err = r.SearchSubj(inp.Subj[0])
		if IsSubjectClaimMapping(err) {
			// The user does not appear to exist. So we initialize the user object.
			{
				inp.Crtd = now
				inp.User = objectid.Random(objectid.Time(now))
				inp.Name.Time = now
			}

			// Here we create the mapping between external subject claim and internal
			// user ID.
			{
				err = r.red.Simple().Create().Element(useCla(inp.Subj[0]), inp.User.String())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// The onboarding process is rather automatic and might be integrated with
			// OAuth flows. It may happen that a user joins with a name that is
			// already taken, without having the chance to correct such an
			// unintentional mistake. In that case we append the user ID to the user
			// name to simply have a unique name for mapping. Later the user may chose
			// to update their name to something that is available and more suitable
			// for them.
			{
				exi, err := r.red.Simple().Exists().Multi(useNam(keyfmt.Indx(inp.Name.Data)))
				if err != nil {
					return nil, tracer.Mask(err)
				}

				if exi == 1 {
					inp.Name.Data = fmt.Sprintf("%s-%s", inp.Name.Data, inp.User.String())
				}
			}

			// Since we want to be able to search for users by their name, we have to
			// create a separate mapping that allows us to go from user name to user
			// ID.
			{
				err = r.red.Simple().Create().Element(useNam(keyfmt.Indx(inp.Name.Data)), inp.User.String())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// Now create the mapping between internal user ID and internal user
			// object.
			{
				err = r.red.Simple().Create().Element(useObj(inp.User), musStr(inp))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			return inp, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		} else if out.Imag != inp.Imag || out.Name.Data != inp.Name.Data {
			// The user exists and we update it due to changes in profile picture
			// and/or username.

			// The onboarding process is rather automatic and might be integrated with
			// OAuth flows. It may happen that a user joins with a name that is
			// already taken without having the chance to correct such an
			// unintentional mistake. In that case we append the user ID to the user
			// name to simply have a unique name for mapping. Later the user may chose
			// to update their name to something that is available and more suitable
			// to them.
			{
				exi, err := r.red.Simple().Exists().Multi(useNam(keyfmt.Indx(inp.Name.Data)))
				if err != nil {
					return nil, tracer.Mask(err)
				}

				if exi == 1 {
					inp.Name.Data = fmt.Sprintf("%s-%s", inp.Name.Data, inp.User.String())
				}
			}

			// The user's mapping between name and ID is updated by first creating the
			// new reference and then deleting the old one.
			{
				err = r.red.Simple().Create().Element(useNam(keyfmt.Indx(inp.Name.Data)), out.User.String())
				if err != nil {
					return nil, tracer.Mask(err)
				}

				_, err = r.red.Simple().Delete().Multi(useNam(keyfmt.Indx(out.Name.Data)))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				out.Imag = inp.Imag
				out.Name.Data = inp.Name.Data
				inp.Name.Time = now
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
