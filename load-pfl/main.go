package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type footprint struct {
	nbnLocationIdentifier 		string
	gnafPersistentIdentifier 	string
	rolloutRegionIdentifier 	string
	distributionAreaIdentifier 	string
	formattedAddressString 		string
	serviceClass 				int
	serviceClassDescription 	string
	serviceClassReason 			string
	readyForServiceDate			string
	disconnectionDate 			string
	unitNumber 					string
	unitTypeCode 				string
	levelNumber 				string
	levelTypeCode 				string
	addressSiteName 			string
	roadNumber1 				string
	roadNumber2 				string
	lotNumber 					string
	roadName 					string
	roadSuffixCode 				string
	roadTypeCode				string
	localityName				string
	secondaryComplexName		string
	complexRoadNumber1			string
	complexRoadNumber2			string
	complexRoadName				string
	complexRoadTypeCode			string
	complexRoadSuffixCode		string
	postcode					string
	stateTerritoryCode			string
	latitude					float64
	longitude					float64
	isComplexPremiseYn			bool
	serviceLevelRegion			string
	serviceType					string
	listingType					string
	technologyType				string
	isEarlyAccessAvailable		bool
	poiIdentifier				string
	poiName						string
	transitionalPoiIdentifier	string
	transitionalPoiName			string
	connectivityServicingAreaIdentifier	string
	connectivityServicingAreaName	string
	transitionalConnectivityServicingAreaIdentifier	string
	transitionalConnectivityServicingAreaName	string
	newDevelopmentsChargeApplies	bool
	deltaType					string
	lastUpdatedTimestamp		string
	isCriticalServiceFlagYn		bool
}

var m map[int]string

func main() {
	fileName := path.Join("..", "data", "SR_HFL_PFL_Delta_CSV_20190409", "SR_SAT_ADDR_HFL_PFL_Delta_20190409_201407.csv")
	fmt.Println(fileName)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','
	r.Comment = '#'

	headers, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	m = ConvertHeadersToMap(headers)
	fmt.Println(m)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)

		pfl := ConvertRecordToFootprint(record)
	}
}

func ConvertRecordToFootprint(record []string) footprint {

	var f footprint
	for i, field := range record {
		f.(m[i]) = field
	}

	return f
}

// snakeCaseToCamelCase takes a snake_case string and returns a camelCase string
func SnakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
	isToUpper := false

	for _, v := range inputUnderScoreStr {
		if isToUpper {
			camelCase += strings.ToUpper(string(v))
			isToUpper = false
		} else {
			if v == '_' {
				isToUpper = true
			} else {
				camelCase += strings.ToLower(string(v))
			}
		}
	}
	return
}

func ConvertHeadersToMap (fields []string) map[int]string {

	m:= make(map[int]string)

	for i, field := range fields {
		m[i] = SnakeCaseToCamelCase(field)
	}

	return m
}