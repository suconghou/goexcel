package route

import (
	"net/http"
	"regexp"

	"github.com/suconghou/goexcel/xlsx"
)

// 路由定义
type routeInfo struct {
	Reg     *regexp.Regexp
	Handler func(http.ResponseWriter, *http.Request, []string) error
}

// Rules route
var Rules = []routeInfo{
	{regexp.MustCompile(`^/xlsx/(.+)\.xlsx$`), xlsx.Export},
}
