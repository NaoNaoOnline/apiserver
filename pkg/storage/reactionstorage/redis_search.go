package reactionstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchBltn() []*Object {
	return []*Object{

		//
		// 1. row facial expressions
		//

		{
			Html: "️😍", // https://emojipedia.org/smiling-face-with-heart-eyes
			Kind: "bltn",
			Name: "Smiling Face With Heart Eyes",
			User: objectid.System(),
		},
		{
			Html: "😂", // https://emojipedia.org/face-with-tears-of-joy
			Kind: "bltn",
			Name: "Face With Tears Of Joy",
			User: objectid.System(),
		},
		{
			Html: "😲", // https://emojipedia.org/astonished-face
			Kind: "bltn",
			Name: "Astonished Face",
			User: objectid.System(),
		},
		{
			Html: "🥳", // https://emojipedia.org/partying-face
			Kind: "bltn",
			Name: "Partying Face",
			User: objectid.System(),
		},
		{
			Html: "😎", // https://emojipedia.org/smiling-face-with-sunglasses
			Kind: "bltn",
			Name: "Smiling Face With Sunglasses",
			User: objectid.System(),
		},
		{
			Html: "🫡", // https://emojipedia.org/saluting-face
			Kind: "bltn",
			Name: "Saluting Face",
			User: objectid.System(),
		},

		//
		// 2. row hand gestures
		//

		{
			Html: "👍", // https://emojipedia.org/thumbs-up
			Kind: "bltn",
			Name: "Thumbs Up",
			User: objectid.System(),
		},
		{
			Html: "💪", // https://emojipedia.org/flexed-biceps
			Kind: "bltn",
			Name: "Flexed Biceps",
			User: objectid.System(),
		},
		{
			Html: "👏", // https://emojipedia.org/clapping-hands
			Kind: "bltn",
			Name: "Clapping Hands",
			User: objectid.System(),
		},
		{
			Html: "✊", // https://emojipedia.org/raised-fist
			Kind: "bltn",
			Name: "Raised Fist",
			User: objectid.System(),
		},
		{
			Html: "🤝", // https://emojipedia.org/handshake
			Kind: "bltn",
			Name: "Handshake",
			User: objectid.System(),
		},
		{
			Html: "🙏", // https://emojipedia.org/folded-hands
			Kind: "bltn",
			Name: "Folded Hands",
			User: objectid.System(),
		},

		//
		// 3. row animate objects
		//

		{
			Html: "🚀", // https://emojipedia.org/rocket
			Kind: "bltn",
			Name: "Rocket",
			User: objectid.System(),
		},
		{
			Html: "💡", // https://emojipedia.org/light-bulb
			Kind: "bltn",
			Name: "Light Bulb",
			User: objectid.System(),
		},
		{
			Html: "👑", // https://emojipedia.org/crown
			Kind: "bltn",
			Name: "Crown",
			User: objectid.System(),
		},
		{
			Html: "⭐", // https://emojipedia.org/star
			Kind: "bltn",
			Name: "Star",
			User: objectid.System(),
		},
		{
			Html: "🦄", // https://emojipedia.org/unicorn
			Kind: "bltn",
			Name: "Unicorn",
			User: objectid.System(),
		},
		{
			Html: "🤖", // https://emojipedia.org/robot
			Kind: "bltn",
			Name: "Robot",
			User: objectid.System(),
		},

		//
		// 4. row various expressions
		//

		{
			Html: "💦", // https://emojipedia.org/sweat-droplets
			Kind: "bltn",
			Name: "Sweat Droplets",
			User: objectid.System(),
		},
		{
			Html: "🔥", // https://emojipedia.org/fire
			Kind: "bltn",
			Name: "Fire",
			User: objectid.System(),
		},
		{
			Html: "👀", // https://emojipedia.org/eyes
			Kind: "bltn",
			Name: "Eyes",
			User: objectid.System(),
		},
		{
			Html: "✅", // https://emojipedia.org/check-mark-button
			Kind: "bltn",
			Name: "Check Mark Button",
			User: objectid.System(),
		},
		{
			Html: "❗", // https://emojipedia.org/exclamation-mark
			Kind: "bltn",
			Name: "Exclamation Mark",
			User: objectid.System(),
		},
		{
			Html: "💯", // https://emojipedia.org/hundred-points
			Kind: "bltn",
			Name: "Hundred Points",
			User: objectid.System(),
		},
	}
}

func (r *Redis) SearchKind(inp []string) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range inp {
		if x != "bltn" && x != "user" {
			return nil, tracer.Mask(reactionKindInvalidError)
		}

		// val will result in a list of all reaction IDs grouped under the given
		// reaction kind, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(rctKin(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next reaction kind, if any.
		if len(val) == 0 {
			continue
		}

		var jsn []string
		{
			jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(val, keyfmt.ReactionObject)...)
			if simple.IsNotFound(err) {
				return nil, tracer.Maskf(reactionObjectNotFoundError, "%v", val)
			} else if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		for _, x := range jsn {
			var obj *Object
			{
				obj = &Object{}
			}

			if x != "" {
				err = json.Unmarshal([]byte(x), obj)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			out = append(out, obj)
		}
	}

	return out, nil
}
