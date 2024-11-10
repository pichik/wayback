package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"

	"github.com/pichik/go-modules/output"
	"github.com/pichik/go-modules/utils/request"
)

func (iTool *Wayback) SetupFlags() {

	for _, u := range iTool.utils {
		iTool.toolData.Utils = append(iTool.toolData.Utils, u.SetupFlags()...)
	}

	iTool.toolData.Name = "wayback"
	iTool.toolData.AName = `
	█░█░█ ▄▀█ █▄█ █▄▄ ▄▀█ █▀▀ █▄▀
	▀▄▀▄▀ █▀█ ░█░ █▄█ █▀█ █▄▄ █░█
	`
	iTool.toolData.Description = "Urls collected from wayback machine prepared for fuzz.\nClear 'wayback/wb-urls' from useless, repetive endpoints and fuzz it."
	request.SetToolDir(iTool.toolData.Name)
}

func (iTool *Wayback) SetupInput(urls []string) {
	for _, u := range iTool.utils {
		u.SetupData()
	}

	url := urls[0]

	request.FlowStart(pagination(url), *iTool)
}

func pagination(search string) []request.ParsedUrl {
	url := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=%s&output=json&fl=timestamp,original,mimetype,statuscode,length&filter=!mimetype:(warc|image)/.*&collapse=timestamp:8&pageSize=1", search)
	// maxPages := getPages(url)

	// var urls []string
	// for page := 0; page < maxPages; page++ {
	// 	urls = append(urls, fmt.Sprintf("%s&page=%d", url, page))
	// }
	return request.ParseUrls([]string{url})
}

func getPages(url string) int {
	url = fmt.Sprintf("%s&showNumPages=true", url)
	requestData := request.RequestBase
	requestData.ParsedUrl = request.ParseUrl(url)
	request.CreateRequest(&requestData)

	numberRegex := regexp.MustCompile(`[0-9]*`)
	responsePages := numberRegex.FindString(requestData.ResponseBody)

	pages, err := strconv.Atoi(responsePages)
	if err != nil {
		fmt.Printf("%s%s\nayayayay[%d] %s%s", output.Red, err, requestData.ResponseStatus, requestData.ParsedUrl.Url, output.White)
	}
	return pages
}

func (iTool Wayback) Results(requestData request.RequestData, m *sync.Mutex) {
	if requestData.ResponseContentLength == 0 {
		return
	}
	res := []WB{}
	if err := json.Unmarshal(requestData.ResponseBodyBytes, &res); err != nil {
		log.Fatal(err)
	}
	m.Lock()
	request.CustomOutputs(urlProcessing(res), "wb-urls")
	m.Unlock()
}

func urlProcessing(wbs []WB) []string {
	var urls []string

	for _, wb := range wbs {
		if wb.Statuscode != "200" {
			continue
		}

		url := fmt.Sprintf("https://web.archive.org/web/%s/%s", wb.Timestamp, wb.Original)
		urls = append(urls, url)
	}
	return urls
}

type WB struct {
	Timestamp  string
	Original   string
	Mimetype   string
	Statuscode string
	Length     string
}

// custom unmarshal to WB struct
func (r *WB) UnmarshalJSON(p []byte) error {
	var tmp []interface{}
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	r.Timestamp = tmp[0].(string)
	r.Original = tmp[1].(string)
	r.Mimetype = tmp[2].(string)
	r.Statuscode = tmp[3].(string)
	r.Length = tmp[4].(string)

	return nil
}
