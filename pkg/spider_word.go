package pkg

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pmars/beego/logs"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Process() {
	initMysqlEngine()
	var n int64
	words := make([]*Word, 0)
	for n == 0 || len(words) == 100 {
		words, _ = MGetWordImage(n)
		n = words[len(words)-1].Id
		for _, word := range words {
			if word.Audio == "" || word.TransShort == "" || word.Symbol == "" {
				res, err := http.Get("http://www.iciba.com/index.php?callback=jQuery19005692850544180779_1563859725655&a=getWordMean&c=search&list=1%2C2%2C3%2C4%2C5%2C8%2C9%2C10%2C12%2C13%2C14%2C15%2C18%2C21%2C22%2C24%2C3003%2C3004%2C3005&word=" + word.Word + "&_=1563859725656")
				if err != nil {
					logs.Error("http.Get ERROR: %v", err)
					res.Body.Close()
					continue
				}
				bytes, err := ioutil.ReadAll(res.Body)

				str := string(bytes)
				logs.Debug("res:%v ", str)
				str = strings.TrimLeft(str, "jQuery19005692850544180779_1563859725655(")
				str = strings.TrimRight(str, ")")

				logs.Debug("res1: ", str)

				r := &Res{}
				err = json.Unmarshal([]byte(str), r)
				if err != nil {
					logs.Debug("Unmarshal: %v", err)
					res.Body.Close()
					continue
				}
				res.Body.Close()
				if r.BaesInfo != nil && len(r.BaesInfo.Symbols) > 0 {
					if word.Audio == "" {

						mp3 := ""
						if mp3 == "" && r.BaesInfo.Symbols[0].PhAmMp3 != "" {
							mp3 = r.BaesInfo.Symbols[0].PhAmMp3
						}
						if mp3 == "" && r.BaesInfo.Symbols[0].PhEnMp3 != "" {
							mp3 = r.BaesInfo.Symbols[0].PhEnMp3
						}
						if mp3 == "" && r.BaesInfo.Symbols[0].PhTtsMp3 != "" {
							mp3 = r.BaesInfo.Symbols[0].PhTtsMp3
						}
						logs.Debug("mp3: ", mp3)
						if mp3 == "" {
							logs.Error("no audio")
						}
						mp3 = strings.ReplaceAll(mp3, `\`, "")

						logs.Debug("mp3:%v ", mp3)

						res2, err := http.Get(mp3)
						if err != nil {
							logs.Error("http.Get: %v", err)
							continue
						}
						path := "./static/audio/" + word.Word + ".mp3"
						fs, err := os.Create(path)
						MUpdateWordAudio(word.Word+".mp3", word.Id)
						io.Copy(fs, res2.Body)
						res2.Body.Close()
						fs.Close()
					}

					if word.Symbol == "" {
						symbol := ""
						if symbol == "" && r.BaesInfo.Symbols[0].PhAm != "" {
							symbol = r.BaesInfo.Symbols[0].PhAm
						}
						if symbol == "" && r.BaesInfo.Symbols[0].PhEn != "" {
							symbol = r.BaesInfo.Symbols[0].PhEn
						}
						if symbol == "" && r.BaesInfo.Symbols[0].PhOther != "" {
							symbol = r.BaesInfo.Symbols[0].PhOther
						}
						logs.Debug("symbol: ", symbol)
						if symbol == "" {
							logs.Error("no symbol")

						}
						if symbol != "" {
							MUpdateWordSymbol("["+symbol+"]", word.Id)
						}

					}

					if word.TransShort == "" {
						trans := ""
						transS := ""
						transA := make([]string, 0)
						for _, v := range r.BaesInfo.Symbols[0].Parts {
							transA = append(transA, v.Part+" "+strings.Join(v.Means, "，"))
						}
						trans = strings.Join(transA, "\n")
						if len(r.BaesInfo.Symbols[0].Parts) > 0 {
							transS = strings.Join(r.BaesInfo.Symbols[0].Parts[0].Means, "，")
						}
						logs.Debug("trans: ", trans)
						if trans == "" {
							logs.Error("no trans")

						}
						MUpdateWordTrans(transS, trans, word.Id)
					}
				}
			}
		}
	}

}

type WY struct {
	Datas []*Data
}

type Data struct {
	Url string `json:"Url"`
}
type Res struct {
	BaesInfo *BaesInfo `json:"baesInfo"`
}
type BaesInfo struct {
	Symbols []*Symbols `json:"symbols"`
}
type Symbols struct {
	PhEn     string    `json:"ph_en"`
	PhOther  string    `json:"ph_other"`
	PhAm     string    `json:"ph_am"`
	PhEnMp3  string    `json:"ph_en_mp3"`
	PhAmMp3  string    `json:"ph_am_mp3"`
	PhTtsMp3 string    `json:"ph_tts_mp3"`
	Parts    []PartObj `json:"parts"`
}

type PartObj struct {
	Part  string   `json:"part"`
	Means []string `json:"means"`
}
type Word struct {
	Id          int64  `json:"id,string" redis:"id"`
	Status      int64  `json:"status,string" redis:"status"`
	Word        string `json:"word" redis:"word"`
	Symbol      string `json:"symbol" redis:"symbol"`
	Discription string `json:"discription" redis:"discription"`
	Translation string `json:"translation" redis:"translation"`
	TransShort  string `json:"trans_short" redis:"trans_short"`
	Image       string `json:"image" redis:"image"`
	Audio       string `json:"audio" redis:"audio"`
	Video       string `json:"video" redis:"video"`
	Level       string `json:"level" redis:"level"`
	Exist       bool   `json:"exist,string" redis:"exist"` // 是否是新词，是否在用户单词本里面
}

var sqlbyWord1 = `select * from dr_word where word =?`
var sqlbyWord = `select * from dr_word where word COLLATE utf8mb4_bin=?`
var sql7 = `update dr_word  set translation=? ,discription=? where id=?`
var sql8 = `insert into dr_word    (word,translation  ,discription) values(?,?,?)`

var sql1 = `select * from dr_word  where id>? and status =0 and audio = '' limit 100`
var sql4 = `select * from dr_word  where id>?  limit 100`
var sql2 = `update dr_word  set audio=? ,status=status+2 where id=?`

var sql5 = `update dr_word  set symbol=? ,status=status+1 where id=?`

var sql6 = `update dr_word  set trans_short=? ,translation=? ,status=status+8 where id=?`

var sql3 = `update dr_word  set image=? ,status=4 where id=?`

func MGetWord(n int64) ([]*Word, error) {
	wordList := make([]*Word, 0)
	err := Engine.SQL(sql1, n).Find(&wordList)
	if err != nil {
		logs.Error("MGetWord ERROR %v ", err)
	}
	return wordList, err
}

func MGetWordImage(n int64) ([]*Word, error) {
	wordList := make([]*Word, 0)
	err := Engine.SQL(sql4, n).Find(&wordList)
	if err != nil {
		logs.Error("MGetWord ERROR %v ", err)
	}
	return wordList, err
}

func MUpdateWordAudio(mp3 string, id int64) error {
	_, err := Engine.Exec(sql2, mp3, id)
	if err != nil {
		logs.Error("MUpdateWord  wordId%v  ERROR:%v", id, err)
		return err
	}
	return err
}

func MUpdateWordSymbol(symbol string, id int64) error {
	_, err := Engine.Exec(sql5, symbol, id)
	if err != nil {
		logs.Error("MUpdateWordSymbol  wordId%v  ERROR:%v", id, err)
		return err
	}
	return err
}

func MUpdateWordTrans(transS, trans string, id int64) error {
	_, err := Engine.Exec(sql6, transS, trans, id)
	if err != nil {
		logs.Error("MUpdateWordTrans  wordId%v  ERROR:%v", id, err)
		return err
	}
	return err
}

func MUpdateWordImage(image string, id int64) error {
	_, err := Engine.Exec(sql3, image, id)
	if err != nil {
		logs.Error("MUpdateWordImage  wordId%v  ERROR:%v", id, err)
		return err
	}
	return err
}

func MGetWordByword(word string) ([]*Word, error) {
	wordList := make([]*Word, 0)
	err := Engine.SQL(sqlbyWord, word).Find(&wordList)
	if err != nil {
		logs.Error("MGetWord ERROR %v ", err)
	}
	return wordList, err
}

func MGetWordByword1(word string) ([]*Word, error) {
	wordList := make([]*Word, 0)
	err := Engine.SQL(sqlbyWord1, word).Find(&wordList)
	if err != nil {
		logs.Error("MGetWordByword1 ERROR %v ", err)
	}
	return wordList, err
}

func MUpdateWord(trans, des string, id int64) error {
	_, err := Engine.Exec(sql7, trans, des, id)
	if err != nil {
		logs.Error("MUpdateWord  wordId%v  ERROR:%v", id, err)
		return err
	}
	return err
}

func MInsertWord(trans, des, word string) error {
	_, err := Engine.Exec(sql8, word, trans, des)
	if err != nil {
		logs.Error("MInsertWord  wordId%v  ERROR:%v", word, err)
		return err
	}
	return err
}
