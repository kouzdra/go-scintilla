package scintilla

// #cgo pkg-config: gmodule-2.0 gtk+-2.0
// #cgo LDFLAGS: -lm -lstdc++ -L/home/msk/local/lib -lscintilla
// #cgo CFLAGS: -I${SRCDIR}/../scintilla/include
// #include "../../../mattn/go-gtk/gtk/gtk.go.h"
// #include "scintilla.go.h"
import "C"
import (
	"unsafe"
	"log"
	//"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/kouzdra/go-scintilla/gtk/consts"
)

var _ = log.Printf

type Color uint32
type Style uint32
type Pos   uint32

type arg   uint32

//-----------------------------------------------------------------------
// GtkScintilla
//-----------------------------------------------------------------------

type Scintilla struct {
	gtk.Container
	Styling *Styling
	Handlers *Handlers
	Id int
}

var sciMap = make (map [int] *Scintilla, 128)
var lastId = 1

func (v *Scintilla) toNativeScintilla() *C.ScintillaObject {
	return C.toGtkScintilla(unsafe.Pointer(v.GWidget))
}

func NewScintilla() *Scintilla {
	w := *gtk.WidgetFromNative(unsafe.Pointer(C._gtk_scintilla_new()))
	sci := &Scintilla{Container:gtk.Container{w}, Handlers:&Handlers{}, Id: lastId}
	lastId ++
	sci.SetIdentifier (sci.Id)
	//sci.GetIdentifier ()
	sciMap [sci.Id] = sci
	sci.Styling = &Styling{sci}
	return sci
}

func (sci *Scintilla) sendMessage (msg uint, wParam arg, lParam arg) uint {
	v := uint (C._gtk_scintilla_send_message(
		sci.toNativeScintilla (), C.guint(msg), C.guintptr(uint (wParam)), C.guintptr(uint (lParam))))
	//log.Printf ("SCI: MSG=%d wP=%d lP=%d => %d", msg, wParam, lParam, v)
	return v
}

//export gtk_sci_notification_handler
func gtk_sci_notification_handler(sciGtk *C.ScintillaObject, id int, scn *C.SCNotification) {
	code := scn.nmhdr.code
	sci := sciMap [id]
	//log.Printf ("SCI NOTIFY: %d\n", code);
	switch code {
	case consts.SCN_MODIFIED:
		if h := sci.Handlers.OnModify; h != nil {
			length := uint (scn.length)
			//log.Printf (">>> Length=%d\n", length)
			var text string
			if scn.text != nil {
				text = C.GoStringN (scn.text, C.int (length))
			}
			//log.Printf (">>> Text=[%s]\n", text)
				
			h (uint (scn.modificationType), Pos (scn.position), length, int (scn.linesAdded),
				text,
				uint (scn.line), uint (scn.foldLevelNow), uint (scn.foldLevelPrev))
		}
	}
}

type Handlers struct {
	OnModify func (uint, Pos, uint, int, string, uint, uint, uint)
}

//////////////////////////////////////////////////////////////////////////////

func (sci *Scintilla) GetCharAt(pos Pos) byte {
	return byte (sci.sendMessage (C.SCI_GETCHARAT, arg (pos), 0))
}

//------------------------------------------------------------

func (sci *Scintilla) SetText (text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	sci.sendMessage (consts.SCI_SETTEXT, 0, gstring2arg (ptr))
}

func (sci *Scintilla) SetIdentifier(id int) {
	sci.sendMessage (consts.SCI_SETIDENTIFIER, arg (id), 0)
}

func (sci *Scintilla) GetIdentifier() int {
	return int (sci.sendMessage (consts.SCI_GETIDENTIFIER, 0, 0))
}

func (sci *Scintilla) SetLexer(lex uint) {
	sci.sendMessage (consts.SCI_SETLEXER, arg (lex), 0)
}

func (sci *Scintilla) SetLexerLanguage(lang string) {
	ptr := C.CString(lang)
	defer cfree(ptr)
	sci.sendMessage (consts.SCI_SETLEXER, 0, gstring2arg (ptr))
}

//----------------------------------- aux ----------------------------------


func gstring(s *C.char) *C.gchar { return C.toGstr(s) }
func gstring2arg(s *C.char) arg { return arg (C.toGstrUint(s)) }
func cstring(s *C.gchar) *C.char { return C.toCstr(s) }
func gostring(s *C.gchar) string { return C.GoString(cstring(s)) }

func cfree(s *C.char) { C.freeCstr(s) }

/*func gslist(l *glib.SList) *C.GSList {
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
}*/

//////////////////////////////////////////////////////////////////////

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
	//ff := C.make_string (1024)
	//s.sci.sendMessage (consts.SCI_STYLEGETFONT, arg (style), gstring2arg (ff))
	//log.Printf ("OLD FONT: \"%s\"\n", gostring (C.toGstr (ff)))
	//cfree (ff)

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


