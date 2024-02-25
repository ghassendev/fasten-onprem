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

type FhirObservation struct {
	models.ResourceBase
	// Reference to the service request.
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// The classification of the type of observation
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Code that identifies the allergy or intolerance
	   * [Condition](condition.html): Code for the condition
	   * [DeviceRequest](devicerequest.html): Code for what is being requested/ordered
	   * [DiagnosticReport](diagnosticreport.html): The code for the report, as opposed to codes for the atomic results, which are the names on the observation resource referred to from the result
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a condition code
	   * [List](list.html): What the purpose of this list is
	   * [Medication](medication.html): Returns medications for a specific code
	   * [MedicationAdministration](medicationadministration.html): Return administrations of this medication code
	   * [MedicationDispense](medicationdispense.html): Returns dispenses of this medicine code
	   * [MedicationRequest](medicationrequest.html): Return prescriptions of this medication code
	   * [MedicationStatement](medicationstatement.html): Return statements of this medication code
	   * [Observation](observation.html): The code of the observation type
	   * [Procedure](procedure.html): A code to identify a  procedure
	   * [ServiceRequest](servicerequest.html): What is being requested/ordered
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Code datatypes.JSON `gorm:"column:code;type:text;serializer:json" json:"code,omitempty"`
	// The code of the observation type or component type
	// https://hl7.org/fhir/r4/search.html#token
	ComboCode datatypes.JSON `gorm:"column:comboCode;type:text;serializer:json" json:"comboCode,omitempty"`
	// The reason why the expected value in the element Observation.value[x] or Observation.component.value[x] is missing.
	// https://hl7.org/fhir/r4/search.html#token
	ComboDataAbsentReason datatypes.JSON `gorm:"column:comboDataAbsentReason;type:text;serializer:json" json:"comboDataAbsentReason,omitempty"`
	// The value or component value of the observation, if the value is a CodeableConcept
	// https://hl7.org/fhir/r4/search.html#token
	ComboValueConcept datatypes.JSON `gorm:"column:comboValueConcept;type:text;serializer:json" json:"comboValueConcept,omitempty"`
	// The value or component value of the observation, if the value is a Quantity, or a SampledData (just search on the bounds of the values in sampled data)
	// https://hl7.org/fhir/r4/search.html#quantity
	ComboValueQuantity datatypes.JSON `gorm:"column:comboValueQuantity;type:text;serializer:json" json:"comboValueQuantity,omitempty"`
	// The component code of the observation type
	// https://hl7.org/fhir/r4/search.html#token
	ComponentCode datatypes.JSON `gorm:"column:componentCode;type:text;serializer:json" json:"componentCode,omitempty"`
	// The reason why the expected value in the element Observation.component.value[x] is missing.
	// https://hl7.org/fhir/r4/search.html#token
	ComponentDataAbsentReason datatypes.JSON `gorm:"column:componentDataAbsentReason;type:text;serializer:json" json:"componentDataAbsentReason,omitempty"`
	// The value of the component observation, if the value is a CodeableConcept
	// https://hl7.org/fhir/r4/search.html#token
	ComponentValueConcept datatypes.JSON `gorm:"column:componentValueConcept;type:text;serializer:json" json:"componentValueConcept,omitempty"`
	// The value of the component observation, if the value is a Quantity, or a SampledData (just search on the bounds of the values in sampled data)
	// https://hl7.org/fhir/r4/search.html#quantity
	ComponentValueQuantity datatypes.JSON `gorm:"column:componentValueQuantity;type:text;serializer:json" json:"componentValueQuantity,omitempty"`
	// The reason why the expected value in the element Observation.value[x] is missing.
	// https://hl7.org/fhir/r4/search.html#token
	DataAbsentReason datatypes.JSON `gorm:"column:dataAbsentReason;type:text;serializer:json" json:"dataAbsentReason,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Date first version of the resource instance was recorded
	   * [CarePlan](careplan.html): Time period plan covers
	   * [CareTeam](careteam.html): Time period team covers
	   * [ClinicalImpression](clinicalimpression.html): When the assessment was documented
	   * [Composition](composition.html): Composition editing time
	   * [Consent](consent.html): When this Consent was created or indexed
	   * [DiagnosticReport](diagnosticreport.html): The clinically relevant time of the report
	   * [Encounter](encounter.html): A date within the period the Encounter lasted
	   * [EpisodeOfCare](episodeofcare.html): The provided date search value falls within the episode of care's period
	   * [FamilyMemberHistory](familymemberhistory.html): When history was recorded or last updated
	   * [Flag](flag.html): Time period when flag is active
	   * [Immunization](immunization.html): Vaccination  (non)-Administration Date
	   * [List](list.html): When the list was prepared
	   * [Observation](observation.html): Obtained date/time. If the obtained element is a period, a date that falls in the period
	   * [Procedure](procedure.html): When the procedure was performed
	   * [RiskAssessment](riskassessment.html): When was assessment made?
	   * [SupplyRequest](supplyrequest.html): When the request was made
	*/
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	// Related measurements the observation is made from
	// https://hl7.org/fhir/r4/search.html#reference
	DerivedFrom datatypes.JSON `gorm:"column:derivedFrom;type:text;serializer:json" json:"derivedFrom,omitempty"`
	// The Device that generated the observation data.
	// https://hl7.org/fhir/r4/search.html#reference
	Device datatypes.JSON `gorm:"column:device;type:text;serializer:json" json:"device,omitempty"`
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
	// The focus of an observation when the focus is not the patient of record.
	// https://hl7.org/fhir/r4/search.html#reference
	Focus datatypes.JSON `gorm:"column:focus;type:text;serializer:json" json:"focus,omitempty"`
	// Related resource that belongs to the Observation group
	// https://hl7.org/fhir/r4/search.html#reference
	HasMember datatypes.JSON `gorm:"column:hasMember;type:text;serializer:json" json:"hasMember,omitempty"`
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
	// The method used for the observation
	// https://hl7.org/fhir/r4/search.html#token
	Method datatypes.JSON `gorm:"column:method;type:text;serializer:json" json:"method,omitempty"`
	// Part of referenced event
	// https://hl7.org/fhir/r4/search.html#reference
	PartOf datatypes.JSON `gorm:"column:partOf;type:text;serializer:json" json:"partOf,omitempty"`
	// Who performed the observation
	// https://hl7.org/fhir/r4/search.html#reference
	Performer datatypes.JSON `gorm:"column:performer;type:text;serializer:json" json:"performer,omitempty"`
	// Specimen used for this observation
	// https://hl7.org/fhir/r4/search.html#reference
	Specimen datatypes.JSON `gorm:"column:specimen;type:text;serializer:json" json:"specimen,omitempty"`
	// The status of the observation
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The subject that the observation is about
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// The value of the observation, if the value is a CodeableConcept
	// https://hl7.org/fhir/r4/search.html#token
	ValueConcept datatypes.JSON `gorm:"column:valueConcept;type:text;serializer:json" json:"valueConcept,omitempty"`
	// The value of the observation, if the value is a date or period of time
	// https://hl7.org/fhir/r4/search.html#date
	ValueDate *time.Time `gorm:"column:valueDate;type:datetime" json:"valueDate,omitempty"`
	// The value of the observation, if the value is a Quantity, or a SampledData (just search on the bounds of the values in sampled data)
	// https://hl7.org/fhir/r4/search.html#quantity
	ValueQuantity datatypes.JSON `gorm:"column:valueQuantity;type:text;serializer:json" json:"valueQuantity,omitempty"`
	// The value of the observation, if the value is a string, and also searches in CodeableConcept.text
	// https://hl7.org/fhir/r4/search.html#string
	ValueString datatypes.JSON `gorm:"column:valueString;type:text;serializer:json" json:"valueString,omitempty"`
}

func (s *FhirObservation) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"basedOn":                   "reference",
		"category":                  "token",
		"code":                      "token",
		"comboCode":                 "token",
		"comboDataAbsentReason":     "token",
		"comboValueConcept":         "token",
		"comboValueQuantity":        "quantity",
		"componentCode":             "token",
		"componentDataAbsentReason": "token",
		"componentValueConcept":     "token",
		"componentValueQuantity":    "quantity",
		"dataAbsentReason":          "token",
		"date":                      "date",
		"derivedFrom":               "reference",
		"device":                    "reference",
		"encounter":                 "reference",
		"focus":                     "reference",
		"hasMember":                 "reference",
		"id":                        "keyword",
		"identifier":                "token",
		"language":                  "token",
		"metaLastUpdated":           "date",
		"metaProfile":               "reference",
		"metaTag":                   "token",
		"metaVersionId":             "keyword",
		"method":                    "token",
		"partOf":                    "reference",
		"performer":                 "reference",
		"sort_date":                 "date",
		"source_id":                 "keyword",
		"source_resource_id":        "keyword",
		"source_resource_type":      "keyword",
		"source_uri":                "keyword",
		"specimen":                  "reference",
		"status":                    "token",
		"subject":                   "reference",
		"text":                      "keyword",
		"valueConcept":              "token",
		"valueDate":                 "date",
		"valueQuantity":             "quantity",
		"valueString":               "string",
	}
	return searchParameters
}
func (s *FhirObservation) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting BasedOn
	basedOnResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.basedOn')")
	if err == nil && basedOnResult.String() != "undefined" {
		s.BasedOn = []byte(basedOnResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.category')")
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting Code
	codeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.code | AllergyIntolerance.reaction.substance | Condition.code | (DeviceRequest.codeCodeableConcept) | DiagnosticReport.code | FamilyMemberHistory.condition.code | List.code | Medication.code | (MedicationAdministration.medicationCodeableConcept) | (MedicationDispense.medicationCodeableConcept) | (MedicationRequest.medicationCodeableConcept) | (MedicationStatement.medicationCodeableConcept) | Observation.code | Procedure.code | ServiceRequest.code')")
	if err == nil && codeResult.String() != "undefined" {
		s.Code = []byte(codeResult.String())
	}
	// extracting ComboCode
	comboCodeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.code | Observation.component.code')")
	if err == nil && comboCodeResult.String() != "undefined" {
		s.ComboCode = []byte(comboCodeResult.String())
	}
	// extracting ComboDataAbsentReason
	comboDataAbsentReasonResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.dataAbsentReason | Observation.component.dataAbsentReason')")
	if err == nil && comboDataAbsentReasonResult.String() != "undefined" {
		s.ComboDataAbsentReason = []byte(comboDataAbsentReasonResult.String())
	}
	// extracting ComboValueConcept
	comboValueConceptResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, '(Observation.valueCodeableConcept) | (Observation.component.valueCodeableConcept)')")
	if err == nil && comboValueConceptResult.String() != "undefined" {
		s.ComboValueConcept = []byte(comboValueConceptResult.String())
	}
	// extracting ComboValueQuantity
	comboValueQuantityResult, err := vm.RunString("extractCatchallSearchParameters(fhirResource, '(Observation.valueQuantity) | (Observation.valueSampledData) | (Observation.component.valueQuantity) | (Observation.component.valueSampledData)')")
	if err == nil && comboValueQuantityResult.String() != "undefined" {
		s.ComboValueQuantity = []byte(comboValueQuantityResult.String())
	}
	// extracting ComponentCode
	componentCodeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.component.code')")
	if err == nil && componentCodeResult.String() != "undefined" {
		s.ComponentCode = []byte(componentCodeResult.String())
	}
	// extracting ComponentDataAbsentReason
	componentDataAbsentReasonResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.component.dataAbsentReason')")
	if err == nil && componentDataAbsentReasonResult.String() != "undefined" {
		s.ComponentDataAbsentReason = []byte(componentDataAbsentReasonResult.String())
	}
	// extracting ComponentValueConcept
	componentValueConceptResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, '(Observation.component.valueCodeableConcept)')")
	if err == nil && componentValueConceptResult.String() != "undefined" {
		s.ComponentValueConcept = []byte(componentValueConceptResult.String())
	}
	// extracting ComponentValueQuantity
	componentValueQuantityResult, err := vm.RunString("extractCatchallSearchParameters(fhirResource, '(Observation.component.valueQuantity) | (Observation.component.valueSampledData)')")
	if err == nil && componentValueQuantityResult.String() != "undefined" {
		s.ComponentValueQuantity = []byte(componentValueQuantityResult.String())
	}
	// extracting DataAbsentReason
	dataAbsentReasonResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.dataAbsentReason')")
	if err == nil && dataAbsentReasonResult.String() != "undefined" {
		s.DataAbsentReason = []byte(dataAbsentReasonResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'AllergyIntolerance.recordedDate | CarePlan.period | CareTeam.period | ClinicalImpression.date | Composition.date | Consent.dateTime | DiagnosticReport.effectiveDateTime | DiagnosticReport.effectivePeriod | Encounter.period | EpisodeOfCare.period | FamilyMemberHistory.date | Flag.period | (Immunization.occurrenceDateTime) | List.date | Observation.effectiveDateTime | Observation.effectivePeriod | Observation.effectiveTiming | Observation.effectiveInstant | Procedure.performedDateTime | Procedure.performedPeriod | Procedure.performedString | Procedure.performedAge | Procedure.performedRange | (RiskAssessment.occurrenceDateTime) | SupplyRequest.authoredOn')")
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
	// extracting DerivedFrom
	derivedFromResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.derivedFrom')")
	if err == nil && derivedFromResult.String() != "undefined" {
		s.DerivedFrom = []byte(derivedFromResult.String())
	}
	// extracting Device
	deviceResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.device')")
	if err == nil && deviceResult.String() != "undefined" {
		s.Device = []byte(deviceResult.String())
	}
	// extracting Encounter
	encounterResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Composition.encounter | DeviceRequest.encounter | DiagnosticReport.encounter | DocumentReference.context.encounter.where(resolve() is Encounter) | Flag.encounter | List.encounter | NutritionOrder.encounter | Observation.encounter | Procedure.encounter | RiskAssessment.encounter | ServiceRequest.encounter | VisionPrescription.encounter')")
	if err == nil && encounterResult.String() != "undefined" {
		s.Encounter = []byte(encounterResult.String())
	}
	// extracting Focus
	focusResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.focus')")
	if err == nil && focusResult.String() != "undefined" {
		s.Focus = []byte(focusResult.String())
	}
	// extracting HasMember
	hasMemberResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.hasMember')")
	if err == nil && hasMemberResult.String() != "undefined" {
		s.HasMember = []byte(hasMemberResult.String())
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
	// extracting Method
	methodResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.method')")
	if err == nil && methodResult.String() != "undefined" {
		s.Method = []byte(methodResult.String())
	}
	// extracting PartOf
	partOfResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.partOf')")
	if err == nil && partOfResult.String() != "undefined" {
		s.PartOf = []byte(partOfResult.String())
	}
	// extracting Performer
	performerResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.performer')")
	if err == nil && performerResult.String() != "undefined" {
		s.Performer = []byte(performerResult.String())
	}
	// extracting Specimen
	specimenResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.specimen')")
	if err == nil && specimenResult.String() != "undefined" {
		s.Specimen = []byte(specimenResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Observation.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Observation.subject')")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	// extracting ValueConcept
	valueConceptResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, '(Observation.valueCodeableConcept)')")
	if err == nil && valueConceptResult.String() != "undefined" {
		s.ValueConcept = []byte(valueConceptResult.String())
	}
	// extracting ValueDate
	valueDateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, '(Observation.valueDateTime) | (Observation.valuePeriod)')")
	if err == nil && valueDateResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, valueDateResult.String()); err == nil {
			s.ValueDate = &t
		} else if t, err = time.Parse("2006-01-02", valueDateResult.String()); err == nil {
			s.ValueDate = &t
		} else if t, err = time.Parse("2006-01", valueDateResult.String()); err == nil {
			s.ValueDate = &t
		} else if t, err = time.Parse("2006", valueDateResult.String()); err == nil {
			s.ValueDate = &t
		}
	}
	// extracting ValueQuantity
	valueQuantityResult, err := vm.RunString("extractCatchallSearchParameters(fhirResource, '(Observation.valueQuantity) | (Observation.valueSampledData)')")
	if err == nil && valueQuantityResult.String() != "undefined" {
		s.ValueQuantity = []byte(valueQuantityResult.String())
	}
	// extracting ValueString
	valueStringResult, err := vm.RunString("extractStringSearchParameters(fhirResource, '(Observation.valueString) | (Observation.valueCodeableConcept).text')")
	if err == nil && valueStringResult.String() != "undefined" {
		s.ValueString = []byte(valueStringResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirObservation) TableName() string {
	return "fhir_observation"
}
