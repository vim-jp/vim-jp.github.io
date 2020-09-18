---
layout: blog
category: blog
title: Dropped migemo from ctrlp.vim
---
ctrlp.vim から migemo 検索機能がドロップされました。
migemo 検索機能は cmigemo を使用して正規表現パターンを生成し、ctrlp.vim の正規表現検索に流し込む実装となっていました。
しかしながら汎用性がなく、かつ [migemogrep ](https://github.com/peco/migemogrep) を使い matcher として事で代用出来てしまう事から本日ドロップされました。

[Drop migemo by mattn - Pull Request #19 - ctrlpvim/ctrlp.vim](https://github.com/ctrlpvim/ctrlp.vim/pull/19)

なお、ctrlp.vim から migemogrep を使う場合は以下の様に matcher を用意します。

```vim
function! MigemoMatch(items, str, limit, mmode, ispath, crfile, regex)
  let tmp = tempname()
  try
    if a:str =~ '^\s*$'
      return copy(a:items)
    endif
    call writefile(split(iconv(join(a:items, "\n"), &encoding, 'utf-8'), "\n"), tmp)
    return split(iconv(system(
    \  printf('migemogrep %s %s',
    \    shellescape(a:str),
    \    shellescape(tmp))), 'utf-8', &encoding), "\n")
  catch
    return []
  finally
    call delete(tmp)
  endtry
endfunction
let g:ctrlp_match_func = {'match' : 'MigemoMatch' }
```

