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

type FhirDocumentReference struct {
	models.ResourceBase
	// Who/what authenticated the document
	// https://hl7.org/fhir/r4/search.html#reference
	Authenticator datatypes.JSON `gorm:"column:authenticator;type:text;serializer:json" json:"authenticator,omitempty"`
	// Who and/or what authored the document
	// https://hl7.org/fhir/r4/search.html#reference
	Author datatypes.JSON `gorm:"column:author;type:text;serializer:json" json:"author,omitempty"`
	// Categorization of document
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	// Mime type of the content, with charset etc.
	// https://hl7.org/fhir/r4/search.html#token
	Contenttype datatypes.JSON `gorm:"column:contenttype;type:text;serializer:json" json:"contenttype,omitempty"`
	// Organization which maintains the document
	// https://hl7.org/fhir/r4/search.html#reference
	Custodian datatypes.JSON `gorm:"column:custodian;type:text;serializer:json" json:"custodian,omitempty"`
	// When this document reference was created
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	// Human-readable description
	// https://hl7.org/fhir/r4/search.html#string
	Description datatypes.JSON `gorm:"column:description;type:text;serializer:json" json:"description,omitempty"`
	/*
	   Multiple Resources:

	   * [Composition](composition.html): Context of the Composition
	   * [DeviceRequest](devicerequest.html): Encounter during which request was created
	   * [DiagnosticReport](diagnosticreport.html): The Encounter when the order was made
	   * [DocumentReference](documentreference.html): Context of the document  content
	   * [Flag](flag.html): Alert relevant during encounter
	   * [List](list.html): Context in which list created
	   * [NutritionOrder](nutritionorder.html): Return nutrition orders with this encounter identifier
	   * [Observation](observation.html): Encounter related to the observation
	   * [Procedure](procedure.html): Encounter created as part of
	   * [RiskAssessment](riskassessment.html): Where was assessment performed?
	   * [ServiceRequest](servicerequest.html): An encounter in which this request is made
	   * [VisionPrescription](visionprescription.html): Return prescriptions with this encounter identifier
	*/
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
	// Main clinical acts documented
	// https://hl7.org/fhir/r4/search.html#token
	Event datatypes.JSON `gorm:"column:event;type:text;serializer:json" json:"event,omitempty"`
	// Kind of facility where patient was seen
	// https://hl7.org/fhir/r4/search.html#token
	Facility datatypes.JSON `gorm:"column:facility;type:text;serializer:json" json:"facility,omitempty"`
	// Format/content rules for the document
	// https://hl7.org/fhir/r4/search.html#token
	Format datatypes.JSON `gorm:"column:format;type:text;serializer:json" json:"format,omitempty"`
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
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// Uri where the data can be found
	// https://hl7.org/fhir/r4/search.html#uri
	Location string `gorm:"column:location;type:text" json:"location,omitempty"`
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
	// Time of service that is being documented
	// https://hl7.org/fhir/r4/search.html#date
	Period *time.Time `gorm:"column:period;type:datetime" json:"period,omitempty"`
	// Related identifiers or resources
	// https://hl7.org/fhir/r4/search.html#reference
	Related datatypes.JSON `gorm:"column:related;type:text;serializer:json" json:"related,omitempty"`
	// Target of the relationship
	// https://hl7.org/fhir/r4/search.html#reference
	Relatesto datatypes.JSON `gorm:"column:relatesto;type:text;serializer:json" json:"relatesto,omitempty"`
	// replaces | transforms | signs | appends
	// https://hl7.org/fhir/r4/search.html#token
	Relation datatypes.JSON `gorm:"column:relation;type:text;serializer:json" json:"relation,omitempty"`
	// Document security-tags
	// https://hl7.org/fhir/r4/search.html#token
	SecurityLabel datatypes.JSON `gorm:"column:securityLabel;type:text;serializer:json" json:"securityLabel,omitempty"`
	// Additional details about where the content was created (e.g. clinical specialty)
	// https://hl7.org/fhir/r4/search.html#token
	Setting datatypes.JSON `gorm:"column:setting;type:text;serializer:json" json:"setting,omitempty"`
	// current | superseded | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Who/what is the subject of the document
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): allergy | intolerance - Underlying mechanism (if known)
	   * [Composition](composition.html): Kind of composition (LOINC if possible)
	   * [DocumentManifest](documentmanifest.html): Kind of document set
	   * [DocumentReference](documentreference.html): Kind of document (LOINC if possible)
	   * [Encounter](encounter.html): Specific type of encounter
	   * [EpisodeOfCare](episodeofcare.html): Type/class  - e.g. specialist referral, disease management
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirDocumentReference) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"authenticator":        "reference",
		"author":               "reference",
		"category":             "token",
		"contenttype":          "token",
		"custodian":            "reference",
		"date":                 "date",
		"description":          "string",
		"encounter":            "reference",
		"event":                "token",
		"facility":             "token",
		"format":               "token",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"location":             "uri",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"period":               "date",
		"related":              "reference",
		"relatesto":            "reference",
		"relation":             "token",
		"securityLabel":        "token",
		"setting":              "token",
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
func (s *FhirDocumentReference) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Authenticator
	authenticatorResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'DocumentReference.authenticator')")
	if err == nil && authenticatorResult.String() != "undefined" {
		s.Authenticator = []byte(authenticatorResult.String())
	}
	// extracting Author
	authorResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'DocumentReference.author')")
	if err == nil && authorResult.String() != "undefined" {
		s.Author = []byte(authorResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.category')")
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting Contenttype
	contenttypeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.content.attachment.contentType')")
	if err == nil && contenttypeResult.String() != "undefined" {
		s.Contenttype = []byte(contenttypeResult.String())
	}
	// extracting Custodian
	custodianResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'DocumentReference.custodian')")
	if err == nil && custodianResult.String() != "undefined" {
		s.Custodian = []byte(custodianResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'DocumentReference.date')")
	if err == nil && dateResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, dateResult.String()); err == nil {
			s.Date = &t
		} else if t, err = time.Parse("2006-01-02", dateResult.String()); err == nil {
			s.Date = &t
		} else if t, err = time.Parse("2006-01", dateResult.String()); err == nil {
			s.Date = &t
		} else if t, err = time.Parse("2006", dateResult.String()); err == nil {
			s.Date = &t
		}
	}
	// extracting Description
	descriptionResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'DocumentReference.description')")
	if err == nil && descriptionResult.String() != "undefined" {
		s.Description = []byte(descriptionResult.String())
	}
	// extracting Encounter
	encounterResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Composition.encounter | DeviceRequest.encounter | DiagnosticReport.encounter | DocumentReference.context.encounter.where(resolve() is Encounter) | Flag.encounter | List.encounter | NutritionOrder.encounter | Observation.encounter | Procedure.encounter | RiskAssessment.encounter | ServiceRequest.encounter | VisionPrescription.encounter')")
	if err == nil && encounterResult.String() != "undefined" {
		s.Encounter = []byte(encounterResult.String())
	}
	// extracting Event
	eventResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.context.event')")
	if err == nil && eventResult.String() != "undefined" {
		s.Event = []byte(eventResult.String())
	}
	// extracting Facility
	facilityResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.context.facilityType')")
	if err == nil && facilityResult.String() != "undefined" {
		s.Facility = []byte(facilityResult.String())
	}
	// extracting Format
	formatResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.content.format')")
	if err == nil && formatResult.String() != "undefined" {
		s.Format = []byte(formatResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting Location
	locationResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'DocumentReference.content.attachment.url')")
	if err == nil && locationResult.String() != "undefined" {
		s.Location = locationResult.String()
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
	// extracting Period
	periodResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'DocumentReference.context.period')")
	if err == nil && periodResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, periodResult.String()); err == nil {
			s.Period = &t
		} else if t, err = time.Parse("2006-01-02", periodResult.String()); err == nil {
			s.Period = &t
		} else if t, err = time.Parse("2006-01", periodResult.String()); err == nil {
			s.Period = &t
		} else if t, err = time.Parse("2006", periodResult.String()); err == nil {
			s.Period = &t
		}
	}
	// extracting Related
	relatedResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'DocumentReference.context.related')")
	if err == nil && relatedResult.String() != "undefined" {
		s.Related = []byte(relatedResult.String())
	}
	// extracting Relatesto
	relatestoResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'DocumentReference.relatesTo.target')")
	if err == nil && relatestoResult.String() != "undefined" {
		s.Relatesto = []byte(relatestoResult.String())
	}
	// extracting Relation
	relationResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.relatesTo.code')")
	if err == nil && relationResult.String() != "undefined" {
		s.Relation = []byte(relationResult.String())
	}
	// extracting SecurityLabel
	securityLabelResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.securityLabel')")
	if err == nil && securityLabelResult.String() != "undefined" {
		s.SecurityLabel = []byte(securityLabelResult.String())
	}
	// extracting Setting
	settingResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.context.practiceSetting')")
	if err == nil && settingResult.String() != "undefined" {
		s.Setting = []byte(settingResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'DocumentReference.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'DocumentReference.subject')")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	// extracting Type
	typeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.type | Composition.type | DocumentManifest.type | DocumentReference.type | Encounter.type | EpisodeOfCare.type')")
	if err == nil && typeResult.String() != "undefined" {
		s.Type = []byte(typeResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirDocumentReference) TableName() string {
	return "fhir_document_reference"
}
