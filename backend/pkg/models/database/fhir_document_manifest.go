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

type FhirDocumentManifest struct {
	models.ResourceBase
	// Who and/or what authored the DocumentManifest
	// https://hl7.org/fhir/r4/search.html#reference
	Author datatypes.JSON `gorm:"column:author;type:text;serializer:json" json:"author,omitempty"`
	// When this document manifest created
	// https://hl7.org/fhir/r4/search.html#date
	Created *time.Time `gorm:"column:created;type:datetime" json:"created,omitempty"`
	// Human-readable description (title)
	// https://hl7.org/fhir/r4/search.html#string
	Description datatypes.JSON `gorm:"column:description;type:text;serializer:json" json:"description,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): External ids for this item
	   * [CarePlan](careplan.html): External Ids for this plan
	   * [CareTeam](careteam.html): External Ids for this team
	   * [Composition](composition.html): Version-independent identifier for the Composition
	   * [Condition](condition.html): A unique identifier of the condition record
	   * [Consent](consent.html): Identifier for this record (external references)
	   * [DetectedIssue](detectedissue.html): Unique id for the detected issue
	   * [DeviceRequest](devicerequest.html): Business identifier for request/order
	   * [DiagnosticReport](diagnosticreport.html): An identifier for the report
	   * [DocumentManifest](documentmanifest.html): Unique Identifier for the set of documents
	   * [DocumentReference](documentreference.html): Master Version Specific Identifier
	   * [Encounter](encounter.html): Identifier(s) by which this encounter is known
	   * [EpisodeOfCare](episodeofcare.html): Business Identifier(s) relevant for this EpisodeOfCare
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a record identifier
	   * [Goal](goal.html): External Ids for this goal
	   * [ImagingStudy](imagingstudy.html): Identifiers for the Study, such as DICOM Study Instance UID and Accession number
	   * [Immunization](immunization.html): Business identifier
	   * [List](list.html): Business identifier
	   * [MedicationAdministration](medicationadministration.html): Return administrations with this external identifier
	   * [MedicationDispense](medicationdispense.html): Returns dispenses with this external identifier
	   * [MedicationRequest](medicationrequest.html): Return prescriptions with this external identifier
	   * [MedicationStatement](medicationstatement.html): Return statements with this external identifier
	   * [NutritionOrder](nutritionorder.html): Return nutrition orders with this external identifier
	   * [Observation](observation.html): The unique id for a particular observation
	   * [Procedure](procedure.html): A unique identifier for a procedure
	   * [RiskAssessment](riskassessment.html): Unique identifier for the assessment
	   * [ServiceRequest](servicerequest.html): Identifiers assigned to this order
	   * [SupplyDelivery](supplydelivery.html): External identifier
	   * [SupplyRequest](supplyrequest.html): Business Identifier for SupplyRequest
	   * [VisionPrescription](visionprescription.html): Return prescriptions with this external identifier
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Items in manifest
	// https://hl7.org/fhir/r4/search.html#reference
	Item datatypes.JSON `gorm:"column:item;type:text;serializer:json" json:"item,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Intended to get notified about this set of documents
	// https://hl7.org/fhir/r4/search.html#reference
	Recipient datatypes.JSON `gorm:"column:recipient;type:text;serializer:json" json:"recipient,omitempty"`
	// Identifiers of things that are related
	// https://hl7.org/fhir/r4/search.html#token
	RelatedId datatypes.JSON `gorm:"column:relatedId;type:text;serializer:json" json:"relatedId,omitempty"`
	// Related Resource
	// https://hl7.org/fhir/r4/search.html#reference
	RelatedRef datatypes.JSON `gorm:"column:relatedRef;type:text;serializer:json" json:"relatedRef,omitempty"`
	// The source system/application/software
	// https://hl7.org/fhir/r4/search.html#uri
	Source string `gorm:"column:source;type:text" json:"source,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// current | superseded | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The subject of the set of documents
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

func (s *FhirDocumentManifest) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"author":      "reference",
		"created":     "date",
		"description": "string",
		"identifier":  "token",
		"item":        "reference",
		"language":    "token",
		"lastUpdated": "date",
		"profile":     "reference",
		"recipient":   "reference",
		"relatedId":   "token",
		"relatedRef":  "reference",
		"source":      "uri",
		"sourceUri":   "uri",
		"status":      "token",
		"subject":     "reference",
		"tag":         "token",
		"text":        "string",
		"type":        "special",
	}
	return searchParameters
}
func (s *FhirDocumentManifest) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	authorResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'DocumentManifest.author'))")
	if err == nil && authorResult.String() != "undefined" {
	}
	// extracting Created
	createdResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'DocumentManifest.created')[0]")
	if err == nil && createdResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, createdResult.String())
		if err == nil {
			s.Created = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", createdResult.String())
			if err == nil {
				s.Created = &d
			}
		}
	}
	// extracting Description
	descriptionResult, err := vm.RunString(` 
							DescriptionResult = window.fhirpath.evaluate(fhirResource, 'DocumentManifest.description')
							DescriptionProcessed = DescriptionResult.reduce((accumulator, currentValue) => {
								if (typeof currentValue === 'string') {
									//basic string
									accumulator.push(currentValue)
								} else if (currentValue.family  || currentValue.given) {
									//HumanName http://hl7.org/fhir/R4/datatypes.html#HumanName
									var humanNameParts = []
									if (currentValue.prefix) {
										humanNameParts = humanNameParts.concat(currentValue.prefix)
									}
									if (currentValue.given) {	
										humanNameParts = humanNameParts.concat(currentValue.given)
									}	
									if (currentValue.family) {	
										humanNameParts.push(currentValue.family)	
									}	
									if (currentValue.suffix) {	
										humanNameParts = humanNameParts.concat(currentValue.suffix)	
									}
									accumulator.push(humanNameParts.join(" "))
								} else if (currentValue.city || currentValue.state || currentValue.country || currentValue.postalCode) {
									//Address http://hl7.org/fhir/R4/datatypes.html#Address
									var addressParts = []		
									if (currentValue.line) {
										addressParts = addressParts.concat(currentValue.line)
									}
									if (currentValue.city) {
										addressParts.push(currentValue.city)
									}	
									if (currentValue.state) {	
										addressParts.push(currentValue.state)
									}	
									if (currentValue.postalCode) {
										addressParts.push(currentValue.postalCode)
									}	
									if (currentValue.country) {
										addressParts.push(currentValue.country)	
									}	
									accumulator.push(addressParts.join(" "))
								} else {
									//string, boolean
									accumulator.push(currentValue)
								}
								return accumulator
							}, [])
						
							if(DescriptionProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(DescriptionProcessed)
							}
						 `)
	if err == nil && descriptionResult.String() != "undefined" {
		s.Description = []byte(descriptionResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')
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
	// extracting Item
	itemResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'DocumentManifest.content'))")
	if err == nil && itemResult.String() != "undefined" {
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
	// extracting Profile
	profileResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'meta.profile'))")
	if err == nil && profileResult.String() != "undefined" {
	}
	// extracting Recipient
	recipientResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'DocumentManifest.recipient'))")
	if err == nil && recipientResult.String() != "undefined" {
	}
	// extracting RelatedId
	relatedIdResult, err := vm.RunString(` 
							RelatedIdResult = window.fhirpath.evaluate(fhirResource, 'DocumentManifest.related.identifier')
							RelatedIdProcessed = RelatedIdResult.reduce((accumulator, currentValue) => {
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
						
				
							if(RelatedIdProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(RelatedIdProcessed)
							}
						 `)
	if err == nil && relatedIdResult.String() != "undefined" {
		s.RelatedId = []byte(relatedIdResult.String())
	}
	// extracting RelatedRef
	relatedRefResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'DocumentManifest.related.ref'))")
	if err == nil && relatedRefResult.String() != "undefined" {
	}
	// extracting Source
	sourceResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'DocumentManifest.source')[0]")
	if err == nil && sourceResult.String() != "undefined" {
		s.Source = sourceResult.String()
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'DocumentManifest.status')
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
	subjectResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'DocumentManifest.subject'))")
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
func (s *FhirDocumentManifest) TableName() string {
	return "fhir_document_manifest"
}
