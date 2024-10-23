package folder_test

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

var (
	org1 = uuid.Must(uuid.NewV4())
	org2 = uuid.Must(uuid.NewV4())
	org3 = uuid.Must(uuid.NewV4())
	org4 = uuid.Must(uuid.NewV4())
)

// Test startup dataset containing 2 organizations and no duplicate folder names
func noSameFolderNameSample() []folder.Folder {
	return []folder.Folder{
		{Name: "alpha", OrgId: org1, Paths: "alpha"},
		{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
		{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
		{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
		{Name: "golf", OrgId: org1, Paths: "golf"},
	}
}

// Test startup dataset containing 3 organizations and some duplicate folder names across organizations
func mixedFolderNameSample() []folder.Folder {
	return []folder.Folder{
		{Name: "alpha", OrgId: org1, Paths: "alpha"},
		{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
		{Name: "echo", OrgId: org1, Paths: "echo"},
		{Name: "alpha", OrgId: org2, Paths: "alpha"},
		{Name: "bravo", OrgId: org2, Paths: "alpha.bravo"},
		{Name: "romeo", OrgId: org3, Paths: "romeo"},
		{Name: "foxtrot", OrgId: org3, Paths: "romeo.foxtrot"},
	}
}