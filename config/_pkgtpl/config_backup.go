// +build ignore

package backup

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
			ID: "system",
			Groups: element.NewGroupSlice(
				&element.Group{
					ID:        "backup",
					Label:     `Scheduled Backup Settings`,
					SortOrder: 500,
					Scope:     scope.NewPerm(scope.DefaultID),
					Fields: element.NewFieldSlice(
						&element.Field{
							// Path: system/backup/enabled
							ID:        "enabled",
							Label:     `Enable Scheduled Backup`,
							Type:      element.TypeSelect,
							SortOrder: 10,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
							// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
						},

						&element.Field{
							// Path: system/backup/type
							ID:        "type",
							Label:     `Backup Type`,
							Type:      element.TypeSelect,
							SortOrder: 20,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
							// SourceModel: Otnegam\Backup\Model\Config\Source\Type
						},

						&element.Field{
							// Path: system/backup/time
							ID:        "time",
							Label:     `Start Time`,
							Type:      element.TypeTime,
							SortOrder: 30,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
						},

						&element.Field{
							// Path: system/backup/frequency
							ID:        "frequency",
							Label:     `Frequency`,
							Type:      element.TypeSelect,
							SortOrder: 40,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
							// BackendModel: Otnegam\Backup\Model\Config\Backend\Cron
							// SourceModel: Otnegam\Cron\Model\Config\Source\Frequency
						},

						&element.Field{
							// Path: system/backup/maintenance
							ID:        "maintenance",
							Label:     `Maintenance Mode`,
							Comment:   element.LongText(`Please put your store into maintenance mode during backup.`),
							Type:      element.TypeSelect,
							SortOrder: 50,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID),
							// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
						},
					),
				},
			),
		},
	)
	Backend = NewBackend(ConfigStructure)
}
