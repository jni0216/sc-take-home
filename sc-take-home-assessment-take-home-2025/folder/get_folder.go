package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

// Approach and reasoning
//
// The code readability is prioritized over a one pass or two pass approach.
// Using GetFoldersByOrgID over f.orgId = orgID adheres to many good coding principles (DRY, SRP, etc).
// A more general approach is used that will work even if order of folder array input is changed.
//
// Unaccounted for edge cases missing in specification
//
// Same folder name in same organization at different height are technically different folders.
//
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders := f.folders

	// A flag checks for folder in all folders. If also in correct organization, save the folder
	var folderInAll bool
	var parentFolder *Folder

	for i := range folders {
		if folders[i].Name == name {
			folderInAll = true
			if folders[i].OrgId == orgID {
				parentFolder = &folders[i]
				break
			}
		}
	}

	if !folderInAll {
		return nil, errors.New("folder does not exist")
	}
	if parentFolder == nil {
		return nil, errors.New("folder does not exist in the specified organization")
	}

	// If any folder prefix matches our parent folder, it is a subfolder/child of our parent folder
	prefix := parentFolder.Paths + "."
	childFolders := []Folder{}
	orgFolders := f.GetFoldersByOrgID(orgID)

    for _, f := range orgFolders {
        if strings.HasPrefix(f.Paths, prefix) {
            childFolders = append(childFolders, f)
        }
    }

    return childFolders, nil
}


