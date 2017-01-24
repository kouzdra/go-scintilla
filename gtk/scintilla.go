package scintilla

// #cgo pkg-config: gmodule-2.0 gtk+-2.0
// #cgo LDFLAGS: -lm -lstdc++ -L/home/msk/local/lib -lscintilla
// #cgo CFLAGS: -I${SRCDIR}/../scintilla/include
// #include "../../../mattn/go-gtk/gtk/gtk.go.h"
// #include "scintilla.go.h"
import "C"
import "unsafe"
import "log"
// import "github.com/mattn/go-gtk/glib"
import "github.com/mattn/go-gtk/gtk"
import "github.com/kouzdra/go-scintilla/gtk/consts"

var _ = log.Printf

type Color uint32
type Style uint32
type Pos   uint32

type Arg   uint32


//----------------------------------- aux ----------------------------------

func    bool2arg(b bool) Arg {
	if b { return Arg (1) } else { return Arg (0) }
}

func make_string (count uint) *C.char { return C.make_string (C.int (count)) }
func gstring2arg(s *C.char) Arg { return Arg (C.toGstrUint(s)) }
func gostring(s *C.gchar) string { return gocstring(C.toCstr(s)) }
func gocstring(s *C.char) string { return C.GoString(s) }
func gostringN(s *C.gchar, n int) string { return gocstringN(C.toCstr(s), n) }
func gocstringN(s *C.char, n int) string { return C.GoStringN(s, C.int (n)) }
func cstring(s string) *C.char { return C.CString(s) }
func cfree(s *C.char) { C.freeCstr(s) }

//-----------------------------------------------------------------------
// GtkScintilla
//-----------------------------------------------------------------------

type Scintilla struct {
	gtk.Container
	Indic    *   Indic
	Styling  * Styling
	Handlers *Handlers
}

var sciMap = make (map [int] *Scintilla, 128)
var lastId = 1

func (v *Scintilla) toNativeScintilla() *C.ScintillaObject {
	return C.toGtkScintilla(unsafe.Pointer(v.GWidget))
}

func NewScintilla() *Scintilla {
	w := *gtk.WidgetFromNative(unsafe.Pointer(C._gtk_scintilla_new()))
	id := lastId
	sci := &Scintilla{Container:gtk.Container{w}, Handlers:&Handlers{}}
	lastId ++
	sci.SetIdentifier (id)
	//sci.GetIdentifier ()
	sciMap [id] = sci
	sci.Indic   = &  Indic{sci}
	sci.Styling = &Styling{sci}
	return sci
}

func (sci *Scintilla) SendMessage (msg uint, wParam Arg, lParam Arg) uint {
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

func (sci *Scintilla) SetPhasesDraw(count int) {
	sci.SendMessage (consts.SCI_SETPHASESDRAW, Arg (count), 0)
}

func (sci *Scintilla) GetPhasesDraw() int {
	return int (sci.SendMessage (consts.SCI_GETPHASESDRAW, 0, 0))
}

//------------------------------------------------------------

func (sci *Scintilla) GetCharAt(pos Pos) byte {
	return byte (sci.SendMessage (C.SCI_GETCHARAT, Arg (pos), 0))
}


func (sci *Scintilla) SetText (text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	sci.SendMessage (consts.SCI_SETTEXT, 0, gstring2arg (ptr))
}
//---

func (sci *Scintilla) SetIdentifier(id int) {
	sci.SendMessage (consts.SCI_SETIDENTIFIER, Arg (id), 0)
}

func (sci *Scintilla) GetIdentifier() int {
	return int (sci.SendMessage (consts.SCI_GETIDENTIFIER, 0, 0))
}

//---

func (sci *Scintilla) SetLexer(lex uint) {
	sci.SendMessage (consts.SCI_SETLEXER, Arg (lex), 0)
}

func (sci *Scintilla) SetLexerLanguage(lang string) {
	ptr := C.CString(lang)
	defer cfree(ptr)
	sci.SendMessage (consts.SCI_SETLEXER, 0, gstring2arg (ptr))
}
