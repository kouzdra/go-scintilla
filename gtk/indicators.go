package scintilla

import "C"
import "github.com/kouzdra/go-scintilla/gtk/consts"

type Indic struct {
	Sci *Scintilla
}

func (i *Indic) SetUnder (indic uint, u bool) {
	i.Sci.SendMessage (consts.SCI_INDICSETUNDER, Arg (indic), bool2arg (u))
}

func (i *Indic) GetUnder (indic uint) bool {
	return i.Sci.SendMessage (consts.SCI_INDICGETUNDER, Arg (indic), 0) != 0
}

//--------

func (i *Indic) SetStyle (indic uint, style uint) {
	i.Sci.SendMessage (consts.SCI_INDICSETSTYLE, Arg (indic), Arg (style))
}

func (i *Indic) SetFg (indic uint, color Color) {
	i.Sci.SendMessage (consts.SCI_INDICSETFORE, Arg (indic), Arg (color))
}

//--------

func (i *Indic) SetRange (indic uint, beg Pos, end Pos) {
	i.Sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
	i.Sci.SendMessage (consts.SCI_INDICATORFILLRANGE, Arg (beg), Arg (int (end) - int (beg)))
}

func (i *Indic) ClearRange (indic uint, beg Pos, end Pos) {
	i.Sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
	i.Sci.SendMessage (consts.SCI_INDICATORCLEARRANGE, Arg (beg), Arg (int (end) - int (beg)))
}

func (i *Indic) ClearIndic (indic uint) {
	i.Sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
	i.Sci.SendMessage (consts.SCI_INDICATORCLEARRANGE, Arg (0), Arg (i.Sci.GetTextLength ()))
}

func (i *Indic) Clear () {
	length := Arg (i.Sci.GetTextLength ())
	for indic := 0; indic <= consts.INDIC_MAX; indic ++ {
		i.Sci.SendMessage (consts.SCI_SETINDICATORCURRENT, Arg (indic), 0)
		i.Sci.SendMessage (consts.SCI_INDICATORCLEARRANGE, Arg (0), length)
	}
}

