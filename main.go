package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	//arrays.ArraysDs()
	//arrays.TwoDArrayDs()
	//arrays.DynamicArray()
	//arrays.Rotation()
	//arrays.SparseArrays()
	//arrays.ArrayManipulation()
	fitness()
}

func fitness() {
	filePath := "output.json"
	data, err := readFromFile(filePath)
	if err != nil {
		panic(err)
	}

	var fitness AppleFitness
	if err := json.Unmarshal(data, &fitness); err != nil {
		panic(err)
	}

	// TODO sort.Sort 구현할수 있도록
	type result struct {
		dt time.Time
		wt float64
	}

	var fitResult []result
	for _, component := range fitness.ClinicalDocument.Entry.Organizer.Component {
		obvTime, err := time.Parse(
			"20060102150405-0700",
			component.Observation.EffectiveTime.Low.Value,
		)
		if err != nil {
			panic(err)
		}
		f, err := strconv.ParseFloat(component.Observation.Value.Value, 64)
		if err != nil {
			panic(err)
		}
		fitResult = append(fitResult, result{obvTime, f})
		//// "20150527083200+0900"
		//component.Observation.EffectiveTime.Low
		//// "79.8"
		//component.Observation.Value.Value
		//// "kg"
		//component.Observation.Value.Unit
		//break
	}
	fmt.Println(fitResult)
}

func readFromFile(filePath string) ([]byte, error) {
	// Read file to byte slice
	return ioutil.ReadFile(filePath)
}

type AppleFitness struct {
	ClinicalDocument struct {
		XmlnsXsi          string `json:"xmlns$xsi"`
		XsiSchemaLocation string `json:"xsi$schemaLocation"`
		Xmlns             string `json:"xmlns"`
		XmlnsCda          string `json:"xmlns$cda"`
		XmlnsSdtc         string `json:"xmlns$sdtc"`
		XmlnsFhir         string `json:"xmlns$fhir"`
		RealmCode         struct {
			Code string `json:"code"`
		} `json:"realmCode"`
		TypeID struct {
			Root      string `json:"root"`
			Extension string `json:"extension"`
		} `json:"typeId"`
		TemplateID struct {
			Root string `json:"root"`
		} `json:"templateId"`
		ID struct {
			Extension string `json:"extension"`
			Root      string `json:"root"`
		} `json:"id"`
		Code struct {
			CodeSystem     string `json:"codeSystem"`
			CodeSystemName string `json:"codeSystemName"`
			Code           string `json:"code"`
			DisplayName    string `json:"displayName"`
		} `json:"code"`
		Title struct {
			T string `json:"$t"`
		} `json:"title"`
		EffectiveTime struct {
			Value string `json:"value"`
		} `json:"effectiveTime"`
		ConfidentialityCode struct {
			Code       string `json:"code"`
			CodeSystem string `json:"codeSystem"`
		} `json:"confidentialityCode"`
		RecordTarget struct {
			PatientRole struct {
				ID struct {
					Root       string `json:"root"`
					NullFlavor string `json:"nullFlavor"`
				} `json:"id"`
				Patient struct {
					Name struct {
						Use string `json:"use"`
						T   string `json:"$t"`
					} `json:"name"`
					AdministrativeGenderCode struct {
						Code        string `json:"code"`
						CodeSystem  string `json:"codeSystem"`
						DisplayName string `json:"displayName"`
					} `json:"administrativeGenderCode"`
					BirthTime struct {
						Value string `json:"value"`
					} `json:"birthTime"`
				} `json:"patient"`
			} `json:"patientRole"`
		} `json:"recordTarget"`
		Entry struct {
			TypeCode  string `json:"typeCode"`
			Organizer struct {
				ClassCode  string `json:"classCode"`
				MoodCode   string `json:"moodCode"`
				TemplateID struct {
					Root string `json:"root"`
				} `json:"templateId"`
				ID struct {
					Root string `json:"root"`
				} `json:"id"`
				Code struct {
					Code           string `json:"code"`
					CodeSystem     string `json:"codeSystem"`
					CodeSystemName string `json:"codeSystemName"`
					DisplayName    string `json:"displayName"`
				} `json:"code"`
				StatusCode struct {
					Code string `json:"code"`
				} `json:"statusCode"`
				EffectiveTime struct {
					Low struct {
						Value string `json:"value"`
					} `json:"low"`
					High struct {
						Value string `json:"value"`
					} `json:"high"`
				} `json:"effectiveTime"`
				Component []struct {
					Observation struct {
						ClassCode  string `json:"classCode"`
						MoodCode   string `json:"moodCode"`
						TemplateID struct {
							Root string `json:"root"`
						} `json:"templateId"`
						ID struct {
							Root string `json:"root"`
						} `json:"id"`
						Code struct {
							Code           string `json:"code"`
							CodeSystem     string `json:"codeSystem"`
							CodeSystemName string `json:"codeSystemName"`
							DisplayName    string `json:"displayName"`
						} `json:"code"`
						Text struct {
							SourceName struct {
								T string `json:"$t"`
							} `json:"sourceName"`
							Value struct {
								T string `json:"$t"`
							} `json:"value"`
							Type struct {
								T string `json:"$t"`
							} `json:"type"`
							Unit struct {
								T string `json:"$t"`
							} `json:"unit"`
							MetadataEntry struct {
								Key struct {
									T string `json:"$t"`
								} `json:"key"`
								Value struct {
									T string `json:"$t"`
								} `json:"value"`
							} `json:"metadataEntry"`
						} `json:"text"`
						StatusCode struct {
							Code string `json:"code"`
						} `json:"statusCode"`
						EffectiveTime struct {
							Low struct {
								Value string `json:"value"`
							} `json:"low"`
							High struct {
								Value string `json:"value"`
							} `json:"high"`
						} `json:"effectiveTime"`
						Value struct {
							XsiType string `json:"xsi$type"`
							Value   string `json:"value"`
							Unit    string `json:"unit"`
						} `json:"value"`
						InterpretationCode struct {
							Code       string `json:"code"`
							CodeSystem string `json:"codeSystem"`
						} `json:"interpretationCode"`
					} `json:"observation"`
				} `json:"component"`
			} `json:"organizer"`
		} `json:"entry"`
	} `json:"ClinicalDocument"`
}
