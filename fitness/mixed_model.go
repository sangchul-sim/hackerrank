package fitness

type AppleFitnessAll struct {
	HealthData struct {
		Workout []struct {
			CreationDate  string `json:"-creationDate"`
			EndDate       string `json:"-endDate"`
			MetadataEntry struct {
				Key   string `json:"-key"`
				Value string `json:"-value"`
			} `json:"MetadataEntry"`
			WorkoutActivityType   string `json:"-workoutActivityType"`
			TotalDistanceUnit     string `json:"-totalDistanceUnit"`
			TotalEnergyBurned     string `json:"-totalEnergyBurned"`
			SourceName            string `json:"-sourceName"`
			StartDate             string `json:"-startDate"`
			Duration              string `json:"-duration"`
			DurationUnit          string `json:"-durationUnit"`
			TotalDistance         string `json:"-totalDistance"`
			TotalEnergyBurnedUnit string `json:"-totalEnergyBurnedUnit"`
			SourceVersion         string `json:"-sourceVersion,omitempty"`
		} `json:"Workout"`
		Locale     string `json:"-locale"`
		ExportDate struct {
			Value string `json:"-value"`
		} `json:"ExportDate"`
		Me struct {
			HKCharacteristicTypeIdentifierFitzpatrickSkinType string `json:"-HKCharacteristicTypeIdentifierFitzpatrickSkinType"`
			HKCharacteristicTypeIdentifierDateOfBirth         string `json:"-HKCharacteristicTypeIdentifierDateOfBirth"`
			HKCharacteristicTypeIdentifierBiologicalSex       string `json:"-HKCharacteristicTypeIdentifierBiologicalSex"`
			HKCharacteristicTypeIdentifierBloodType           string `json:"-HKCharacteristicTypeIdentifierBloodType"`
		} `json:"Me"`
		Record []struct {
			Value         string `json:"-value"`
			MetadataEntry struct {
				Key   string `json:"-key"`
				Value string `json:"-value"`
			} `json:"MetadataEntry,omitempty"`
			Type          string `json:"-type"`
			SourceName    string `json:"-sourceName"`
			Unit          string `json:"-unit,omitempty"`
			CreationDate  string `json:"-creationDate"`
			StartDate     string `json:"-startDate"`
			EndDate       string `json:"-endDate"`
			SourceVersion string `json:"-sourceVersion,omitempty"`
			Device        string `json:"-device,omitempty"`
		} `json:"Record"`
	} `json:"HealthData"`
}
