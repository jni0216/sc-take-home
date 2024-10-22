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
// The code readability is prioritized over a one pass solution for the use case.
// Using GetFoldersByOrgID over f.orgId = orgID adheres to many good coding principles (DRY, SRP, etc).
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders := f.folders

	// Flags check if folder exists in All Folders and Org. If exists in Org, save the folder
	var foundFolderInAll bool
	var foundFolderInOrg bool
	var parentFolder *Folder

	for _, f := range folders {
		if f.Name == name {
			foundFolderInAll = true
			if f.OrgId == orgID {
				foundFolderInOrg = true
				parentFolder = &f
				break
			}
		}
	}

	if !foundFolderInAll {
		return nil, errors.New("folder does not exist")
	}
	if !foundFolderInOrg {
		return nil, errors.New("folder does not exist in the specified organization")
	}

	// Prefix match/find for all child folders for our parent folder in the organization
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


