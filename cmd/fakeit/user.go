package fakeit

import (
	"fmt"
	"net/http"

	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createUser(sto *storage.Storage, obj *userstorage.Object) error {
	{
		_, err := sto.User().Create(obj)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return nil
}

func (r *run) randomUser(sto *storage.Storage, fak *gofakeit.Faker) *userstorage.Object {
	var err error

	// We want to produce a fake profile picture for our fake user. Here we fetch
	// a random image URL from some image provider. For that to work we use a
	// custom HTTP client that is not following redirects, because the random
	// image URL we are looking for will be set in the location header of the HTTP
	// response below.
	var cli *http.Client
	{
		cli = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	var res *http.Response
	{
		res, err = cli.Get(fak.ImageURL(24, 24))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var obj *userstorage.Object
	{
		obj = &userstorage.Object{
			Imag: res.Header.Get("location"),
			Name: fak.Username(),
			Subj: []string{
				fmt.Sprintf("google-oauth2|%06d", fak.Number(0, 99999)),
			},
		}
	}

	return obj
}
