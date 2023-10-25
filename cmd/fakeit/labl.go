package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createLabl(sto *storage.Storage, obj ...*labelstorage.Object) error {
	{
		_, err := sto.Labl().Create(obj)
		if labelstorage.IsLabelObjectAlreadyExists(err) {
			// fall through
		} else if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return nil
}

func (r *run) randomLabl(sto *storage.Storage, fak *gofakeit.Faker) *labelstorage.Object {
	var err error

	var kin string
	{
		kin = fak.RandomString([]string{
			"cate",
			"host",
		})
	}

	var nam string
	if kin == "cate" {
		nam = fak.RandomString([]string{
			"Bitcoin",
			"Crypto",
			"DAOs",
			"EAS",
			"ETF",
			"Ethereum",
			"Internet Money",
			"Layer 2",
			"MEV",
			"Macro",
			"NFT",
			"PoS",
			"Payments",
			"Podcast",
			"Regulation",
			"Rollups",
			"Trading",
			"Waifus",
			"Web3",
		})
	} else {
		nam = fak.RandomString([]string{
			"Aave",
			"Arbitrum",
			"Bankless",
			"Banteg",
			"Chainlink",
			"Cobie",
			"Cred",
			"DonAlt",
			"Flashbots",
			"FooBar",
			"Icebergy",
			"IporLabs",
			"Jason",
			"Kain",
			"Ledger",
			"Mike",
			"Nic Carter",
			"Optimism",
			"Racer",
			"Sassal",
			"StarkNet",
			"Superfluid",
			"Synthetix",
			"Vance",
		})
	}

	var use userstorage.Slicer
	{
		use, err = sto.User().SearchFake()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(use)
	}

	var obj *labelstorage.Object
	{
		obj = &labelstorage.Object{
			Kind: kin,
			Name: nam,
			User: use.IDs()[0],
		}
	}

	return obj
}
