package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		want    []folder.Folder
	}{
		{
			name: "OrgID contains 5 folders",
			orgID: org1,
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "echo"},
			},
		},
		{
			name: "OrgID contains 2 folders",
			orgID: org2,
			want: []folder.Folder{
				{Name: "alpha", OrgId: org2, Paths: "alpha"},
				{Name: "bravo", OrgId: org2, Paths: "alpha.bravo"},
			},
		},
		{
			name:  "OrgID contains no folders",
			orgID: org4,
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(getFolderSample())
			res := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		want    []folder.Folder
		folderName string
		expectError bool
		errorMessage string
	}{
		{
			name: "Does not exist in all folders",
			orgID: org1,
			folderName: "invalid",
			expectError: true,
			errorMessage: "folder does not exist",
		},
		{
			name: "Does not exist in the organization",
			orgID: org2,
			folderName: "charlie",
			expectError: true,
			errorMessage: "folder does not exist in the specified organization",
		},
		{
			name: "Unique folder name with children",
			orgID: org3,
			folderName: "romeo",
			want: []folder.Folder{
				{Name: "foxtrot", OrgId: org3, Paths: "romeo.foxtrot"},
			},
			expectError: false,
		},
		{
			name: "Unique folder name with no child folders",
			orgID: org1,
			folderName: "delta",
			want: []folder.Folder{},
			expectError : false,
		},
		{
			name: "Duplicate folder name",
			orgID: org1,
			folderName: "bravo",
			want: []folder.Folder{
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
			},
			expectError: false,
		},
		{
			name: "Duplicate folder name (Alternative)",
			orgID: org2,
			folderName: "bravo",
			want: []folder.Folder{},
			expectError: false,
		},
		{
			name: "Same organization different head node folder",
			orgID: org1,
			folderName: "alpha",
			want: []folder.Folder{
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
			},
			expectError: false,
		},
		{
			name: "Same organization different head node folder (Alternative)",
			orgID: org1,
			folderName: "echo",
			want: []folder.Folder{},
			expectError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(getFolderSample())
			res, err := f.GetAllChildFolders(tt.orgID, tt.folderName)
			
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, res)
				assert.EqualError(t, err, tt.errorMessage)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, res)
			}
		})
	}
}

