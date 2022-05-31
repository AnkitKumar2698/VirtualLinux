package main

import (

	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main(){
	a:=app.New()
	w:=a.NewWindow("Calculator")
	output :=""
	input:=widget.NewLabel(output);
	showHistory:=false;
	history:=""
	historylabel:=widget.NewLabel(history)
	var historyarr [] string 
    historyBtn:=widget.NewButton("History",func(){
		if !showHistory {
         for i:=(len(historyarr)-1);i>=0;i-- {
			history+=historyarr[i]
			history+="\n"
		 }	 
		}else{
			history=""
		}
		showHistory=!showHistory
		historylabel.SetText(history); 
	})
	backBtn:=widget.NewButton("Back",func(){
		if output=="error"{
			output=""
			input.SetText(output)	
		}
		if len(output)>0{
		 output=output[:len(output)-1]
		 input.SetText(output)
		}		
	})
	clearBtn:=widget.NewButton("Clear",func(){
		output=""
		input.SetText(output)
	})
	openBtn:=widget.NewButton("(",func(){
		output=output+"("
		input.SetText(output)
	})
	clostBtn:=widget.NewButton(")",func(){
		output=output+")"
		input.SetText(output)
	})
	devideBtn:=widget.NewButton("/",func(){
		output=output+"/"
		input.SetText(output)	
	})
	nineBtn:=widget.NewButton("9",func(){
		output=output+"9"
		input.SetText(output)
	})
	eightBtn:=widget.NewButton("8",func(){
		output=output+"8"
		input.SetText(output)
	})
	sevenBtn:=widget.NewButton("7",func(){
		output=output+"7"
		input.SetText(output)
	})
	multiplyBtn:=widget.NewButton("*",func(){
		output=output+"*"
		input.SetText(output)
	})
	sixBtn:=widget.NewButton("6",func(){
		output=output+"6"
		input.SetText(output)
	})
	fiveBtn:=widget.NewButton("5",func(){
		output=output+"5"
		input.SetText(output)
	})
	fourBtn:=widget.NewButton("4",func(){
		output=output+"4"
		input.SetText(output)
	})
	minusBtn:=widget.NewButton("-",func(){
		output=output+"-"
		input.SetText(output)
	})
	threeBtn:=widget.NewButton("3",func(){
		output=output+"3"
		input.SetText(output)
	})
	twoBtn:=widget.NewButton("2",func(){
		output=output+"2"
		input.SetText(output)
	})
	oneBtn:=widget.NewButton("1",func(){
		output=output+"1"
		input.SetText(output)
	})
	plusBtn:=widget.NewButton("+",func(){
		output=output+"+"
		input.SetText(output)
	})
	dotBtn:=widget.NewButton(".",func(){
		output=output+"."
		input.SetText(output)
	})
	zeroBtn:=widget.NewButton("0",func(){
		output=output+"0"
		input.SetText(output)
	})
	equalBtn:=widget.NewButton("=",func(){
		
		expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil{
			result, err := expression.Evaluate(nil);
			if err==nil{
				
				res:=strconv.FormatFloat(result.(float64),'f',-1,64)
				reshistory:=output+"="+res
				historyarr=append(historyarr,reshistory)
				output=res
			}
		}else{
			output="error"
		}

		input.SetText(output)
	})
	w.SetContent(container.NewVBox(
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
	w.ShowAndRun()
}