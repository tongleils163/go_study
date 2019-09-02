package pkg

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/pmars/beego/logs"
	"github.com/shamsher31/goimgtype"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

func imageCompress(
	localPath string,
	to string,
	Quality int,
	format string) bool {
	/** 读取文件 */
	file_origin, err := os.Open(localPath)
	defer file_origin.Close()
	if err != nil {
		logs.Error("os.Open(file)错误 err:%v", err)
		return false
	}
	var origin image.Image
	var config image.Config
	var typeImage int
	format = strings.ToLower(format)
	/** jpg 格式 */
	if format == "jpg" || format == "jpeg" {
		typeImage = 1
		origin, err = jpeg.Decode(file_origin)
		if err != nil {
			logs.Error("jpeg.Decode(file_origin) err:%v file:%v", err, localPath)
			return false
		}
		temp, err := os.Open(localPath)
		if err != nil {
			logs.Error("os.Open(temp) err:%v", err)
			return false
		}
		config, err = jpeg.DecodeConfig(temp)
		if err != nil {
			logs.Error("jpeg.DecodeConfig(temp) err:%v", err)
			return false
		}
	} else if format == "png" {
		typeImage = 0
		fileorigin, err := os.Open(localPath)
		defer fileorigin.Close()
		origin, err = png.Decode(fileorigin)
		if err != nil {
			logs.Error("png.Decode(fileorigin) err:", err)
			return false
		}
		temp, err := os.Open(localPath)
		if err != nil {
			logs.Error("os.Open(temp) err:%v", err)
			return false
		}
		config, err = png.DecodeConfig(temp)
		if err != nil {
			logs.Error("png.DecodeConfig(temp) err:%v", err)
			return false
		}
	}
	//logs.Info("typeImage:", typeImage)
	/** 做等比缩放 */
	// width := uint(base) /** 基准 */
	// height := uint(base * config.Height / config.Width)
	//logs.Info("config.Height:", config.Height, "config.With:", config.Width)
	var scale float64
	radio := 400.0 / 400.0
	//fmt.Println("radio:", radio)
	radioImage := float64(config.Width) / float64(config.Height)
	//logs.Info("radioImage:", radioImage)
	if radioImage > radio {
		// 按宽缩放
		scale = 400.0 / float64(config.Width)
	} else {
		// 按高缩放
		scale = 400.0 / float64(config.Height)
	}
	//logs.Info("scale:", scale)
	width := uint(float64(config.Width) * scale)
	height := uint(float64(config.Height) * scale)
	//logs.Info("width:", width, "height:", height)
	// 缩略图
	// canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	canvas := resize.Resize(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	defer file_out.Close()
	if err != nil {
		logs.Error("os create err:%v", err)
		return false
	}
	if typeImage == 0 {
		err = png.Encode(file_out, canvas)
		if err != nil {
			logs.Error("压缩图片失败 err:%v", err)
			return false
		}
	} else {
		err = jpeg.Encode(file_out, canvas, &jpeg.Options{Quality})
		if err != nil {
			logs.Error("压缩图片失败 err:%v", err)
			return false
		}
	}
	return true
}

func CompressImage(path1, path2 string) {
	files, _ := ioutil.ReadDir(path1)
	for _, v := range files {
		// fmt.Println(v.Name())
		datatype, _ := imgtype.Get(path1 + "/" + v.Name())
		// if datatype=="image/gif"{
		// 	fmt.Println( v.Name())
		// }
		if datatype != "image/jpeg" && datatype != "image/png" {
			logs.Error(" :%v", datatype)
		}
		imageCompress(path1+"/"+v.Name(), path2+"/"+v.Name(), 90, "jpeg")
	}
}

/** 是否是图片 */
func isPictureFormat(path string) (string, string, string) {
	temp := strings.Split(path, ".")
	if len(temp) <= 1 {
		return "", "", ""
	}
	index := len(temp)
	logs.Info(" index:", index, "temp:", temp)
	mapRule := make(map[string]int64)
	mapRule["jpg"] = 1
	mapRule["png"] = 1
	mapRule["jpeg"] = 1
	/** 添加其他格式 */
	if mapRule[temp[index-1]] == 1 {
		return path, temp[index-1], temp[index-2]
	} else {
		return "", "", ""
	}
}

func SCompressImage(localPath string) {
	// localPath := "./10.png"
	// outPath := "./27.png"
	outPath := localPath
	oldPath := localPath
	imgExt, err := GetImgExts(localPath)
	if err != nil {
		logs.Error("err:", err)
	}
	path, img_fmat, temp1 := isPictureFormat(localPath)
	logs.Info("imgExt:", imgExt, "path:", path, "temp1:", temp1, "img_fmat:", img_fmat)
	var newPath string
	isChange := false
	if imgExt == ".jpeg" {
		// 生成jpeg的文件
		tempStr := strings.Split(localPath, ".")
		for i, temp := range tempStr {
			if i > 0 {
				if i == len(tempStr)-1 {
					newPath += ".jpeg"
					break
				} else {
					newPath += "." + temp
				}
			}
			isChange = true
		}
		logs.Info("newPath:", newPath)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			logs.Error("111 os rename err:", err)
		}
	}
	// 处理后缀虽然是png的图片但实际不是但图片
	if isChange {
		if !imageCompress(
			newPath,
			newPath,
			75,
			"jpeg") {
			logs.Error("生成缩略图失败")
		} else {
			err := os.Rename(newPath, outPath)
			if err != nil {
				logs.Error("end os rename err:", err)
			} else {
				logs.Info("生成缩略图成功 " + outPath)
			}
		}
	} else {
		if !imageCompress(
			localPath,
			outPath,
			75,
			"png") {
			logs.Error("生成缩略图失败")
		} else {
			logs.Info("生成缩略图成功 " + outPath)
		}
	}
}

func GetImgExts(file string) (ext string, err error) {
	var headerByte []byte
	headerByte = make([]byte, 8)
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer fd.Close()
	_, err = fd.Read(headerByte)
	if err != nil {
		return "", err
	}
	xStr := fmt.Sprintf("%x", headerByte)
	switch {
	case xStr == "89504e470d0a1a0a":
		ext = ".png"
	case xStr == "0000010001002020":
		ext = ".ico"
	case xStr == "0000020001002020":
		ext = ".cur"
	case xStr[:12] == "474946383961" || xStr[:12] == "474946383761":
		ext = ".gif"
	case xStr[:10] == "0000020000" || xStr[:10] == "0000100000":
		ext = ".tga"
	case xStr[:8] == "464f524d":
		ext = ".iff"
	case xStr[:8] == "52494646":
		ext = ".ani"
	case xStr[:4] == "4d4d" || xStr[:4] == "4949":
		ext = ".tiff"
	case xStr[:4] == "424d":
		ext = ".bmp"
	case xStr[:4] == "ffd8":
		ext = ".jpeg"
	case xStr[:2] == "0a":
		ext = ".pcx"
	default:
		ext = ""
	}
	return ext, nil
}
