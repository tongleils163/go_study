package pkg

import (
	"github.com/pmars/beego/logs"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"strings"
)

func ImportWord(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, v := range files {
		xlFile, err := xlsx.OpenFile(path + "/" + v.Name())
		if v.Name() == ".DS_Store" {
			continue
		}
		if err != nil {
			logs.Error("ImportWord openFile err:%v", err)
			// return
		}

		sheet := xlFile.Sheets[0]
		// if sheet.Rows[0].Cells[0].Value != "key_json" ||
		// 	sheet.Rows[0].Cells[1].Value != "display" ||
		// 	sheet.Rows[0].Cells[2].Value != "discription" ||
		// 	sheet.Rows[0].Cells[3].Value != "Chinese discription" {
		// 	logs.Error("ImportWord wrong header, wanted:%v", "key_json, display,discription,Chinese discription ")
		// 	logs.Error("err file:%v", v.Name())
		// 	continue
		// }
		for _, row := range sheet.Rows {
			if len(row.Cells) < 4 || row.Cells[0].Value == "key_json" || strings.Trim(row.Cells[1].Value, " ") == "" {
				continue
			}

			words, err := MGetWordByword(strings.Trim(row.Cells[1].Value, " "))
			if err != nil {
				continue
			}
			if len(words) == 0 {
				words, err = MGetWordByword1(strings.Trim(row.Cells[1].Value, " "))
				if err != nil {
					continue
				}
			}
			if len(words) > 1 {
				logs.Error("ImportWord MGetWordByword is not one:%v, len :%v", strings.Trim(row.Cells[1].Value, " "), len(words))
				continue
			}

			discription := row.Cells[2].Value
			transShort := row.Cells[3].Value

			discription = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(discription,
				"<p>", ""),
				"</p>", ""),
				"<ol>", ""),
				"</ol>", ""),
				"<li>", ""),
				"</li>", "")

			transShort = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(transShort,
				"<p>", ""),
				"</p>", ""),
				"<ol>", ""),
				"</ol>", ""),
				"<li>", ""),
				"</li>", "")
			if len(words) == 1 {
				err = MUpdateWord(transShort, discription, words[0].Id)
				if err != nil {
					logs.Error("ImportWord MUpdateWord word:%v, file :%v", words[0].Word, v.Name())

				}
			} else {
				err = MInsertWord(transShort, discription, strings.Trim(row.Cells[1].Value, " "))
				if err != nil {
					logs.Error("ImportWord MInsertWord word:%v, file :%v", words[0].Word, v.Name())
				}
			}
		}

	}
}
