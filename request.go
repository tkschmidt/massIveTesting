package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)



// http://massive.ucsd.edu/ProteoSAFe/proxi/v0.1/spectra?resultType=full&usi=mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2
type MassIVE []struct {
	Status      string   `json:"status"`
	Usi         string   `json:"usi"`
	Intensities []string `json:"intensities"`
	Mzs         []string `json:"mzs"`
}

type RequestResponse struct {
	Body       []byte
	StatusCode int
	Error      error
	RequestUrl string
}

func minimalRequest(endpoint string, usi string) RequestResponse {
	url := endpoint + usi
	// fmt.Println(url)
	proxyReq, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return RequestResponse{Body: body, StatusCode: resp.StatusCode, Error: err, RequestUrl: usi}
}
func worker(endpoint string, usi string, outputChannel chan<- RequestResponse) {
	req := minimalRequest(endpoint, usi)
	outputChannel <- req

}

const endpoint = "http://massive.ucsd.edu/ProteoSAFe/proxi/v0.1/spectra?resultType=full&usi="
const req1 = "mzspec:PXD000394:20130504_EXQ3_MiBa_SA_Fib-2:scan:4234:SGVSRKPAPG/2"
const req2 = "mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2"

func main() {
	numbPtr := flag.Int("numberRequests", 50, "number of requests to two different endpoints each")
	flag.Parse()


	fmt.Printf("%d request against %s and %s, respectivly\n\n", *numbPtr, req1, req2)

	requests := *numbPtr
	outputChannel := make(chan RequestResponse, requests)

	for i := 0; i <= requests; i++ {
		go worker(endpoint, req1, outputChannel )
		go worker(endpoint, req2, outputChannel )
	}

	wrongResponses := 0
	for i := 0; i <= requests; i++ {
		req := <-outputChannel
		if req.StatusCode == 200 {

			res := MassIVE{}
			err := json.Unmarshal(req.Body, &res)
			if err != nil {
				fmt.Println("internal json parsing issue")
			}

			requestPxd := strings.Split(req.RequestUrl, ":")[1]
			responsePxd := strings.Split(res[0].Usi, ":")[1]
			if requestPxd != responsePxd {
				wrongResponses += 1
				fmt.Printf("request %s \nresponse %s\n\n", req.RequestUrl, res[0].Usi)
			}
		}
	}
	fmt.Printf("%d wrong responses", wrongResponses)

}
