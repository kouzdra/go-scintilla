// +build !cgocheck

package gtk

// #cgo pkg-config: gmodule-2.0 gtk+-2.0
// #cgo LDFLAGS: -lm -lstdc++ -L/home/msk/local/lib -lscintilla
// #cgo CFLAGS: -I${SRCDIR}/../scintilla/include
// #include "../../../mattn/go-gtk/gtk/gtk.go.h"
// #include "gtkscintilla.go.h"
import "C"
import (
	"unsafe"
	"log"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//_ = unsafe


func gint(v int) C.gint           { return C.gint(v) }
func from_gint(v C.gint) int           { return int(v) }
func guint(v uint) C.guint        { return C.guint(v) }
func guint16(v uint16) C.guint16  { return C.guint16(v) }
func guint32(v uint32) C.guint32  { return C.guint32(v) }
func gulong(v uint) C.gulong    { return C.gulong(v) }
func glong(v int) C.glong       { return C.glong(v) }
func from_glong(v C.glong) int  { return int (v) }
func gdouble(v float64) C.gdouble { return C.gdouble(v) }
func gsize_t(v C.size_t) C.gint   { return C.gint(v) }

func guintptr(v uint) C.guintptr        { return C.guintptr(v) }
func gintptr(v int) C.gintptr        { return C.gintptr(v) }

func from_gintptr(v C.gintptr) int           { return int(v) }
func from_guintptr(v C.guintptr) uint           { return uint(v) }


func gstring(s *C.char) *C.gchar { return C.toGstr(s) }
func gstring2uint(s *C.char) uint { return uint (C.toGstrUint(s)) }
func cstring(s *C.gchar) *C.char { return C.toCstr(s) }
func gostring(s *C.gchar) string { return C.GoString(cstring(s)) }

func gslist(l *glib.SList) *C.GSList {
	if l == nil {
		return nil
	}
	return C.to_gslist(unsafe.Pointer(l.ToSList()))
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

func cfree(s *C.char) { C.freeCstr(s) }

//-----------------------------------------------------------------------
// GtkScintilla
//-----------------------------------------------------------------------

type Scintilla struct {
	gtk.Container
}

func (v *Scintilla) ToNativeScintilla() *C.ScintillaObject {
	return C.toGtkScintilla(unsafe.Pointer(v.GWidget))
}

func NewScintilla() *Scintilla {
	cw := C._gtk_scintilla_new()
	w := *gtk.WidgetFromNative(unsafe.Pointer(cw))
	return &Scintilla{gtk.Container{w}}
}

func (sci *Scintilla) SendMessage (msg uint, wParam uint, lParam uint) uint {
	v := from_guintptr (C._gtk_scintilla_send_message(sci.ToNativeScintilla (), guint(msg), guintptr(wParam), guintptr(lParam)))
	log.Printf ("SCI: MSG=%d wP=%d lP=%d => %d", msg, wParam, lParam, v)
	return v
}

func (sci *Scintilla) StartStyling (pos uint) {
	sci.SendMessage (C.SCI_STARTSTYLING, pos, 0)
}

func (sci *Scintilla) SetStyling (length uint, style uint) {
	sci.SendMessage (C.SCI_SETSTYLING, length, style)
}

func (sci *Scintilla) GetEndStyled () uint {
	return sci.SendMessage (C.SCI_GETENDSTYLED, 0, 0)
}

func (sci *Scintilla) StyleResetDefault () {
	sci.SendMessage (C.SCI_STYLERESETDEFAULT, 0, 0)
}

func (sci *Scintilla) StyleSetFg (style uint, color uint) {
	sci.SendMessage (C.SCI_STYLESETFORE, style, color)
}

func (sci *Scintilla) StyleSetBg (style uint, color uint) {
	sci.SendMessage (C.SCI_STYLESETBACK, style, color)
}


func (sci *Scintilla) StyleSetUnderline (style uint, u bool) {

	var uu uint
	if u {
		uu = 1
	} else {
		uu = 0
	}
	sci.SendMessage (C.SCI_STYLESETUNDERLINE, style, uu)
}

func (sci *Scintilla) GetCharAt(pos uint) byte {
	return byte (sci.SendMessage (C.SCI_GETCHARAT, pos, 0))
}

func (sci *Scintilla) GetStyleAt(pos uint) uint {
	return sci.SendMessage (C.SCI_GETSTYLEAT, pos, 0)
}

//------------------------------------------------------------

func (sci *Scintilla) SetText (text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	sci.SendMessage (C.SCI_SETTEXT, 0, gstring2uint (ptr))
}

func (sci *Scintilla) SetLexer(lex uint) {
	sci.SendMessage (C.SCI_SETLEXER, lex, 0)
}

func (sci *Scintilla) SetLexerLanguage(lang string) {
	ptr := C.CString(lang)
	defer cfree(ptr)
	sci.SendMessage (C.SCI_SETLEXER, 0, gstring2uint (ptr))
}
