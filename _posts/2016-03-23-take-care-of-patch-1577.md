---
layout: post
category: blog
title: Patch 1577 の使用にご注意下さい
---

[Patch 1577](https://groups.google.com/forum/#!topic/vim_dev/yAWqYrURjjk)

Patch 1577 にて Vim script の新しい API として partial argument と呼ばれる機能が実装されました。

この API は

```vim
function! MyFunc(arg1, arg2, arg3)
  echo a:arg1 a:arg2 a:arg3
endfunction

let Cb = function('MyFunc', ["foo", "bar"])

call Cb("zzz")
```

関数 MyFunc の呼び出し引数を関数リファレンス Cb にホールドする事でカリー化を実現する為の物ですが、この実装に際してバグが混入し多くの Vim plugin が動作しなくなってしまいました。

現在は修正パッチがリリースされ正しく動作する様に復旧出来ましたが、復旧に至るまでに多くのパッチをリリースする必要がありました。

<http://ftp.vim.org/vim/patches/7.4/7.4.1577>

<http://ftp.vim.org/vim/patches/7.4/7.4.1580>

<http://ftp.vim.org/vim/patches/7.4/7.4.1581>

<http://ftp.vim.org/vim/patches/7.4/7.4.1582>

<http://ftp.vim.org/vim/patches/7.4/7.4.1585>

<http://ftp.vim.org/vim/patches/7.4/7.4.1586>

<http://ftp.vim.org/vim/patches/7.4/7.4.1589>

<http://ftp.vim.org/vim/patches/7.4/7.4.1597>

<http://ftp.vim.org/vim/patches/7.4/7.4.1605>

<http://ftp.vim.org/vim/patches/7.4/7.4.1606>

<http://ftp.vim.org/vim/patches/7.4/7.4.1607>

<http://ftp.vim.org/vim/patches/7.4/7.4.1608>

<http://ftp.vim.org/vim/patches/7.4/7.4.1638>

<http://ftp.vim.org/vim/patches/7.4/7.4.1639>

この 1577 から 1639 までの Vim では多くの Vim plugin が動作しない可能性があります。また著名な Vim plugin へ誤動作の問い合わせが幾つかあったと見られています。 

ご使用になられる場合は 1577 未満か、1639 以降のバージョンを使われる事をお奨めします。また各ディストリビューションのメンテナをされておられる方におかれましてはこの間のバージョンのパッケージングを避けて頂ける様、お願い致します。
