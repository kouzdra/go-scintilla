GOCDEFSDIR = ../go-tools
GOCDEFS = $(GOCDEFSDIR)/gocdefs

all:: consts
	go build all.go

clean:
	go clean all.go

install::
	tar xzf scintilla.tar.gz
	cd scintilla/gtk; make


$(GOCDEFS): $(GOCDEFS).go
	$(MAKE) -C $(GOCDEFSDIR) gocdefs

consts: $(GOCDEFS) gtk/consts/consts.go

gtk/consts/consts.go::
	cat $@.tpl >$@
	@$(GOCDEFS) "SC_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCWS_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCVS_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCTD_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCI_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCK_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCN_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCEN_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCMOD_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "SCFIND_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "STYLE_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "EDGE_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "CARET_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "CARETSTYLE_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "ANNOTATION_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "UNDO_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "VISIBLE_" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "INDIC" <scintilla/include/Scintilla.h >>$@
	@$(GOCDEFS) "KEYWORDSET_MAX" <scintilla/include/Scintilla.h >>$@
	$(GOCDEFS) "INVALID_POSITION" <scintilla/include/Scintilla.h >>$@
