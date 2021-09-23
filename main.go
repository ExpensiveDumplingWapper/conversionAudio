package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	id3 "github.com/mikkyang/id3-go"
)

//no tag
var ErrFormat = errors.New("No TAG!")

// type Mp3 struct {
// 	Name string
// 	Size int64
// 	//the mp3 attr
// 	Title   string
// 	Artist  string
// 	Album   string
// 	Year    string
// 	Comment string
// 	Genre   uint8
// }

// var

func main() {
	path := "/Users/lzz/Documents/094熊逸说苏轼/"
	// f, err := os.Open(path)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// defer f.Close()
	// names, err := f.Readdirnames(-1)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	names := make([]string, 100)
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Print(err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			pathname := path + fi.Name() + "/"
			names = append(names, GetAllFile(pathname)...)
			// names = append(names, GetAllFile(pathname+fi.Name()+"/")...)
		} else {
			names = append(names, path+fi.Name())
		}
	}

	for _, name := range names {
		a := strings.Split(name, ".")
		fileExt := strings.ToLower(a[len(a)-1])
		// if fileExt == "mp3" || fileExt == "m4a" {   // m4a 文件会导致文件异常
		if fileExt == "mp3" {
			tmp := strings.Split(name, "/")
			tmpName := strings.ToLower(tmp[len(tmp)-1])
			mp3File, _ := id3.Open(name)
			defer mp3File.Close()
			mp3File.SetTitle(tmpName)
			mp3File.SetAlbum("苏轼")
			mp3File.SetArtist("苏轼")
			fmt.Println(mp3File.Title())
		} else if fileExt == "m4a" {
			// cmd := exec.Command("/usr/local/Cellar/ffmpeg/4.4_2/bin/ffmpeg ", " -i "+pathTmp+tmpName+" -f mp3 -acodec libmp3lame -y "+pathTmp+tmpNameMp3+".mp3")
			// cmd := exec.Command("ffmpeg ", " -i "+pathTmp+tmpName, " -f mp3 ", "-acodec libmp3lame ", " -y "+pathTmp+tmpNameMp3+".mp3")
		}
	}
	//fmt.Printf("%v",names)
}

func GetAllFile(path string) []string {
	rd, err := ioutil.ReadDir(path)
	fmt.Print(err)
	names := make([]string, 100)
	for _, fi := range rd {
		if fi.IsDir() {
			pathname := path + fi.Name() + "/"
			names = append(names, GetAllFile(pathname)...)
		} else {
			names = append(names, path+fi.Name())
		}
	}
	return names
}
