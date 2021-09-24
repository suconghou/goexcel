package xlsx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"goexcel/util"

	"github.com/xuri/excelize/v2"
)

// excelData data
type excelData map[string]struct {
	Categories []string               `json:"categories"`
	Values     [][]interface{}        `json:"values"`
	DataMaps   map[string]interface{} `json:"dataMaps"`
}

type resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Export excel max 10MB json payload
func Export(w http.ResponseWriter, r *http.Request, match []string) error {
	var data *excelData
	if err := parse(w, r.Body, 10485760, &data); err != nil {
		_, err = util.JSONPut(w, resp{-1, err.Error()})
		return err
	}
	return excel(data, w)
}

func excel(data *excelData, w io.Writer) error {
	f := excelize.NewFile()
	for s, item := range *data {
		for i, label := range item.Categories {
			k := fmt.Sprintf("%s%d", column(i), 1)
			if err := f.SetCellValue(s, k, label); err != nil {
				return err
			}
		}
		for i, vv := range item.Values {
			for j, v := range vv {
				k := fmt.Sprintf("%s%d", column(j), i+2)
				if err := f.SetCellValue(s, k, v); err != nil {
					return err
				}
			}
		}
		for k, v := range item.DataMaps {
			if err := f.SetCellValue(s, k, v); err != nil {
				return err
			}
		}
	}
	return f.Write(w)
}

func parse(w http.ResponseWriter, r io.ReadCloser, n int64, v interface{}) error {
	bs, err := io.ReadAll(http.MaxBytesReader(w, r, n))
	if err == nil {
		if len(bs) <= 2 {
			err = fmt.Errorf("bad request")
		}
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, &v)
}

func column(i int) string {
	var (
		ret = ""
		fn  func(n int)
	)
	// fn 从下标1计算, 1=A 26=Z
	fn = func(x int) {
		x = x - 1
		var a = x / 26
		var r = x % 26
		ret = string(rune(65+r)) + ret
		if a > 0 {
			fn(a)
		}
	}
	// 我们的函数是0=A,25=Z,所以此处+1
	fn(i + 1)
	return ret
}
