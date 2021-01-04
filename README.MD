# goexcel

```
docker run --name goexcel -m 60m -d -p 7781:6060 --restart always suconghou/tools:goexcel
```


POST `^/xlsx/(.+)\.xlsx$` JSON数据格式

将直接输出生成的文件流

支持两种模式

## 1. 普通表单模式,上下结构 

最多支持26列

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


两种模式可以同时使用,第二种模式优先级更高.

