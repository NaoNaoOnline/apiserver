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

	// EventCreator is used to store all the IDs of users who create events.
	//
	//     event creator          user IDs
	//                      ->
	//     eve/sym/use            1234,5678
	//
	EventCreator = "eve/sym/use"

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

	// EventReference is used to store a self referential event ID. We need this
	// for instance to resolve event keys to event values when collecting event
	// objects based on rules for custom lists.
	//
	//     event ID              event ID
	//                     ->
	//     eve/ref/1234          1234
	//
	EventReference = "eve/ref/%s"

	// EventTime is used to resolve from event times to their respective event
	// IDs.
	//
	//     event time           event IDs
	//                    ->
	//     eve/sym/tim          1234,5678
	//
	EventTime = "eve/sym/tim"

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

	// LikeDescription is used to store all the user IDs that have reacted to a
	// description in the form of a like.
	//
	//     description ID          user IDs
	//                       ->
	//     lik/des/1234            1234,5678
	//
	LikeDescription = "lik/des/%s"

	// LikeMapping is used to store all the indications of users liking a
	// description.
	//
	//     user ID / description ID          0 / 1
	//                                 ->
	//     lik/use/1234/des/5678             1
	//
	LikeMapping = "lik/use/%s/des/%s"

	// LikeUser is used to store all the event IDs that a user reacted to in the
	// form of a description like.
	//
	//     user ID               event IDs
	//                     ->
	//     lik/use/1234          1234,5678
	//
	LikeUser = "lik/use/%s"

	// LinkEvent is used to store all the user IDs that have clicked on the event
	// link while it was actively listed. Events that have already happened are
	// not active and thus do not track clicks.
	//
	//     event ID              user IDs
	//                     ->
	//     lin/eve/1234          1234,5678
	//
	LinkEvent = "lin/eve/%s"

	// LinkMapping is used to store all the indications of users visiting an
	// event.
	//
	//     user ID / event ID             0 / 1
	//                              ->
	//     lin/use/1234/eve/5678          1
	//
	LinkMapping = "lin/use/%s/eve/%s"

	// LinkUser is used to store all the event IDs that a user visited in the form
	// of a link click.
	//
	//     user ID               event IDs
	//                     ->
	//     lik/use/1234          1234,5678
	//
	LinkUser = "lin/use/%s"

	// ListObject is used to store our internal representation of a list object.
	//
	//     list ID               list object
	//                     ->
	//     lis/obj/1234          {"key": "val"}
	//
	ListObject = "lis/obj/%s"

	// ListUser is used to store all the IDs of lists created by a specific user.
	//
	//     user ID               list IDs
	//                     ->
	//     lis/use/5678          1234,5678
	//
	ListUser = "lis/use/%s"

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

	// RuleEvent is used to store all the IDs of rules defining single event IDs
	// in their resource list.
	//
	//     event ID               rule IDs
	//                     ->
	//     rul/eve/5678          1234,5678
	//
	RuleEvent = "rul/eve/%s"

	// RuleList is used to store all the IDs of rules mapped to a specific list.
	//
	//     list ID               rule IDs
	//                     ->
	//     rul/lis/5678          1234,5678
	//
	RuleList = "rul/lis/%s"

	// RuleObject is used to store our internal representation of a rule object.
	//
	//     rule ID               rule object
	//                     ->
	//     rul/obj/1234          {"key": "val"}
	//
	RuleObject = "rul/obj/%s"

	// RuleUser is used to store all the IDs of rules created by a specific user.
	//
	//     user ID               rule IDs
	//                     ->
	//     rul/use/5678          1234,5678
	//
	RuleUser = "rul/use/%s"

	// SubscriptionObject is used to store our internal representation of a
	// Subscription object.
	//
	//     subscription ID          subscription object
	//                        ->
	//     sub/obj/1234             {"key": "val"}
	//
	SubscriptionObject = "sub/obj/%s"

	// SubscriptionPayer is used to store all the IDs of subscriptions payed by a
	// specific user. The user IDs here are pointing to the subscription IDs that
	// they themselves paid for. Note that receiving a subscription and paying for
	// it may not be the same thing, since subscriptions can be gifted.
	//
	//     user ID               subscription IDs
	//                     ->
	//     sub/pay/5678          1234,5678
	//
	SubscriptionPayer = "sub/pay/%s"

	// SubscriptionReceiver is used to store all the IDs of subscriptions received
	// by a specific user. The user IDs here are pointing to the subscription IDs
	// that they themselves received. Note that receiving a subscription and
	// paying for it may not be the same thing, since subscriptions can be gifted.
	//
	//     user ID               subscription IDs
	//                     ->
	//     sub/rec/5678          1234,5678
	//
	SubscriptionReceiver = "sub/rec/%s"

	// SubscriptionUser is used to store all the IDs of subscriptions created by a
	// specific user.
	//
	//     user ID               subscription IDs
	//                     ->
	//     sub/use/5678          1234,5678
	//
	SubscriptionUser = "sub/use/%s"

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

	// WalletAddress is used to store wallet mappings between wallet address and
	// internal user representations. Note that the value here is an object ID
	// pair. The first element is the user ID. The second element is the wallet
	// ID.
	//
	//     wallet address          user ID / wallet ID
	//                       ->
	//     wal/add/0x5678          1234,5678
	//
	WalletAddress = "wal/add/%s"

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
