---
layout: post
category: blog
title: Windows上でruby拡張を有効にする際のビルド方法が変更されました。
---

Patch 7.3.620 によりWindows上でruby拡張を有効にする際のビルド方法が変更されました。

これまでは以下の様にRUBY\_VERを指定する必要があり、このRUBY\_VERはリンクライブラリのパスを決定するパラメータとして使用されていました。

     nmake -f Make_mvc.mak GUI=yes IME=yes MBYTE=yes
       ICONV=yes PERL=C:\strawberry\perl DYNAMIC_PERL=yes
       PERL_VER=512 PYTHON=c:\python27 DYNAMIC_PYTHON=yes
       PYTHON_VER=27 RUBY=c:\ruby192 DYNAMIC_RUBY=yes RUBY_VER=191
       RUBY_VER_LONG=1.9.1 CSCOPE=yes MSVCVER=6.0 NETBEANS=yes
       DEBUG=no
 
つまりrubyのAPIバージョンを指定していました。しかしこの値はvim内部でのrubyバージョンの扱いにも使用されており

    # if defined(DYNAMIC_RUBY_VER) && DYNAMIC_RUBY_VER >= 19

の様なコードに影響を及ぼしていました。つまり、rubyのバージョンとAPIバージョンが混同して使われてきました。
今回の修正によって、今後はRUBY\_VERはrubyのバージョンそのものを、またRUBY\_VER\_LONGにAPIのバージョンを指定する事になりました。
 
     nmake -f Make_mvc.mak GUI=yes IME=yes MBYTE=yes
       ICONV=yes PERL=C:\strawberry\perl DYNAMIC_PERL=yes
       PERL_VER=512 PYTHON=c:\python27 DYNAMIC_PYTHON=yes
       PYTHON_VER=27 RUBY=c:\ruby193 DYNAMIC_RUBY=yes RUBY_VER=19
       RUBY_VER_LONG=1.9.1 CSCOPE=yes MSVCVER=6.0 NETBEANS=yes
       DEBUG=no
 
実際には、RUBY\_VER\_LONGはその値を使ってincludeパスに使用したり、ドットを削除した値を用いてランタイムライブラリ(msvcrt-ruby193.dll)の名前を決定する材料としても使用されます。

これまでWindows上でruby拡張を有効にしてvimをビルドしてこられた方には以上の変更が必要となります。お気を付け下さい。

参考

* [MinGWを使ってのビルド方法](http://vim-jp.org/docs/build_windows_mingw.html)
* [VisualStudio10を使ってのビルド方法](http://vim-jp.org/docs/build_windows_msvc.html)
