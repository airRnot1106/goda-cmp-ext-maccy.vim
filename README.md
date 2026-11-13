# goda-cmp-ext-maccy.vim

[Maccy](https://github.com/p0deje/Maccy) completion for [goda-cmp.vim](https://github.com/Shougo/goda-cmp.vim).

This source completes recent clips copied in Maccy, the macOS clipboard manager. It is a port of [ddc-source-maccy](https://github.com/airRnot1106/ddc-source-maccy).

## Dependencies

- [goda-cmp.vim](https://github.com/Shougo/goda-cmp.vim)
- [goda.vim](https://github.com/Shougo/goda.vim)
- [Maccy](https://github.com/p0deje/Maccy)
- Neovim, because accepting a clip uses `OnCompleteDone`

The source reads Maccy's database directly, so the `sqlite3` command is not needed.

## Install

The source is a Go package that your goda-cmp configuration imports.

```sh
cd ~/.config/nvim/goda-cmp
go get github.com/airRnot1106/goda-cmp-ext-maccy.vim
```

## Configuration

```go
import "github.com/airRnot1106/goda-cmp-ext-maccy.vim/ext/source/maccy"

b.AddSource(maccy.New(maccy.Params{
	RecentMs:            300000,
	CacheTtlMs:          2000,
	Types:               []string{"public.utf8-plain-text"},
	DBPath:              "~/Library/Containers/org.p0deje.Maccy/Data/Library/Application Support/Maccy/Storage.sqlite",
	MaxByteLength:       10000,
	Dedupe:              true,
	TrimTrailingNewline: true,
	ExcludeAfterUse:     true,
	MaxUseCount:         1,
}))
```

Add `"maccy"` to `Sources`, then rebuild with `:call goda_cmp#build()`.

See `:help goda-cmp-ext-maccy`.
