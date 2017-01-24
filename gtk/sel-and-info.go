package scintilla

//import "github.com/mattn/go-gtk/gtk"
import "github.com/kouzdra/go-scintilla/gtk/consts"


func (sci *Scintilla) LinesOnScreen() int {
	return int (sci.SendMessage (consts.SCI_LINESONSCREEN, 0, 0))
}

func (sci *Scintilla) GetModify() bool {
	return int (sci.SendMessage (consts.SCI_GETTEXTLENGTH, 0, 0)) != 0
}

func (sci *Scintilla) GetTextLength() int {
	return int (sci.SendMessage (consts.SCI_GETTEXTLENGTH, 0, 0))
}

func (sci *Scintilla) GetLength() int {
	return int (sci.SendMessage (consts.SCI_GETLENGTH, 0, 0))
}

func (sci *Scintilla) GetLinesCount() int {
	return int (sci.SendMessage (consts.SCI_GETLINECOUNT, 0, 0))
}

func (sci *Scintilla) GetCurrentPos() Pos {
	return Pos (sci.SendMessage (consts.SCI_GETCURRENTPOS, 0, 0))
}

func (sci *Scintilla) GetSelText() string {
	length := sci.SendMessage (consts.SCI_GETSELTEXT, 0, 0)
	if length == 0 { return "" }
	res := make_string (length+1)
	defer cfree (res)
	sci.SendMessage (consts.SCI_GETSELTEXT, 0, gstring2arg (res))
	return gocstring (res)
}
