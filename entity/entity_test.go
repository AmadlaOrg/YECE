package entity

import (
	"fmt"
	"github.com/AmadlaOrg/hery/storage"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// FIXME:
/*func TestSetEntity(t *testing.T) {
	entityService := NewEntityService()

	tests := []struct {
		name            string
		serviceEntities []Entity
		inputEntityUri  string
		expectedEntity  Entity
		expectedErr     error
		hasError        bool
	}{
		{
			name: "Valid: entity",
			serviceEntities: []Entity{
				{
					Id:              "97d4b783-f448-483c-8111-380d6082ae1c",
					Entity:          "github.com/AmadlaOrg/Entity@v0.0.0-20240924093300-abcd1234efgh",
					Name:            "Entity",
					RepoUrl:         "https://github.com/AmadlaOrg/Entity",
					Origin:          "github.com/AmadlaOrg/Entity",
					Version:         "latest",
					LatestVersion:   true,
					IsPseudoVersion: true,
					AbsPath:         "/home/user/.hery/collection/amadla/entity/github.com/AmadlaOrg/Entity@v.0.0.0-20240924093300-abcd1234efgh",
					Have:            true,
					Hash:            "",
					Exist:           true,
					Schema:          &jsonschema.Schema{},
					Config: map[string]any{
						"name":        "Entity",
						"description": "The meta Entity definition.",
						"category":    "General",
						"tags": []string{
							"main",
							"master",
						},
					},
				},
				{
					Id:              "56d4b783-c448-983c-8911-400d6082ae1c",
					Entity:          "github.com/AmadlaOrg/Entity@v0.0.0-20240924093300-abcd1234efgh",
					Name:            "Entity",
					RepoUrl:         "https://github.com/AmadlaOrg/Entity",
					Origin:          "github.com/AmadlaOrg/Entity",
					Version:         "latest",
					LatestVersion:   false,
					IsPseudoVersion: true,
					AbsPath:         "/home/user/.hery/amadla/entity/github.com/AmadlaOrg/Entity@v.0.0.0-20240924093300-abcd1234efgh",
					Have:            true,
					Hash:            "",
					Exist:           true,
					Schema:          &jsonschema.Schema{},
					Config: map[string]any{
						"name":        "Entity",
						"description": "The meta Entity definition.",
						"category":    "General",
						"tags": []string{
							"main",
							"master",
						},
					},
				},
			},
			inputEntityUri: "github.com/AmadlaOrg/Entity@latest",
			expectedEntity: Entity{},
			expectedErr:    nil,
			hasError:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entityService.Entities = tt.serviceEntities

			got, err := entityService.GetEntity(tt.inputEntityUri)
			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedEntity, got)
		})
	}
}*/

// FIXME:
/*func TestGetEntity(t *testing.T) {
	entityService := NewEntityService()
	tests := []struct {
		name            string
		serviceEntities []Entity
		inputEntityUri  string
		expectedEntity  Entity
		expectedErr     error
		hasError        bool
	}{
		{
			name: "Valid: entity",
			serviceEntities: []Entity{
				{
					Id:              "97d4b783-f448-483c-8111-380d6082ae1c",
					Entity:          "github.com/AmadlaOrg/Entity@v0.0.0-20240924093300-abcd1234efgh",
					Name:            "Entity",
					RepoUrl:         "https://github.com/AmadlaOrg/Entity",
					Origin:          "github.com/AmadlaOrg/Entity",
					Version:         "latest",
					LatestVersion:   true,
					IsPseudoVersion: true,
					AbsPath:         "/home/user/.hery/collection/amadla/entity/github.com/AmadlaOrg/Entity@v.0.0.0-20240924093300-abcd1234efgh",
					Have:            true,
					Hash:            "",
					Exist:           true,
					Schema:          &jsonschema.Schema{},
					Config: map[string]any{
						"name":        "Entity",
						"description": "The meta Entity definition.",
						"category":    "General",
						"tags": []string{
							"main",
							"master",
						},
					},
				},
				{
					Id:              "12c4b793-d458-756f-8151-740d6082ae1f",
					Entity:          "github.com/AmadlaOrg/Entity@v0.0.0-20230924093300-abcd1234efgh",
					Name:            "Entity",
					RepoUrl:         "https://github.com/AmadlaOrg/Entity",
					Origin:          "github.com/AmadlaOrg/Entity",
					Version:         "latest",
					LatestVersion:   false,
					IsPseudoVersion: true,
					AbsPath:         "/home/user/.hery/collection/amadla/entity/github.com/AmadlaOrg/Entity@v.0.0.0-20230924093300-abcd1234efgh",
					Have:            true,
					Hash:            "",
					Exist:           true,
					Schema:          &jsonschema.Schema{},
					Config: map[string]any{
						"name":        "Entity",
						"description": "The meta Entity definition.",
						"category":    "General",
						"tags": []string{
							"main",
							"master",
						},
					},
				},
				{
					Id:              "98d4b682-c758-943c-8911-560d9022ae3c",
					Entity:          "github.com/AmadlaOrg/QAFixturesEntityMultipleTagVersion@v2.1.0",
					Name:            "QAFixturesEntityMultipleTagVersion",
					RepoUrl:         "https://github.com/AmadlaOrg/QAFixturesEntityMultipleTagVersion",
					Origin:          "github.com/AmadlaOrg/QAFixturesEntityMultipleTagVersion",
					Version:         "v2.1.0",
					LatestVersion:   true,
					IsPseudoVersion: false,
					AbsPath:         "/home/user/.hery/amadla/entity/github.com/AmadlaOrg/QAFixturesEntityMultipleTagVersion@v2.1.0",
					Have:            true,
					Hash:            "",
					Exist:           true,
					Schema:          &jsonschema.Schema{},
					Config: map[string]any{
						"name":        "QAFixturesEntityMultipleTagVersion",
						"description": "Entity Multiple Tag Version definitions.",
						"category":    "QA",
						"tags": []string{
							"QA",
							"fixture",
							"test",
						},
					},
				},
			},
			inputEntityUri: "github.com/AmadlaOrg/Entity@latest",
			expectedEntity: Entity{},
			expectedErr:    nil,
			hasError:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entityService.Entities = tt.serviceEntities
			got, err := entityService.GetEntity(tt.inputEntityUri)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedEntity, got)
		})
	}
}*/

func TestFindEntityDir(t *testing.T) {
	entityService := NewEntityService()

	// Setup test directories
	basePath := "/tmp/.hery/test/entity"
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		t.Fatal("cannot create test directory")
	}
	defer func() {
		err := os.RemoveAll("/tmp/.hery")
		if err != nil {
			t.Fatal("cannot remove test directory")
		}
	}() // Clean up after tests

	tests := []struct {
		name        string
		paths       storage.AbsPaths
		entityVals  Entity
		setupFunc   func()
		expected    string
		expectedErr error
	}{
		{
			name: "Exact version match",
			paths: storage.AbsPaths{
				Entities: basePath,
			},
			entityVals: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0",
				Origin:  "github.com/AmadlaOrg",
				Entity:  filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0"),
			},
			setupFunc: func() {
				err := os.MkdirAll(filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0"), os.ModePerm)
				if err != nil {
					t.Fatal("cannot create test directory")
				}
			},
			expected:    filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0"),
			expectedErr: nil,
		},
		{
			name: "Pseudo version match",
			paths: storage.AbsPaths{
				Entities: basePath,
			},
			entityVals: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0-20240726095222-c7e9911d38b2",
				Origin:  "github.com/AmadlaOrg",
			},
			setupFunc: func() {
				err := os.MkdirAll(filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0-20240726095222-c7e9911d38b2"), os.ModePerm)
				if err != nil {
					t.Fatal("cannot create test directory")
				}
			},
			expected:    filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0-20240726095222-c7e9911d38b2"),
			expectedErr: nil,
		},
		/*{
			name: "No matching exact version",
			paths: storage.AbsPaths{
				Entities: basePath,
			},
			entityVals: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0",
				Origin:  "github.com/AmadlaOrg",
				Entity:  filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0"),
			},
			setupFunc: func() {},
			expected:  "",
			expectedErr: errors.Join(
				errtypes.NotFoundError,
				fmt.Errorf("no matching directory found for exact version: %s", filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0")),
			),
		},
		{
			name: "No matching pseudo version",
			paths: storage.AbsPaths{
				Entities: basePath,
			},
			entityVals: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0-20240726095222-c7e9911d38b2",
				Origin:  "github.com/AmadlaOrg",
			},
			setupFunc: func() {},
			expected:  "",
			expectedErr: errors.Join(
				errtypes.NotFoundError,
				fmt.Errorf("no matching directories found for pattern: %s", filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0-*c7e9911d38b2")),
			),
		},
		{
			name: "Multiple matching pseudo versions",
			paths: storage.AbsPaths{
				Entities: basePath,
			},
			entityVals: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0-20240726095222-c7e9911d38b2",
				Origin:  "github.com/AmadlaOrg",
			},
			setupFunc: func() {
				err := os.MkdirAll(filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0-20240726095222-c7e9911d38b2"), os.ModePerm)
				if err != nil {
					t.Fatal("cannot create test directory")
				}
				err = os.MkdirAll(filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0-20240726095322-c7e9911d38b2"), os.ModePerm)
				if err != nil {
					t.Fatal("cannot create test directory")
				}
			},
			expected: "",
			expectedErr: errors.Join(
				errtypes.MultipleFoundError,
				fmt.Errorf("multiple matching directories found for pattern: %s", filepath.Join(basePath, "github.com/AmadlaOrg", "EntityApplication@v0.0.0-*c7e9911d38b2")),
			),
		},*/
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Setup test case
			test.setupFunc()

			result, err := entityService.FindEntityDir(test.paths, test.entityVals)
			if test.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expectedErr.Error())
			}
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCheckDuplicateEntity(t *testing.T) {
	entityService := NewEntityService()

	tests := []struct {
		name       string
		entities   []Entity
		entityMeta Entity
		expected   error
	}{
		{
			name: "Exact version match",
			entities: []Entity{
				{
					Name:    "EntityApplication",
					Version: "v0.0.0",
					Origin:  "github.com/AmadlaOrg",
				},
			},
			entityMeta: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0",
				Origin:  "github.com/AmadlaOrg",
			},
			expected: fmt.Errorf("duplicate entity found: %v", Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0",
				Origin:  "github.com/AmadlaOrg",
			}),
		},
		{
			name: "Pseudo version match",
			entities: []Entity{
				{
					Name:    "EntityApplication",
					Version: "v0.0.0-20240726095222-c7e9911d38b2",
					Origin:  "github.com/AmadlaOrg",
				},
			},
			entityMeta: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0-20240726095222-c7e9911d38b2",
				Origin:  "github.com/AmadlaOrg",
			},
			expected: fmt.Errorf("duplicate entity found: %v", Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0-20240726095222-c7e9911d38b2",
				Origin:  "github.com/AmadlaOrg",
			}),
		},
		{
			name: "No match",
			entities: []Entity{
				{
					Name:    "EntityApplication",
					Version: "v0.0.0",
					Origin:  "github.com/AmadlaOrg",
				},
			},
			entityMeta: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.1",
				Origin:  "github.com/AmadlaOrg",
			},
			expected: nil,
		},
		{
			name: "Pseudo version no match",
			entities: []Entity{
				{
					Name:    "EntityApplication",
					Version: "v0.0.0-20240726095222-c7e9911d38b2",
					Origin:  "github.com/AmadlaOrg",
				},
			},
			entityMeta: Entity{
				Name:    "EntityApplication",
				Version: "v0.0.0-20240726095222-c889911d00b2",
				Origin:  "github.com/AmadlaOrg",
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := entityService.CheckDuplicateEntity(test.entities, test.entityMeta)
			if test.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expected.Error())
			}
		})
	}
}

func TestGeneratePseudoVersionPattern(t *testing.T) {
	entityService := NewEntityService()

	tests := []struct {
		name         string
		inputName    string
		inputVersion string
		expected     string
	}{
		{
			name:         "Basic pseudo version",
			inputName:    "EntityApplication",
			inputVersion: "v0.0.0-20231231235959-1234567890ab",
			expected:     "EntityApplication@v0.0.0-*-1234567890ab",
		},
		{
			name:         "Another pseudo version",
			inputName:    "AnotherEntity",
			inputVersion: "v1.2.3-20230101010101-abcdefabcdef",
			expected:     "AnotherEntity@v1.2.3-*-abcdefabcdef",
		},
		{
			name:         "Different pseudo version format",
			inputName:    "ThirdEntity",
			inputVersion: "v2.0.0-20231231235959-1234567890ab",
			expected:     "ThirdEntity@v2.0.0-*-1234567890ab",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := entityService.GeneratePseudoVersionPattern(test.inputName, test.inputVersion)
			assert.Equal(t, test.expected, result)
		})
	}
}

/*func createTestDirectoryStructure(t *testing.T, root string) {
	// Create test directories and files
	err := os.MkdirAll(filepath.Join(root, "origin1", "entity1@v1.0.0"), 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(filepath.Join(root, "origin1", "entity2@v2.1.0"), 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(filepath.Join(root, "origin2", "entity3@v0.9.1"), 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(filepath.Join(root, "origin2", "entity4@v1.1.0"), 0755)
	assert.NoError(t, err)

	// Create some files (to be ignored by the crawler)
	_, err = os.Create(filepath.Join(root, "origin1", "file1.txt"))
	assert.NoError(t, err)
	_, err = os.Create(filepath.Join(root, "origin2", "file2.txt"))
	assert.NoError(t, err)
}

func TestCrawlDirectoriesParallel(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Set up the test directory structure
	createTestDirectoryStructure(t, tmpDir)

	// Define the expected entities
	expectedEntities := map[string]Entity{
		"entity1": {Origin: "origin1", Version: "v1.0.0"},
		"entity2": {Origin: "origin1", Version: "v2.1.0"},
		"entity3": {Origin: "origin2", Version: "v0.9.1"},
		"entity4": {Origin: "origin2", Version: "v1.1.0"},
	}

	// Run the function under test
	entities, err := CrawlDirectoriesParallel(tmpDir)
	assert.NoError(t, err)
	assert.Equal(t, expectedEntities, entities)
}*/
