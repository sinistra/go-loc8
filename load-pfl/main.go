package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Footprint struct {
	nbnLocationIdentifier                           string
	gnafPersistentIdentifier                        string
	rolloutRegionIdentifier                         string
	distributionAreaIdentifier                      string
	formattedAddressString                          string
	serviceClass                                    int
	serviceClassDescription                         string
	serviceClassReason                              string
	readyForServiceDate                             string
	disconnectionDate                               string
	unitNumber                                      string
	unitTypeCode                                    string
	levelNumber                                     string
	levelTypeCode                                   string
	addressSiteName                                 string
	roadNumber1                                     string
	roadNumber2                                     string
	lotNumber                                       string
	roadName                                        string
	roadSuffixCode                                  string
	roadTypeCode                                    string
	localityName                                    string
	secondaryComplexName                            string
	complexRoadNumber1                              string
	complexRoadNumber2                              string
	complexRoadName                                 string
	complexRoadTypeCode                             string
	complexRoadSuffixCode                           string
	postcode                                        string
	stateTerritoryCode                              string
	latitude                                        float64
	longitude                                       float64
	isComplexPremiseYn                              bool
	serviceLevelRegion                              string
	serviceType                                     string
	listingType                                     string
	technologyType                                  string
	isEarlyAccessAvailable                          bool
	poiIdentifier                                   string
	poiName                                         string
	transitionalPoiIdentifier                       string
	transitionalPoiName                             string
	connectivityServicingAreaIdentifier             string
	connectivityServicingAreaName                   string
	transitionalConnectivityServicingAreaIdentifier string
	transitionalConnectivityServicingAreaName       string
	newDevelopmentsChargeApplies                    bool
	deltaType                                       string
	lastUpdatedTimestamp                            string
	isCriticalServiceFlagYn                         bool
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

		pfl := ConvertRecordToFootprint(record)
		fmt.Println(pfl)
	}
}

func ConvertRecordToFootprint(record []string) Footprint {

	var f Footprint
	var err error
	for i, field := range record {
		fieldName := m[i]
		switch fieldName {
		case "nbnLocationIdentifier":
			f.nbnLocationIdentifier = field
		case "gnafPersistentIdentifier":
			f.gnafPersistentIdentifier = field
		case "rolloutRegionIdentifier":
			f.rolloutRegionIdentifier = field
		case "distributionAreaIdentifier":
			f.distributionAreaIdentifier = field
		case "formattedAddressString":
			f.formattedAddressString = field
		case "serviceClass":
			f.serviceClass, err = strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
		case "serviceClassDescription":
			f.serviceClassDescription = field
		case "serviceClassReason":
			f.serviceClassReason = field
		case "readyForServiceDate":
			rfsDate, err := time.Parse("2006-01-02", field)
			if err != nil {
				log.Fatal(err)
			}
			f.readyForServiceDate = rfsDate.Format("2006-01-02")
		case "disconnectionDate":
			if field != "" {
				disconnectionDate, err := time.Parse("2006-01-02", field)
				if err != nil {
					log.Fatal(err)
				}
				f.disconnectionDate = disconnectionDate.Format("2006-01-02")
			} else {
				f.disconnectionDate = ""
			}
		case "unitNumber":
			f.unitNumber = field
		case "unitTypeCode":
			f.unitTypeCode = field
		case "levelNumber":
			f.levelNumber = field
		case "levelTypeCode":
			f.levelTypeCode = field
		case "addressSiteName":
			f.addressSiteName = field
		case "roadNumber1":
			f.roadNumber1 = field
		case "roadNumber2":
			f.roadNumber2 = field
		case "lotNumber":
			f.lotNumber = field
		case "roadName":
			f.roadName = field
		case "roadSuffixCode":
			f.roadSuffixCode = field
		case "roadTypeCode":
			f.roadTypeCode = field
		case "localityName":
			f.localityName = field
		case "secondaryComplexName":
			f.secondaryComplexName = field
		case "complexRoadNumber1":
			f.complexRoadNumber1 = field
		case "complexRoadNumber2":
			f.complexRoadNumber2 = field
		case "complexRoadName":
			f.complexRoadName = field
		case "complexRoadTypeCode":
			f.complexRoadTypeCode = field
		case "complexRoadSuffixCode":
			f.complexRoadSuffixCode = field
		case "postcode":
			f.postcode = field
		case "stateTerritoryCode":
			f.stateTerritoryCode = field
		case "latitude":
			f.latitude, err = strconv.ParseFloat(field, 64)
			if err != nil {
				log.Fatal(err)
			}
		case "longitude":
			f.longitude, err = strconv.ParseFloat(field, 64)
			if err != nil {
				log.Fatal(err)
			}
		case "isComplexPremiseYn":
			if field != "" {
				f.isComplexPremiseYn, err = strconv.ParseBool(field)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				f.isComplexPremiseYn = false
			}
		case "serviceLevelRegion":
			f.serviceLevelRegion = field
		case "serviceType":
			f.serviceType = field
		case "listingType":
			f.listingType = field
		case "technologyType":
			f.technologyType = field
		case "isEarlyAccessAvailable":
			if field != "" {
				f.isEarlyAccessAvailable, err = strconv.ParseBool(field)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				f.isEarlyAccessAvailable = false
			}
		case "poiIdentifier":
			f.poiIdentifier = field
		case "poiName":
			f.poiName = field
		case "transitionalPoiIdentifier":
			f.transitionalPoiIdentifier = field
		case "transitionalPoiName":
			f.transitionalPoiName = field
		case "connectivityServicingAreaIdentifier":
			f.connectivityServicingAreaIdentifier = field
		case "connectivityServicingAreaName":
			f.connectivityServicingAreaName = field
		case "transitionalConnectivityServicingAreaIdentifier":
			f.transitionalConnectivityServicingAreaIdentifier = field
		case "transitionalConnectivityServicingAreaName":
			f.transitionalConnectivityServicingAreaName = field
		case "newDevelopmentsChargeApplies":
			if field != "" {
				f.newDevelopmentsChargeApplies, err = strconv.ParseBool(field)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				f.newDevelopmentsChargeApplies = false
			}
		case "deltaType":
			f.deltaType = field
		case "lastUpdatedTimestamp":
			f.lastUpdatedTimestamp = field
		case "isCriticalServiceFlagYn":
			if field == "Y" {
				f.isCriticalServiceFlagYn = true
			} else if field == "N" {
				f.isCriticalServiceFlagYn = false
			} else {
				log.Fatal("unknown boolean value:", field)
			}
		default:
			log.Fatal("unknown column", i, fieldName)
		}
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

func ConvertHeadersToMap(fields []string) map[int]string {

	m := make(map[int]string)

	for i, field := range fields {
		m[i] = SnakeCaseToCamelCase(field)
	}

	return m
}
