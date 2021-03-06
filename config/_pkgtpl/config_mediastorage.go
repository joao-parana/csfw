// +build ignore

package mediastorage

import (
	"github.com/corestoreio/csfw/config/element"
	"github.com/corestoreio/csfw/store/scope"
)

// ConfigStructure global configuration structure for this package.
// Used in frontend and backend. See init() for details.
var ConfigStructure element.SectionSlice

func init() {
	ConfigStructure = element.MustNewConfiguration(
		&element.Section{
			ID:        "system",
			SortOrder: 900,
			Scope:     scope.PermAll,
			Groups: element.NewGroupSlice(
				&element.Group{
					ID:        "media_storage_configuration",
					Label:     `Storage Configuration for Media`,
					SortOrder: 900,
					Scope:     scope.PermAll,
					Fields: element.NewFieldSlice(
						&element.Field{
							// Path: system/media_storage_configuration/media_storage
							ID:        "media_storage",
							Label:     `Media Storage`,
							Type:      element.TypeSelect,
							SortOrder: 100,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
							// SourceModel: Otnegam\MediaStorage\Model\Config\Source\Storage\Media\Storage
						},

						&element.Field{
							// Path: system/media_storage_configuration/media_database
							ID:        "media_database",
							Label:     `Select Media Database`,
							Type:      element.TypeSelect,
							SortOrder: 200,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
							// BackendModel: Otnegam\MediaStorage\Model\Config\Backend\Storage\Media\Database
							// SourceModel: Otnegam\MediaStorage\Model\Config\Source\Storage\Media\Database
						},

						&element.Field{
							// Path: system/media_storage_configuration/synchronize
							ID:        "synchronize",
							Comment:   element.LongText(`After selecting a new media storage location, press the Synchronize button to transfer all media to that location. Media will not be available in the new location until the synchronization process is complete.`),
							Type:      element.TypeButton,
							SortOrder: 300,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
						},

						&element.Field{
							// Path: system/media_storage_configuration/configuration_update_time
							ID:        "configuration_update_time",
							Label:     `Environment Update Time`,
							Type:      element.TypeText,
							SortOrder: 400,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
						},
					),
				},
			),
		},
	)
	Backend = NewBackend(ConfigStructure)
}
