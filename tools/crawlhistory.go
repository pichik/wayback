package tools

// import (
// 	"sync"

// 	"../misc"
// 	"../utils"
// )

// func (iTool *Crawler) Setup(urls []string) {

// 	utils.FlowStart(misc.ParseUrls(urls), *iTool)
// }

// func (iTool Crawler) Results(requestData misc.RequestData, m *sync.Mutex) {
// 	// utils.FlowResults(requestData, m)

// 	foundData, completeUrls, incompleteUrls := misc.GetData(requestData.ResponseBody, &requestData.ParsedUrl)

// 	if requestData.ResponseHeaders.Get("Location") != "" {
// 		data, comp, incomp := misc.GetData(requestData.ResponseHeaders.Get("Location"), &requestData.ParsedUrl)
// 		foundData = append(foundData, data...)
// 		completeUrls = append(completeUrls, comp...)
// 		incompleteUrls = append(incompleteUrls, incomp...)
// 	}

// 	urlToSave := requestData.ParsedUrl.Url
// 	// m.Lock()
// 	misc.AddToTested(urlToSave)
// 	misc.DataOutput(foundData, misc.GetUrls(completeUrls), misc.GetUrls(incompleteUrls))

// 	if requestData.ResponseStatus != 404 && requestData.ResponseStatus != 405 && requestData.ResponseContentLength != 0 {
// 		misc.ResponseOutput(requestData)
// 	}
// }
