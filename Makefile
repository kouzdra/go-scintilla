GOCDEFSDIR = ../go-tools
GOCDEFS = $(GOCDEFSDIR)/gocdefs

all:: gtk/consts/consts.go
	go build all.go

$(CGODEFS): $(GOCDEFS).go
	$(MAKE) -C $(GOCDEFSDIR) gocdefs

consts: gtk/consts/consts.go

gtk/consts/consts.go::
	cat $@.tpl >$@
	$(GOCDEFS) "SCI_" <scintilla/include/Scintilla.h >>$@
	$(GOCDEFS) "SCN_" <scintilla/include/Scintilla.h >>$@
	$(GOCDEFS) "SC_" <scintilla/include/Scintilla.h >>$@
