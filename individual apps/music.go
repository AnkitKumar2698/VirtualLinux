package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)
func main() {
	a := app.New()
	w := a.NewWindow("Music Player By Ankit")
	current:=""
	currentlyPlaying:=widget.NewLabel(current)
	closeBtn:=widget.NewButton("Stop Player",func(){
		speaker.Close()
	})

	var musicfileName [] string
	img:=canvas.NewImageFromFile("music.png")
	img.FillMode=canvas.ImageFillOriginal
	musicFiles := binding.BindStringList(
		&[]string{},
	)
	filepath := "C:\\Users\\Ankit Sharma\\Music"
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	for _, singlefile := range files {
		if !singlefile.IsDir() {
			val := fmt.Sprintf(singlefile.Name())
			musicFiles.Append(val)
			musicfileName=append(musicfileName, singlefile.Name())
		}
	}
	startBtn:=widget.NewButton("Start Player",func(){
		f, _ := os.Open("C:\\Users\\Ankit Sharma\\Music\\"+current)
		streamer, format, _ := mp3.Decode(f)
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.Play(streamer)
		currentlyPlaying.SetText("Currently Playing :"+current)	
	})
	musicList := widget.NewListWithData(musicFiles,
		func() fyne.CanvasObject {
			return widget.NewLabel("Here is the song name that you want to listen Pleaseselect")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	musicList.OnSelected = func(id widget.ListItemID) {
		f, _ := os.Open("C:\\Users\\Ankit Sharma\\Music\\"+musicfileName[id])
		streamer, format, _ := mp3.Decode(f)
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.Play(streamer)
		currentlyPlaying.SetText("Currently Playing :"+musicfileName[id])
	    current=musicfileName[id]
	}
	w.Resize(fyne.NewSize(500,700))
	top:=container.NewVBox(
		img,
		container.NewGridWithColumns(2,
			startBtn,
			closeBtn,
		),
		currentlyPlaying,
	)
    content:=container.NewVSplit(top,musicList)
	content.SetOffset(0.1)
	w.SetContent(content)
	w.ShowAndRun()
    
}