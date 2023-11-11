package objectstate

const (
	// Deleted is the status returned when a delete operation was successful.
	Deleted String = "deleted"

	// Dropped is the status returned when an update operation was processed, but
	// found to be not applicable. For instance, a user click on an event link may
	// occur multiple times, even though we do not want to count duplicates for
	// such a metric. Instead of returning an error, because the operation in and
	// of itself is legitimate, we return without any underlying modification.
	Dropped String = "dropped"

	// Started is the status returned when an update operation was initiated for
	// thorough bachground processing.
	Started String = "started"

	// Updated is the status returned when an update operation was successful.
	Updated String = "updated"

	// Waiting is the status returned when an update operation is still in
	// progress.
	Waiting String = "waiting"
)

type String string

func (s String) String() string {
	return string(s)
}
