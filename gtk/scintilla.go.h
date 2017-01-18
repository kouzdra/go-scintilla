#ifndef GO_GTKSCINTILLA_H
#define GO_GTKSCINTILLA_H

#include <gtk/gtk.h>
#include <gdk/gdk.h>
#include <stdlib.h>
#define GTK
#include <Scintilla.h>
#include <ScintillaWidget.h>
#include <stdio.h>

/*static inline gchar** make_strings(int count) {
	return (gchar**)malloc(sizeof(gchar*) * count);
}

static inline void destroy_strings(gchar** strings) {
	free(strings);
}

static inline void set_string(gchar** strings, int n, gchar* str) {
	strings[n] = str;
}

static inline GObject* toGObject(void* o) { return G_OBJECT(o); }
static inline gchar* toGstr(const char* s) { return (gchar*)s; }
static inline char* toCstr(const gchar* s) { return (char*)s; }
static inline gchar** nextGstr(gchar** s) { return (s+1); }
static inline void freeCstr(char* s) { free(s); }
*/
//static inline gchar* toGstr(const char* s) { return (gchar*)s; }
//static inline char* toCstr(const gchar* s) { return (char*)s; }

static inline char* make_string(int count) {
	return (char*)malloc(count);
}

static inline guintptr toGstrUint(const char* s) { return (guintptr)s; }

static ScintillaObject* toGtkScintilla(void* w) { return SCINTILLA(w); }

static void _gtk_sci_set_notification_handler (GtkWidget * obj)
{
  extern void gtk_sci_notification_handler(ScintillaObject *sci, int id, SCNotification *scn);
  SCINTILLA_OBJECT_GET_CLASS(obj) -> notify = gtk_sci_notification_handler;
}

static GtkWidget* _gtk_scintilla_new()
{
  GtkWidget* const sci = scintilla_object_new();
  _gtk_sci_set_notification_handler (sci);
  return sci;
}

static guintptr _gtk_scintilla_send_message(ScintillaObject *sci, guint msg, guintptr wParam, guintptr lParam)
{
  return scintilla_object_send_message(SCINTILLA (sci), msg, wParam, lParam);
}

#endif
