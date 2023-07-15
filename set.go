package safety

import "sync"

// Set is a thread-safe set[any]struct{} implementation.
type Set struct {
	sync.RWMutex
	set map[any]struct{}
}

// NewSet will return a pointer to a new Set instance.
func NewSet() *Set {
	return &Set{
		set: map[any]struct{}{},
	}
}

// Add will put a new entry into the set.
func (m *Set) Add(entry any) {
	m.Lock()
	defer m.Unlock()

	m.set[entry] = struct{}{}
}

// Clear will delete all set entries
func (m *Set) Clear() {
	m.Lock()
	defer m.Unlock()

	for entry := range m.set {
		delete(m.set, entry)
	}
}

// Delete will delete a set entry.
func (m *Set) Delete(entry any) (ok bool) {
	m.Lock()
	defer m.Unlock()

	if _, ok = m.set[entry]; ok {
		delete(m.set, entry)
	}

	return
}

// Get will return a snapshot of the set. There is no guarantee that
// the items will remain in the set for any amount of time.
func (m *Set) Get() (entries []any) {
	m.RLock()
	defer m.RUnlock()

	for entry := range m.set {
		entries = append(entries, entry)
	}

	return
}

// Has will return whether or not the provided entry exists.
func (m *Set) Has(entry any) (ok bool) {
	m.RLock()
	defer m.RUnlock()

	_, ok = m.set[entry]
	return
}

// Range will loop over the set and run the specified function for
// each entry. The return value determines whether or not to break the
// loop. You should not add or delete entries within Range, and you
// should avoid calling other Set functions or you may cause deadlock.
// Range should be safe to nest for any read operations.
func (m *Set) Range(f func(entry any) bool) {
	m.RLock()
	defer m.RUnlock()

	for entry := range m.set {
		if f(entry) {
			break
		}
	}
}
