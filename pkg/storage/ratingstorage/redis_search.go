package ratingstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
)

func (r *Redis) Search() ([]*Object, error) {
	return []*Object{

		//
		// 1. row facial expressions
		//

		{
			Crtd: time.Unix(1692392942, 0),
			Html: "️😍", // https://emojipedia.org/smiling-face-with-heart-eyes
			Name: "smiling-face-with-heart-eyes",
			Rtng: scoreid.String("1692392942673667"),
		},
		{
			Crtd: time.Unix(1692393021, 0),
			Html: "😂", // https://emojipedia.org/face-with-tears-of-joy
			Name: "face-with-tears-of-joy",
			Rtng: scoreid.String("1692393021407686"),
		},
		{
			Crtd: time.Unix(1692394796, 0),
			Html: "😲", // https://emojipedia.org/astonished-face
			Name: "astonished-face",
			Rtng: scoreid.String("1692394796326052"),
		},
		{
			Crtd: time.Unix(1692393087, 0),
			Html: "🥳", // https://emojipedia.org/partying-face
			Name: "partying-face",
			Rtng: scoreid.String("1692393087605581"),
		},
		{
			Crtd: time.Unix(1692393078, 0),
			Html: "😎", // https://emojipedia.org/smiling-face-with-sunglasses
			Name: "smiling-face-with-sunglasses",
			Rtng: scoreid.String("1692393078554976"),
		},
		{
			Crtd: time.Unix(1692393028, 0),
			Html: "🫡", // https://emojipedia.org/saluting-face
			Name: "saluting-face",
			Rtng: scoreid.String("1692393028348327"),
		},

		//
		// 2. row hand gestures
		//

		{
			Crtd: time.Unix(1692393035, 0),
			Html: "👍", // https://emojipedia.org/thumbs-up
			Name: "thumbs-up",
			Rtng: scoreid.String("1692393035303485"),
		},
		{
			Crtd: time.Unix(1692393047, 0),
			Html: "💪", // https://emojipedia.org/flexed-biceps
			Name: "flexed-biceps",
			Rtng: scoreid.String("1692393047919758"),
		},
		{
			Crtd: time.Unix(1692393053, 0),
			Html: "👏", // https://emojipedia.org/clapping-hands
			Name: "clapping-hands",
			Rtng: scoreid.String("1692393053200333"),
		},
		{
			Crtd: time.Unix(1692393068, 0),
			Html: "✊", // https://emojipedia.org/raised-fist
			Name: "raised-fist",
			Rtng: scoreid.String("1692393068586868"),
		},
		{
			Crtd: time.Unix(1692393073, 0),
			Html: "🤝", // https://emojipedia.org/handshake
			Name: "handshake",
			Rtng: scoreid.String("1692393073988751"),
		},
		{
			Crtd: time.Unix(1692394815, 0),
			Html: "🙏", // https://emojipedia.org/folded-hands
			Name: "folded-hands",
			Rtng: scoreid.String("1692394815339622"),
		},

		//
		// 3. row animate objects
		//

		{
			Crtd: time.Unix(1692392933, 0),
			Html: "🚀", // https://emojipedia.org/rocket
			Name: "rocket",
			Rtng: scoreid.String("1692392933890022"),
		},
		{
			Crtd: time.Unix(1692392959, 0),
			Html: "💡", // https://emojipedia.org/light-bulb
			Name: "light-bulb",
			Rtng: scoreid.String("1692392959842025"),
		},
		{
			Crtd: time.Unix(1692393041, 0),
			Html: "👑", // https://emojipedia.org/crown
			Name: "crown",
			Rtng: scoreid.String("1692393041522806"),
		},
		{
			Crtd: time.Unix(1692392978, 0),
			Html: "⭐", // https://emojipedia.org/star
			Name: "star",
			Rtng: scoreid.String("1692392978"),
		},
		{
			Crtd: time.Unix(1692392985, 0),
			Html: "🦄", // https://emojipedia.org/unicorn
			Name: "unicorn",
			Rtng: scoreid.String("1692392985448935"),
		},
		{
			Crtd: time.Unix(1692394828, 0),
			Html: "🤖", // https://emojipedia.org/robot
			Name: "robot",
			Rtng: scoreid.String("1692394828509033"),
		},

		//
		// 4. row various expressions
		//

		{
			Crtd: time.Unix(1692393094, 0),
			Html: "💦", // https://emojipedia.org/sweat-droplets
			Name: "sweat-droplets",
			Rtng: scoreid.String("1692393094788405"),
		},
		{
			Crtd: time.Unix(1692392918, 0),
			Html: "🔥", // https://emojipedia.org/fire
			Name: "fire",
			Rtng: scoreid.String("1692392918537493"),
		},
		{
			Crtd: time.Unix(1692393000, 0),
			Html: "👀", // https://emojipedia.org/eyes
			Name: "eyes",
			Rtng: scoreid.String("1692393000623173"),
		},
		{
			Crtd: time.Unix(1692394843, 0),
			Html: "✅", // https://emojipedia.org/check-mark-button
			Name: "check-mark-button",
			Rtng: scoreid.String("1692394843604468"),
		},
		{
			Crtd: time.Unix(1692392966, 0),
			Html: "❗", // https://emojipedia.org/exclamation-mark
			Name: "exclamation-mark",
			Rtng: scoreid.String("1692392966745970"),
		},
		{
			Crtd: time.Unix(1692393010, 0),
			Html: "💯", // https://emojipedia.org/hundred-points
			Name: "hundred-points",
			Rtng: scoreid.String("1692393010008146"),
		},
	}, nil
}
