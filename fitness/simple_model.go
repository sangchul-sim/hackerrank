package fitness

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

type AppleFitness2 struct {
	ClinicalDocument struct {
		Sdtc       string `json:"-sdtc"`
		TemplateID struct {
			Root string `json:"-root"`
		} `json:"templateId"`
		Code struct {
			Code           string `json:"-code"`
			DisplayName    string `json:"-displayName"`
			CodeSystem     string `json:"-codeSystem"`
			CodeSystemName string `json:"-codeSystemName"`
		} `json:"code"`
		SchemaLocation string `json:"-schemaLocation"`
		EffectiveTime  struct {
			Value string `json:"-value"`
		} `json:"effectiveTime"`
		Xsi   string `json:"-xsi"`
		Cda   string `json:"-cda"`
		Fhir  string `json:"-fhir"`
		Entry struct {
			TypeCode  string `json:"-typeCode"`
			Organizer struct {
				MoodCode   string `json:"-moodCode"`
				TemplateID struct {
					Root string `json:"-root"`
				} `json:"templateId"`
				ID struct {
					Root string `json:"-root"`
				} `json:"id"`
				Code struct {
					Code           string `json:"-code"`
					CodeSystem     string `json:"-codeSystem"`
					CodeSystemName string `json:"-codeSystemName"`
					DisplayName    string `json:"-displayName"`
				} `json:"code"`
				StatusCode struct {
					Code string `json:"-code"`
				} `json:"statusCode"`
				EffectiveTime struct {
					Low struct {
						Value string `json:"-value"`
					} `json:"low"`
					High struct {
						Value string `json:"-value"`
					} `json:"high"`
				} `json:"effectiveTime"`
				Component []struct {
					Observation struct {
						MoodCode   string `json:"-moodCode"`
						StatusCode struct {
							Code string `json:"-code"`
						} `json:"statusCode"`
						EffectiveTime struct {
							Low struct {
								Value string `json:"-value"`
							} `json:"low"`
							High struct {
								Value string `json:"-value"`
							} `json:"high"`
						} `json:"effectiveTime"`
						Value struct {
							Type  string `json:"-type"`
							Value string `json:"-value"`
							Unit  string `json:"-unit"`
						} `json:"value"`
						InterpretationCode struct {
							Code       string `json:"-code"`
							CodeSystem string `json:"-codeSystem"`
						} `json:"interpretationCode"`
						ClassCode  string `json:"-classCode"`
						TemplateID struct {
							Root string `json:"-root"`
						} `json:"templateId"`
						ID struct {
							Root string `json:"-root"`
						} `json:"id"`
						Code struct {
							Code           string `json:"-code"`
							CodeSystem     string `json:"-codeSystem"`
							CodeSystemName string `json:"-codeSystemName"`
							DisplayName    string `json:"-displayName"`
						} `json:"code"`
						Text struct {
							MetadataEntry struct {
								Key   string `json:"key"`
								Value string `json:"value"`
							} `json:"metadataEntry"`
							SourceName string `json:"sourceName"`
							Value      string `json:"value"`
							Type       string `json:"type"`
							Unit       string `json:"unit"`
						} `json:"text"`
					} `json:"observation"`
				} `json:"component"`
				ClassCode string `json:"-classCode"`
			} `json:"organizer"`
		} `json:"entry"`
		Xmlns     string `json:"-xmlns"`
		RealmCode struct {
			Code string `json:"-code"`
		} `json:"realmCode"`
		TypeID struct {
			Root      string `json:"-root"`
			Extension string `json:"-extension"`
		} `json:"typeId"`
		ID struct {
			Extension string `json:"-extension"`
			Root      string `json:"-root"`
		} `json:"id"`
		Title               string `json:"title"`
		ConfidentialityCode struct {
			Code       string `json:"-code"`
			CodeSystem string `json:"-codeSystem"`
		} `json:"confidentialityCode"`
		RecordTarget struct {
			PatientRole struct {
				ID struct {
					Root       string `json:"-root"`
					NullFlavor string `json:"-nullFlavor"`
				} `json:"id"`
				Patient struct {
					Name struct {
						Content string `json:"#content"`
						Use     string `json:"-use"`
					} `json:"name"`
					AdministrativeGenderCode struct {
						Code        string `json:"-code"`
						CodeSystem  string `json:"-codeSystem"`
						DisplayName string `json:"-displayName"`
					} `json:"administrativeGenderCode"`
					BirthTime struct {
						Value string `json:"-value"`
					} `json:"birthTime"`
				} `json:"patient"`
			} `json:"patientRole"`
		} `json:"recordTarget"`
	} `json:"ClinicalDocument"`
}
