package scintilla

import "C"
import "github.com/kouzdra/go-scintilla/gtk/consts"

func (sci *Scintilla) IndicSetUnder (indic uint, u bool) {
	sci.SendMessage (consts.SCI_INDICSETUNDER, Arg (indic), bool2arg (u))
}

func (sci *Scintilla) IndicGetUnder (indic uint) bool {
	return sci.SendMessage (consts.SCI_INDICGETUNDER, Arg (indic), 0) != 0
}

//--------

func (sci *Scintilla) IndicSetStyle (indic uint, style uint) {
	sci.SendMessage (consts.SCI_INDICSETSTYLE, Arg (indic), Arg (style))
}

func (sci *Scintilla) IndicSetFg (indic uint, color Color) {
	sci.SendMessage (consts.SCI_INDICSETFORE, Arg (indic), Arg (color))
}

//--------

func (sci *Scintilla) IndicSetRange (indic uint, beg Pos, end Pos) {
	sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
	sci.SendMessage (consts.SCI_INDICATORFILLRANGE, Arg (beg), Arg (int (end) - int (beg)))
}

func (sci *Scintilla) IndicClearRange (indic uint, beg Pos, end Pos) {
	sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
	sci.SendMessage (consts.SCI_INDICATORCLEARRANGE, Arg (beg), Arg (int (end) - int (beg)))
}

func (sci *Scintilla) IndicClear (indic uint) {
	sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
	sci.SendMessage (consts.SCI_INDICATORCLEARRANGE, Arg (0), Arg (sci.GetTextLength ()))
}

func (sci *Scintilla) IndicClearAll () {
	length := Arg (sci.GetTextLength ())
	for indic := 0; indic <= consts.INDIC_MAX; indic ++ {
		sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
		sci.SendMessage (consts.SCI_INDICATORCLEARRANGE, Arg (0), length)
	}
}

