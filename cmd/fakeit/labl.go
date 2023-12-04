package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
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
			"AA",
			"Account Abstraction",
			"Bitcoin",
			"Crypto",
			"cybertruck",
			"DAOs",
			"Degens",
			"EAS",
			"ETF",
			"Ethereum",
			"Ethereum Virtual Machine",
			"EVM",
			"Internet Money",
			"Layer 2",
			"MEV",
			"Macro",
			"NFT",
			"PoS",
			"Payments",
			"Podcast",
			"Real World Assets",
			"Regulation",
			"Rollups",
			"RWA",
			"The Journey Man",
			"Trading",
			"Waifus",
			"VC",
			"Web3",
		})
	} else {
		nam = fak.RandomString([]string{
			"Aave",
			"Arbitrum",
			"Bankless",
			"Banteg",
			"bytes032",
			"Chainlink",
			"Cobie",
			"Cred",
			"David from Bankless",
			"DonAlt",
			"Flashbots",
			"Fofty",
			"FooBar",
			"HanSolar.eth",
			"hashverse",
			"Icebergy",
			"IporLabs",
			"Jason",
			"Kain",
			"Kieran.eth",
			"Ledger",
			"Michael Anderson",
			"Mippo",
			"Nic Carter",
			"Optimism",
			"Pentoshi",
			"Privy",
			"Racer",
			"RSA.eth",
			"Santi",
			"Sassal",
			"Senator Cynthia Lummis",
			"StarkNet",
			"Superfluid",
			"Synthetix",
			"Sisyphus",
			"Trader Koala",
			"Vance Spencer",
			"zkLink",
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
			Name: objectfield.String{
				Data: nam,
			},
			User: objectfield.ID{
				Data: use.User()[0],
			},
		}
	}

	return obj
}
