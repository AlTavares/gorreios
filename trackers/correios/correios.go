package correios

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/AlTavares/gotracker"

	"github.com/sirupsen/logrus"

	"github.com/moul/http2curl"
)

const (
	host = "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
)

var log = logrus.New()

var httpClient = &http.Client{
	Timeout: time.Second * 3000000,
}

type payload struct {
	XMLName    xml.Name `xml:"rastroObjeto"`
	ObjectType string   `xml:"tipo"`
	Result     string   `xml:"resultado"`
	Objects    string   `xml:"objetos"`
	Language   string   `xml:"lingua"`
}

type Language string

const (
	PortugueseLanguage = Language("101")
	EnglishLanguage    = Language("102")
)

type ObjectType string

const (
	ListObjectType     = ObjectType("L")
	IntervalObjectType = ObjectType("F")
)

type ResultScope string

const (
	AllResultScope       = ResultScope("T")
	LastEventResultScope = ResultScope("U")
)

func init() {
	log.SetLevel(gotracker.LoggerLevel)
}

func GetTrackingInfo(objects ...string) (trackingInfo TrackingInfo, err error) {
	return GetTrackingInfoWithOptions(PortugueseLanguage, ListObjectType, AllResultScope, objects...)
}

func GetTrackingInfoWithOptions(language Language, objectType ObjectType, resultScope ResultScope, objects ...string) (trackingInfo TrackingInfo, err error) {

	data, err := buildPayloadData(language, objectType, resultScope, objects)
	if err != nil {
		log.WithField("error", err).Error("Error building payload")
		return
	}
	response, err := makeRequestWithData(data)
	if err != nil {
		log.WithField("error", err).Error("Error making request")
		return
	}
	defer response.Body.Close()
	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithField("error", err).Error("Error reading response body")
		return
	}
	err = json.Unmarshal(jsonData, &trackingInfo)
	return
}

func codesFromArray(objects []string) string {
	codes := strings.Join(objects, "")
	return strings.ToUpper(codes)
}

func buildPayloadData(language Language, objectType ObjectType, resultScope ResultScope, objects []string) ([]byte, error) {
	payload := payload{
		ObjectType: string(objectType),
		Result:     string(resultScope),
		Objects:    codesFromArray(objects),
		Language:   string(language),
	}
	return xml.Marshal(payload)
}

func makeRequestWithData(data []byte) (*http.Response, error) {
	request, err := http.NewRequest("POST", host, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/xml")
	request.Header.Add("SOAPAction", "buscaEventosLista")
	curl, _ := http2curl.GetCurlCommand(request)
	log.Debug(curl)
	return httpClient.Do(request)
}
