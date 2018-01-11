package composite

import (
	"strconv"
)

type entry interface {
	getName() string
	getSize() int
	PrintList(prefix string) string
	Add(entry entry)
}

type defaultEntry struct {
	entry
	name string
}

func (de *defaultEntry) getName() string {
	return de.name
}

func (de *defaultEntry) print(entry entry) string {
	return entry.getName() + " (" + strconv.Itoa(entry.getSize()) + ")\n"
}

type file struct {
	*defaultEntry
	size int
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) PrintList(prefix string) string {
	return prefix + "/" + f.print(f)
}

func (f *file) Add(entry entry) {}

type directory struct {
	*defaultEntry
	dir []entry
}

func (d *directory) getSize() int {
	size := 0
	for _, dir := range d.dir {
		size += dir.getSize()
	}
	return size
}

func (d *directory) Add(entry entry) {
	d.dir = append(d.dir, entry)
}

func (d *directory) PrintList(prefix string) string {
	list := prefix + "/" + d.print(d)
	for _, dir := range d.dir {
		list += dir.PrintList(prefix + "/" + d.getName())
	}
	return list
}

// NewFile returns a file instance
func NewFile(name string, size int) *file {
	return &file{
		defaultEntry: &defaultEntry{name: name},
		size:         size,
	}
}

// NewDirectory returns a directory instance
func NewDirectory(name string) *directory {
	return &directory{defaultEntry: &defaultEntry{name: name}}
}
