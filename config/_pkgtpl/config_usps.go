// +build ignore

package usps

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
			ID: "carriers",
			Groups: element.NewGroupSlice(
				&element.Group{
					ID:        "usps",
					Label:     `USPS`,
					SortOrder: 110,
					Scope:     scope.PermAll,
					Fields: element.NewFieldSlice(
						&element.Field{
							// Path: carriers/usps/active
							ID:        "active",
							Label:     `Enabled for Checkout`,
							Type:      element.TypeSelect,
							SortOrder: 10,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   false,
							// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
						},

						&element.Field{
							// Path: carriers/usps/active_rma
							ID:        "active_rma",
							Label:     `Enabled for RMA`,
							Type:      element.TypeSelect,
							SortOrder: 15,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   false,
							// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
						},

						&element.Field{
							// Path: carriers/usps/gateway_url
							ID:        "gateway_url",
							Label:     `Gateway URL`,
							Type:      element.TypeText,
							SortOrder: 20,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   `http://production.shippingapis.com/ShippingAPI.dll`,
						},

						&element.Field{
							// Path: carriers/usps/gateway_secure_url
							ID:        "gateway_secure_url",
							Label:     `Secure Gateway URL`,
							Type:      element.TypeText,
							SortOrder: 30,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   `https://secure.shippingapis.com/ShippingAPI.dll`,
						},

						&element.Field{
							// Path: carriers/usps/title
							ID:        "title",
							Label:     `Title`,
							Type:      element.TypeText,
							SortOrder: 40,
							Visible:   element.VisibleYes,
							Scope:     scope.PermAll,
							Default:   `United States Postal Service`,
						},

						&element.Field{
							// Path: carriers/usps/userid
							ID:        "userid",
							Label:     `User ID`,
							Type:      element.TypeObscure,
							SortOrder: 50,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   nil,
							// BackendModel: Otnegam\Config\Model\Config\Backend\Encrypted @todo Otnegam\Config\Model\Config\Backend\Encrypted
						},

						&element.Field{
							// Path: carriers/usps/password
							ID:        "password",
							Label:     `Password`,
							Type:      element.TypeObscure,
							SortOrder: 53,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   nil,
							// BackendModel: Otnegam\Config\Model\Config\Backend\Encrypted @todo Otnegam\Config\Model\Config\Backend\Encrypted
						},

						&element.Field{
							// Path: carriers/usps/mode
							ID:        "mode",
							Label:     `Mode`,
							Type:      element.TypeSelect,
							SortOrder: 54,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							// SourceModel: Otnegam\Shipping\Model\Config\Source\Online\Mode
						},

						&element.Field{
							// Path: carriers/usps/shipment_requesttype
							ID:        "shipment_requesttype",
							Label:     `Packages Request Type`,
							Type:      element.TypeSelect,
							SortOrder: 55,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   false,
							// SourceModel: Otnegam\Shipping\Model\Config\Source\Online\Requesttype
						},

						&element.Field{
							// Path: carriers/usps/container
							ID:        "container",
							Label:     `Container`,
							Type:      element.TypeSelect,
							SortOrder: 60,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   `VARIABLE`,
							// SourceModel: Otnegam\Usps\Model\Source\Container
						},

						&element.Field{
							// Path: carriers/usps/size
							ID:        "size",
							Label:     `Size`,
							Type:      element.TypeSelect,
							SortOrder: 70,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   `REGULAR`,
							// SourceModel: Otnegam\Usps\Model\Source\Size
						},

						&element.Field{
							// Path: carriers/usps/width
							ID:        "width",
							Label:     `Width`,
							Type:      element.TypeText,
							SortOrder: 73,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
						},

						&element.Field{
							// Path: carriers/usps/length
							ID:        "length",
							Label:     `Length`,
							Type:      element.TypeText,
							SortOrder: 72,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
						},

						&element.Field{
							// Path: carriers/usps/height
							ID:        "height",
							Label:     `Height`,
							Type:      element.TypeText,
							SortOrder: 74,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
						},

						&element.Field{
							// Path: carriers/usps/girth
							ID:        "girth",
							Label:     `Girth`,
							Type:      element.TypeText,
							SortOrder: 76,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
						},

						&element.Field{
							// Path: carriers/usps/machinable
							ID:        "machinable",
							Label:     `Machinable`,
							Type:      element.TypeSelect,
							SortOrder: 80,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   `true`,
							// SourceModel: Otnegam\Usps\Model\Source\Machinable
						},

						&element.Field{
							// Path: carriers/usps/max_package_weight
							ID:        "max_package_weight",
							Label:     `Maximum Package Weight (Please consult your shipping carrier for maximum supported shipping weight)`,
							Type:      element.TypeText,
							SortOrder: 90,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   70,
						},

						&element.Field{
							// Path: carriers/usps/handling_type
							ID:        "handling_type",
							Label:     `Calculate Handling Fee`,
							Type:      element.TypeSelect,
							SortOrder: 100,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   `F`,
							// SourceModel: Otnegam\Shipping\Model\Source\HandlingType
						},

						&element.Field{
							// Path: carriers/usps/handling_action
							ID:        "handling_action",
							Label:     `Handling Applied`,
							Type:      element.TypeSelect,
							SortOrder: 110,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   `O`,
							// SourceModel: Otnegam\Shipping\Model\Source\HandlingAction
						},

						&element.Field{
							// Path: carriers/usps/handling_fee
							ID:        "handling_fee",
							Label:     `Handling Fee`,
							Type:      element.TypeText,
							SortOrder: 120,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
						},

						&element.Field{
							// Path: carriers/usps/allowed_methods
							ID:         "allowed_methods",
							Label:      `Allowed Methods`,
							Type:       element.TypeMultiselect,
							SortOrder:  130,
							Visible:    element.VisibleYes,
							Scope:      scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							CanBeEmpty: true,
							Default:    `0_FCLE,0_FCL,0_FCP,1,2,3,4,6,7,13,16,17,22,23,25,27,28,33,34,35,36,37,42,43,53,55,56,57,61,INT_1,INT_2,INT_4,INT_6,INT_7,INT_8,INT_9,INT_10,INT_11,INT_12,INT_13,INT_14,INT_15,INT_16,INT_20,INT_26`,
							// SourceModel: Otnegam\Usps\Model\Source\Method
						},

						&element.Field{
							// Path: carriers/usps/free_method
							ID:        "free_method",
							Label:     `Free Method`,
							Type:      element.TypeSelect,
							SortOrder: 140,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							// SourceModel: Otnegam\Usps\Model\Source\Freemethod
						},

						&element.Field{
							// Path: carriers/usps/free_shipping_enable
							ID:        "free_shipping_enable",
							Label:     `Free Shipping Amount Threshold`,
							Type:      element.TypeSelect,
							SortOrder: 1500,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							// SourceModel: Otnegam\Config\Model\Config\Source\Enabledisable
						},

						&element.Field{
							// Path: carriers/usps/free_shipping_subtotal
							ID:        "free_shipping_subtotal",
							Label:     `Free Shipping Amount Threshold`,
							Type:      element.TypeText,
							SortOrder: 160,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
						},

						&element.Field{
							// Path: carriers/usps/specificerrmsg
							ID:        "specificerrmsg",
							Label:     `Displayed Error Message`,
							Type:      element.TypeTextarea,
							SortOrder: 170,
							Visible:   element.VisibleYes,
							Scope:     scope.PermAll,
							Default:   `This shipping method is currently unavailable. If you would like to ship using this shipping method, please contact us.`,
						},

						&element.Field{
							// Path: carriers/usps/sallowspecific
							ID:        "sallowspecific",
							Label:     `Ship to Applicable Countries`,
							Type:      element.TypeSelect,
							SortOrder: 180,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							Default:   false,
							// SourceModel: Otnegam\Shipping\Model\Config\Source\Allspecificcountries
						},

						&element.Field{
							// Path: carriers/usps/specificcountry
							ID:         "specificcountry",
							Label:      `Ship to Specific Countries`,
							Type:       element.TypeMultiselect,
							SortOrder:  190,
							Visible:    element.VisibleYes,
							Scope:      scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							CanBeEmpty: true,
							// SourceModel: Otnegam\Directory\Model\Config\Source\Country
						},

						&element.Field{
							// Path: carriers/usps/debug
							ID:        "debug",
							Label:     `Debug`,
							Type:      element.TypeSelect,
							SortOrder: 200,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
						},

						&element.Field{
							// Path: carriers/usps/showmethod
							ID:        "showmethod",
							Label:     `Show Method if Not Applicable`,
							Type:      element.TypeSelect,
							SortOrder: 210,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
							// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
						},

						&element.Field{
							// Path: carriers/usps/sort_order
							ID:        "sort_order",
							Label:     `Sort Order`,
							Type:      element.TypeText,
							SortOrder: 220,
							Visible:   element.VisibleYes,
							Scope:     scope.NewPerm(scope.DefaultID, scope.WebsiteID),
						},
					),
				},
			),
		},

		// Hidden Configuration, may be visible somewhere else ...
		&element.Section{
			ID: "carriers",
			Groups: element.NewGroupSlice(
				&element.Group{
					ID: "usps",
					Fields: element.NewFieldSlice(
						&element.Field{
							// Path: carriers/usps/cutoff_cost
							ID:      `cutoff_cost`,
							Type:    element.TypeHidden,
							Visible: element.VisibleNo,
						},

						&element.Field{
							// Path: carriers/usps/free_method
							ID:      `free_method`,
							Type:    element.TypeHidden,
							Visible: element.VisibleNo,
						},

						&element.Field{
							// Path: carriers/usps/handling
							ID:      `handling`,
							Type:    element.TypeHidden,
							Visible: element.VisibleNo,
						},

						&element.Field{
							// Path: carriers/usps/methods
							ID:      `methods`,
							Type:    element.TypeHidden,
							Visible: element.VisibleNo,
						},

						&element.Field{
							// Path: carriers/usps/model
							ID:      `model`,
							Type:    element.TypeHidden,
							Visible: element.VisibleNo,
							Default: `Otnegam\Usps\Model\Carrier`,
						},

						&element.Field{
							// Path: carriers/usps/isproduction
							ID:      `isproduction`,
							Type:    element.TypeHidden,
							Visible: element.VisibleNo,
							Default: false,
						},

						&element.Field{
							// Path: carriers/usps/is_online
							ID:      `is_online`,
							Type:    element.TypeHidden,
							Visible: element.VisibleNo,
							Default: true,
						},
					),
				},
			),
		},
	)
	Backend = NewBackend(ConfigStructure)
}
