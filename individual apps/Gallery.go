package main

import (
	"fmt"
	// "image"
	"io/ioutil"
	"log"
	"strings"
	// "path/filepath"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/widget"
)

func main(){
	a:=app.New();
	w:=a.NewWindow("Gallery")
	filepath:="C:\\test"
	files, err := ioutil.ReadDir(filepath)
    if err != nil {
        log.Fatal(err)
    }
	tabs:=container.NewAppTabs()

    for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())	
		if !file.IsDir() {
			expression:=  strings.Split(file.Name(), ".")[1]
		if expression=="png" || expression=="jpg" ||expression=="jpeg" {	
			image := canvas.NewImageFromFile(filepath+"\\"+file.Name())
			tabs.Append(container.NewTabItem(file.Name(),image))
		  } 
		}
    }
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs) 
	w.ShowAndRun()
}