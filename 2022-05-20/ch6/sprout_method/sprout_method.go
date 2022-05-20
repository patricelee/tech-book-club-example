package sproutmethod

// ----- legacy code -----
type TransactionGateWay struct {
	EntriesManager []Entry
}

func (t TransactionGateWay) PostEntries(entries []Entry) {
	for _, entry := range entries {
		entry.postDate()
	}
}

type Entry struct{}

func (e Entry) postDate() {}

// ----- legacy code end-----

