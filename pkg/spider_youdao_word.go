package pkg

import (
	"encoding/json"
	"github.com/pmars/beego/logs"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func SpiderYoudao() {

	initMysqlEngine()
	var n int64
	words := make([]*Word, 0)
	for n == 0 || len(words) == 100 {
		words, _ = MGetWordImage(n)
		n = words[len(words)-1].Id
		for _, word := range words {
			if word.Image == "" {
				res, err := http.Get("https://dict.youdao.com/ugc/wordjson/" + word.Word)
				if err != nil {
					logs.Error("http.Get ERROR: %v", err)
					res.Body.Close()
					continue
				}
				bytes, err := ioutil.ReadAll(res.Body)
				str := string(bytes)
				if str != "[]" {
					logs.Debug("res:%v ", str)
					r := make([]Data, 0)
					err = json.Unmarshal(bytes, &r)
					if err != nil {
						logs.Debug("Unmarshal: %v", err)
						res.Body.Close()
						continue
					}

					if len(r) > 0 {
						logs.Debug("url: ", r[0].Url)
						res2, err := http.Get(r[0].Url)
						if err != nil {
							logs.Error("http.Get: %v", err)
							continue
						}
						path := "./static/image/" + word.Word + ".png"
						fs, err := os.Create(path)
						MUpdateWordImage(word.Word+".png", word.Id)
						io.Copy(fs, res2.Body)
						res2.Body.Close()
						fs.Close()
					}
				}
			}
			// if word.Status == 0 && word.Audio == "" {
			// 	res, err := http.Get("http://www.iciba.com/index.php?callback=jQuery19005692850544180779_1563859725655&a=getWordMean&c=search&list=1%2C2%2C3%2C4%2C5%2C8%2C9%2C10%2C12%2C13%2C14%2C15%2C18%2C21%2C22%2C24%2C3003%2C3004%2C3005&word=" + word.Word + "&_=1563859725656")
			// 	if err != nil {
			// 		logs.Error("http.Get ERROR: %v", err)
			// 		res.Body.Close()
			// 		continue
			// 	}
			// 	bytes, err := ioutil.ReadAll(res.Body)
			//
			// 	str := string(bytes)
			// 	logs.Debug("res:%v ", str)
			// 	str = strings.TrimLeft(str, "jQuery19005692850544180779_1563859725655(")
			// 	str = strings.TrimRight(str, ")")
			//
			// 	logs.Debug("res1: ", str)
			//
			// 	r := &Res{}
			// 	err = json.Unmarshal([]byte(str), r)
			// 	if err != nil {
			// 		logs.Debug("Unmarshal: %v", err)
			// 		res.Body.Close()
			// 		continue
			// 	}
			// 	res.Body.Close()
			// 	if r.BaesInfo != nil && len(r.BaesInfo.Symbols) > 0 {
			// 		mp3 := ""
			// 		if mp3 == "" && r.BaesInfo.Symbols[0].PhAmMp3 != "" {
			// 			mp3 = r.BaesInfo.Symbols[0].PhAmMp3
			// 		}
			// 		if mp3 == "" && r.BaesInfo.Symbols[0].PhEnMp3 != "" {
			// 			mp3 = r.BaesInfo.Symbols[0].PhEnMp3
			// 		}
			// 		if mp3 == "" && r.BaesInfo.Symbols[0].PhTtsMp3 != "" {
			// 			mp3 = r.BaesInfo.Symbols[0].PhTtsMp3
			// 		}
			// 		logs.Debug("mp3: ", mp3)
			// 		if mp3 == "" {
			// 			logs.Error("no audio")
			// 			continue
			// 		}
			// 		mp3 = strings.ReplaceAll(mp3, `\`, "")
			//
			// 		logs.Debug("mp3:%v ", mp3)
			//
			// 		res2, err := http.Get(mp3)
			// 		if err != nil {
			// 			logs.Error("http.Get: %v", err)
			// 			continue
			// 		}
			// 		path := "./static/audio/" + word.Word + ".mp3"
			// 		fs, err := os.Create(path)
			// 		MUpdateWord( word.Word + ".mp3", word.Id)
			// 		io.Copy(fs, res2.Body)
			// 		res2.Body.Close()
			// 		fs.Close()
			// 	}
			//
			// }
		}
	}

}
