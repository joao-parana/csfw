// +build ignore

package customer

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
	// CustomerAccountShareScope => Share Customer Accounts.
	// Path: customer/account_share/scope
	// BackendModel: Otnegam\Customer\Model\Config\Share
	// SourceModel: Otnegam\Customer\Model\Config\Share
	CustomerAccountShareScope model.Str

	// CustomerCreateAccountAutoGroupAssign => Enable Automatic Assignment to Customer Group.
	// Path: customer/create_account/auto_group_assign
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerCreateAccountAutoGroupAssign model.Bool

	// CustomerCreateAccountTaxCalculationAddressType => Tax Calculation Based On.
	// Path: customer/create_account/tax_calculation_address_type
	// SourceModel: Otnegam\Customer\Model\Config\Source\Address\Type
	CustomerCreateAccountTaxCalculationAddressType model.Str

	// CustomerCreateAccountDefaultGroup => Default Group.
	// Path: customer/create_account/default_group
	// SourceModel: Otnegam\Customer\Model\Config\Source\Group
	CustomerCreateAccountDefaultGroup model.Str

	// CustomerCreateAccountVivDomesticGroup => Group for Valid VAT ID - Domestic.
	// Path: customer/create_account/viv_domestic_group
	// SourceModel: Otnegam\Customer\Model\Config\Source\Group
	CustomerCreateAccountVivDomesticGroup model.Str

	// CustomerCreateAccountVivIntraUnionGroup => Group for Valid VAT ID - Intra-Union.
	// Path: customer/create_account/viv_intra_union_group
	// SourceModel: Otnegam\Customer\Model\Config\Source\Group
	CustomerCreateAccountVivIntraUnionGroup model.Str

	// CustomerCreateAccountVivInvalidGroup => Group for Invalid VAT ID.
	// Path: customer/create_account/viv_invalid_group
	// SourceModel: Otnegam\Customer\Model\Config\Source\Group
	CustomerCreateAccountVivInvalidGroup model.Str

	// CustomerCreateAccountVivErrorGroup => Validation Error Group.
	// Path: customer/create_account/viv_error_group
	// SourceModel: Otnegam\Customer\Model\Config\Source\Group
	CustomerCreateAccountVivErrorGroup model.Str

	// CustomerCreateAccountVivOnEachTransaction => Validate on Each Transaction.
	// Path: customer/create_account/viv_on_each_transaction
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerCreateAccountVivOnEachTransaction model.Bool

	// CustomerCreateAccountVivDisableAutoGroupAssignDefault => Default Value for Disable Automatic Group Changes Based on VAT ID.
	// Path: customer/create_account/viv_disable_auto_group_assign_default
	// BackendModel: Otnegam\Customer\Model\Config\Backend\CreateAccount\DisableAutoGroupAssignDefault
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerCreateAccountVivDisableAutoGroupAssignDefault model.Bool

	// CustomerCreateAccountVatFrontendVisibility => Show VAT Number on Storefront.
	// To show VAT number on Storefront, set Show VAT Number on Storefront option
	// to Yes.
	// Path: customer/create_account/vat_frontend_visibility
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerCreateAccountVatFrontendVisibility model.Bool

	// CustomerCreateAccountEmailDomain => Default Email Domain.
	// Path: customer/create_account/email_domain
	CustomerCreateAccountEmailDomain model.Str

	// CustomerCreateAccountEmailTemplate => Default Welcome Email.
	// Email template chosen based on theme fallback when "Default" option is
	// selected.
	// Path: customer/create_account/email_template
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Template
	CustomerCreateAccountEmailTemplate model.Str

	// CustomerCreateAccountEmailNoPasswordTemplate => Default Welcome Email Without Password.
	// This email will be sent instead of the Default Welcome Email, if a customer
	// was created without password.  Email template chosen based on theme
	// fallback when "Default" option is selected.
	// Path: customer/create_account/email_no_password_template
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Template
	CustomerCreateAccountEmailNoPasswordTemplate model.Str

	// CustomerCreateAccountEmailIdentity => Email Sender.
	// Path: customer/create_account/email_identity
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Identity
	CustomerCreateAccountEmailIdentity model.Str

	// CustomerCreateAccountConfirm => Require Emails Confirmation.
	// Path: customer/create_account/confirm
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerCreateAccountConfirm model.Bool

	// CustomerCreateAccountEmailConfirmationTemplate => Confirmation Link Email.
	// Email template chosen based on theme fallback when "Default" option is
	// selected.
	// Path: customer/create_account/email_confirmation_template
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Template
	CustomerCreateAccountEmailConfirmationTemplate model.Str

	// CustomerCreateAccountEmailConfirmedTemplate => Welcome Email.
	// This email will be sent instead of the Default Welcome Email, after account
	// confirmation.  Email template chosen based on theme fallback when "Default"
	// option is selected.
	// Path: customer/create_account/email_confirmed_template
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Template
	CustomerCreateAccountEmailConfirmedTemplate model.Str

	// CustomerCreateAccountGenerateHumanFriendlyId => Generate Human-Friendly Customer ID.
	// Path: customer/create_account/generate_human_friendly_id
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerCreateAccountGenerateHumanFriendlyId model.Bool

	// CustomerPasswordForgotEmailTemplate => Forgot Email Template.
	// Email template chosen based on theme fallback when "Default" option is
	// selected.
	// Path: customer/password/forgot_email_template
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Template
	CustomerPasswordForgotEmailTemplate model.Str

	// CustomerPasswordRemindEmailTemplate => Remind Email Template.
	// Email template chosen based on theme fallback when "Default" option is
	// selected.
	// Path: customer/password/remind_email_template
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Template
	CustomerPasswordRemindEmailTemplate model.Str

	// CustomerPasswordResetPasswordTemplate => Reset Password Template.
	// Email template chosen based on theme fallback when "Default" option is
	// selected.
	// Path: customer/password/reset_password_template
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Template
	CustomerPasswordResetPasswordTemplate model.Str

	// CustomerPasswordForgotEmailIdentity => Password Template Email Sender.
	// Path: customer/password/forgot_email_identity
	// SourceModel: Otnegam\Config\Model\Config\Source\Email\Identity
	CustomerPasswordForgotEmailIdentity model.Str

	// CustomerPasswordResetLinkExpirationPeriod => Recovery Link Expiration Period (days).
	// Please enter a number 1 or greater in this field.
	// Path: customer/password/reset_link_expiration_period
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Password\Link\Expirationperiod
	CustomerPasswordResetLinkExpirationPeriod model.Str

	// CustomerAddressStreetLines => Number of Lines in a Street Address.
	// Leave empty for default (2). Valid range: 1-4
	// Path: customer/address/street_lines
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Address\Street
	CustomerAddressStreetLines model.Str

	// CustomerAddressPrefixShow => Show Prefix.
	// The title that goes before name (Mr., Mrs., etc.)
	// Path: customer/address/prefix_show
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Show\Address
	// SourceModel: Otnegam\Config\Model\Config\Source\Nooptreq
	CustomerAddressPrefixShow model.Str

	// CustomerAddressPrefixOptions => Prefix Dropdown Options.
	// Semicolon (;) separated values.Put semicolon in the beginning for empty
	// first option.Leave empty for open text field.
	// Path: customer/address/prefix_options
	CustomerAddressPrefixOptions model.Str

	// CustomerAddressMiddlenameShow => Show Middle Name (initial).
	// Always optional.
	// Path: customer/address/middlename_show
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Show\Address
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerAddressMiddlenameShow model.Bool

	// CustomerAddressSuffixShow => Show Suffix.
	// The suffix that goes after name (Jr., Sr., etc.)
	// Path: customer/address/suffix_show
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Show\Address
	// SourceModel: Otnegam\Config\Model\Config\Source\Nooptreq
	CustomerAddressSuffixShow model.Str

	// CustomerAddressSuffixOptions => Suffix Dropdown Options.
	// Semicolon (;) separated values.Put semicolon in the beginning for empty
	// first option.Leave empty for open text field.
	// Path: customer/address/suffix_options
	CustomerAddressSuffixOptions model.Str

	// CustomerAddressDobShow => Show Date of Birth.
	// Path: customer/address/dob_show
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Show\Customer
	// SourceModel: Otnegam\Config\Model\Config\Source\Nooptreq
	CustomerAddressDobShow model.Str

	// CustomerAddressTaxvatShow => Show Tax/VAT Number.
	// Path: customer/address/taxvat_show
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Show\Customer
	// SourceModel: Otnegam\Config\Model\Config\Source\Nooptreq
	CustomerAddressTaxvatShow model.Str

	// CustomerAddressGenderShow => Show Gender.
	// Path: customer/address/gender_show
	// BackendModel: Otnegam\Customer\Model\Config\Backend\Show\Customer
	// SourceModel: Otnegam\Config\Model\Config\Source\Nooptreq
	CustomerAddressGenderShow model.Str

	// CustomerStartupRedirectDashboard => Redirect Customer to Account Dashboard after Logging in.
	// Customer will stay on the current page if "No" is selected.
	// Path: customer/startup/redirect_dashboard
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	CustomerStartupRedirectDashboard model.Bool

	// CustomerAddressTemplatesText => Text.
	// Path: customer/address_templates/text
	CustomerAddressTemplatesText model.Str

	// CustomerAddressTemplatesOneline => Text One Line.
	// Path: customer/address_templates/oneline
	CustomerAddressTemplatesOneline model.Str

	// CustomerAddressTemplatesHtml => HTML.
	// Path: customer/address_templates/html
	CustomerAddressTemplatesHtml model.Str

	// CustomerAddressTemplatesPdf => PDF.
	// Path: customer/address_templates/pdf
	CustomerAddressTemplatesPdf model.Str

	// CustomerOnlineCustomersOnlineMinutesInterval => Online Minutes Interval.
	// Leave empty for default (15 minutes).
	// Path: customer/online_customers/online_minutes_interval
	CustomerOnlineCustomersOnlineMinutesInterval model.Str

	// GeneralStoreInformationValidateVatNumber => .
	// Path: general/store_information/validate_vat_number
	GeneralStoreInformationValidateVatNumber model.Str

	// GeneralRestrictionAutocompleteOnStorefront => Enable Autocomplete on login/forgot password forms.
	// Path: general/restriction/autocomplete_on_storefront
	// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
	GeneralRestrictionAutocompleteOnStorefront model.Bool
}

// NewBackend initializes the global Backend variable. See init()
func NewBackend(cfgStruct element.SectionSlice) *PkgBackend {
	return (&PkgBackend{}).init(cfgStruct)
}

func (pp *PkgBackend) init(cfgStruct element.SectionSlice) *PkgBackend {
	pp.Lock()
	defer pp.Unlock()
	pp.CustomerAccountShareScope = model.NewStr(`customer/account_share/scope`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountAutoGroupAssign = model.NewBool(`customer/create_account/auto_group_assign`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountTaxCalculationAddressType = model.NewStr(`customer/create_account/tax_calculation_address_type`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountDefaultGroup = model.NewStr(`customer/create_account/default_group`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountVivDomesticGroup = model.NewStr(`customer/create_account/viv_domestic_group`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountVivIntraUnionGroup = model.NewStr(`customer/create_account/viv_intra_union_group`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountVivInvalidGroup = model.NewStr(`customer/create_account/viv_invalid_group`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountVivErrorGroup = model.NewStr(`customer/create_account/viv_error_group`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountVivOnEachTransaction = model.NewBool(`customer/create_account/viv_on_each_transaction`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountVivDisableAutoGroupAssignDefault = model.NewBool(`customer/create_account/viv_disable_auto_group_assign_default`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountVatFrontendVisibility = model.NewBool(`customer/create_account/vat_frontend_visibility`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountEmailDomain = model.NewStr(`customer/create_account/email_domain`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountEmailTemplate = model.NewStr(`customer/create_account/email_template`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountEmailNoPasswordTemplate = model.NewStr(`customer/create_account/email_no_password_template`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountEmailIdentity = model.NewStr(`customer/create_account/email_identity`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountConfirm = model.NewBool(`customer/create_account/confirm`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountEmailConfirmationTemplate = model.NewStr(`customer/create_account/email_confirmation_template`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountEmailConfirmedTemplate = model.NewStr(`customer/create_account/email_confirmed_template`, model.WithConfigStructure(cfgStruct))
	pp.CustomerCreateAccountGenerateHumanFriendlyId = model.NewBool(`customer/create_account/generate_human_friendly_id`, model.WithConfigStructure(cfgStruct))
	pp.CustomerPasswordForgotEmailTemplate = model.NewStr(`customer/password/forgot_email_template`, model.WithConfigStructure(cfgStruct))
	pp.CustomerPasswordRemindEmailTemplate = model.NewStr(`customer/password/remind_email_template`, model.WithConfigStructure(cfgStruct))
	pp.CustomerPasswordResetPasswordTemplate = model.NewStr(`customer/password/reset_password_template`, model.WithConfigStructure(cfgStruct))
	pp.CustomerPasswordForgotEmailIdentity = model.NewStr(`customer/password/forgot_email_identity`, model.WithConfigStructure(cfgStruct))
	pp.CustomerPasswordResetLinkExpirationPeriod = model.NewStr(`customer/password/reset_link_expiration_period`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressStreetLines = model.NewStr(`customer/address/street_lines`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressPrefixShow = model.NewStr(`customer/address/prefix_show`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressPrefixOptions = model.NewStr(`customer/address/prefix_options`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressMiddlenameShow = model.NewBool(`customer/address/middlename_show`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressSuffixShow = model.NewStr(`customer/address/suffix_show`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressSuffixOptions = model.NewStr(`customer/address/suffix_options`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressDobShow = model.NewStr(`customer/address/dob_show`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressTaxvatShow = model.NewStr(`customer/address/taxvat_show`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressGenderShow = model.NewStr(`customer/address/gender_show`, model.WithConfigStructure(cfgStruct))
	pp.CustomerStartupRedirectDashboard = model.NewBool(`customer/startup/redirect_dashboard`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressTemplatesText = model.NewStr(`customer/address_templates/text`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressTemplatesOneline = model.NewStr(`customer/address_templates/oneline`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressTemplatesHtml = model.NewStr(`customer/address_templates/html`, model.WithConfigStructure(cfgStruct))
	pp.CustomerAddressTemplatesPdf = model.NewStr(`customer/address_templates/pdf`, model.WithConfigStructure(cfgStruct))
	pp.CustomerOnlineCustomersOnlineMinutesInterval = model.NewStr(`customer/online_customers/online_minutes_interval`, model.WithConfigStructure(cfgStruct))
	pp.GeneralStoreInformationValidateVatNumber = model.NewStr(`general/store_information/validate_vat_number`, model.WithConfigStructure(cfgStruct))
	pp.GeneralRestrictionAutocompleteOnStorefront = model.NewBool(`general/restriction/autocomplete_on_storefront`, model.WithConfigStructure(cfgStruct))

	return pp
}
