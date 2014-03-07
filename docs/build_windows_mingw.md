---
layout: docs
title: MinGWを使ってのビルド方法
---

1.  MinGWのインストール

    [MinGW | Minimalist GNU for Windows](http://www.mingw.org/)からお使いのコンピューターにあったコンパイラをダウンロードしてインストールします。

2.  Mercurialのインストール

    [Mercurial SCM](http://mercurial.selenic.com/)からMercurialをインストールします。

3.  ソース取得

    コマンドプロンプトから以下を実行します。

        hg clone https://vim.googlecode.com/hg/ vim

    `hg clone`を実行した後にソースが更新された場合は、以下のコマンドで最新のソースを取得できます。

        hg pull
        hg update

    あるいは2つをまとめて、以下のコマンドでもOKです。

        hg pull -u

4.  コンパイル

    `vim/src`フォルダに移動し以下のコマンドを実行します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes
          ICONV=yes DEBUG=no

    ※実際は1行

    もしPerl拡張やRuby拡張、Python拡張を使う場合は以下の様に指定します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes
          ICONV=yes PERL=C:\strawberry\perl DYNAMIC_PERL=yes
          PERL_VER=512 PYTHON=c:\python27 DYNAMIC_PYTHON=yes
          PYTHON_VER=27 RUBY=c:\ruby193 DYNAMIC_RUBY=yes RUBY_VER=19
          RUBY_VER_LONG=1.9.1 CSCOPE=yes NETBEANS=yes
          DEBUG=no

    ※実際は1行
