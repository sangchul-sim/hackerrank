package fitness

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	xj "github.com/basgys/goxml2json"
)

type dailySimpleItem struct {
	dt  time.Time
	val float64
}
type dailyMixedItem struct {
	dt  time.Time
	val float64
}

type simpleResult []dailySimpleItem

func (p simpleResult) Len() int           { return len(p) }
func (p simpleResult) Less(i, j int) bool { return p[i].dt.UnixNano() < p[j].dt.UnixNano() }
func (p simpleResult) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type mixedResult []dailyMixedItem

func (p mixedResult) Len() int           { return len(p) }
func (p mixedResult) Less(i, j int) bool { return p[i].dt.UnixNano() < p[j].dt.UnixNano() }
func (p mixedResult) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// TODO google REST API 로 연결하기 (https://developers.google.com/fit/rest/v1/get-started)
// {"web":{"client_id":"112404397721-8hv9hetfoqkeq9bjr5b6ojrlmjtj96ho.apps.googleusercontent.com","project_id":"my-fitness-project-192408","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://accounts.google.com/o/oauth2/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"x8utO7s_2yQxXJ2wPYROUTZ1","redirect_uris":["http://localhost:1337/oauthcallback"]}}
func Fitness() {
	//simpleData()
	mixedData()
}

func mixedData() {
	const (
		RecordTypeWalkingRunningDistance = "HKQuantityTypeIdentifierDistanceWalkingRunning"
		RecordTypeSwimmingDistanceMeter  = "HKQuantityTypeIdentifierDistanceSwimming"
		RecordTypeSleepAnalysisAsleep    = "HKCategoryTypeIdentifierSleepAnalysis"
		RecordTypeSwimmingDistanceKilo   = "HKWorkoutActivityTypeSwimming"
	)
	//filePath := "/Users/semtleie/Documents/fitness/data_sample.xml"
	filePath := "/Users/semtleie/Documents/fitness/data.xml"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	j, err := xj.Convert(bufio.NewReader(file))
	if err != nil {
		panic("That's embarrassing...")
	}
	data := j.Bytes()

	var fitness AppleFitnessAll
	if err := json.Unmarshal(data, &fitness); err != nil {
		panic(err)
	}

	//var fitResult mixedResult

	var (
		swimmingResult mixedResult
		walkingResult  mixedResult
		sleepResult    mixedResult
	)

	for _, item := range fitness.HealthData.Record {
		switch item.Type {
		default:
			continue
		case RecordTypeWalkingRunningDistance:
			t, err := time.Parse(
				"2006-01-02 15:04:05 -0700",
				item.StartDate,
			)
			if err != nil {
				panic(err)
			}
			val, err := strconv.ParseFloat(item.Value, 64)
			if err != nil {
				panic(err)
			}
			walkingResult = append(walkingResult, dailyMixedItem{t, val})
		case RecordTypeSwimmingDistanceMeter:
			t, err := time.Parse(
				"2006-01-02 15:04:05 -0700",
				item.StartDate,
			)
			if err != nil {
				panic(err)
			}
			val, err := strconv.ParseFloat(item.Value, 64)
			if err != nil {
				panic(err)
			}
			swimmingResult = append(swimmingResult, dailyMixedItem{t, val / 1000})
		case RecordTypeSleepAnalysisAsleep:
			t1, err := time.Parse(
				"2006-01-02 15:04:05 -0700",
				item.StartDate,
			)
			if err != nil {
				panic(err)
			}
			t2, err := time.Parse(
				"2006-01-02 15:04:05 -0700",
				item.EndDate,
			)
			if err != nil {
				panic(err)
			}
			duration := t2.Sub(t1)
			sleepResult = append(sleepResult, dailyMixedItem{t1, duration.Minutes()})
		}
	}
	for _, item := range fitness.HealthData.Workout {
		switch item.WorkoutActivityType {
		default:
			continue
		case RecordTypeSwimmingDistanceKilo:
			t, err := time.Parse(
				"2006-01-02 15:04:05 -0700",
				item.StartDate,
			)
			if err != nil {
				panic(err)
			}
			val, err := strconv.ParseFloat(item.MetadataEntry.Value, 64)
			if err != nil {
				panic(err)
			}
			swimmingResult = append(swimmingResult, dailyMixedItem{t, val})
		}
	}

	sort.Sort(swimmingResult)
	//sort.Sort(walkingResult)
	//sort.Sort(sleepResult)
	for _, val := range swimmingResult {
		fmt.Println(val.dt.String(), val.val)
	}
}

func simpleData() {
	filePath := "/Users/semtleie/Documents/fitness/data_cda.xml"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	j, err := xj.Convert(bufio.NewReader(file))
	if err != nil {
		panic("That's embarrassing...")
	}
	data := j.Bytes()

	//var fitness AppleFitness
	var fitness AppleFitness2
	//var fitness AppleFitnessAll
	if err := json.Unmarshal(data, &fitness); err != nil {
		panic(err)
	}

	var fitResult simpleResult
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
		fitResult = append(fitResult, dailySimpleItem{obvTime, f})
	}

	sort.Sort(fitResult)
	for i, val := range fitResult {
		fmt.Printf("%d\t%f\t%s\n", i, val.val, val.dt.String())
	}
}

func readFromFile(filePath string) ([]byte, error) {
	// Read file to byte slice
	return ioutil.ReadFile(filePath)
}
