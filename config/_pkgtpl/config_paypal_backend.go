// +build ignore

package paypal

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
	// PaypalGeneralMerchantCountry => Merchant Country.
	// If not specified, Default Country from General Config will be used
	// Path: paypal/general/merchant_country
	// BackendModel: Otnegam\Paypal\Model\System\Config\Backend\MerchantCountry
	// SourceModel: Otnegam\Paypal\Model\System\Config\Source\MerchantCountry
	PaypalGeneralMerchantCountry model.Str
}

// NewBackend initializes the global Backend variable. See init()
func NewBackend(cfgStruct element.SectionSlice) *PkgBackend {
	return (&PkgBackend{}).init(cfgStruct)
}

func (pp *PkgBackend) init(cfgStruct element.SectionSlice) *PkgBackend {
	pp.Lock()
	defer pp.Unlock()
	pp.PaypalGeneralMerchantCountry = model.NewStr(`paypal/general/merchant_country`, model.WithConfigStructure(cfgStruct))

	return pp
}
