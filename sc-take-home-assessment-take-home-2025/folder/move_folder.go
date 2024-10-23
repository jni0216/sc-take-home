package folder

import (
	"errors"
	"strings"
)

// Approach and reasoning
//
// We will take a faster implementation using 2 passes to implement move folder
// due to typically requiring fast desired behaviour when moving things.
//
// Unaccounted for edge cases missing in specification
//
// No way to differentiate if folders have the same name in different locations.
// Can't move a folder to the root node (start of an organization)
//
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders

	// Get the folders and perform error checking
	var sourceFolder *Folder
	var destinationFolder *Folder

	for i := range folders {
		if folders[i].Name == name {
			sourceFolder = &folders[i]
		}
		if folders[i].Name == dst {
			destinationFolder = &folders[i]
		}
	}

	if sourceFolder == nil {
		return nil, errors.New("source folder does not exist")
	}
	if destinationFolder == nil {
		return nil, errors.New("destination folder does not exist")
	}
	if sourceFolder == destinationFolder {
		return nil, errors.New("cannot move a folder to itself")
	}
	if sourceFolder.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}
	if strings.HasPrefix(destinationFolder.Paths, sourceFolder.Paths) {
		return nil, errors.New("cannot move a folder to a child of itself")
	}

	// For all folders, if it is the source folder (e.g. a.b), or a child of the source folder (e.g. a.b...)
	// Replace prefix with destination folder path (e.g. c.x.b) then append the rest of the path (e.g. c.x.b...)
	oldPathPrefix := sourceFolder.Paths
	newPathPrefix := destinationFolder.Paths + "." + sourceFolder.Name

	for i := range folders {
		if strings.HasPrefix(folders[i].Paths, oldPathPrefix) {
			folders[i].Paths = newPathPrefix + strings.TrimPrefix(folders[i].Paths, oldPathPrefix)
		}
	}

	return folders, nil
}
