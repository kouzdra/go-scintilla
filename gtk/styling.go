package scintilla

import "C"
import "github.com/kouzdra/go-scintilla/gtk/consts"

type Styling struct {
	sci *Scintilla
}

func (s *Styling) Clear () {
	s.sci.sendMessage (consts.SCI_CLEARDOCUMENTSTYLE, 0, 0)
}

func (s *Styling) Start (pos Pos) {
	s.sci.sendMessage (consts.SCI_STARTSTYLING, arg (pos), 0)
}

func (s *Styling) Set (length uint, style Style) {
	s.sci.sendMessage (consts.SCI_SETSTYLING, arg (length), arg (style))
}

func (s *Styling) Range (style Style, bg, en Pos) {
	s.Start (bg)
	s.Set (uint (int (en)-int (bg)), style)
}

func (s *Styling) GetEnd () uint {
	return s.sci.sendMessage (consts.SCI_GETENDSTYLED, 0, 0)
}

func (s *Styling) ResetDefault () {
	s.sci.sendMessage (consts.SCI_STYLERESETDEFAULT, 0, 0)
}

func (s *Styling) SetFont (style Style, font string) {
	ptr := C.CString(font)
	defer cfree(ptr)
	s.sci.sendMessage (consts.SCI_STYLESETFONT, arg (style), gstring2arg (ptr))
}

func (s *Styling) SetFg (style Style, color Color) {
	s.sci.sendMessage (consts.SCI_STYLESETFORE, arg (style), arg (color))
}

func (s *Styling) SetBg (style Style, color Color) {
	s.sci.sendMessage (consts.SCI_STYLESETBACK, arg (style), arg (color))
}


func (s *Styling) SetUnderline (style Style, u bool) {
	var uu uint
	if u { uu = 1 } else { uu = 0 }
	s.sci.sendMessage (consts.SCI_STYLESETUNDERLINE, arg (style), arg (uu))
}

func (s *Styling) SetItalic (style Style, i bool) {
	var ii uint
	if i { ii = 1 } else { ii = 0 }
	s.sci.sendMessage (consts.SCI_STYLESETITALIC, arg (style), arg (ii))
}

func (s *Styling) SetBold (style Style, b bool) {
	var bb uint
	if b { bb = 1 } else { bb = 0 }
	s.sci.sendMessage (consts.SCI_STYLESETBOLD, arg (style), arg (bb))
}

func (s *Styling) GetAt(pos Pos) Style {
	return Style (s.sci.sendMessage (consts.SCI_GETSTYLEAT, arg (pos), 0))
}

