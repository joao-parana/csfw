// +build ignore

package tax

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
	// TaxClassesShippingTaxClass => Tax Class for Shipping.
	// Path: tax/classes/shipping_tax_class
	// SourceModel: Otnegam\Tax\Model\TaxClass\Source\Product
	TaxClassesShippingTaxClass model.Str

	// TaxClassesDefaultProductTaxClass => Default Tax Class for Product.
	// Path: tax/classes/default_product_tax_class
	// BackendModel: Otnegam\Tax\Model\Config\TaxClass
	// SourceModel: Otnegam\Tax\Model\TaxClass\Source\Product
	TaxClassesDefaultProductTaxClass model.Str

	// TaxClassesDefaultCustomerTaxClass => Default Tax Class for Customer.
	// Path: tax/classes/default_customer_tax_class
	// SourceModel: Otnegam\Tax\Model\TaxClass\Source\Customer
	TaxClassesDefaultCustomerTaxClass model.Str

	// TaxCalculationAlgorithm => Tax Calculation Method Based On.
	// Path: tax/calculation/algorithm
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Algorithm
	TaxCalculationAlgorithm model.Str

	// TaxCalculationBasedOn => Tax Calculation Based On.
	// Path: tax/calculation/based_on
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\Config\Source\Basedon
	TaxCalculationBasedOn model.Str

	// TaxCalculationPriceIncludesTax => Catalog Prices.
	// This sets whether catalog prices entered from Otnegam Admin include tax.
	// Path: tax/calculation/price_includes_tax
	// BackendModel: Otnegam\Tax\Model\Config\Price\IncludePrice
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\PriceType
	TaxCalculationPriceIncludesTax model.Str

	// TaxCalculationShippingIncludesTax => Shipping Prices.
	// This sets whether shipping amounts entered from Otnegam Admin or obtained
	// from gateways include tax.
	// Path: tax/calculation/shipping_includes_tax
	// BackendModel: Otnegam\Tax\Model\Config\Price\IncludePrice
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\PriceType
	TaxCalculationShippingIncludesTax model.Str

	// TaxCalculationApplyAfterDiscount => Apply Customer Tax.
	// Path: tax/calculation/apply_after_discount
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Apply
	TaxCalculationApplyAfterDiscount model.Str

	// TaxCalculationDiscountTax => Apply Discount On Prices.
	// Apply discount on price including tax is calculated based on store tax if
	// "Apply Tax after Discount" is selected.
	// Path: tax/calculation/discount_tax
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\PriceType
	TaxCalculationDiscountTax model.Str

	// TaxCalculationApplyTaxOn => Apply Tax On.
	// Path: tax/calculation/apply_tax_on
	// SourceModel: Otnegam\Tax\Model\Config\Source\Apply\On
	TaxCalculationApplyTaxOn model.Str

	// TaxCalculationCrossBorderTradeEnabled => Enable Cross Border Trade.
	// When catalog price includes tax, enable this setting to fix the price no
	// matter what the customer's tax rate.
	// Path: tax/calculation/cross_border_trade_enabled
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	TaxCalculationCrossBorderTradeEnabled model.Bool

	// TaxDefaultsCountry => Default Country.
	// Path: tax/defaults/country
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Country
	TaxDefaultsCountry model.Str

	// TaxDefaultsRegion => Default State.
	// Path: tax/defaults/region
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Region
	TaxDefaultsRegion model.Str

	// TaxDefaultsPostcode => Default Post Code.
	// Path: tax/defaults/postcode
	TaxDefaultsPostcode model.Str

	// TaxDisplayType => Display Product Prices In Catalog.
	// Path: tax/display/type
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxDisplayType model.Str

	// TaxDisplayShipping => Display Shipping Prices.
	// Path: tax/display/shipping
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxDisplayShipping model.Str

	// TaxCartDisplayPrice => Display Prices.
	// Path: tax/cart_display/price
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxCartDisplayPrice model.Str

	// TaxCartDisplaySubtotal => Display Subtotal.
	// Path: tax/cart_display/subtotal
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxCartDisplaySubtotal model.Str

	// TaxCartDisplayShipping => Display Shipping Amount.
	// Path: tax/cart_display/shipping
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxCartDisplayShipping model.Str

	// TaxCartDisplayGrandtotal => Include Tax In Order Total.
	// Path: tax/cart_display/grandtotal
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	TaxCartDisplayGrandtotal model.Bool

	// TaxCartDisplayFullSummary => Display Full Tax Summary.
	// Path: tax/cart_display/full_summary
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	TaxCartDisplayFullSummary model.Bool

	// TaxCartDisplayZeroTax => Display Zero Tax Subtotal.
	// Path: tax/cart_display/zero_tax
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	TaxCartDisplayZeroTax model.Bool

	// TaxSalesDisplayPrice => Display Prices.
	// Path: tax/sales_display/price
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxSalesDisplayPrice model.Str

	// TaxSalesDisplaySubtotal => Display Subtotal.
	// Path: tax/sales_display/subtotal
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxSalesDisplaySubtotal model.Str

	// TaxSalesDisplayShipping => Display Shipping Amount.
	// Path: tax/sales_display/shipping
	// BackendModel: Otnegam\Tax\Model\Config\Notification
	// SourceModel: Otnegam\Tax\Model\System\Config\Source\Tax\Display\Type
	TaxSalesDisplayShipping model.Str

	// TaxSalesDisplayGrandtotal => Include Tax In Order Total.
	// Path: tax/sales_display/grandtotal
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	TaxSalesDisplayGrandtotal model.Bool

	// TaxSalesDisplayFullSummary => Display Full Tax Summary.
	// Path: tax/sales_display/full_summary
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	TaxSalesDisplayFullSummary model.Bool

	// TaxSalesDisplayZeroTax => Display Zero Tax Subtotal.
	// Path: tax/sales_display/zero_tax
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	TaxSalesDisplayZeroTax model.Bool
}

// NewBackend initializes the global Backend variable. See init()
func NewBackend(cfgStruct element.SectionSlice) *PkgBackend {
	return (&PkgBackend{}).init(cfgStruct)
}

func (pp *PkgBackend) init(cfgStruct element.SectionSlice) *PkgBackend {
	pp.Lock()
	defer pp.Unlock()
	pp.TaxClassesShippingTaxClass = model.NewStr(`tax/classes/shipping_tax_class`, model.WithConfigStructure(cfgStruct))
	pp.TaxClassesDefaultProductTaxClass = model.NewStr(`tax/classes/default_product_tax_class`, model.WithConfigStructure(cfgStruct))
	pp.TaxClassesDefaultCustomerTaxClass = model.NewStr(`tax/classes/default_customer_tax_class`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationAlgorithm = model.NewStr(`tax/calculation/algorithm`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationBasedOn = model.NewStr(`tax/calculation/based_on`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationPriceIncludesTax = model.NewStr(`tax/calculation/price_includes_tax`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationShippingIncludesTax = model.NewStr(`tax/calculation/shipping_includes_tax`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationApplyAfterDiscount = model.NewStr(`tax/calculation/apply_after_discount`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationDiscountTax = model.NewStr(`tax/calculation/discount_tax`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationApplyTaxOn = model.NewStr(`tax/calculation/apply_tax_on`, model.WithConfigStructure(cfgStruct))
	pp.TaxCalculationCrossBorderTradeEnabled = model.NewBool(`tax/calculation/cross_border_trade_enabled`, model.WithConfigStructure(cfgStruct))
	pp.TaxDefaultsCountry = model.NewStr(`tax/defaults/country`, model.WithConfigStructure(cfgStruct))
	pp.TaxDefaultsRegion = model.NewStr(`tax/defaults/region`, model.WithConfigStructure(cfgStruct))
	pp.TaxDefaultsPostcode = model.NewStr(`tax/defaults/postcode`, model.WithConfigStructure(cfgStruct))
	pp.TaxDisplayType = model.NewStr(`tax/display/type`, model.WithConfigStructure(cfgStruct))
	pp.TaxDisplayShipping = model.NewStr(`tax/display/shipping`, model.WithConfigStructure(cfgStruct))
	pp.TaxCartDisplayPrice = model.NewStr(`tax/cart_display/price`, model.WithConfigStructure(cfgStruct))
	pp.TaxCartDisplaySubtotal = model.NewStr(`tax/cart_display/subtotal`, model.WithConfigStructure(cfgStruct))
	pp.TaxCartDisplayShipping = model.NewStr(`tax/cart_display/shipping`, model.WithConfigStructure(cfgStruct))
	pp.TaxCartDisplayGrandtotal = model.NewBool(`tax/cart_display/grandtotal`, model.WithConfigStructure(cfgStruct))
	pp.TaxCartDisplayFullSummary = model.NewBool(`tax/cart_display/full_summary`, model.WithConfigStructure(cfgStruct))
	pp.TaxCartDisplayZeroTax = model.NewBool(`tax/cart_display/zero_tax`, model.WithConfigStructure(cfgStruct))
	pp.TaxSalesDisplayPrice = model.NewStr(`tax/sales_display/price`, model.WithConfigStructure(cfgStruct))
	pp.TaxSalesDisplaySubtotal = model.NewStr(`tax/sales_display/subtotal`, model.WithConfigStructure(cfgStruct))
	pp.TaxSalesDisplayShipping = model.NewStr(`tax/sales_display/shipping`, model.WithConfigStructure(cfgStruct))
	pp.TaxSalesDisplayGrandtotal = model.NewBool(`tax/sales_display/grandtotal`, model.WithConfigStructure(cfgStruct))
	pp.TaxSalesDisplayFullSummary = model.NewBool(`tax/sales_display/full_summary`, model.WithConfigStructure(cfgStruct))
	pp.TaxSalesDisplayZeroTax = model.NewBool(`tax/sales_display/zero_tax`, model.WithConfigStructure(cfgStruct))

	return pp
}
