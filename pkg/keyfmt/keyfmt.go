package keyfmt

const (
	// DescriptionEvent is used to store all the IDs of descriptions mapped to a
	// specific event.
	//
	//     event ID              description IDs
	//                     ->
	//     des/eve/5678          1234,5678
	//
	DescriptionEvent = "des/eve/%s"
	// DescriptionObject is used to store our internal representation of a
	// description object.
	//
	//     description ID          description object
	//                       ->
	//     des/obj/1234            {"key": "val"}
	//
	DescriptionObject = "des/obj/%s"
	// DescriptionUser is used to store all the IDs of descriptions created by a
	// specific user.
	//
	//     user ID               description IDs
	//                     ->
	//     des/use/5678          1234,5678
	//
	DescriptionUser = "des/use/%s"
	// EventLabel is used to store all the IDs of events mapped to a specific
	// label.
	//
	//     label ID              event IDs
	//                     ->
	//     eve/lab/5678          1234,5678
	//
	EventLabel = "eve/lab/%s"
	// EventObject is used to store our internal representation of an event
	// object.
	//
	//     event ID              event object
	//                     ->
	//     eve/obj/1234          {"key": "val"}
	//
	EventObject = "eve/obj/%s"
	// EventTime is used to store the timestamps of event objects, which resolve
	// to their respective event IDs.
	//
	//     kind time            timestamps
	//                    ->
	//     eve/kin/tim          1234,5678
	//
	EventTime = "eve/kin/tim"
	// EventUser is used to store all the IDs of events created by a specific
	// user.
	//
	//     user ID               event IDs
	//                     ->
	//     eve/use/5678          1234,5678
	//
	EventUser = "eve/use/%s"
	// LabelCategory is used to store all the IDs of category labels.
	//
	//     kind category          label IDs
	//                      ->
	//     lab/kin/cat            1234,5678
	//
	LabelCategory = "lab/kin/cat"
	// LabelHost is used to store all the IDs of host labels.
	//
	//     kind host            label IDs
	//                    ->
	//     lab/kin/hos          1234,5678
	//
	LabelHost = "lab/kin/hos"
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
	//     user ID               label IDs
	//                     ->
	//     lab/use/5678          1234,5678
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
	//     user ID               user object
	//                     ->
	//     use/obj/5678          {"key": "val"}
	//
	UserObject = "use/obj/%s"
	// VoteDescription is used to store all the IDs of votes mapped to a specific
	// description.
	//
	//     description ID          vote IDs
	//                       ->
	//     vot/des/5678            1234,5678
	//
	VoteDescription = "vot/des/%s"
	// VoteObject is used to store our internal representation of a vote object.
	//
	//     vote ID               vote object
	//                     ->
	//     vot/obj/1234          {"key": "val"}
	//
	VoteObject = "vot/obj/%s"
	// VoteUser is used to store all the IDs of votes created by a specific
	// user.
	//
	//     user ID               vote IDs
	//                     ->
	//     vot/use/5678          1234,5678
	//
	VoteUser = "vot/use/%s"
	// VoteUserEvent is used to store all the IDs of votes mapped to a specific
	// user/event combination.
	//
	//     user ID / event ID             vote IDs
	//                              ->
	//     vot/use/1234/eve/5678          1234,5678
	//
	VoteUserEvent = "vot/use/%s/eve/%s"
)
