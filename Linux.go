package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("ALinux")
	fullScreen := true
	w.Resize(fyne.NewSize(1300, 700))
	dark := false
	// startMenu:=false
	w.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyF11 {
			w.SetFullScreen(fullScreen)
			fullScreen = !fullScreen
		}
	})

	iconNotepad := canvas.NewImageFromFile("Golang-Crazy-Project\\images\\Notepad.png")
	iconNotepad.FillMode = canvas.ImageFillContain
	notepadBtn := widget.NewButton("      \n            \n        ", func() {
		notepad(a)
	})
	notepadbt := container.NewPadded(iconNotepad, notepadBtn)

	iconWeather := canvas.NewImageFromFile("Golang-Crazy-Project\\images\\weather.jpg")
	iconWeather.FillMode = canvas.ImageFillContain
	weatherbt := widget.NewButton("      \n            \n        ", func() {
		weather(a)
	})
	WeatherBtn := container.NewPadded(iconWeather, weatherbt)

	iconCalc := canvas.NewImageFromFile("Golang-Crazy-Project\\imagescalculator.png")
	iconCalc.FillMode = canvas.ImageFillContain
	calcbt := widget.NewButton("      \n            \n        ", func() {
		calculator(a)
	})
	calculatorBtn := container.NewPadded(iconCalc, calcbt)

	iconGallery := canvas.NewImageFromFile("Golang-Crazy-Project\\images\\gallery.png")
	iconGallery.FillMode = canvas.ImageFillContain
	gallerybt := widget.NewButton("      \n            \n        ", func() {
		gallery(a)
	})
	galleryBtn := container.NewPadded(iconGallery, gallerybt)

	iconMusic := canvas.NewImageFromFile("Golang-Crazy-Project\\images\\music.jpg")
	iconMusic.FillMode = canvas.ImageFillContain
	musicbt := widget.NewButton("      \n            \n        ", func() {
		musicPlayer(a)
	})
	musicBtn := container.NewPadded(iconMusic, musicbt)
	iconSettings := canvas.NewImageFromFile("Golang-Crazy-Project\\images\\settings.png")
	iconSettings.FillMode = canvas.ImageFillContain
	settingbt := widget.NewButton("      \n            \n        ", func() {
		setting(a)
	})
	settingBtn := container.NewPadded(iconSettings, settingbt)
    iconTerminal:= canvas.NewImageFromFile("Golang-Crazy-Project\\images\\terminal.png")
	iconTerminal.FillMode = canvas.ImageFillContain
	terminalbt := widget.NewButton("      \n            \n        ", func() {
		setting(a)
	})
	TerminalBtn := container.NewPadded(iconTerminal,terminalbt)

	menuItem1 := &fyne.MenuItem{
		Label: "Notepad",
		Action: func() {
			notepad(a)
		},
	}
	menuItem2 := &fyne.MenuItem{
		Label: "Gallery",
		Action: func() {
			gallery(a)
		},
	}
	menuItem3 := &fyne.MenuItem{
		Label: "Calculator",
		Action: func() {
			calculator(a)
		},
	}
	menuItem4 := &fyne.MenuItem{
		Label: "Music",
		Action: func() {
			musicPlayer(a)
		},
	}
	menuItem5 := &fyne.MenuItem{
		Label: "Weather",
		Action: func() {
			weather(a)
		},
	}
	menuItem6 := &fyne.MenuItem{
		Label: "Full Screen",
		Action: func() {
			w.SetFullScreen(fullScreen)
			fullScreen = !fullScreen
		},
	}
	menuItem7 := &fyne.MenuItem{
		Label: "Dark Mode",
		Action: func() {
			dark = !dark
			if dark {
				a.Settings().SetTheme(theme.DarkTheme())
			} else {
				a.Settings().SetTheme(theme.LightTheme())
			}
		},
	}
	menuItems := []*fyne.MenuItem{menuItem1, menuItem2, menuItem3, menuItem4, menuItem5, menuItem6, menuItem7}
	menu := &fyne.Menu{
		Label: "Main Menu",
		Items: menuItems,
	}
	mainMenu := fyne.NewMainMenu(menu)
	w.SetMainMenu(mainMenu)
	sidebar := container.NewVBox(notepadbt,
		WeatherBtn,
		galleryBtn,
		calculatorBtn,
		musicBtn,
		settingBtn,
	TerminalBtn,)
	img := canvas.NewImageFromFile("Alinux.jpg")
	img.FillMode = canvas.ImageFillStretch
	content := container.NewHSplit(sidebar, img)
	content.SetOffset(0.07)
	w.SetContent(content)
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}

func gallery(a fyne.App) {
	w1 := a.NewWindow("Gallery")
	filepath := "Golang-Crazy-Project\\images"
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			expression := strings.Split(file.Name(), ".")[1]
			if expression == "png" || expression == "jpg" || expression == "jpeg" {
				image := canvas.NewImageFromFile(filepath + "\\" + file.Name())
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}
	tabs.SetTabLocation(container.TabLocationLeading)
	w1.Resize(fyne.NewSize(500, 400))
	w1.SetContent(tabs)
	w1.CenterOnScreen()
	w1.Show()
}

func notepad(a fyne.App) {
	w3 := a.NewWindow("NotePad By Ankit")
	var count int = 0
	openedfiles := binding.BindStringList(
		&[]string{},
	)
	input := widget.NewMultiLineEntry()
	var inputtedtext = make(map[int]string)
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {

		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			openfiledialog := dialog.NewFileOpen(
				func(r fyne.URIReadCloser, _ error) {
					ReadData, _ := ioutil.ReadAll(r)

					output := fyne.NewStaticResource("Open"+strconv.Itoa(count), ReadData)
					val := fmt.Sprintf(output.StaticName)
					openedfiles.Append(val)
					inputtedtext[count] = string(output.StaticContent)
					count++
				}, w3)
			openfiledialog.SetFilter(
				storage.NewExtensionFileFilter([]string{".txt"}),
			)
			openfiledialog.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			saveFileDialog := dialog.NewFileSave(
				func(uc fyne.URIWriteCloser, _ error) {
					textdata := []byte(input.Text)
					uc.Write(textdata)
				}, w3)
			saveFileDialog.SetFileName("New File" + strconv.Itoa(count) + ".txt")
			saveFileDialog.Show()
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

			val := fmt.Sprintf("File %d", count)
			openedfiles.Append(val)
			inputtedtext[count] = ""
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
	listSide.OnSelected = func(id widget.ListItemID) {
		if inputtedtext != nil {
			_, present := inputtedtext[id]
			if present {
				input.SetText(inputtedtext[id])
			}
		}

	}

	listSide.OnUnselected = func(id widget.ListItemID) {
		inputtedtext[id] = input.Text
	}
	list := container.New(layout.NewHBoxLayout(), listSide)
	split := container.NewHSplit(list, input)
	split.SetOffset(0.12)
	side := container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, split)
	w3.Resize(fyne.NewSize(600, 500))
	w3.SetContent(side)
	w3.CenterOnScreen()
	w3.Show()
}

func calculator(a fyne.App) {
	w4 := a.NewWindow("Calculator")
	output := ""
	input := widget.NewLabel(output)
	showHistory := false
	history := ""
	historylabel := widget.NewLabel(history)
	var historyarr []string
	historyBtn := widget.NewButton("History", func() {
		if !showHistory {
			for i := (len(historyarr) - 1); i >= 0; i-- {
				history += historyarr[i]
				history += "\n"
			}
		} else {
			history = ""
		}
		showHistory = !showHistory
		historylabel.SetText(history)
	})
	backBtn := widget.NewButton("Back", func() {
		if output == "error" {
			output = ""
			input.SetText(output)
		}
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	openBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	clostBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	devideBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	nineBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	multiplyBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	minusBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	plusBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() {

		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {

				res := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				reshistory := output + "=" + res
				historyarr = append(historyarr, reshistory)
				output = res
			}
		} else {
			output = "error"
		}

		input.SetText(output)
	})
	w4.SetContent(container.NewVBox(
		input,
		historylabel,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
		),
		container.NewGridWithColumns(4,
			clearBtn,
			openBtn,
			clostBtn,
			devideBtn,
		),
		container.NewGridWithColumns(4,
			sevenBtn,
			eightBtn,
			nineBtn,
			multiplyBtn,
		),
		container.NewGridWithColumns(4,
			fourBtn,
			fiveBtn,
			sixBtn,
			minusBtn,
		),
		container.NewGridWithColumns(4,
			oneBtn,
			twoBtn,
			threeBtn,
			plusBtn,
		),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				dotBtn,
				zeroBtn,
			),
			equalBtn,
		),
	))
	w4.CenterOnScreen()
	w4.Show()
}

func weather(a fyne.App) {
	w5 := a.NewWindow("Weather")
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Karnal&Appid=d0c9519b1dbd29fc02c4333c9faaaaa8")

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	weather, err := UnmarshalWelcome(body)
	img := canvas.NewImageFromFile("weather.jpg")
	img.FillMode = canvas.ImageFillOriginal
	lable1 := canvas.NewText("Weather Details In Karnal", color.Black)
	lable1.TextStyle = fyne.TextStyle{Bold: true}
	label2 := canvas.NewText(fmt.Sprintf("Country: %s", weather.Sys.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed: %.2f ", weather.Wind.Speed), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("Temperature: %2f k ", weather.Main.Temp), color.Black)
	// label5:=canvas.NewText("Country ",weather.Welcome.timezone)
	w5.SetContent(container.NewVBox(
		lable1,
		img,
		label2,
		label3,
		label4,
	))
	w5.CenterOnScreen()
	w5.Show()
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
type diagonal struct {
}

func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}
func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(20, 20)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}
func musicPlayer(a fyne.App) {
	w6 := a.NewWindow("Music Player By Ankit")
	current := ""
	currentlyPlaying := widget.NewLabel(current)
	closeBtn := widget.NewButton("Stop Player", func() {
		speaker.Close()
	})

	var musicfileName []string
	img := canvas.NewImageFromFile("music.png")
	img.FillMode = canvas.ImageFillOriginal
	musicFiles := binding.BindStringList(
		&[]string{},
	)
	filepath := "Golang-Crazy-Project\\music"
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	for _, singlefile := range files {
		if !singlefile.IsDir() {
			val := fmt.Sprintf(singlefile.Name())
			musicFiles.Append(val)
			musicfileName = append(musicfileName, singlefile.Name())
		}
	}
	startBtn := widget.NewButton("Start Player", func() {
		f, _ := os.Open("Golang-Crazy-Project\\music" + current)
		streamer, format, _ := mp3.Decode(f)
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.Play(streamer)
		currentlyPlaying.SetText("Currently Playing :" + current)
	})
	musicList := widget.NewListWithData(musicFiles,
		func() fyne.CanvasObject {
			return widget.NewLabel("Here is the song name that you want to listen Pleaseselect")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	musicList.OnSelected = func(id widget.ListItemID) {
		f, _ := os.Open("Golang-Crazy-Project\\music" + musicfileName[id])
		streamer, format, _ := mp3.Decode(f)
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.Play(streamer)
		currentlyPlaying.SetText("Currently Playing :-" + musicfileName[id])
		current = musicfileName[id]
	}
	w6.Resize(fyne.NewSize(500, 600))
	top := container.NewVBox(
		img,
		currentlyPlaying,
		container.NewGridWithColumns(2,
			startBtn,
			closeBtn,
		),
	)
	content := container.NewVSplit(top, musicList)
	content.SetOffset(0.1)
	w6.SetContent(content)
	w6.CenterOnScreen()
	w6.Show()
	w6.SetOnClosed(func() {
		speaker.Close()
	})
}
func setting(a fyne.App) {
	w7 := a.NewWindow("Settings")
	iconLight := canvas.NewImageFromFile("Golang-Crazy-Project\\images\\light.png")
	iconLight.FillMode = canvas.ImageFillContain
	buttonLight := widget.NewButton("      \n            \n        ", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	lightBtn := container.NewPadded(iconLight, buttonLight)

	iconBlack := canvas.NewImageFromFile("Golang-Crazy-Project\\images\\black.png")
	iconBlack.FillMode = canvas.ImageFillContain
	buttonBlack := widget.NewButton("      \n             \n    \n    ", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	blackBtn := container.NewPadded(iconBlack, buttonBlack)

	content := container.NewVBox(container.NewHBox(lightBtn, blackBtn))
	w7.Resize(fyne.NewSize(190, 130))
	w7.SetContent(content)
	w7.CenterOnScreen()
	w7.Show()
}
