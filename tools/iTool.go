package tools

import (
	"sync"

	"github.com/pichik/go-modules/tool"
	"github.com/pichik/go-modules/utils/request"
)

type Crawler struct {
	toolData *tool.ToolData
	utils    []request.IUtil
}

type Wayback struct {
	toolData *tool.ToolData
	utils    []request.IUtil
}

type ITool interface {
	SetupFlags()
	SetupInput(urls []string)
}

type IFlowTool interface {
	Results(requestData request.RequestData, m *sync.Mutex)
}

func GetTool(t string) (ITool, tool.Tool) {
	newTool := tool.GetTool(t)
	data := tool.CreateFlagSet(newTool)

	var u []request.IUtil
	switch newTool.Name {
	// case misc.Crawler:
	// 	u = append(u, utils.RequestFlow{UtilData: &misc.UtilData{}})
	// 	u = append(u, utils.Filter{UtilData: &misc.UtilData{}})
	// 	return &Crawler{toolData: data, utils: u}, newTool

	case "wayback":
		u = append(u, request.RequestFlow{UtilData: &tool.UtilData{}})
		u = append(u, request.Filter{UtilData: &tool.UtilData{}})
		return &Wayback{toolData: data, utils: u}, newTool

	default:
		return nil, newTool
	}
}
