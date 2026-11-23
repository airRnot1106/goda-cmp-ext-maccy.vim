// Command e2ecfg is the configuration of the end-to-end tests: the test UI and
// the maccy source reading $MACCY_DB, set up as the documentation recommends.
package main

import (
	"os"

	"github.com/Shougo/goda-cmp.vim/cmp"
	"github.com/Shougo/goda-cmp.vim/cmp/cmptest"

	"github.com/airRnot1106/goda-cmp-ext-maccy.vim/ext/source/maccy"
)

func main() {
	b := cmp.New()
	b.AddUI(cmptest.NewUI("test", "i"))
	b.AddSource(maccy.New(maccy.Params{
		RecentMs:      300000,
		Types:         []string{"public.utf8-plain-text"},
		DBPath:        os.Getenv("MACCY_DB"),
		MaxByteLength: 10000,
	}))
	b.PatchGlobal(cmp.Config{
		UI:      "test",
		Sources: []string{"maccy"},
		SourceOptions: map[string]cmp.SourceOption{
			"maccy": {Matchers: []string{}, Sorters: []string{}, Converters: []string{}, MinAutoCompleteLength: new(0)},
		},
		AutoCompleteDelay: new(10),
	})
	b.Main()
}
