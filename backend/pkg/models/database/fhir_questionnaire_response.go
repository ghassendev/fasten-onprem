// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirQuestionnaireResponse struct {
	models.ResourceBase
	// The author of the questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	Author datatypes.JSON `gorm:"column:author;type:text;serializer:json" json:"author,omitempty"`
	// When the questionnaire response was last changed
	// https://hl7.org/fhir/r4/search.html#date
	Authored *time.Time `gorm:"column:authored;type:datetime" json:"authored,omitempty"`
	// Plan/proposal/order fulfilled by this questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// Encounter associated with the questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
	// The unique identifier for the questionnaire response
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// Procedure or observation this questionnaire response was performed as a part of
	// https://hl7.org/fhir/r4/search.html#reference
	PartOf datatypes.JSON `gorm:"column:partOf;type:text;serializer:json" json:"partOf,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// The questionnaire the answers are provided for
	// https://hl7.org/fhir/r4/search.html#reference
	Questionnaire datatypes.JSON `gorm:"column:questionnaire;type:text;serializer:json" json:"questionnaire,omitempty"`
	// The individual providing the information reflected in the questionnaire respose
	// https://hl7.org/fhir/r4/search.html#reference
	Source datatypes.JSON `gorm:"column:source;type:text;serializer:json" json:"source,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// The status of the questionnaire response
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The subject of the questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirQuestionnaireResponse) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"author":        "reference",
		"authored":      "date",
		"basedOn":       "reference",
		"encounter":     "reference",
		"identifier":    "token",
		"language":      "token",
		"lastUpdated":   "date",
		"partOf":        "reference",
		"profile":       "reference",
		"questionnaire": "reference",
		"source":        "reference",
		"sourceUri":     "uri",
		"status":        "token",
		"subject":       "reference",
		"tag":           "token",
		"text":          "string",
		"type":          "special",
	}
	return searchParameters
}
func (s *FhirQuestionnaireResponse) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Author
	authorResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.author'))")
	if err == nil && authorResult.String() != "undefined" {
	}
	// extracting Authored
	authoredResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.authored')[0]")
	if err == nil && authoredResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, authoredResult.String())
		if err == nil {
			s.Authored = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", authoredResult.String())
			if err == nil {
				s.Authored = &d
			}
		}
	}
	// extracting BasedOn
	basedOnResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.basedOn'))")
	if err == nil && basedOnResult.String() != "undefined" {
	}
	// extracting Encounter
	encounterResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.encounter'))")
	if err == nil && encounterResult.String() != "undefined" {
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.identifier')
							IdentifierProcessed = IdentifierResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(IdentifierProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(IdentifierProcessed)
							}
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'language')
							LanguageProcessed = LanguageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(LanguageProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LanguageProcessed)
							}
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", lastUpdatedResult.String())
			if err == nil {
				s.LastUpdated = &d
			}
		}
	}
	// extracting PartOf
	partOfResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.partOf'))")
	if err == nil && partOfResult.String() != "undefined" {
	}
	// extracting Profile
	profileResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'meta.profile'))")
	if err == nil && profileResult.String() != "undefined" {
	}
	// extracting Questionnaire
	questionnaireResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.questionnaire'))")
	if err == nil && questionnaireResult.String() != "undefined" {
	}
	// extracting Source
	sourceResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.source'))")
	if err == nil && sourceResult.String() != "undefined" {
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.status')
							StatusProcessed = StatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(StatusProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(StatusProcessed)
							}
						 `)
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'QuestionnaireResponse.subject'))")
	if err == nil && subjectResult.String() != "undefined" {
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'meta.tag')
							TagProcessed = TagResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(TagProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(TagProcessed)
							}
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirQuestionnaireResponse) TableName() string {
	return "fhir_questionnaire_response"
}
