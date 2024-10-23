package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		folderSource string
		folderDestination string
		expectError bool
		errorMessage string
		want    []folder.Folder
	}{
		{
			name: "Invalid Move - Source folder does not exist",
			folderSource: "invalid",
			folderDestination: "delta",
			expectError: true,
			errorMessage: "source folder does not exist",
		},
		{
			name: "Invalid Move - Destination folder does not exist",
			folderSource: "bravo",
			folderDestination: "invalid",
			expectError: true,
			errorMessage: "destination folder does not exist",
		},
		{
			name: "Invalid Move - Cannot move a folder to itself",
			folderSource: "bravo",
			folderDestination: "bravo",
			expectError: true,
			errorMessage: "cannot move a folder to itself",
		},
		{
			name: "Invalid Move - Cannot move a folder to a different organization",
			folderSource: "bravo",
			folderDestination: "foxtrot",
			expectError: true,
			errorMessage: "cannot move a folder to a different organization",
		},
		{
			name: "Invalid Move - Cannot move a folder to a child of itself",
			folderSource: "bravo",
			folderDestination: "charlie",
			expectError: true,
			errorMessage: "cannot move a folder to a child of itself",
		},
		{
			name: "Valid Move - Move folder within same tree in an organization",
			folderSource: "bravo",
			folderDestination: "delta",
			expectError: false,
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.delta.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.delta.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
				{Name: "golf", OrgId: org1, Paths: "golf"},
			},
		},
		{
			name: "Valid Move - Move folder to different tree in an organization",
			folderSource: "bravo",
			folderDestination: "golf",
			expectError: false,
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "golf.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "golf.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
				{Name: "golf", OrgId: org1, Paths: "golf"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f:= folder.NewDriver(noSameFolderNameSample())
			res, err := f.MoveFolder(tt.folderSource, tt.folderDestination)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, res)
				assert.EqualError(t, err, tt.errorMessage)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.want, res) // Order doesn't matter
			}
		})
	}
}
