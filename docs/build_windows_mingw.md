---
layout: docs
title: MinGWを使ってのビルド方法
---

1.  MinGWのインストール

    [MinGW \| Minimalist GNU for Windows](http://www.mingw.org/)からお使いのコンピューターにあったコンパイラをダウンロードしてインストールします。

2.  Gitのインストール

    [Git](https://git-scm.com/)からGitをインストールします。

3.  ソース取得

    コマンドプロンプトから以下を実行します。

        git clone https://github.com/vim/vim.git

    `git clone`を実行した後にソースが更新された場合は、以下のコマンドで最新のソースを取得できます。

        git fetch
        git merge

    ローカルでの変更がない場合、あるいは2つをまとめて以下のコマンドでもOKです。

        git pull

    特定のバージョンを指定して取得する場合は、以下のコマンドを実行します。

        git checkout v7.4.393

4.  コンパイル

    `vim/src`フォルダに移動し以下のコマンドを実行します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ^
          ICONV=yes DEBUG=no

    もしPerl拡張やRuby拡張、Python拡張を使う場合は以下の様に指定します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes ^
          PERL=C:\strawberry\perl DYNAMIC_PERL=yes PERL_VER=512 ^
          PYTHON=c:\python27 DYNAMIC_PYTHON=yes PYTHON_VER=27 ^
          RUBY=c:\ruby193 DYNAMIC_RUBY=yes RUBY_VER=19 ^
          RUBY_VER_LONG=1.9.1 CSCOPE=yes NETBEANS=yes ^
          ARCH=x86-64 STATIC_STDCPLUS=yes DEBUG=no

    msys2 付属の perl/python/ruby を使ってビルドする場合は以下の様に指定します。

        mingw32-make -f Make_ming.mak GUI=yes IME=yes MBYTE=yes ICONV=yes ^
          PERL=c:/mingw64/mingw64 DYNAMIC_PERL=yes PERL_VER=522 ^
            PERLEXE=c:/msys64/mingw64/bin/perl.exe ^
            PERLLIB=c:/msys64/mingw64/lib/perl5/core_perl ^
          PYTHON=c:/msys64/mingw64 DYNAMIC_PYTHON=yes PYTHON_VER=27 ^
            PYTHON_HOME=c:/msys64/mingw64 ^
            PYTHONINC=-Ic:/msys64/mingw64/include/python2.7 ^
            DYNAMIC_PYTHON_DLL=libpython2.7.dll ^
          RUBY=c:/msys64/mingw64 DYNAMIC_RUBY=yes RUBY_VER=22 RUBY_VER_LONG=2.2.0 ^
          CSCOPE=yes NETBEANS=yes ARCH=x86-64 STATIC_STDCPLUS=yes DEBUG=no

    注意1) 複数行に跨るのでバッチファイルに記載して実行して下さい。

    注意2) ARCH は i686 または x86-64 を指定して下さい。

    Vim 7.4.393以降で使えるようになったDirectWriteを有効にするには、MinGWの派生版である[MinGW-w64](http://mingw-w64.sourceforge.net/)を使ってコンパイルする必要があります。
    32bit版は `DIRECTX=yes ARCH=i686`、64bit版は `DIRECTX=yes ARCH=x86-64` を指定する必要があります。
    なお、DirectWriteを有効にすると、デフォルトでは実行時に libstdc++-6.dll, libgcc\_s\_sjlj-1.dll が必要となりますが、`STATIC_STDCPLUS=yes` を指定することでこれらのライブラリをスタティックリンクすることが出来ます。
