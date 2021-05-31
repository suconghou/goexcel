package xlsx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/suconghou/vfs/util"
)

var (
	indexMap = map[int]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}
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

// Export excel
func Export(w http.ResponseWriter, r *http.Request, match []string) error {
	bs, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 10485760))
	if err != nil {
		if len(bs) <= 2 {
			err = fmt.Errorf("bad request")
		}
	}
	if err != nil {
		util.JSONPut(w, resp{-1, err.Error()})
		return err
	}
	var data *excelData
	err = json.Unmarshal(bs, &data)
	if err != nil {
		util.JSONPut(w, resp{-2, err.Error()})
		return err
	}
	return excel(data, w)
}

func excel(data *excelData, w io.Writer) error {
	f := excelize.NewFile()
	for s, item := range *data {
		for i, label := range item.Categories {
			k := fmt.Sprintf("%s%d", indexMap[i+1], 1)
			f.SetCellValue(s, k, label)
		}
		for i, vv := range item.Values {
			for j, v := range vv {
				k := fmt.Sprintf("%s%d", indexMap[j+1], i+2)
				f.SetCellValue(s, k, v)
			}
		}
		for k, v := range item.DataMaps {
			f.SetCellValue(s, k, v)
		}
	}
	return f.Write(w)
}
