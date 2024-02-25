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

type FhirSpecimen struct {
	models.ResourceBase
	// The accession number associated with the specimen
	// https://hl7.org/fhir/r4/search.html#token
	Accession datatypes.JSON `gorm:"column:accession;type:text;serializer:json" json:"accession,omitempty"`
	// The code for the body site from where the specimen originated
	// https://hl7.org/fhir/r4/search.html#token
	Bodysite datatypes.JSON `gorm:"column:bodysite;type:text;serializer:json" json:"bodysite,omitempty"`
	// The date the specimen was collected
	// https://hl7.org/fhir/r4/search.html#date
	Collected *time.Time `gorm:"column:collected;type:datetime" json:"collected,omitempty"`
	// Who collected the specimen
	// https://hl7.org/fhir/r4/search.html#reference
	Collector datatypes.JSON `gorm:"column:collector;type:text;serializer:json" json:"collector,omitempty"`
	// The kind of specimen container
	// https://hl7.org/fhir/r4/search.html#token
	Container datatypes.JSON `gorm:"column:container;type:text;serializer:json" json:"container,omitempty"`
	// The unique identifier associated with the specimen container
	// https://hl7.org/fhir/r4/search.html#token
	ContainerId datatypes.JSON `gorm:"column:containerId;type:text;serializer:json" json:"containerId,omitempty"`
	// The unique identifier associated with the specimen
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
	// The parent of the specimen
	// https://hl7.org/fhir/r4/search.html#reference
	Parent datatypes.JSON `gorm:"column:parent;type:text;serializer:json" json:"parent,omitempty"`
	// available | unavailable | unsatisfactory | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The subject of the specimen
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// The specimen type
	// https://hl7.org/fhir/r4/search.html#token
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirSpecimen) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"accession":            "token",
		"bodysite":             "token",
		"collected":            "date",
		"collector":            "reference",
		"container":            "token",
		"containerId":          "token",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"parent":               "reference",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"subject":              "reference",
		"text":                 "keyword",
		"type":                 "token",
	}
	return searchParameters
}
func (s *FhirSpecimen) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Accession
	accessionResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Specimen.accessionIdentifier')")
	if err == nil && accessionResult.String() != "undefined" {
		s.Accession = []byte(accessionResult.String())
	}
	// extracting Bodysite
	bodysiteResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Specimen.collection.bodySite')")
	if err == nil && bodysiteResult.String() != "undefined" {
		s.Bodysite = []byte(bodysiteResult.String())
	}
	// extracting Collected
	collectedResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'Specimen.collection.collectedDateTime | Specimen.collection.collectedPeriod')")
	if err == nil && collectedResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, collectedResult.String()); err == nil {
			s.Collected = &t
		} else if t, err = time.Parse("2006-01-02", collectedResult.String()); err == nil {
			s.Collected = &t
		} else if t, err = time.Parse("2006-01", collectedResult.String()); err == nil {
			s.Collected = &t
		} else if t, err = time.Parse("2006", collectedResult.String()); err == nil {
			s.Collected = &t
		}
	}
	// extracting Collector
	collectorResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Specimen.collection.collector')")
	if err == nil && collectorResult.String() != "undefined" {
		s.Collector = []byte(collectorResult.String())
	}
	// extracting Container
	containerResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Specimen.container.type')")
	if err == nil && containerResult.String() != "undefined" {
		s.Container = []byte(containerResult.String())
	}
	// extracting ContainerId
	containerIdResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Specimen.container.identifier')")
	if err == nil && containerIdResult.String() != "undefined" {
		s.ContainerId = []byte(containerIdResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Specimen.identifier')")
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
	// extracting Parent
	parentResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Specimen.parent')")
	if err == nil && parentResult.String() != "undefined" {
		s.Parent = []byte(parentResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Specimen.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Specimen.subject')")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	// extracting Type
	typeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Specimen.type')")
	if err == nil && typeResult.String() != "undefined" {
		s.Type = []byte(typeResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirSpecimen) TableName() string {
	return "fhir_specimen"
}
