package maccy

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/Shougo/goda-cmp.vim/cmp/cmptest/e2e"
)

func TestRealAcceptingAClipKeepsWhatWasTyped(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a binary and starts an editor")
	}
	f := newFixture(t)
	at := float64(time.Now().Add(-time.Second).UnixMilli())/1000 - coreDataEpoch
	if _, err := f.db.Exec(`INSERT INTO ZHISTORYITEM (Z_PK, ZLASTCOPIEDAT) VALUES (1, ?)`, at); err != nil {
		t.Fatal(err)
	}
	if _, err := f.db.Exec(`INSERT INTO ZHISTORYITEMCONTENT (ZITEM, ZTYPE, ZVALUE) VALUES (1, ?, ?)`, plainText, []byte("https://example.com")); err != nil {
		t.Fatal(err)
	}
	bin := e2e.BuildConfig(t, "testdata/e2ecfg")
	vimrc, _ := filepath.Abs("testdata/accept.vim")
	e2e.Run(t, func(t *testing.T, ed e2e.Editor) {
		if ed.Name != "nvim" {
			t.Skip("OnCompleteDone needs Neovim")
		}
		steps := `[
		  ["i", 0.3], ["see fo", 1.0], ["o", 1.0],
		  ["\u000c", 0.3], ["\u000e\u0019", 0.5], ["\u001d", 0.3],
		  ["\u001b", 0.3]
		]`
		ds := e2e.Drive(t, ed, e2e.Options{Bin: bin, File: filepath.Join(t.TempDir(), "text.txt"), Steps: steps, Vimrc: vimrc,
			Env: []string{"GODA_DUMP_UI=test", "MACCY_DB=" + f.path}})
		if len(ds) < 1 {
			t.Fatal("nothing recorded")
		}
		if want := "see foohttps://example.com"; ds[0].Line != want {
			t.Errorf("line %q, want %q", ds[0].Line, want)
		}
	})
}
