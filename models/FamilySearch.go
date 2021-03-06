package models

type FamilySearch struct {
	ChildAndParentsRelationships []ChildAndParentsRelationship `json:"childAndParentsRelationships"`
	Discussions                  []Discussion                  `json:"discussions"`
	Users                        []User                        `json:"users"`
	Merges                       []Merge                       `json:"merges"`
	MergeAnalyses                []MergeAnalysis               `json:"mergeAnalyses"`
	Features                     []FeatureSet                  `json:"features"`
	Lang                         string                        `json:"lang"`
	Description                  string                        `json:"description"`
	Profile                      string                        `json:"profile"`
	Attribution                  Attribution                   `json:"attribution"`
	Persons                      []Person                      `json:"persons"`
	Relationships                []Relationship                `json:"relationships"`
	SourceDescriptions           []SourceDescription           `json:"sourceDescriptions"`
	Agents                       []Agent                       `json:"agents"`
	Events                       []Event                       `json:"events"`
	Places                       []PlaceDescription            `json:"places"`
	Documents                    []Document                    `json:"documents"`
	Collections                  []Collection                  `json:"collections"`
	Fields                       []Field                       `json:"fields"`
	RecordDescriptors            []RecordDescriptor            `json:"recordDescriptors"`
	Links                        map[string]Link               `json:"links"`
	Id                           string                        `json:"id"`
}
