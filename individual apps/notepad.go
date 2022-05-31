package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	// "strings"

	// "image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"

	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main(){
	a:=app.New()
    w:=a.NewWindow("NotePad By Ankit")
	var count int =0
	 openedfiles:=binding.BindStringList(
		&[]string{},
	)
	input:=widget.NewMultiLineEntry()
   var inputtedtext= make(map[int]string)
	bar:=widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(),func(){ 
		
		 }),
		widget.NewToolbarAction(theme.FolderOpenIcon(),func(){
			openfiledialog:=dialog.NewFileOpen(
			  	func(r fyne.URIReadCloser,_ error){
					  ReadData,_:=ioutil.ReadAll(r)
					 
					  output:=fyne.NewStaticResource("Open"+strconv.Itoa(count),ReadData)
					  val := fmt.Sprintf(output.StaticName)
					  openedfiles.Append(val)
					  inputtedtext[count]=string(output.StaticContent)
					  count++
				  },w)
				  openfiledialog.SetFilter(
					storage.NewExtensionFileFilter([]string{".txt"}),
				  )	
				  openfiledialog.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(),func(){
			saveFileDialog:=dialog.NewFileSave(
				func(uc fyne.URIWriteCloser, _ error){
					textdata:=[]byte(input.Text)
					uc.Write(textdata)
				},w)
				saveFileDialog.SetFileName("New File"+strconv.Itoa(count)+".txt")
				saveFileDialog.Show()
		}), 
		widget.NewToolbarAction(theme.ContentAddIcon(),func(){
			
				val := fmt.Sprintf("File %d",count)
				openedfiles.Append(val)
				inputtedtext[count]=""
				count++
		}),
	)
	listSide := widget.NewListWithData(openedfiles,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
     listSide.OnSelected=func(id widget.ListItemID){
		 if inputtedtext!=nil{
         _,present:=inputtedtext[id]
		  if present {
			input.SetText(inputtedtext[id])
		  }
	}
		
	}

	 listSide.OnUnselected=func(id widget.ListItemID){
		 inputtedtext[id]=input.Text
	 }
	list:=container.New(layout.NewHBoxLayout(),listSide)
	split:=container.NewHSplit(list,input)
	split.SetOffset(0.12)
	side:=container.New(layout.NewBorderLayout(bar,nil,nil,nil),bar,split)
	w.Resize(fyne.NewSize(600,500))
	w.SetContent(side) 
	w.ShowAndRun()
}