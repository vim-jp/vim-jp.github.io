---
layout: post
category: blog
title: Patch 1577 の使用にご注意下さい
---

## TL;DR

安定したVimを使いたいのならば、7.4.1558 以前か 1639 以降のバージョンを使ってください。

## 詳細

[Patch 1559][1] において Vim script に partial argument と言う新機能が実装されました。

この機能は、既存の `function()` 関数を拡張する形で提供され、以下のように利用します。

```vim
function! MyFunc(arg1, arg2, arg3)
  echo a:arg1 a:arg2 a:arg3
endfunction

let Cb = function('MyFunc', ["foo", "bar"])

call Cb("zzz")
" foo bar zzz
```

関数 `MyFunc` の呼び出し引数の幾つかを、関数リファレンス `Cb` にホールドする事でカリー化(正確には部分適用)を実現する為の物ですが、この実装に際して複数のバグが混入し、多くの Vim plugin が動作しなくなっていました。

**現在では正しく動作するように復旧しました**が、そこに至るまでには以下に示す多くの修正パッチがリリースされています。

*   <http://ftp.vim.org/vim/patches/7.4/7.4.1577>
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

この 1559 から 1638 までの Vim では、多くの Vim plugin が動作しない可能性があります。現に、本不具合が原因と見られる誤動作が、著名な Vim plugin の誤動作と誤解され、各々へ報告されているのが観測されました。

Vim script を正常に利用するには、7.4.1558 以前か 7.4.1639 以降のバージョンの Vim をご利用いただくことを強く推奨いたします。また各ディストリビューションの Vim 関連パッケージのメンテナをされている方々には、この間のバージョン(7.4.1559 から 7.4.1638) を避けてパッケージングするよう、重ねてお願い申し上げます。

[1]:https://groups.google.com/d/msg/vim_dev/oaKOwnslCZY/60J7AjUdGAAJ
