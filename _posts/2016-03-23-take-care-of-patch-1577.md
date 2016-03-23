---
layout: post
category: blog
title: Patch 7.4.1577 の使用にご注意ください
---

## TL;DR

安定したVimを使いたいのならば、7.4.1576 以前か 1639 以降のバージョンを使ってください。

## 詳細

[Patch 7.4.1559][1] において Vim script に partial argument と言う新機能が実装されました。

この機能は、既存の `function()` 関数を拡張する形で提供され、以下のように利用します。

```vim
function! MyFunc(arg1, arg2, arg3)
  echo a:arg1 a:arg2 a:arg3
endfunction

let Cb = function('MyFunc', ["foo", "bar"])

call Cb("zzz")
" foo bar zzz
```

関数 `MyFunc` の呼び出し引数の幾つかを、関数リファレンス `Cb` にホールドする事で部分適用を実現する為の物で、辞書関数の場合には `self` を含めることもできます。

```vim
function! MyFunc() dict
  echo self.value
endfunction

let Cb = function('MyFunc', {'value': 'foo'})

call Cb()
" foo
```

この時点では、この `Partial` は `function()` に追加された引数を渡した場合のみ生成されていたため、過去に作られたスクリプトに影響はありませんでした。

ですが、[Patch 7.4.1577][2] において、辞書関数にアクセスすると `Partial` が返されるように変更が行われました。

```vim
let dict = {'value': 'foo'}
fnction! dict.Func()
  return self.value
endfunction

let Cb = dict.Func

" 7.4.1576 以前では Cb はただの Funcref だったので dict を明示しないと呼べなかったが、
" 7.4.1577 以降では Cb は dict をホールドした Partial なので、直接呼べる
call Cb()
" foo
```

この実装に際して、今まで `Funcref` が返ってきていたところが `Partial` に変わったことにより、複数のバグが混入し、多くの Vim plugin が動作しなくなっていました。

**現在では正しく動作するように復旧しました**が、そこに至るまでには以下に示す多くの修正パッチがリリースされています。

*   <http://ftp.vim.org/vim/patches/7.4/7.4.1580>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1581>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1582>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1585>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1586>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1589>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1597>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1605>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1606>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1607>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1608>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1638>
*   <http://ftp.vim.org/vim/patches/7.4/7.4.1639>

この 7.4.1577 から 7.4.1638 までの Vim では、多くの Vim plugin が動作しない可能性があります。現に、本不具合が原因と見られる誤動作が、著名な Vim plugin の誤動作と誤解され、各々へ報告されているのが観測されました。

Vim script を正常に利用するには、7.4.1576 以前か 7.4.1639 以降のバージョンの Vim をご利用いただくことを強く推奨いたします。また各ディストリビューションの Vim 関連パッケージのメンテナをされている方々には、この間のバージョン(7.4.1576 から 7.4.1638) を避けてパッケージングするよう、重ねてお願い申し上げます。

[1]:https://groups.google.com/d/topic/vim_dev/oaKOwnslCZY/discussion
[2]:https://groups.google.com/d/topic/vim_dev/yAWqYrURjjk/discussion

