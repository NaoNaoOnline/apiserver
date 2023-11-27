package objectstate

const (
	// Created is the status used when a resource was just created.
	Created String = "created"

	// Deleted is the status used when a delete operation was successful.
	Deleted String = "deleted"

	// Dropped is the status used when an update operation was processed, but
	// found to be not applicable. For instance, a user click on an event link may
	// occur multiple times, even though we do not want to count duplicates for
	// such a metric. Instead of returning an error, because the operation in and
	// of itself is legitimate, we return without any underlying modification.
	Dropped String = "dropped"

	// Failure is the status used when an update operation failed to complete as
	// intended. The update operation here intended to reconcile a complex task,
	// potentially resulting in a multitude of outcomes. Failure is then one
	// possible final state.
	Failure String = "failure"

	// Started is the status used when an update operation was initiated for
	// thorough bachground processing.
	Started String = "started"

	// Success is the status used when an update operation was successful. The
	// update operation here intended to reconcile a complex task, potentially
	// resulting in a multitude of outcomes. Success is then one possible final
	// state.
	Success String = "success"

	// Updated is the status used when an update operation was successful. The
	// update operation here intended to update a resource and succeeded doing so.
	Updated String = "updated"

	// Waiting is the status used when an update operation is still in progress.
	Waiting String = "waiting"
)

type String string

func (s String) String() string {
	return string(s)
}
