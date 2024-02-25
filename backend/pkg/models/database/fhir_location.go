// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirLocation struct {
	models.ResourceBase
	// A (part of the) address of the location
	// https://hl7.org/fhir/r4/search.html#string
	Address datatypes.JSON `gorm:"column:address;type:text;serializer:json" json:"address,omitempty"`
	// A city specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressCity datatypes.JSON `gorm:"column:addressCity;type:text;serializer:json" json:"addressCity,omitempty"`
	// A country specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressCountry datatypes.JSON `gorm:"column:addressCountry;type:text;serializer:json" json:"addressCountry,omitempty"`
	// A postal code specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressPostalcode datatypes.JSON `gorm:"column:addressPostalcode;type:text;serializer:json" json:"addressPostalcode,omitempty"`
	// A state specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressState datatypes.JSON `gorm:"column:addressState;type:text;serializer:json" json:"addressState,omitempty"`
	// A use code specified in an address
	// https://hl7.org/fhir/r4/search.html#token
	AddressUse datatypes.JSON `gorm:"column:addressUse;type:text;serializer:json" json:"addressUse,omitempty"`
	// Technical endpoints providing access to services operated for the location
	// https://hl7.org/fhir/r4/search.html#reference
	Endpoint datatypes.JSON `gorm:"column:endpoint;type:text;serializer:json" json:"endpoint,omitempty"`
	// An identifier for the location
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	MetaLastUpdated *time.Time `gorm:"column:metaLastUpdated;type:datetime" json:"metaLastUpdated,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	MetaProfile datatypes.JSON `gorm:"column:metaProfile;type:text;serializer:json" json:"metaProfile,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	MetaTag datatypes.JSON `gorm:"column:metaTag;type:text;serializer:json" json:"metaTag,omitempty"`
	// Tags applied to this resource
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	MetaVersionId string `gorm:"column:metaVersionId;type:text" json:"metaVersionId,omitempty"`
	// A portion of the location's name or alias
	// https://hl7.org/fhir/r4/search.html#string
	Name datatypes.JSON `gorm:"column:name;type:text;serializer:json" json:"name,omitempty"`
	// Searches for locations (typically bed/room) that have an operational status (e.g. contaminated, housekeeping)
	// https://hl7.org/fhir/r4/search.html#token
	OperationalStatus datatypes.JSON `gorm:"column:operationalStatus;type:text;serializer:json" json:"operationalStatus,omitempty"`
	// Searches for locations that are managed by the provided organization
	// https://hl7.org/fhir/r4/search.html#reference
	Organization datatypes.JSON `gorm:"column:organization;type:text;serializer:json" json:"organization,omitempty"`
	// A location of which this location is a part
	// https://hl7.org/fhir/r4/search.html#reference
	Partof datatypes.JSON `gorm:"column:partof;type:text;serializer:json" json:"partof,omitempty"`
	// Searches for locations with a specific kind of status
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A code for the type of location
	// https://hl7.org/fhir/r4/search.html#token
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirLocation) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"address":              "string",
		"addressCity":          "string",
		"addressCountry":       "string",
		"addressPostalcode":    "string",
		"addressState":         "string",
		"addressUse":           "token",
		"endpoint":             "reference",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"name":                 "string",
		"operationalStatus":    "token",
		"organization":         "reference",
		"partof":               "reference",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"text":                 "keyword",
		"type":                 "token",
	}
	return searchParameters
}
func (s *FhirLocation) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// compile the searchParametersExtractor library
	searchParametersExtractorJsProgram, err := goja.Compile("searchParameterExtractor.js", searchParameterExtractorJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// add the searchParametersExtractor library in the goja vm
	_, err = vm.RunProgram(searchParametersExtractorJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Address
	addressResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Location.address')")
	if err == nil && addressResult.String() != "undefined" {
		s.Address = []byte(addressResult.String())
	}
	// extracting AddressCity
	addressCityResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Location.address.city')")
	if err == nil && addressCityResult.String() != "undefined" {
		s.AddressCity = []byte(addressCityResult.String())
	}
	// extracting AddressCountry
	addressCountryResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Location.address.country')")
	if err == nil && addressCountryResult.String() != "undefined" {
		s.AddressCountry = []byte(addressCountryResult.String())
	}
	// extracting AddressPostalcode
	addressPostalcodeResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Location.address.postalCode')")
	if err == nil && addressPostalcodeResult.String() != "undefined" {
		s.AddressPostalcode = []byte(addressPostalcodeResult.String())
	}
	// extracting AddressState
	addressStateResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Location.address.state')")
	if err == nil && addressStateResult.String() != "undefined" {
		s.AddressState = []byte(addressStateResult.String())
	}
	// extracting AddressUse
	addressUseResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Location.address.use')")
	if err == nil && addressUseResult.String() != "undefined" {
		s.AddressUse = []byte(addressUseResult.String())
	}
	// extracting Endpoint
	endpointResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Location.endpoint')")
	if err == nil && endpointResult.String() != "undefined" {
		s.Endpoint = []byte(endpointResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Location.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'meta.lastUpdated')")
	if err == nil && metaLastUpdatedResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01-02", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		}
	}
	// extracting MetaProfile
	metaProfileResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'meta.profile')")
	if err == nil && metaProfileResult.String() != "undefined" {
		s.MetaProfile = []byte(metaProfileResult.String())
	}
	// extracting MetaTag
	metaTagResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'meta.tag')")
	if err == nil && metaTagResult.String() != "undefined" {
		s.MetaTag = []byte(metaTagResult.String())
	}
	// extracting MetaVersionId
	metaVersionIdResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.versionId')")
	if err == nil && metaVersionIdResult.String() != "undefined" {
		s.MetaVersionId = metaVersionIdResult.String()
	}
	// extracting Name
	nameResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Location.name | Location.alias')")
	if err == nil && nameResult.String() != "undefined" {
		s.Name = []byte(nameResult.String())
	}
	// extracting OperationalStatus
	operationalStatusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Location.operationalStatus')")
	if err == nil && operationalStatusResult.String() != "undefined" {
		s.OperationalStatus = []byte(operationalStatusResult.String())
	}
	// extracting Organization
	organizationResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Location.managingOrganization')")
	if err == nil && organizationResult.String() != "undefined" {
		s.Organization = []byte(organizationResult.String())
	}
	// extracting Partof
	partofResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Location.partOf')")
	if err == nil && partofResult.String() != "undefined" {
		s.Partof = []byte(partofResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Location.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	// extracting Type
	typeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Location.type')")
	if err == nil && typeResult.String() != "undefined" {
		s.Type = []byte(typeResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirLocation) TableName() string {
	return "fhir_location"
}
