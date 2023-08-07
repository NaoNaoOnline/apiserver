package keyfmt

const (
	// DescriptionEvent is used to store all the IDs of descriptions mapped to a
	// specific event.
	DescriptionEvent = "des/eve/%s"
	// DescriptionObject is used to store our internal representation of a
	// description object.
	DescriptionObject = "des/obj/%s"
	// DescriptionUser is used to store all the IDs of descriptions created by a
	// specific user.
	DescriptionUser = "des/use/%s"
	// EventObject is used to store our internal representation of an event
	// object.
	EventObject = "eve/obj/%s"
	// LabelCategory is used to store all the IDs of category labels.
	//
	//     kind category          label object IDs
	//                      ->
	//     lab/cat                1234,5678
	//
	LabelCategory = "lab/cat"
	// LabelHost is used to store all the IDs of host labels.
	//
	//     kind host          label object IDs
	//                  ->
	//     lab/hos            1234,5678
	//
	LabelHost = "lab/hos"
	// LabelObject is used to store our internal representation of a label object.
	//
	//     label ID              label object
	//                     ->
	//     lab/obj/1234          {"key": "val"}
	//
	LabelObject = "lab/obj/%s"
	// LabelUser is used to store all the IDs of labels created by a specific
	// user.
	//
	//     internal user ID          label ID
	//                         ->
	//     lab/use/5678              1234,5678
	//
	LabelUser = "lab/use/%s"
	// UserClaim is used to store user mappings between external and internal user
	// representations. An external representation might be an OAuth subject claim
	// provided with an access token when authenticating via Google. This subject
	// claim would become part of the key used here. The internal user
	// representation is our own unified user ID, which would then become the
	// value stored using the created subject claim key.
	//
	//     external subject claim              internal user ID
	//                                   ->
	//     use/sub/google-oauth2|1234          5678
	//
	UserClaim = "use/sub/%s"
	// UserObject is used to store our internal representation of a user object.
	//
	//     internal user ID          user object
	//                         ->
	//     use/obj/5678              {"key": "val"}
	//
	UserObject = "use/obj/%s"
)
