" Ctrl-L hands what the test UI was asked to show to Vim's own completion, so
" that Ctrl-N Ctrl-Y inserts the first item from the column the engine gave.
function! s:accept() abort
  let s = goda_cmp_testui#test#state()
  call complete(s.col, s.items)
endfunction
inoremap <C-l> <Cmd>call <SID>accept()<CR>
