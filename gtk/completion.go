package scintilla

import "github.com/kouzdra/go-scintilla/gtk/consts"

func (sci *Scintilla) AutoCShow(len uint, list string) byte {
	ptr := cstring(list)
	defer cfree(ptr)
	return byte (sci.SendMessage (consts.SCI_AUTOCSHOW, Arg (len), gstring2arg (ptr)))
}

func (sci *Scintilla) AutoCGetSeparator() byte {
	return byte (sci.SendMessage (consts.SCI_AUTOCGETSEPARATOR, 0, 0))
}

func (sci *Scintilla) AutoCSetSeparator(sep byte) {
	sci.SendMessage (consts.SCI_AUTOCSETSEPARATOR, Arg (sep), 0)
}


func (sci *Scintilla) AutoCGetTypeSeparator() byte {
	return byte (sci.SendMessage (consts.SCI_AUTOCGETTYPESEPARATOR, 0, 0))
}

func (sci *Scintilla) AutoCSetTypeSeparator(sep byte) {
	sci.SendMessage (consts.SCI_AUTOCSETTYPESEPARATOR, Arg (sep), 0)
}


func (sci *Scintilla) AutoCGetDropRestOfWold() bool {
	return sci.SendMessage (consts.SCI_AUTOCGETDROPRESTOFWORD, 0, 0) != 0
}

func (sci *Scintilla) AutoCSetDropRestOfWord(drop bool) {
	sci.SendMessage (consts.SCI_AUTOCSETDROPRESTOFWORD, bool2arg (drop), 0)
}
