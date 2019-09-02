package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/pmars/beego/logs"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func ProcessBaidu() {
	initMysqlEngine()
	url := "https://fanyi.baidu.com/v2transapi?from=en&to=zh&query=%v&simple_means_flag=3&sign=%v&token=%v"
	var n int64
	words := make([]*Word, 0)
	for n == 0 || len(words) == 100 {
		words, _ = MGetWordImage(n)
		n = words[len(words)-1].Id
		for _, word := range words {
			if word.Status == 0 && word.Image == "" {

				sign := CallJs(word.Word)
				if sign == "" {
					logs.Error("http.Get: %v")
					continue
				}

				// req, _ := http.NewRequest("GET", "https://fanyi.baidu.com/", nil)
				// req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
				// req.Header.Add("Cookie", "BAIDUID=53029BFDEE93DC1D575CBC0D2B5846CE:FG=1; BIDUPSID=53029BFDEE93DC1D575CBC0D2B5846CE; PSTM=1549007813; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; __cfduid=ded90d635bb6e627640e2ca6d724a1d421556513220; to_lang_often=%5B%7B%22value%22%3A%22en%22%2C%22text%22%3A%22%u82F1%u8BED%22%7D%2C%7B%22value%22%3A%22zh%22%2C%22text%22%3A%22%u4E2D%u6587%22%7D%5D; from_lang_often=%5B%7B%22value%22%3A%22hu%22%2C%22text%22%3A%22%u5308%u7259%u5229%u8BED%22%7D%2C%7B%22value%22%3A%22zh%22%2C%22text%22%3A%22%u4E2D%u6587%22%7D%2C%7B%22value%22%3A%22en%22%2C%22text%22%3A%22%u82F1%u8BED%22%7D%5D; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=1425_21104_29523_29518_29099_29567_29220_26350_22160; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1564135317,1565316432,1565860881,1566203519; CHINA_PINYIN_SWITCH=0; DOUBLE_LANG_SWITCH=0; sideAdClose=18127; Hm_lvt_afd111fa62852d1f37001d1f980b6800=1566204189; BDSFRCVID=rBPOJeC62Ce1xxJwSPza2L93FgPIUgQTH6aofF6TXigrJZBrU4KgEG0PHM8g0Ku-2G8IogKK0mOTHUFF_2uxOjjg8UtVJeC6EG0P3J; H_BDCLCKID_SF=JJktoIIXtCvqKROkq4cE-t4hMMoXetJyaR3H5fJvWJ5TMCoJbjQW2MIW0n3E56jrfR7l0MJYQR3GShPC-tntQ6K7WMrMbfTz0jrP_pvj3l02VhnEe-t2ynQDD4tOB4RMW20jWl7mWU5PsxA45J7cM4IseboJLfT-0bc4KKJxthF0HPonHj8hD53-3D; Hm_lpvt_afd111fa62852d1f37001d1f980b6800=1566262829; MCITY=-131%3A; delPer=0; PSINO=1; ZD_ENTRY=baidu; locale=zh; yjs_js_security_passport=f1bb76747c1e4d750460d8c1d6cfc11c5f5ff8d2_1566288466_js; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1566292881; __yjsv5_shitong=1.0_7_054c1c0c8ddb51c396a0042ebecd864e5c51_300_1566292881617_101.254.182.132_ec88e850")
				// req.Header.Add("Host", "fanyi.baidu.com")
				// req.Header.Add("Origin", "https://fanyi.baidu.com")
				// req.Header.Add("Referer", "https://fanyi.baidu.com/?aldtype=16047")
				//
				// r1, err := (&http.Client{}).Do(req)
				// if err != nil {
				// 	logs.Error("http.Get: %v", err)
				// 	continue
				// }
				// respByte, _ := ioutil.ReadAll(r1.Body)
				// r1.Body.Close()
				//
				// html := string(respByte)
				// start := strings.Index(html, "token: '")
				// token := string([]byte(html)[start+len("token: ")+1:])
				//
				// end := strings.Index(token, "'")
				// token = string([]byte(token)[:end])
				// if token == "" {
				// 	logs.Error("http.Get: %v", err)
				// 	continue
				// }
				token := "e1b50799dcdf006c766a04ce80f0b9c0"

				url := fmt.Sprintf(url, word.Word, sign, token)

				req2, err := http.NewRequest("GET", url, nil)
				if err != nil {
					logs.Error("http.Get: %v", err)
					continue
				}
				req2.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
				req2.Header.Add("Cookie", "BAIDUID=53029BFDEE93DC1D575CBC0D2B5846CE:FG=1; BIDUPSID=53029BFDEE93DC1D575CBC0D2B5846CE; PSTM=1549007813; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; __cfduid=ded90d635bb6e627640e2ca6d724a1d421556513220; to_lang_often=%5B%7B%22value%22%3A%22en%22%2C%22text%22%3A%22%u82F1%u8BED%22%7D%2C%7B%22value%22%3A%22zh%22%2C%22text%22%3A%22%u4E2D%u6587%22%7D%5D; from_lang_often=%5B%7B%22value%22%3A%22hu%22%2C%22text%22%3A%22%u5308%u7259%u5229%u8BED%22%7D%2C%7B%22value%22%3A%22zh%22%2C%22text%22%3A%22%u4E2D%u6587%22%7D%2C%7B%22value%22%3A%22en%22%2C%22text%22%3A%22%u82F1%u8BED%22%7D%5D; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=1425_21104_29523_29518_29099_29567_29220_26350_22160; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1564135317,1565316432,1565860881,1566203519; CHINA_PINYIN_SWITCH=0; DOUBLE_LANG_SWITCH=0; sideAdClose=18127; Hm_lvt_afd111fa62852d1f37001d1f980b6800=1566204189; BDSFRCVID=rBPOJeC62Ce1xxJwSPza2L93FgPIUgQTH6aofF6TXigrJZBrU4KgEG0PHM8g0Ku-2G8IogKK0mOTHUFF_2uxOjjg8UtVJeC6EG0P3J; H_BDCLCKID_SF=JJktoIIXtCvqKROkq4cE-t4hMMoXetJyaR3H5fJvWJ5TMCoJbjQW2MIW0n3E56jrfR7l0MJYQR3GShPC-tntQ6K7WMrMbfTz0jrP_pvj3l02VhnEe-t2ynQDD4tOB4RMW20jWl7mWU5PsxA45J7cM4IseboJLfT-0bc4KKJxthF0HPonHj8hD53-3D; Hm_lpvt_afd111fa62852d1f37001d1f980b6800=1566262829; MCITY=-131%3A; delPer=0; PSINO=1; ZD_ENTRY=baidu; locale=zh; yjs_js_security_passport=f1bb76747c1e4d750460d8c1d6cfc11c5f5ff8d2_1566288466_js; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1566292881; __yjsv5_shitong=1.0_7_054c1c0c8ddb51c396a0042ebecd864e5c51_300_1566292881617_101.254.182.132_ec88e850")
				req2.Header.Add("Host", "fanyi.baidu.com")
				req2.Header.Add("Origin", "https://fanyi.baidu.com")
				req2.Header.Add("Referer", "https://fanyi.baidu.com/?aldtype=16047")
				resp2, err := (&http.Client{}).Do(req2)
				if err != nil {
					logs.Error("http.Get: %v", err)
					continue
				}
				respByte2, err := ioutil.ReadAll(resp2.Body)
				if err != nil {
					logs.Error("http.Get: %v", err)
					continue
				}
				fmt.Println(string(respByte2))
				r := Baidu_word{}
				err = json.Unmarshal(respByte2, &r)
				resp2.Body.Close()

				if r.Dict_result.Baike_img_url != "" {
					logs.Debug("url: ", r.Dict_result.Baike_img_url)
					res3, err := http.Get(r.Dict_result.Baike_img_url)
					if err != nil {
						logs.Error("http.Get: %v", err)
						continue
					}
					path := "./static/image/" + word.Word + ".png"
					fs, err := os.Create(path)
					MUpdateWordImage(word.Word+".png", word.Id)
					io.Copy(fs, res3.Body)
					res3.Body.Close()
					fs.Close()
				}
			}
		}
	}
}

type BaiduBody struct {
	From              string `json:"from"`
	To                string `json:"to"`
	Query             string `json:"query"`
	Simple_means_flag int    `json:"baike_img_url"`
	Sign              string `json:"sign"`
	Token             string `json:"token"`
}

type Baidu_word struct {
	Dict_result Dict_result `json:"dict_result"`
}

type Dict_result struct {
	Baike_img_url string       `json:"baike_img_url"`
	Simple_means  Simple_means `json:"simple_means"`
}

type Simple_means struct {
	SymbolsList []Symbol `json:"symbols"`
	Word_means  string   `json:"word_means"`
}

type Symbol struct {
	Ph_en string `json:"ph_en"`
	Ph_am string `json:"ph_am"`
	parts []struct {
		Part  string `json:"part"`
		Means string `json:"means"`
	}
}
