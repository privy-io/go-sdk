// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/privy-io/go-sdk"
	"github.com/privy-io/go-sdk/internal/testutil"
	"github.com/privy-io/go-sdk/option"
)

func TestOrganizationKYBList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Organizations.KYB.List(context.TODO(), "organization_id")
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationKYBInitiateLinksWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Organizations.KYB.InitiateLinks(
		context.TODO(),
		"organization_id",
		privyclient.OrganizationKYBInitiateLinksParams{
			KYBLinksRequestBody: privyclient.KYBLinksRequestBody{
				Email:             "dev@stainless.com",
				Provider:          privyclient.KyxProviderBridge,
				BusinessName:      privyclient.String("x"),
				ClientAgreementID: privyclient.String("x"),
				Endorsements:      []privyclient.KyxEndorsementName{"sepa"},
				Environment:       privyclient.KyxEnvironmentProduction,
				RedirectUri:       privyclient.String("https://example.com"),
			},
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationKYBInitiateTosWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Organizations.KYB.InitiateTos(
		context.TODO(),
		"organization_id",
		privyclient.OrganizationKYBInitiateTosParams{
			KYBTosRequestBody: privyclient.KYBTosRequestBody{
				Email:        "dev@stainless.com",
				Provider:     privyclient.KyxProviderBridge,
				BusinessName: privyclient.String("x"),
				Environment:  privyclient.KyxEnvironmentProduction,
			},
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationKYBSubmitWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := privyclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAppID("My App ID"),
		option.WithAppSecret("My App Secret"),
	)
	_, err := client.Organizations.KYB.Submit(
		context.TODO(),
		"organization_id",
		privyclient.OrganizationKYBSubmitParams{
			KYBSubmitRequestBody: privyclient.KYBSubmitRequestBody{
				Data: privyclient.KYBSubmitData{
					AccountPurpose:       privyclient.String("treasury_management"),
					AccountPurposeOther:  privyclient.String("x"),
					ActingAsIntermediary: privyclient.Bool(true),
					AssociatedPersons: []privyclient.KYBAssociatedPerson{{
						DateOfBirth:  "7321-69-10",
						Email:        "dev@stainless.com",
						FirstName:    "x",
						HasControl:   true,
						HasOwnership: true,
						IdentifyingInformation: []privyclient.VerificationDocument{{
							IssuingCountry: "xxx",
							Type:           "type",
							Description:    privyclient.String("description"),
							Expiration:     privyclient.String("expiration"),
							ImageBack:      privyclient.String("image_back"),
							ImageFront:     privyclient.String("image_front"),
							Number:         privyclient.String("number"),
						}},
						IsSigner: true,
						LastName: "xx",
						ResidentialAddress: privyclient.VerificationAddress{
							City:        "x",
							Country:     "xxx",
							StreetLine1: "xxxx",
							PostalCode:  privyclient.String("x"),
							StreetLine2: privyclient.String("x"),
							Subdivision: privyclient.String("x"),
						},
						Documents: []privyclient.KYBIndividualDocument{{
							File:        "x",
							Purposes:    []privyclient.KYBIndividualDocumentPurpose{"proof_of_address"},
							Description: privyclient.String("x"),
						}},
						IsDirector:          privyclient.Bool(true),
						MiddleName:          privyclient.String("x"),
						Nationalities:       []string{"xxx"},
						OwnershipPercentage: privyclient.Int(0),
						Phone:               privyclient.String("phone"),
						PlaceOfBirth: privyclient.KYBPlaceOfBirth{
							Country: "xxx",
							City:    privyclient.String("x"),
						},
						RelationshipEstablishedAt: privyclient.String("7321-69-10"),
						Title:                     privyclient.String("x"),
						TransliteratedFirstName:   privyclient.String("x"),
						TransliteratedLastName:    privyclient.String("x"),
						TransliteratedMiddleName:  privyclient.String("x"),
						TransliteratedResidentialAddress: privyclient.VerificationAddress{
							City:        "x",
							Country:     "xxx",
							StreetLine1: "xxxx",
							PostalCode:  privyclient.String("x"),
							StreetLine2: privyclient.String("x"),
							Subdivision: privyclient.String("x"),
						},
					}},
					BusinessDescription:              privyclient.String("x"),
					BusinessIndustry:                 []string{"x"},
					BusinessLegalName:                privyclient.String("x"),
					BusinessTradeName:                privyclient.String("x"),
					BusinessType:                     privyclient.String("llc"),
					ComplianceScreeningExplanation:   privyclient.String("x"),
					ConductsMoneyServices:            privyclient.Bool(true),
					ConductsMoneyServicesDescription: privyclient.String("x"),
					ConductsMoneyServicesUsingBridge: privyclient.Bool(true),
					Documents: []privyclient.KYBBusinessDocument{{
						File:        "x",
						Purposes:    []privyclient.KYBDocumentPurpose{"business_formation"},
						Description: privyclient.String("x"),
					}},
					Email:                            privyclient.String("dev@stainless.com"),
					EstimatedAnnualRevenueUsd:        privyclient.String("1000000_9999999"),
					ExpectedMonthlyPaymentsUsd:       privyclient.Int(0),
					HasForeignTaxRegistration:        privyclient.Bool(true),
					HasMaterialIntermediaryOwnership: privyclient.Bool(true),
					HighRiskActivities:               []privyclient.KYBHighRiskActivity{"none_of_the_above"},
					HighRiskActivitiesExplanation:    privyclient.String("x"),
					IdentifyingInformation: []privyclient.VerificationDocument{{
						IssuingCountry: "xxx",
						Type:           "type",
						Description:    privyclient.String("description"),
						Expiration:     privyclient.String("expiration"),
						ImageBack:      privyclient.String("image_back"),
						ImageFront:     privyclient.String("image_front"),
						Number:         privyclient.String("number"),
					}},
					IncorporationDate:             privyclient.String("7321-69-10"),
					IsDao:                         privyclient.Bool(true),
					OperatesInProhibitedCountries: privyclient.Bool(true),
					OtherWebsites:                 []string{"string"},
					OwnershipThreshold:            privyclient.Int(5),
					Phone:                         privyclient.String("phone"),
					PhysicalAddress: privyclient.VerificationAddress{
						City:        "x",
						Country:     "xxx",
						StreetLine1: "xxxx",
						PostalCode:  privyclient.String("x"),
						StreetLine2: privyclient.String("x"),
						Subdivision: privyclient.String("x"),
					},
					PrimaryWebsite: privyclient.String("primary_website"),
					PubliclyTradedListings: []privyclient.KYBPubliclyTradedListing{{
						MarketIdentifierCode: "xxxx",
						StockNumber:          "x",
						Ticker:               "x",
					}},
					RegisteredAddress: privyclient.VerificationAddress{
						City:        "x",
						Country:     "xxx",
						StreetLine1: "xxxx",
						PostalCode:  privyclient.String("x"),
						StreetLine2: privyclient.String("x"),
						Subdivision: privyclient.String("x"),
					},
					RegulatedActivity: privyclient.KYBRegulatedActivity{
						LicenseNumber:                     "x",
						PrimaryRegulatoryAuthorityCountry: "xxx",
						PrimaryRegulatoryAuthorityName:    "x",
						RegulatedActivitiesDescription:    "x",
					},
					SourceOfFunds:                   privyclient.String("sales_of_goods_and_services"),
					SourceOfFundsDescription:        privyclient.String("x"),
					TransliteratedBusinessLegalName: privyclient.String("x"),
					TransliteratedBusinessTradeName: privyclient.String("x"),
					TransliteratedPhysicalAddress: privyclient.VerificationAddress{
						City:        "x",
						Country:     "xxx",
						StreetLine1: "xxxx",
						PostalCode:  privyclient.String("x"),
						StreetLine2: privyclient.String("x"),
						Subdivision: privyclient.String("x"),
					},
					TransliteratedRegisteredAddress: privyclient.VerificationAddress{
						City:        "x",
						Country:     "xxx",
						StreetLine1: "xxxx",
						PostalCode:  privyclient.String("x"),
						StreetLine2: privyclient.String("x"),
						Subdivision: privyclient.String("x"),
					},
				},
				Provider:          privyclient.KyxProviderBridge,
				ClientAgreementID: privyclient.String("x"),
				Endorsements:      []privyclient.KyxEndorsementName{"sepa"},
				Environment:       privyclient.KyxEnvironmentProduction,
			},
		},
	)
	if err != nil {
		var apierr *privyclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
