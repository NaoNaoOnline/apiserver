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

	// EventTime is used to resolve from event times to their respective event
	// IDs.
	//
	//     event time           event IDs
	//                    ->
	//     eve/eve/tim          1234,5678
	//
	EventTime = "eve/eve/tim"

	// EventUser is used to store all the IDs of events created by a specific
	// user.
	//
	//     user ID               event IDs
	//                     ->
	//     eve/use/5678          1234,5678
	//
	EventUser = "eve/use/%s"

	// LabelSystem is used to store all the IDs of system labels.
	//
	//     kind system            label IDs
	//                      ->
	//     lab/kin/sys            1234,5678
	//
	LabelSystem = "lab/kin/sys"

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

	// PolicyBuffer is used to store all chain specific policy records
	// intermittendly in a sorted set. The values here are policy records. The
	// scores here are chain IDs.
	//
	//     policy buffer          policy records
	//                      ->
	//     pol/buf                {"key": "val"}
	//
	PolicyBuffer = "pol/buf"

	// PolicyActive is used to store all active permission states in a simple
	// key-value pair. The value here is a list of currently active permission
	// states, read policy records.
	//
	//     active permissions          policy records
	//                           ->
	//     pol/act                     [{"key": "val"}]
	//
	PolicyActive = "pol/act"

	// ReactionCustom is used to store all the IDs of reactions individually
	// created by users.
	//
	//     kind custom          reaction IDs
	//                    ->
	//     rct/kin/cus          1234,5678
	//
	ReactionCustom = "rct/kin/cus"

	// ReactionObject is used to store our internal representation of an reaction
	// object.
	//
	//     reaction ID           reaction object
	//                     ->
	//     rct/obj/1234          {"key": "val"}
	//
	ReactionObject = "rct/obj/%s"

	// ReactionSystem is used to store all the IDs of reactions natively supported
	// by the system.
	//
	//     kind system          reaction IDs
	//                    ->
	//     rct/kin/sys          1234,5678
	//
	ReactionSystem = "rct/kin/sys"

	// ReactionUser is used to store all the IDs of reactions created by a specific
	// user.
	//
	//     user ID               reaction IDs
	//                     ->
	//     rct/use/5678          1234,5678
	//
	ReactionUser = "rct/use/%s"

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

	// UserName is used to store user mappings between user names and user IDs.
	//
	//     user name            user ID
	//                    ->
	//     use/nam/foo          5678
	//
	UserName = "use/nam/%s"

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

	// WalletAddress is used to store wallet mappings between wallet address and
	// internal user representations.
	//
	//     wallet address          user ID
	//                       ->
	//     wal/add/5678            1234
	//
	WalletAddress = "wal/add/%s"

	// VoteUserEvent is used to store all the IDs of votes mapped to a specific
	// user/event combination.
	//
	//     user ID / event ID             vote IDs
	//                              ->
	//     vot/use/1234/eve/5678          1234,5678
	//
	VoteUserEvent = "vot/use/%s/eve/%s"

	// WalletEthereum is used to store all the IDs of user wallets with kind eth.
	//
	//     kind eth                      wallet IDs
	//                             ->
	//     wal/use/1234/kin/eth          1234,5678
	//
	WalletEthereum = "wal/use/%s/kin/eth"

	// Walletbject is used to store all the IDs of wallets mapped to a specific
	// user/wallet combination.
	//
	//     user ID / wallet ID            wallet object
	//                              ->
	//     wal/use/1234/obj/1234          {"key": "val"}
	//
	WalletObject = "wal/use/%s/obj/%s"

	// WalletUser is used to store all the IDs of wallets created by a specific
	// user.
	//
	//     user ID               wallet IDs
	//                     ->
	//     wal/use/5678          1234,5678
	//
	WalletUser = "wal/use/%s"
)
