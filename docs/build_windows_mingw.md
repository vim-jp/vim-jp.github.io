---
layout: docs
title: MinGWを使ってのビルド方法
---
MinGWをインストールする
-----------------------

1. MinGWのインストール

> [MinGW | Minimalist GNU for Windows](http://www.mingw.org/)からお使いのコンピュータにあったコンパイラをダウンロードしてインストールします。

2. Mercurialのインストール

> [Mercurial SCM](http://mercurial.selenic.com/)からMercurialをインストールします。

3. ソース取得

>  コマンドプロンプトから以下を実行します。

>     hg clone hg clone https://code.google.com/p/vim/

4. コンパイル

> `vim/src`フォルダに移動し以下のコマンドを実行します。

>     mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes DEBUG=no

> もしPerl拡張やRuby拡張、Python拡張を使う場合は以下の様に指定します。

>     mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes PERL=C:\strawberry\perl DYNAMIC_PERL=yes PERL_VER=512 PYTHON=c:\python27 DYNAMIC_PYTHON=yes PYTHON_VER=27 RUBY=c:\ruby192 DYNAMIC_RUBY=yes RUBY_VER=191 RUBY_VER_LONG=1.9.1 CSCOPE=yes MSVCVER=6.0 NETBEANS=yes DEBUG=no

