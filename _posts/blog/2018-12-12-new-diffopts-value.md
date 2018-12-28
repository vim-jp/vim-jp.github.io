---
layout: post
category: blog
title: 新しい diffopts オプションの動作とその注意点について
---
[patch 8.1.0363](https://github.com/vim/vim/commit/c93262b2e3fb043c685bd4014a35a29111dea484) にて diffopts オプションに `internal` という値が追加され、デフォルト値となりました。これまでは `diffopts=filter` という値がデフォルト値になっており差分表示を行う際には外部コマンドの diff が使われていました。今回入った `diffopts=internal` により Vim の内部に実装された処理により diff の表示が行われます。diff コマンドがインストールされていない OS の場合はとても便利になりました。

さて、これまでも Vim の diff 表示は便利に使われてきたのですが、中には diff モードに入った事を継起にキーをマップしたり便利な設定をされていたかもしれません。その際これまでは外部コマンドによるフィルタ処理が行われていたので `FilterWritePre` や `FilterWritePost` というイベントが発火しておりそれを利用されていたと思います。今回入った `internal` は外部コマンドを使わなくなりますので `FilterWritePre` や `FilterWritePost` は発火しません。別途 [8.1.0397](https://github.com/vim/vim/commit/e8fa05b5bc2d6d76bf5af50176a63655d00d1110) で用意された `DiffUpdated` イベントを使って下さい。

```vim
autocmd FilterWritePre * call s:my_diff_settings()
function! s:my_diff_settings() abort
  if &diff
    setlocal nocursorline
    nnoremap <buffer> u     u:diffupdate<CR>
    nnoremap <buffer> <C-r> <C-r>:diffupdate<CR>
  endif
endfunction
```
