// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"encoding/json"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// FiatService contains methods and other services that help with interacting with
// the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFiatService] method instead.
type FiatService struct {
	Options []option.RequestOption
}

// NewFiatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFiatService(opts ...option.RequestOption) (r FiatService) {
	r = FiatService{}
	r.Options = opts
	return
}

// Request body for creating a Bridge external fiat account.
//
// The properties Account, AccountOwnerName, Currency, Provider are required.
type CreateExternalFiatAccountRequestBody struct {
	// Bank account details. The `type` field discriminates which shape applies.
	Account          ExternalFiatAccountDataUnion `json:"account,omitzero" api:"required"`
	AccountOwnerName string                       `json:"account_owner_name" api:"required"`
	Currency         string                       `json:"currency" api:"required"`
	// Discriminator: the external fiat account is orchestrated via Bridge.
	//
	// Any of "bridge".
	Provider CreateExternalFiatAccountRequestBodyProvider `json:"provider,omitzero" api:"required"`
	BankName param.Opt[string]                            `json:"bank_name,omitzero"`
	// Physical address associated with an external fiat account.
	Address ExternalFiatAccountAddress `json:"address,omitzero"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment,omitzero"`
	paramObj
}

func (r CreateExternalFiatAccountRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CreateExternalFiatAccountRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateExternalFiatAccountRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator: the external fiat account is orchestrated via Bridge.
type CreateExternalFiatAccountRequestBodyProvider string

const (
	CreateExternalFiatAccountRequestBodyProviderBridge CreateExternalFiatAccountRequestBodyProvider = "bridge"
)

// Request body for creating a Bridge fiat deposit account linked to a wallet.
//
// The properties Destination, Provider, Source are required.
type CreateFiatDepositAccountRequestBody struct {
	// The destination crypto asset and chain for a fiat deposit account.
	Destination FiatDepositAccountDestination `json:"destination,omitzero" api:"required"`
	// Discriminator: the fiat deposit account is orchestrated via Bridge.
	//
	// Any of "bridge".
	Provider CreateFiatDepositAccountRequestBodyProvider `json:"provider,omitzero" api:"required"`
	// The source fiat currency for a fiat deposit account.
	Source CreateFiatDepositAccountSource `json:"source,omitzero" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment,omitzero"`
	paramObj
}

func (r CreateFiatDepositAccountRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CreateFiatDepositAccountRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFiatDepositAccountRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator: the fiat deposit account is orchestrated via Bridge.
type CreateFiatDepositAccountRequestBodyProvider string

const (
	CreateFiatDepositAccountRequestBodyProviderBridge CreateFiatDepositAccountRequestBodyProvider = "bridge"
)

// The source fiat currency for a fiat deposit account.
//
// The property Currency is required.
type CreateFiatDepositAccountSource struct {
	Currency string `json:"currency" api:"required"`
	paramObj
}

func (r CreateFiatDepositAccountSource) MarshalJSON() (data []byte, err error) {
	type shadow CreateFiatDepositAccountSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFiatDepositAccountSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for initiating a payout (crypto to fiat offramp) from a wallet.
//
// The properties Destination, Source are required.
type CreatePayoutRequestBody struct {
	// The destination bank account for a payout.
	Destination PayoutDestination `json:"destination,omitzero" api:"required"`
	// The source crypto asset, chain, and amount for a payout.
	Source PayoutSource `json:"source,omitzero" api:"required"`
	paramObj
}

func (r CreatePayoutRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow CreatePayoutRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePayoutRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Bridge external fiat account linked to a user.
type ExternalFiatAccount struct {
	ID               string `json:"id" api:"required"`
	AccountOwnerName string `json:"account_owner_name" api:"required"`
	AccountType      string `json:"account_type" api:"required"`
	CreatedAt        string `json:"created_at" api:"required"`
	Currency         string `json:"currency" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// Discriminator: the external fiat account is orchestrated via Bridge.
	//
	// Any of "bridge".
	Provider ExternalFiatAccountProvider `json:"provider" api:"required"`
	UserID   string                      `json:"user_id" api:"required"`
	BankName string                      `json:"bank_name"`
	Last4    string                      `json:"last_4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountOwnerName respjson.Field
		AccountType      respjson.Field
		CreatedAt        respjson.Field
		Currency         respjson.Field
		Environment      respjson.Field
		Provider         respjson.Field
		UserID           respjson.Field
		BankName         respjson.Field
		Last4            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalFiatAccount) RawJSON() string { return r.JSON.raw }
func (r *ExternalFiatAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator: the external fiat account is orchestrated via Bridge.
type ExternalFiatAccountProvider string

const (
	ExternalFiatAccountProviderBridge ExternalFiatAccountProvider = "bridge"
)

// Physical address associated with an external fiat account.
//
// The properties City, Country, StreetLine1 are required.
type ExternalFiatAccountAddress struct {
	City        string            `json:"city" api:"required"`
	Country     string            `json:"country" api:"required"`
	StreetLine1 string            `json:"street_line_1" api:"required"`
	PostalCode  param.Opt[string] `json:"postal_code,omitzero"`
	State       param.Opt[string] `json:"state,omitzero"`
	StreetLine2 param.Opt[string] `json:"street_line_2,omitzero"`
	paramObj
}

func (r ExternalFiatAccountAddress) MarshalJSON() (data []byte, err error) {
	type shadow ExternalFiatAccountAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalFiatAccountAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ExternalFiatAccountDataOfUs(accountNumber string, routingNumber string, type_ ExternalFiatAccountUsDataType) ExternalFiatAccountDataUnion {
	var us ExternalFiatAccountUsData
	us.AccountNumber = accountNumber
	us.RoutingNumber = routingNumber
	us.Type = type_
	return ExternalFiatAccountDataUnion{OfUs: &us}
}

func ExternalFiatAccountDataOfGB(accountNumber string, sortCode string, type_ ExternalFiatAccountGBDataType) ExternalFiatAccountDataUnion {
	var gb ExternalFiatAccountGBData
	gb.AccountNumber = accountNumber
	gb.SortCode = sortCode
	gb.Type = type_
	return ExternalFiatAccountDataUnion{OfGB: &gb}
}

func ExternalFiatAccountDataOfPix(type_ ExternalFiatAccountPixDataType) ExternalFiatAccountDataUnion {
	var pix ExternalFiatAccountPixData
	pix.Type = type_
	return ExternalFiatAccountDataUnion{OfPix: &pix}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExternalFiatAccountDataUnion struct {
	OfUs    *ExternalFiatAccountUsData    `json:",omitzero,inline"`
	OfGB    *ExternalFiatAccountGBData    `json:",omitzero,inline"`
	OfPix   *ExternalFiatAccountPixData   `json:",omitzero,inline"`
	OfIban  *ExternalFiatAccountIbanData  `json:",omitzero,inline"`
	OfSwift *ExternalFiatAccountSwiftData `json:",omitzero,inline"`
	paramUnion
}

func (u ExternalFiatAccountDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUs,
		u.OfGB,
		u.OfPix,
		u.OfIban,
		u.OfSwift)
}
func (u *ExternalFiatAccountDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ExternalFiatAccountDataUnion](
		"type",
		apijson.Discriminator[ExternalFiatAccountUsData]("us"),
		apijson.Discriminator[ExternalFiatAccountGBData]("gb"),
		apijson.Discriminator[ExternalFiatAccountPixData]("pix"),
		apijson.Discriminator[ExternalFiatAccountIbanData]("iban"),
		apijson.Discriminator[ExternalFiatAccountSwiftData]("swift"),
	)
}

// UK bank account data for an external fiat account. Pays out over Faster
// Payments.
//
// The properties AccountNumber, SortCode, Type are required.
type ExternalFiatAccountGBData struct {
	// The 8-digit UK bank account number.
	AccountNumber string `json:"account_number" api:"required"`
	// The 6-digit sort code, without hyphens.
	SortCode string `json:"sort_code" api:"required"`
	// Any of "gb".
	Type ExternalFiatAccountGBDataType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ExternalFiatAccountGBData) MarshalJSON() (data []byte, err error) {
	type shadow ExternalFiatAccountGBData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalFiatAccountGBData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalFiatAccountGBDataType string

const (
	ExternalFiatAccountGBDataTypeGB ExternalFiatAccountGBDataType = "gb"
)

// IBAN bank account data for an external fiat account. Pays out over SEPA.
//
// The properties AccountNumber, Bic, Country, Type are required.
type ExternalFiatAccountIbanData struct {
	// The IBAN. Up to 34 characters, per ISO 13616.
	AccountNumber string `json:"account_number" api:"required"`
	// The BIC/SWIFT code of the beneficiary bank.
	Bic string `json:"bic" api:"required"`
	// Country the account is held in, as an ISO 3166-1 alpha-3 code.
	Country string `json:"country" api:"required"`
	// Any of "iban".
	Type ExternalFiatAccountIbanDataType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ExternalFiatAccountIbanData) MarshalJSON() (data []byte, err error) {
	type shadow ExternalFiatAccountIbanData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalFiatAccountIbanData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalFiatAccountIbanDataType string

const (
	ExternalFiatAccountIbanDataTypeIban ExternalFiatAccountIbanDataType = "iban"
)

// Brazilian Pix account data for an external fiat account. Provide exactly one of
// `pix_key` or `br_code`.
//
// The property Type is required.
type ExternalFiatAccountPixData struct {
	// Any of "pix".
	Type ExternalFiatAccountPixDataType `json:"type,omitzero" api:"required"`
	// The Pix "copia e cola" (copy and paste) BR Code.
	BrCode param.Opt[string] `json:"br_code,omitzero"`
	// Optional CPF/CNPJ associated with the account, digits only.
	DocumentNumber param.Opt[string] `json:"document_number,omitzero"`
	// The Pix key: an EVP (UUID), CPF, CNPJ, Brazilian phone number (+55…), or email
	// address.
	PixKey param.Opt[string] `json:"pix_key,omitzero"`
	paramObj
}

func (r ExternalFiatAccountPixData) MarshalJSON() (data []byte, err error) {
	type shadow ExternalFiatAccountPixData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalFiatAccountPixData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalFiatAccountPixDataType string

const (
	ExternalFiatAccountPixDataTypePix ExternalFiatAccountPixDataType = "pix"
)

// Response containing a single external fiat account.
type ExternalFiatAccountResponse struct {
	// A Bridge external fiat account linked to a user.
	ExternalFiatAccount ExternalFiatAccount `json:"external_fiat_account" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalFiatAccount respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalFiatAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalFiatAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business relationship between the payer and the SWIFT account owner.
type ExternalFiatAccountSwiftCategory string

const (
	ExternalFiatAccountSwiftCategoryClient        ExternalFiatAccountSwiftCategory = "client"
	ExternalFiatAccountSwiftCategoryParentCompany ExternalFiatAccountSwiftCategory = "parent_company"
	ExternalFiatAccountSwiftCategorySubsidiary    ExternalFiatAccountSwiftCategory = "subsidiary"
	ExternalFiatAccountSwiftCategorySupplier      ExternalFiatAccountSwiftCategory = "supplier"
)

// SWIFT bank account data for an external fiat account. Pays out over wire. The
// beneficiary address is required for SWIFT and is supplied as the request's
// top-level `address`.
//
// The properties AccountNumber, Bic, Category, PurposeOfFunds,
// ShortBusinessDescription, Type are required.
type ExternalFiatAccountSwiftData struct {
	AccountNumber string `json:"account_number" api:"required"`
	// The BIC/SWIFT code of the beneficiary bank.
	Bic string `json:"bic" api:"required"`
	// Business relationship between the payer and the SWIFT account owner.
	//
	// Any of "client", "parent_company", "subsidiary", "supplier".
	Category                 ExternalFiatAccountSwiftCategory         `json:"category,omitzero" api:"required"`
	PurposeOfFunds           []ExternalFiatAccountSwiftPurposeOfFunds `json:"purpose_of_funds,omitzero" api:"required"`
	ShortBusinessDescription string                                   `json:"short_business_description" api:"required"`
	// Any of "swift".
	Type ExternalFiatAccountSwiftDataType `json:"type,omitzero" api:"required"`
	// Country the account is held in, as an ISO 3166-1 alpha-3 code.
	Country param.Opt[string] `json:"country,omitzero"`
	paramObj
}

func (r ExternalFiatAccountSwiftData) MarshalJSON() (data []byte, err error) {
	type shadow ExternalFiatAccountSwiftData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalFiatAccountSwiftData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalFiatAccountSwiftDataType string

const (
	ExternalFiatAccountSwiftDataTypeSwift ExternalFiatAccountSwiftDataType = "swift"
)

// Reason funds are sent to a SWIFT account, required for cross-border compliance.
type ExternalFiatAccountSwiftPurposeOfFunds string

const (
	ExternalFiatAccountSwiftPurposeOfFundsIntraGroupTransfer         ExternalFiatAccountSwiftPurposeOfFunds = "intra_group_transfer"
	ExternalFiatAccountSwiftPurposeOfFundsInvoiceForGoodsAndServices ExternalFiatAccountSwiftPurposeOfFunds = "invoice_for_goods_and_services"
)

// US bank account data for an external fiat account.
//
// The properties AccountNumber, RoutingNumber, Type are required.
type ExternalFiatAccountUsData struct {
	AccountNumber string `json:"account_number" api:"required"`
	RoutingNumber string `json:"routing_number" api:"required"`
	// Any of "us".
	Type              ExternalFiatAccountUsDataType `json:"type,omitzero" api:"required"`
	CheckingOrSavings param.Opt[string]             `json:"checking_or_savings,omitzero"`
	paramObj
}

func (r ExternalFiatAccountUsData) MarshalJSON() (data []byte, err error) {
	type shadow ExternalFiatAccountUsData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalFiatAccountUsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalFiatAccountUsDataType string

const (
	ExternalFiatAccountUsDataTypeUs ExternalFiatAccountUsDataType = "us"
)

// A Bridge fiat deposit account linked to a wallet.
type FiatDepositAccount struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"created_at" api:"required"`
	// Bank or payment deposit instructions for a fiat deposit account. Shape varies by
	// source currency.
	DepositInstructions FiatDepositInstructions `json:"deposit_instructions" api:"required"`
	// The destination crypto asset and chain for a fiat deposit account.
	Destination FiatDepositAccountDestinationResp `json:"destination" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `json:"environment" api:"required"`
	// Discriminator: the fiat deposit account is orchestrated via Bridge.
	//
	// Any of "bridge".
	Provider FiatDepositAccountProvider `json:"provider" api:"required"`
	// The source fiat currency and available payment rails for a fiat deposit account.
	Source FiatDepositAccountSource `json:"source" api:"required"`
	// Activation status of a fiat deposit account.
	//
	// Any of "activated", "deactivated".
	Status   FiatDepositAccountStatus `json:"status" api:"required"`
	WalletID string                   `json:"wallet_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		DepositInstructions respjson.Field
		Destination         respjson.Field
		Environment         respjson.Field
		Provider            respjson.Field
		Source              respjson.Field
		Status              respjson.Field
		WalletID            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FiatDepositAccount) RawJSON() string { return r.JSON.raw }
func (r *FiatDepositAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator: the fiat deposit account is orchestrated via Bridge.
type FiatDepositAccountProvider string

const (
	FiatDepositAccountProviderBridge FiatDepositAccountProvider = "bridge"
)

// The destination crypto asset and chain for a fiat deposit account.
type FiatDepositAccountDestinationResp struct {
	// Destination crypto asset (e.g. "usdc").
	Asset string `json:"asset" api:"required"`
	// Destination chain (e.g. "base", "tempo").
	Chain string `json:"chain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		Chain       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FiatDepositAccountDestinationResp) RawJSON() string { return r.JSON.raw }
func (r *FiatDepositAccountDestinationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FiatDepositAccountDestinationResp to a
// FiatDepositAccountDestination.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FiatDepositAccountDestination.Overrides()
func (r FiatDepositAccountDestinationResp) ToParam() FiatDepositAccountDestination {
	return param.Override[FiatDepositAccountDestination](json.RawMessage(r.RawJSON()))
}

// The destination crypto asset and chain for a fiat deposit account.
//
// The properties Asset, Chain are required.
type FiatDepositAccountDestination struct {
	// Destination crypto asset (e.g. "usdc").
	Asset string `json:"asset" api:"required"`
	// Destination chain (e.g. "base", "tempo").
	Chain string `json:"chain" api:"required"`
	paramObj
}

func (r FiatDepositAccountDestination) MarshalJSON() (data []byte, err error) {
	type shadow FiatDepositAccountDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FiatDepositAccountDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a single fiat deposit account.
type FiatDepositAccountResponse struct {
	// A Bridge fiat deposit account linked to a wallet.
	FiatDepositAccount FiatDepositAccount `json:"fiat_deposit_account" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FiatDepositAccount respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FiatDepositAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *FiatDepositAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source fiat currency and available payment rails for a fiat deposit account.
type FiatDepositAccountSource struct {
	Currency     string   `json:"currency" api:"required"`
	PaymentRails []string `json:"payment_rails" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency     respjson.Field
		PaymentRails respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FiatDepositAccountSource) RawJSON() string { return r.JSON.raw }
func (r *FiatDepositAccountSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Activation status of a fiat deposit account.
type FiatDepositAccountStatus string

const (
	FiatDepositAccountStatusActivated   FiatDepositAccountStatus = "activated"
	FiatDepositAccountStatusDeactivated FiatDepositAccountStatus = "deactivated"
)

// Bank or payment deposit instructions for a fiat deposit account. Shape varies by
// source currency.
type FiatDepositInstructions struct {
	AccountHolderName      string   `json:"account_holder_name"`
	AccountNumber          string   `json:"account_number"`
	BankAccountNumber      string   `json:"bank_account_number"`
	BankAddress            string   `json:"bank_address"`
	BankBeneficiaryAddress string   `json:"bank_beneficiary_address"`
	BankBeneficiaryName    string   `json:"bank_beneficiary_name"`
	BankName               string   `json:"bank_name"`
	BankRoutingNumber      string   `json:"bank_routing_number"`
	Bic                    string   `json:"bic"`
	BrCode                 string   `json:"br_code"`
	BreBKey                string   `json:"bre_b_key"`
	Clabe                  string   `json:"clabe"`
	DepositMessage         string   `json:"deposit_message"`
	Iban                   string   `json:"iban"`
	PaymentRails           []string `json:"payment_rails"`
	SortCode               string   `json:"sort_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountHolderName      respjson.Field
		AccountNumber          respjson.Field
		BankAccountNumber      respjson.Field
		BankAddress            respjson.Field
		BankBeneficiaryAddress respjson.Field
		BankBeneficiaryName    respjson.Field
		BankName               respjson.Field
		BankRoutingNumber      respjson.Field
		Bic                    respjson.Field
		BrCode                 respjson.Field
		BreBKey                respjson.Field
		Clabe                  respjson.Field
		DepositMessage         respjson.Field
		Iban                   respjson.Field
		PaymentRails           respjson.Field
		SortCode               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FiatDepositInstructions) RawJSON() string { return r.JSON.raw }
func (r *FiatDepositInstructions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KYBAccountPurpose = string

// A beneficial owner, control person, or signer associated with the business. At
// least one of has_ownership, has_control, or is_signer must be true, and the
// business must have at least one control person and one signer.
//
// The properties DateOfBirth, Email, FirstName, HasControl, HasOwnership,
// IdentifyingInformation, IsSigner, LastName, ResidentialAddress are required.
type KYBAssociatedPerson struct {
	// Date of birth in YYYY-MM-DD format. Must be 18 years or older.
	DateOfBirth string `json:"date_of_birth" api:"required"`
	// Email address.
	Email string `json:"email" api:"required" format:"email"`
	// Legal first name.
	FirstName string `json:"first_name" api:"required"`
	// Whether this person is a control person.
	HasControl bool `json:"has_control" api:"required"`
	// Whether this person owns 25% or more of the business.
	HasOwnership bool `json:"has_ownership" api:"required"`
	// Identifying documents for this person.
	IdentifyingInformation []VerificationDocument `json:"identifying_information,omitzero" api:"required"`
	// Whether this person is a signer for the business.
	IsSigner bool `json:"is_signer" api:"required"`
	// Legal last name.
	LastName string `json:"last_name" api:"required"`
	// A postal address used in KYC and KYB data submission.
	ResidentialAddress VerificationAddress `json:"residential_address,omitzero" api:"required"`
	// Whether this person is a director.
	IsDirector param.Opt[bool] `json:"is_director,omitzero"`
	// Legal middle name.
	MiddleName param.Opt[string] `json:"middle_name,omitzero"`
	// Percentage of the business this person owns.
	OwnershipPercentage param.Opt[int64] `json:"ownership_percentage,omitzero"`
	// Phone number in E.164 format.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Date the relationship with the business was established, in YYYY-MM-DD format.
	RelationshipEstablishedAt param.Opt[string] `json:"relationship_established_at,omitzero"`
	// Job title. Required when has_control is true.
	Title param.Opt[string] `json:"title,omitzero"`
	// Latin-1 transliteration of the first name. Required for non-Latin-1 names.
	TransliteratedFirstName param.Opt[string] `json:"transliterated_first_name,omitzero"`
	// Latin-1 transliteration of the last name. Required for non-Latin-1 names.
	TransliteratedLastName param.Opt[string] `json:"transliterated_last_name,omitzero"`
	// Latin-1 transliteration of the middle name. Required for non-Latin-1 names.
	TransliteratedMiddleName param.Opt[string] `json:"transliterated_middle_name,omitzero"`
	// Supporting documents for this person, such as proof of address.
	Documents []KYBIndividualDocument `json:"documents,omitzero"`
	// ISO 3166-1 alpha-3 codes for all nationalities held.
	Nationalities []string `json:"nationalities,omitzero"`
	// Place of birth for an associated person.
	PlaceOfBirth KYBPlaceOfBirth `json:"place_of_birth,omitzero"`
	// A postal address used in KYC and KYB data submission.
	TransliteratedResidentialAddress VerificationAddress `json:"transliterated_residential_address,omitzero"`
	paramObj
}

func (r KYBAssociatedPerson) MarshalJSON() (data []byte, err error) {
	type shadow KYBAssociatedPerson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBAssociatedPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A supporting document for business verification.
//
// The properties File, Purposes are required.
type KYBBusinessDocument struct {
	// Base64-encoded data URI of the document.
	File string `json:"file" api:"required"`
	// What this document evidences. Supports multiple purposes per file.
	Purposes []KYBDocumentPurpose `json:"purposes,omitzero" api:"required"`
	// Document description. Required when "other" is one of the purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r KYBBusinessDocument) MarshalJSON() (data []byte, err error) {
	type shadow KYBBusinessDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBBusinessDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KYBBusinessType = string

type KYBDocumentPurpose = string

type KYBEstimatedAnnualRevenue = string

type KYBHighRiskActivity = string

// A supporting document for an associated person.
//
// The properties File, Purposes are required.
type KYBIndividualDocument struct {
	// Base64-encoded data URI of the document.
	File string `json:"file" api:"required"`
	// What this document evidences. Supports multiple purposes per file.
	Purposes []KYBIndividualDocumentPurpose `json:"purposes,omitzero" api:"required"`
	// Document description. Required when "other" is one of the purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r KYBIndividualDocument) MarshalJSON() (data []byte, err error) {
	type shadow KYBIndividualDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBIndividualDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KYBIndividualDocumentPurpose = string

// Request body for initiating a hosted KYB flow for an organization.
//
// The properties Email, Provider are required.
type KYBLinksRequestBody struct {
	// Email address for the organization.
	Email string `json:"email" api:"required" format:"email"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider,omitzero" api:"required"`
	// Legal name of the business.
	BusinessName param.Opt[string] `json:"business_name,omitzero"`
	// Client-side agreement ID for ToS acceptance.
	ClientAgreementID param.Opt[string] `json:"client_agreement_id,omitzero"`
	// URI to redirect after completing KYB.
	RedirectUri param.Opt[string] `json:"redirect_uri,omitzero" format:"uri"`
	// Endorsements to request during KYB.
	Endorsements []KyxEndorsementName `json:"endorsements,omitzero"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment,omitzero"`
	paramObj
}

func (r KYBLinksRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow KYBLinksRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBLinksRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Place of birth for an associated person.
//
// The property Country is required.
type KYBPlaceOfBirth struct {
	// ISO 3166-1 alpha-3 country code.
	Country string `json:"country" api:"required"`
	// City of birth.
	City param.Opt[string] `json:"city,omitzero"`
	paramObj
}

func (r KYBPlaceOfBirth) MarshalJSON() (data []byte, err error) {
	type shadow KYBPlaceOfBirth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBPlaceOfBirth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A public exchange listing for the business.
//
// The properties MarketIdentifierCode, StockNumber, Ticker are required.
type KYBPubliclyTradedListing struct {
	// ISO 10383 market identifier code of the listing venue.
	MarketIdentifierCode string `json:"market_identifier_code" api:"required"`
	// ISIN with dashes removed.
	StockNumber string `json:"stock_number" api:"required"`
	// Exchange ticker symbol.
	Ticker string `json:"ticker" api:"required"`
	paramObj
}

func (r KYBPubliclyTradedListing) MarshalJSON() (data []byte, err error) {
	type shadow KYBPubliclyTradedListing
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBPubliclyTradedListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the regulated activity a business is licensed to perform.
//
// The properties LicenseNumber, PrimaryRegulatoryAuthorityCountry,
// PrimaryRegulatoryAuthorityName, RegulatedActivitiesDescription are required.
type KYBRegulatedActivity struct {
	// License number issued by the regulator.
	LicenseNumber string `json:"license_number" api:"required"`
	// ISO 3166-1 alpha-3 country code of the primary regulator.
	PrimaryRegulatoryAuthorityCountry string `json:"primary_regulatory_authority_country" api:"required"`
	// Name of the primary regulator.
	PrimaryRegulatoryAuthorityName string `json:"primary_regulatory_authority_name" api:"required"`
	// Description of the regulated activities performed.
	RegulatedActivitiesDescription string `json:"regulated_activities_description" api:"required"`
	paramObj
}

func (r KYBRegulatedActivity) MarshalJSON() (data []byte, err error) {
	type shadow KYBRegulatedActivity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBRegulatedActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KYBSourceOfFunds = string

// List of KYB status snapshots, one per configured provider/environment.
type KYBStatusListResponse struct {
	KYBStatuses []KYBStatusResponse `json:"kyb_statuses" api:"required"`
	NextCursor  string              `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KYBStatuses respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYBStatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *KYBStatusListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full KYB status for an organization with a given provider.
type KYBStatusResponse struct {
	// Capability statuses for the customer.
	Capabilities KyxCapabilities  `json:"capabilities" api:"required"`
	Endorsements []KyxEndorsement `json:"endorsements" api:"required"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment" api:"required"`
	// Items that will be required in the future.
	FutureRequirementsDue []string `json:"future_requirements_due" api:"required"`
	// Verification status detail for a KYC or KYB check.
	KYB KyxVerificationStatusDetail `json:"kyb" api:"required"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider" api:"required"`
	// Top-level items still needed (e.g. link a bank account).
	RequirementsDue []string `json:"requirements_due" api:"required"`
	// KYC/KYB status for the user.
	Status KyxProviderStatus `json:"status" api:"required"`
	// Terms of Service acceptance status for a KYC or KYB flow.
	Tos KyxTosStatusDetail `json:"tos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities          respjson.Field
		Endorsements          respjson.Field
		Environment           respjson.Field
		FutureRequirementsDue respjson.Field
		KYB                   respjson.Field
		Provider              respjson.Field
		RequirementsDue       respjson.Field
		Status                respjson.Field
		Tos                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYBStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *KYBStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYB verification data for headless submission. Fields are individually optional
// because the provider accepts partial submissions and grants endorsements once
// enough data has arrived; a partial submission can be completed by calling the
// endpoint again.
type KYBSubmitData struct {
	// Primary purpose the business will use the account for. Passthrough to the
	// provider. See the
	// [Bridge customer API reference](https://apidocs.bridge.xyz/platform/customers/customers/api)
	// for accepted values.
	AccountPurpose param.Opt[KYBAccountPurpose] `json:"account_purpose,omitzero"`
	// Free-text purpose. Required when account_purpose is "other".
	AccountPurposeOther param.Opt[string] `json:"account_purpose_other,omitzero"`
	// Whether the business moves funds on behalf of third parties.
	ActingAsIntermediary param.Opt[bool] `json:"acting_as_intermediary,omitzero"`
	// Short summary of what the business does.
	BusinessDescription param.Opt[string] `json:"business_description,omitzero"`
	// Registered legal name as filed with government authorities.
	BusinessLegalName param.Opt[string] `json:"business_legal_name,omitzero"`
	// Public trading name (DBA), if different from the legal name.
	BusinessTradeName param.Opt[string] `json:"business_trade_name,omitzero"`
	// Legal structure of the business. Passthrough to the provider. See the
	// [Bridge customer API reference](https://apidocs.bridge.xyz/platform/customers/customers/api)
	// for accepted values.
	BusinessType param.Opt[KYBBusinessType] `json:"business_type,omitzero"`
	// Description of the AML and sanctions screening controls in place.
	ComplianceScreeningExplanation param.Opt[string] `json:"compliance_screening_explanation,omitzero"`
	// Whether the business conducts money services.
	ConductsMoneyServices param.Opt[bool] `json:"conducts_money_services,omitzero"`
	// Description of the money services conducted.
	ConductsMoneyServicesDescription param.Opt[string] `json:"conducts_money_services_description,omitzero"`
	// Whether money services are conducted through the provider. Requires a
	// flow_of_funds document when true.
	ConductsMoneyServicesUsingBridge param.Opt[bool] `json:"conducts_money_services_using_bridge,omitzero"`
	// Primary business email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Estimated annual revenue of the business, in USD buckets. Passthrough to the
	// provider. See the
	// [Bridge customer API reference](https://apidocs.bridge.xyz/platform/customers/customers/api)
	// for accepted values.
	EstimatedAnnualRevenueUsd param.Opt[KYBEstimatedAnnualRevenue] `json:"estimated_annual_revenue_usd,omitzero"`
	// Expected monthly payment volume in USD. Required for high-risk businesses.
	ExpectedMonthlyPaymentsUsd param.Opt[int64] `json:"expected_monthly_payments_usd,omitzero"`
	// Whether the business is tax-registered outside its country of incorporation.
	HasForeignTaxRegistration param.Opt[bool] `json:"has_foreign_tax_registration,omitzero"`
	// Whether an intermediate entity owner holds 25% or more of the business.
	HasMaterialIntermediaryOwnership param.Opt[bool] `json:"has_material_intermediary_ownership,omitzero"`
	// Explanation of the high-risk activities. Required unless the only value is
	// "none_of_the_above".
	HighRiskActivitiesExplanation param.Opt[string] `json:"high_risk_activities_explanation,omitzero"`
	// Date of incorporation in YYYY-MM-DD format.
	IncorporationDate param.Opt[string] `json:"incorporation_date,omitzero"`
	// Whether the business is a decentralized autonomous organization.
	IsDao param.Opt[bool] `json:"is_dao,omitzero"`
	// Whether the business operates in prohibited jurisdictions.
	OperatesInProhibitedCountries param.Opt[bool] `json:"operates_in_prohibited_countries,omitzero"`
	// Ownership percentage at which a person is treated as a beneficial owner.
	OwnershipThreshold param.Opt[int64] `json:"ownership_threshold,omitzero"`
	// Business phone number in E.164 format.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Primary website. If omitted, a proof_of_nature_of_business document is required.
	PrimaryWebsite param.Opt[string] `json:"primary_website,omitzero"`
	// Primary source of the funds the business will transact with. Passthrough to the
	// provider. See the
	// [Bridge customer API reference](https://apidocs.bridge.xyz/platform/customers/customers/api)
	// for accepted values.
	SourceOfFunds param.Opt[KYBSourceOfFunds] `json:"source_of_funds,omitzero"`
	// Free-text detail on the source of funds. Required for high-risk businesses.
	SourceOfFundsDescription param.Opt[string] `json:"source_of_funds_description,omitzero"`
	// Latin-1 transliteration of the legal name. Required for non-Latin-1 names.
	TransliteratedBusinessLegalName param.Opt[string] `json:"transliterated_business_legal_name,omitzero"`
	// Latin-1 transliteration of the trade name. Required for non-Latin-1 names.
	TransliteratedBusinessTradeName param.Opt[string] `json:"transliterated_business_trade_name,omitzero"`
	// Beneficial owners, control persons, and signers.
	AssociatedPersons []KYBAssociatedPerson `json:"associated_persons,omitzero"`
	// 2022 NAICS codes describing the industries the business operates in.
	BusinessIndustry []string `json:"business_industry,omitzero"`
	// Supporting documents for verification.
	Documents []KYBBusinessDocument `json:"documents,omitzero"`
	// High-risk activities the business engages in.
	HighRiskActivities []KYBHighRiskActivity `json:"high_risk_activities,omitzero"`
	// Business tax and registration identifiers.
	IdentifyingInformation []VerificationDocument `json:"identifying_information,omitzero"`
	// Additional websites and social handles.
	OtherWebsites []string `json:"other_websites,omitzero"`
	// A postal address used in KYC and KYB data submission.
	PhysicalAddress VerificationAddress `json:"physical_address,omitzero"`
	// Public exchange listings for the business.
	PubliclyTradedListings []KYBPubliclyTradedListing `json:"publicly_traded_listings,omitzero"`
	// A postal address used in KYC and KYB data submission.
	RegisteredAddress VerificationAddress `json:"registered_address,omitzero"`
	// Details of the regulated activity a business is licensed to perform.
	RegulatedActivity KYBRegulatedActivity `json:"regulated_activity,omitzero"`
	// A postal address used in KYC and KYB data submission.
	TransliteratedPhysicalAddress VerificationAddress `json:"transliterated_physical_address,omitzero"`
	// A postal address used in KYC and KYB data submission.
	TransliteratedRegisteredAddress VerificationAddress `json:"transliterated_registered_address,omitzero"`
	paramObj
}

func (r KYBSubmitData) MarshalJSON() (data []byte, err error) {
	type shadow KYBSubmitData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBSubmitData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for headless KYB data submission.
//
// The properties Data, Provider are required.
type KYBSubmitRequestBody struct {
	// KYB verification data for headless submission. Fields are individually optional
	// because the provider accepts partial submissions and grants endorsements once
	// enough data has arrived; a partial submission can be completed by calling the
	// endpoint again.
	Data KYBSubmitData `json:"data,omitzero" api:"required"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider,omitzero" api:"required"`
	// Client-side agreement ID for ToS acceptance.
	ClientAgreementID param.Opt[string] `json:"client_agreement_id,omitzero"`
	// Endorsements to request during KYB.
	Endorsements []KyxEndorsementName `json:"endorsements,omitzero"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment,omitzero"`
	paramObj
}

func (r KYBSubmitRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow KYBSubmitRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBSubmitRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for initiating Terms of Service acceptance for an organization.
//
// The properties Email, Provider are required.
type KYBTosRequestBody struct {
	// Email address for the organization.
	Email string `json:"email" api:"required" format:"email"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider,omitzero" api:"required"`
	// Legal name of the business.
	BusinessName param.Opt[string] `json:"business_name,omitzero"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment,omitzero"`
	paramObj
}

func (r KYBTosRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow KYBTosRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYBTosRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for initiating a hosted KYC flow.
//
// The property Provider is required.
type KYCLinksRequestBody struct {
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider,omitzero" api:"required"`
	// Client-side agreement ID for ToS acceptance.
	ClientAgreementID param.Opt[string] `json:"client_agreement_id,omitzero"`
	// Email address for the KYC session.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// URI to redirect the user after completing KYC.
	RedirectUri param.Opt[string] `json:"redirect_uri,omitzero" format:"uri"`
	// Endorsements to request during KYC.
	Endorsements []KyxEndorsementName `json:"endorsements,omitzero"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment,omitzero"`
	paramObj
}

func (r KYCLinksRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow KYCLinksRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYCLinksRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of KYC status snapshots, one per configured provider/environment.
type KYCStatusListResponse struct {
	KYCStatuses []KYCStatusResponse `json:"kyc_statuses" api:"required"`
	NextCursor  string              `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KYCStatuses respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCStatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *KYCStatusListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full KYC status for a user with a given provider.
type KYCStatusResponse struct {
	// Capability statuses for the customer.
	Capabilities KyxCapabilities  `json:"capabilities" api:"required"`
	Endorsements []KyxEndorsement `json:"endorsements" api:"required"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment" api:"required"`
	// Items that will be required in the future.
	FutureRequirementsDue []string `json:"future_requirements_due" api:"required"`
	// Verification status detail for a KYC or KYB check.
	KYC KyxVerificationStatusDetail `json:"kyc" api:"required"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider" api:"required"`
	// Top-level items still needed (e.g. link a bank account).
	RequirementsDue []string `json:"requirements_due" api:"required"`
	// KYC/KYB status for the user.
	Status KyxProviderStatus `json:"status" api:"required"`
	// Terms of Service acceptance status for a KYC or KYB flow.
	Tos KyxTosStatusDetail `json:"tos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities          respjson.Field
		Endorsements          respjson.Field
		Environment           respjson.Field
		FutureRequirementsDue respjson.Field
		KYC                   respjson.Field
		Provider              respjson.Field
		RequirementsDue       respjson.Field
		Status                respjson.Field
		Tos                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *KYCStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC verification data for headless submission.
type KYCSubmitData struct {
	// Date of birth in YYYY-MM-DD format.
	DateOfBirth param.Opt[string] `json:"date_of_birth,omitzero"`
	// Email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Legal first name.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Legal last name.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	// Phone number in E.164 format.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Identifying documents.
	IdentifyingInformation []VerificationDocument `json:"identifying_information,omitzero"`
	// A postal address used in KYC and KYB data submission.
	ResidentialAddress VerificationAddress `json:"residential_address,omitzero"`
	paramObj
}

func (r KYCSubmitData) MarshalJSON() (data []byte, err error) {
	type shadow KYCSubmitData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYCSubmitData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for headless KYC data submission.
//
// The properties Data, Provider are required.
type KYCSubmitRequestBody struct {
	// KYC verification data for headless submission.
	Data KYCSubmitData `json:"data,omitzero" api:"required"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider,omitzero" api:"required"`
	// Client-side agreement ID for ToS acceptance.
	ClientAgreementID param.Opt[string] `json:"client_agreement_id,omitzero"`
	// Endorsements to request during KYC.
	Endorsements []KyxEndorsementName `json:"endorsements,omitzero"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment,omitzero"`
	paramObj
}

func (r KYCSubmitRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow KYCSubmitRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYCSubmitRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Capability statuses for the customer.
type KyxCapabilities struct {
	// Status of a capability. Passthrough from the provider.
	PayinCrypto KyxCapabilityStatus `json:"payin_crypto" api:"required"`
	// Status of a capability. Passthrough from the provider.
	PayinFiat KyxCapabilityStatus `json:"payin_fiat" api:"required"`
	// Status of a capability. Passthrough from the provider.
	PayoutCrypto KyxCapabilityStatus `json:"payout_crypto" api:"required"`
	// Status of a capability. Passthrough from the provider.
	PayoutFiat KyxCapabilityStatus `json:"payout_fiat" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayinCrypto  respjson.Field
		PayinFiat    respjson.Field
		PayoutCrypto respjson.Field
		PayoutFiat   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KyxCapabilities) RawJSON() string { return r.JSON.raw }
func (r *KyxCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KyxCapabilityStatus = string

// An endorsement with its approval status and missing requirements.
type KyxEndorsement struct {
	// Missing requirements, or null if complete.
	Missing []string `json:"missing" api:"required"`
	// Endorsement identifier.
	Name KyxEndorsementName `json:"name" api:"required"`
	// Status of an endorsement. Passthrough from the provider.
	Status KyxEndorsementStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Missing     respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KyxEndorsement) RawJSON() string { return r.JSON.raw }
func (r *KyxEndorsement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KyxEndorsementName = string

type KyxEndorsementStatus = string

// Provider environment (production or sandbox).
type KyxEnvironment string

const (
	KyxEnvironmentProduction KyxEnvironment = "production"
	KyxEnvironmentSandbox    KyxEnvironment = "sandbox"
)

// KYC/KYB provider identifier.
type KyxProvider string

const (
	KyxProviderBridge KyxProvider = "bridge"
)

type KyxProviderStatus = string

// Request body for initiating Terms of Service acceptance.
//
// The property Provider is required.
type KyxTosRequestBody struct {
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider,omitzero" api:"required"`
	// Email for the user. If not provided, falls back to the user's linked email.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment,omitzero"`
	paramObj
}

func (r KyxTosRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow KyxTosRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KyxTosRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a Terms of Service link.
type KyxTosResponse struct {
	// Provider environment (production or sandbox).
	//
	// Any of "production", "sandbox".
	Environment KyxEnvironment `json:"environment" api:"required"`
	// URL for the Terms of Service acceptance page.
	Link string `json:"link" api:"required" format:"uri"`
	// KYC/KYB provider identifier.
	//
	// Any of "bridge".
	Provider KyxProvider `json:"provider" api:"required"`
	// Status of Terms of Service acceptance. Passthrough from the provider.
	Status KyxTosStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Environment respjson.Field
		Link        respjson.Field
		Provider    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KyxTosResponse) RawJSON() string { return r.JSON.raw }
func (r *KyxTosResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KyxTosStatus = string

// Terms of Service acceptance status for a KYC or KYB flow.
type KyxTosStatusDetail struct {
	// Status of Terms of Service acceptance. Passthrough from the provider.
	Status KyxTosStatus `json:"status" api:"required"`
	// ToS acceptance link, if pending.
	Link string `json:"link" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KyxTosStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *KyxTosStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KyxVerificationStatus = string

// Verification status detail for a KYC or KYB check.
type KyxVerificationStatusDetail struct {
	// Status of KYC/KYB verification. Passthrough from the provider.
	Status KyxVerificationStatus `json:"status" api:"required"`
	// Verification link, if applicable.
	Link string `json:"link" format:"uri"`
	// Reasons for rejection, if status is closed or action_required.
	RejectionReasons []string `json:"rejection_reasons"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status           respjson.Field
		Link             respjson.Field
		RejectionReasons respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KyxVerificationStatusDetail) RawJSON() string { return r.JSON.raw }
func (r *KyxVerificationStatusDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of external fiat accounts linked to a user.
type ListExternalFiatAccountsResponse struct {
	ExternalFiatAccounts []ExternalFiatAccount `json:"external_fiat_accounts" api:"required"`
	NextCursor           string                `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalFiatAccounts respjson.Field
		NextCursor           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListExternalFiatAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListExternalFiatAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of fiat deposit accounts linked to a wallet.
type ListFiatDepositAccountsResponse struct {
	FiatDepositAccounts []FiatDepositAccount `json:"fiat_deposit_accounts" api:"required"`
	NextCursor          string               `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FiatDepositAccounts respjson.Field
		NextCursor          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFiatDepositAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListFiatDepositAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of external fiat accounts linked to an organization.
type ListOrganizationExternalFiatAccountsResponse struct {
	ExternalFiatAccounts []OrganizationExternalFiatAccount `json:"external_fiat_accounts" api:"required"`
	NextCursor           string                            `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalFiatAccounts respjson.Field
		NextCursor           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListOrganizationExternalFiatAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListOrganizationExternalFiatAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Bridge external fiat account linked to an organization.
type OrganizationExternalFiatAccount struct {
	ID               string `json:"id" api:"required"`
	AccountOwnerName string `json:"account_owner_name" api:"required"`
	AccountType      string `json:"account_type" api:"required"`
	CreatedAt        string `json:"created_at" api:"required"`
	Currency         string `json:"currency" api:"required"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment    Environment `json:"environment" api:"required"`
	OrganizationID string      `json:"organization_id" api:"required"`
	// Discriminator: the external fiat account is orchestrated via Bridge.
	//
	// Any of "bridge".
	Provider OrganizationExternalFiatAccountProvider `json:"provider" api:"required"`
	BankName string                                  `json:"bank_name"`
	Last4    string                                  `json:"last_4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountOwnerName respjson.Field
		AccountType      respjson.Field
		CreatedAt        respjson.Field
		Currency         respjson.Field
		Environment      respjson.Field
		OrganizationID   respjson.Field
		Provider         respjson.Field
		BankName         respjson.Field
		Last4            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationExternalFiatAccount) RawJSON() string { return r.JSON.raw }
func (r *OrganizationExternalFiatAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator: the external fiat account is orchestrated via Bridge.
type OrganizationExternalFiatAccountProvider string

const (
	OrganizationExternalFiatAccountProviderBridge OrganizationExternalFiatAccountProvider = "bridge"
)

// Response containing a single organization external fiat account.
type OrganizationExternalFiatAccountResponse struct {
	// A Bridge external fiat account linked to an organization.
	ExternalFiatAccount OrganizationExternalFiatAccount `json:"external_fiat_account" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalFiatAccount respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationExternalFiatAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationExternalFiatAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The destination bank account for a payout.
type PayoutDestinationResp struct {
	// The ID of a previously registered external fiat account to pay out to.
	FiatAccountID string `json:"fiat_account_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FiatAccountID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutDestinationResp) RawJSON() string { return r.JSON.raw }
func (r *PayoutDestinationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PayoutDestinationResp to a PayoutDestination.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PayoutDestination.Overrides()
func (r PayoutDestinationResp) ToParam() PayoutDestination {
	return param.Override[PayoutDestination](json.RawMessage(r.RawJSON()))
}

// The destination bank account for a payout.
//
// The property FiatAccountID is required.
type PayoutDestination struct {
	// The ID of a previously registered external fiat account to pay out to.
	FiatAccountID string `json:"fiat_account_id" api:"required"`
	paramObj
}

func (r PayoutDestination) MarshalJSON() (data []byte, err error) {
	type shadow PayoutDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source crypto asset, chain, and amount for a payout.
type PayoutSourceResp struct {
	// Amount to offramp, in the asset's standard units (e.g. "100.00").
	Amount string `json:"amount" api:"required"`
	// Source crypto asset (e.g. "usdc").
	Asset string `json:"asset" api:"required"`
	// Source chain (e.g. "base").
	Chain string `json:"chain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Asset       respjson.Field
		Chain       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutSourceResp) RawJSON() string { return r.JSON.raw }
func (r *PayoutSourceResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PayoutSourceResp to a PayoutSource.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PayoutSource.Overrides()
func (r PayoutSourceResp) ToParam() PayoutSource {
	return param.Override[PayoutSource](json.RawMessage(r.RawJSON()))
}

// The source crypto asset, chain, and amount for a payout.
//
// The properties Amount, Asset, Chain are required.
type PayoutSource struct {
	// Amount to offramp, in the asset's standard units (e.g. "100.00").
	Amount string `json:"amount" api:"required"`
	// Source crypto asset (e.g. "usdc").
	Asset string `json:"asset" api:"required"`
	// Source chain (e.g. "base").
	Chain string `json:"chain" api:"required"`
	paramObj
}

func (r PayoutSource) MarshalJSON() (data []byte, err error) {
	type shadow PayoutSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayoutSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A postal address used in KYC and KYB data submission.
//
// The properties City, Country, StreetLine1 are required.
type VerificationAddress struct {
	// City.
	City string `json:"city" api:"required"`
	// ISO 3166-1 alpha-3 country code.
	Country string `json:"country" api:"required"`
	// Street address line 1.
	StreetLine1 string `json:"street_line_1" api:"required"`
	// Postal code. Required for countries that use them.
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// Street address line 2.
	StreetLine2 param.Opt[string] `json:"street_line_2,omitzero"`
	// ISO 3166-2 state or province code. Required for US addresses.
	Subdivision param.Opt[string] `json:"subdivision,omitzero"`
	paramObj
}

func (r VerificationAddress) MarshalJSON() (data []byte, err error) {
	type shadow VerificationAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerificationAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An identifying document for KYC or KYB verification. Also used for business
// identifiers such as tax and registration numbers, for which the image and
// expiration fields do not apply.
//
// The properties IssuingCountry, Type are required.
type VerificationDocument struct {
	// ISO 3166-1 alpha-3 issuing country code.
	IssuingCountry string `json:"issuing_country" api:"required"`
	// Document type identifier.
	Type string `json:"type" api:"required"`
	// Document description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Document expiration date.
	Expiration param.Opt[string] `json:"expiration,omitzero"`
	// Base64-encoded back image.
	ImageBack param.Opt[string] `json:"image_back,omitzero"`
	// Base64-encoded front image.
	ImageFront param.Opt[string] `json:"image_front,omitzero"`
	// Document number.
	Number param.Opt[string] `json:"number,omitzero"`
	paramObj
}

func (r VerificationDocument) MarshalJSON() (data []byte, err error) {
	type shadow VerificationDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerificationDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
