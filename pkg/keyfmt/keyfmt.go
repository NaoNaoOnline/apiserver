package keyfmt

const (
	// LabelCategory is used to store all the IDs of category labels.
	//
	//     lab/cat    ->    1234,5678
	//
	LabelCategory = "lab/cat"
	// LabelHost is used to store all the IDs of host labels.
	//
	//     lab/hos    ->    1234,5678
	//
	LabelHost = "lab/hos"
	// LabelObject is used to store our internal representation of a label object.
	//
	//     label id                      label object
	//                             ->
	//     lab/1355803846369828          {"key": "val"}
	//
	LabelObject = "lab/%s"
	// LabelUser is used to store all the IDs of labels created by a specific
	// user.
	//
	//     internal user id                  label id
	//                                 ->
	//     lab/use/1257894840369014          1234,5678
	//
	LabelUser = "lab/use/%s"
	// SubjectClaim is used to store user mappings between external and internal
	// user representations. An external representation might be an OAuth subject
	// claim provided with an access token when authenticating via Google. This
	// subject claim would become part of the key used here. The internal user
	// representation is our own unified user ID, which would then become the
	// value stored using the created subject claim key.
	//
	//     external subject claim          internal user id
	//                               ->
	//     sub/google-oauth2|1234          1257894840369014
	//
	SubjectClaim = "sub/%s"
	// UserObject is used to store our internal representation of a user object.
	//
	//     internal user id          user object
	//                         ->
	//     1257894840369014          {"key": "val"}
	//
	UserObject = "use/%s"
)
