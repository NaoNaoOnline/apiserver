package labelstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	// Create persists new label objects, if none exists already with the given
	// name.
	//
	//     @inp[0] the label objects providing label specific information
	//     @out[0] the label objects mapped to their internal label IDs
	//
	Create([]*Object) ([]*Object, error)

	// SearchBltn returns the static list of curated event labels natively
	// supported by the system. This method should only be used to bootstrap the
	// initial system state, never to serve RPC requests.
	//
	//     @out[0] static list of curated event labels natively supported by the system
	//
	SearchBltn() []*Object

	// SearchKind returns the label objects matching the given label kinds, e.g.
	// bltn, cate or host.
	//
	//     @inp[0] the label kinds under which label objects are grouped together
	//     @out[0] the list of label objects matching the given label kinds
	//
	SearchKind([]string) ([]*Object, error)

	// SearchLabl returns the label objects matching the given label IDs.
	//
	//     @inp[0] the label IDs to search for
	//     @out[0] the list of label objects matching the given label IDs
	//
	SearchLabl([]objectid.ID) ([]*Object, error)

	// SearchUser returns the label objects created by the given user.
	//
	//     @inp[0] the user ID used to search labels
	//     @out[0] the list of label objects for the given user
	//
	SearchUser(objectid.ID) ([]*Object, error)
}
