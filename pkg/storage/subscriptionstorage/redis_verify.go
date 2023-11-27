package subscriptionstorage

func (r *Redis) VerifyAddr([]string) ([]bool, error) {
	// TODO VerifyAddr
	//
	//     1. use wallet storage to search for wallet object using wallet address
	//     2. if wallet object is labelled to be used for accounting, then continue
	//     3. use event storage to search for amount of events created by wallet owner
	//     4. if wallet owner created more than X events with Y link clicks, then continue
	//     5. success
	//
	return nil, nil
}
