package keyfmt

const (
	// SubjectClaim is used to store user mappings between external and internal
	// user representations. An external representation might be an OAuth subject
	// claim provided with an access token when authenticating via Google. This
	// subject claim would become part of the key used here. The internal user
	// representation is our own unified user ID, which would then become the
	// value stored using the created subject claim key.
	//
	//     external subject claim          internal user id
	//                               ->
	//     sub:google-oauth2|1234          964295a1-ae56-4b85-af41-1cb1910d7e36
	//
	SubjectClaim = "sub:%s"
	// UserObject is used to store our internal representation of a user,
	// abstracted away from third party mechanisms.
	//
	//     internal user id                               internal user object
	//                                             ->
	//     964295a1-ae56-4b85-af41-1cb1910d7e36           {"key": "val"}
	//
	UserObject = "use:%s"
)
