if exists(':NewEntry')
  finish
endif
let s:basedir = expand('<sfile>:h')

function! s:new_entry()
  let title = input("Title: ")
  if len(title) == 0
    return
  endif
  exec "e" printf("%s/_posts/blog/%s-%s.md", s:basedir, strftime("%Y-%m-%d"), substitute(title, '\s', '-', 'g'))
  call setline(1, [
  \ '---',
  \ 'layout: blog',
  \ 'category: blog',
  \ 'title: ' . title,
  \ '---',
  \ '',
  \])
  normal! G
endfunction
command! NewEntry call <SID>new_entry()
