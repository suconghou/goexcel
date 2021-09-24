# goexcel

```
docker run --name goexcel --log-opt max-size=1m -m 100m -d -p 7781:6060 --restart always suconghou/tools:goexcel
```


POST `^/xlsx/(.+)\.xlsx$` JSON数据格式

将直接输出生成的文件流

支持两种模式

## 1. 普通表单模式,上下结构 


```json
{
	"Sheet1":{
		"categories":[
			"姓名",
			"入学年份",
			"专业",
			"学历层次",
			"工作单位",
			"职务",
			"就餐",
			"手机号"
		],
		"values":[
			["hello","hihi","aa","3434",677,"6767","666"],
			["hello22","hihi22","aa22","343422","67722","676722","66622"]
		]
	}
}
```

## 2. 任意单元格模式

```json
{
	"Sheet1":{
        "dataMaps":{
            "A1":"姓名",
            "B1":"年龄",
            "C1":"手机号",
            "A2":"1",
            "B2":"2",
            "C2":"3"
        }
	}
}
```


两种模式可以同时使用,第二种模式优先级更高,若有重复会覆盖第一种模式单元格的值.

```json
{
	"Sheet1":{
		"categories":[
			"姓名",
			"入学年份",
			"专业",
			"学历层次",
			"工作单位",
			"职务",
			"就餐",
			"手机号"
		],
		"values":[
			["hello","hihi","aa","3434",677,"6767","666"],
			["hello22","hihi22","aa22","343422","67722","676722","66622"]
		],
		"dataMaps":{
            "A5":"A5",
            "B5":"B5",
            "C5":"C5",
            "A9":"A9",
            "B9":"B9",
            "C9":"C9"
        }
	}
}
```


参考 

https://xuri.me/excelize/en/



