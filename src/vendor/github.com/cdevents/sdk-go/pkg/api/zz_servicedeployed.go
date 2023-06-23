// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"fmt"
	"time"
)

var servicedeployedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.3.0/schema/service-deployed-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.service.deployed.0.1.1"],"default":"dev.cdevents.service.deployed.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["service"],"default":"service"},"content":{"properties":{"environment":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]},"artifactId":{"type":"string","minLength":1}},"additionalProperties":false,"type":"object","required":["environment","artifactId"]}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// ServiceDeployed event v0.1.1
	ServiceDeployedEventV1 CDEventType = CDEventType{
		Subject:   "service",
		Predicate: "deployed",
		Version:   "0.1.1",
	}
)

type ServiceDeployedSubjectContent struct {
	ArtifactId string `json:"artifactId" validate:"purl"`

	Environment *Reference `json:"environment"`
}

type ServiceDeployedSubject struct {
	SubjectBase
	Content ServiceDeployedSubjectContent `json:"content"`
}

func (sc ServiceDeployedSubject) GetSubjectType() SubjectType {
	return "service"
}

type ServiceDeployedEvent struct {
	Context Context                `json:"context"`
	Subject ServiceDeployedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ServiceDeployedEvent) GetType() CDEventType {
	return ServiceDeployedEventV1
}

func (e ServiceDeployedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServiceDeployedEvent) GetId() string {
	return e.Context.Id
}

func (e ServiceDeployedEvent) GetSource() string {
	return e.Context.Source
}

func (e ServiceDeployedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceDeployedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceDeployedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceDeployedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ServiceDeployedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ServiceDeployedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ServiceDeployedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ServiceDeployedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ServiceDeployedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceDeployedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceDeployedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceDeployedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceDeployedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ServiceDeployedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ServiceDeployedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), servicedeployedschema
}

// Set subject custom fields

func (e *ServiceDeployedEvent) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

func (e *ServiceDeployedEvent) SetSubjectEnvironment(environment *Reference) {
	e.Subject.Content.Environment = environment
}

// New creates a new ServiceDeployedEvent
func NewServiceDeployedEvent() (*ServiceDeployedEvent, error) {
	e := &ServiceDeployedEvent{
		Context: Context{
			Type:    ServiceDeployedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ServiceDeployedSubject{
			SubjectBase: SubjectBase{
				Type: "service",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
