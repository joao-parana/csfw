// +build ignore

package reports

import (
	"github.com/corestoreio/csfw/config/element"
	"github.com/corestoreio/csfw/config/model"
)

// Backend will be initialized in the init() function together with ConfigStructure.
var Backend *PkgBackend

// PkgBackend just exported for the sake of documentation. See fields
// for more information. The PkgBackend handles the reading and writing
// of configuration values within this package.
type PkgBackend struct {
	model.PkgBackend
	// CatalogRecentlyProductsScope => Show for Current.
	// Path: catalog/recently_products/scope
	// SourceModel: Otnegam\Config\Model\Config\Source\Reports\Scope
	CatalogRecentlyProductsScope model.Str

	// CatalogRecentlyProductsViewedCount => Default Recently Viewed Products Count.
	// Path: catalog/recently_products/viewed_count
	CatalogRecentlyProductsViewedCount model.Str

	// CatalogRecentlyProductsComparedCount => Default Recently Compared Products Count.
	// Path: catalog/recently_products/compared_count
	CatalogRecentlyProductsComparedCount model.Str

	// ReportsDashboardYtdStart => Year-To-Date Starts.
	// Path: reports/dashboard/ytd_start
	ReportsDashboardYtdStart model.Str

	// ReportsDashboardMtdStart => Current Month Starts.
	// Select day of the month.
	// Path: reports/dashboard/mtd_start
	ReportsDashboardMtdStart model.Str
}

// NewBackend initializes the global Backend variable. See init()
func NewBackend(cfgStruct element.SectionSlice) *PkgBackend {
	return (&PkgBackend{}).init(cfgStruct)
}

func (pp *PkgBackend) init(cfgStruct element.SectionSlice) *PkgBackend {
	pp.Lock()
	defer pp.Unlock()
	pp.CatalogRecentlyProductsScope = model.NewStr(`catalog/recently_products/scope`, model.WithConfigStructure(cfgStruct))
	pp.CatalogRecentlyProductsViewedCount = model.NewStr(`catalog/recently_products/viewed_count`, model.WithConfigStructure(cfgStruct))
	pp.CatalogRecentlyProductsComparedCount = model.NewStr(`catalog/recently_products/compared_count`, model.WithConfigStructure(cfgStruct))
	pp.ReportsDashboardYtdStart = model.NewStr(`reports/dashboard/ytd_start`, model.WithConfigStructure(cfgStruct))
	pp.ReportsDashboardMtdStart = model.NewStr(`reports/dashboard/mtd_start`, model.WithConfigStructure(cfgStruct))

	return pp
}
