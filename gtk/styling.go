package scintilla

import "C"
import "github.com/kouzdra/go-scintilla/gtk/consts"

func (sci *Scintilla) StyleClear () {
	sci.SendMessage (consts.SCI_CLEARDOCUMENTSTYLE, 0, 0)
}

func (sci *Scintilla) StyleStart (pos Pos) {
	sci.SendMessage (consts.SCI_STARTSTYLING, Arg (pos), 0)
}

func (sci *Scintilla) StyleSet (length uint, style Style) {
	sci.SendMessage (consts.SCI_SETSTYLING, Arg (length), Arg (style))
}

func (sci *Scintilla) StyleRange (style Style, bg, en Pos) {
	sci.StyleStart (bg)
	sci.StyleSet (uint (int (en)-int (bg)), style)
}

func (sci *Scintilla) StyleGetEnd () uint {
	return sci.SendMessage (consts.SCI_GETENDSTYLED, 0, 0)
}

func (sci *Scintilla) StyleResetDefault () {
	sci.SendMessage (consts.SCI_STYLERESETDEFAULT, 0, 0)
}

func (sci *Scintilla) StyleSetFont (style Style, font string) {
	ptr := C.CString(font)
	defer cfree(ptr)
	sci.SendMessage (consts.SCI_STYLESETFONT, Arg (style), gstring2arg (ptr))
}

func (sci *Scintilla) StyleSetFg (style Style, color Color) {
	sci.SendMessage (consts.SCI_STYLESETFORE, Arg (style), Arg (color))
}

func (sci *Scintilla) StyleSetBg (style Style, color Color) {
	sci.SendMessage (consts.SCI_STYLESETBACK, Arg (style), Arg (color))
}


func (sci *Scintilla) StyleSetUnderline (style Style, u bool) {
	sci.SendMessage (consts.SCI_STYLESETUNDERLINE, Arg (style), bool2arg (u))
}

func (sci *Scintilla) StyleSetItalic (style Style, i bool) {
	sci.SendMessage (consts.SCI_STYLESETITALIC, Arg (style), bool2arg (i))
}

func (sci *Scintilla) StyleSetBold (style Style, b bool) {
	sci.SendMessage (consts.SCI_STYLESETBOLD, Arg (style), bool2arg (b))
}

func (sci *Scintilla) StyleGetAt(pos Pos) Style {
	return Style (sci.SendMessage (consts.SCI_GETSTYLEAT, Arg (pos), 0))
}

