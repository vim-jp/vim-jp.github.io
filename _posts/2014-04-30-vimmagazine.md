---
layout: post
category: vimmagazine
title: Vim Magazine 4 月号
---

## 話題

- [Vimプラグイン読書会 第6回](http://haya14busa.github.io/reading-vimplugin/)

  vim-textobj-userとvim-operator-userを使った自作テキストオブジェクトとオペレーターを作成できるようになるために、[kana/vim-textobj-indent](https://github.com/kana/vim-textobj-indent)と[kana/vim-operator-replace](https://github.com/kana/vim-operator-replace)を読みます。
  5月24日(土) 21:00~ [LingrのVim部屋](http://lingr.com/room/vim)にて

- [第2回 Sapporo.vim](http://atnd.org/events/49000)

  Sapporo.vimは札幌でVimに興味がある！Vimが好きだ！ という方が対象の勉強会です。 皆さんでVimVimしましょう。
  5月25日(日) 札幌

- [momonga.vim #4](http://connpass.com/event/6108/)

  Vimをします（かなり）。
  Vim pluginを作ったり、.vimrcを弄ったり、Vimで何かつくったり。自由にもくもくします。初心者歓迎ですよっ
  6月7日(土) 東京

## 今月の新機能

- 7.4.234 : `v:progpath` が追加されました。
- 7.4.235 : `exepath()` 関数が追加されました。
- 7.4.236 : パッチバージョンが `has("patch-7.4.123")` で確認できるようになりました。
- 7.4.238 : smack (Linux のセキュリティモジュールの一つ) のサポートが追加されました。
- 7.4.241 : `submatch()` にリストを返すための第2引数が追加されました。
- 7.4.242 : `getreg()` にリストを返すための別の引数が追加されました。
- 7.4.243 : `setreg()` がリストを許容するようになりました。
- 7.4.247 : `system()` の input に任意でリストを使うようになりました。
- 7.4.248 : `systemlist()` が追加されました。

## リリース情報

- [7.4.230 : error when using ":options"](http://code.google.com/p/vim/source/detail?r=57ecd7a8c0f052296b41b916eb1ae7f2a9a48b27)
- [7.4.231 : an error in ":options" is not caught by the tests](http://code.google.com/p/vim/source/detail?r=0a295a3c9e473512ad3b006a0fb752ad43d19094)
- [7.4.232 : ":%s/\\n//" uses a lot of memory](http://code.google.com/p/vim/source/detail?r=845608965bd9d0b2755997a7be812746885ff105)
- [7.4.233 : escaping special chars in ":!cmd %" is inconsistant](http://code.google.com/p/vim/source/detail?r=22a1d5762ba3a75984e89dcc47a65498f63a6c2c)
- [7.4.234 : can't get the command that was used to start Vim](http://code.google.com/p/vim/source/detail?r=d2286df8719d6e99c743e3bf6ac14d1f9debc84d)
- [7.4.235 : it is not easy to get the full path of a command](http://code.google.com/p/vim/source/detail?r=5ab2946f7ce560985830fbc3c453bb0f7a01f385)
- [7.4.236 : it's not that easy to check the Vim patch version](http://code.google.com/p/vim/source/detail?r=a44087db72386d080e9da870d751daf498004be8)
- [7.4.237 : (after 7.4.236) has() not checking for specific patch](http://code.google.com/p/vim/source/detail?r=71b165a378ad580818f6d497ecf0f8ad054a9683)
- [7.4.238 : Vim does not support the smack library](http://code.google.com/p/vim/source/detail?r=410ef4f1a3d2f4a6ecad9aaa87dae645d1578a19)
- [7.4.239 : ":e +" does not position cursor at end of the file](http://code.google.com/p/vim/source/detail?r=98bfec9ea7608f312129475d4ca0ae6d1c6c232e)
- [7.4.240 : ":tjump" shows "\\n" as "\\\\n"](http://code.google.com/p/vim/source/detail?r=8d1ba0a23588932d22ad37cbd87ae3bbd4bfeff8)
- [7.4.241 : submatch() does not distinguish between a NL and a NUL](http://code.google.com/p/vim/source/detail?r=a63d0cd691dc925283815d17d62f4e948d723a59)
- [7.4.242 : getreg() does not distinguish between a NL and a NUL character](http://code.google.com/p/vim/source/detail?r=f084024c0ddbba46aabfafa2996c3f7d13080ab6)
- [7.4.243 : cannot use setreg() to add text that includes a NUL](http://code.google.com/p/vim/source/detail?r=9f8fa56f1906f4f634cd602a7a2b4f8631faf526)
- [7.4.244 : (after 7.4.238) the smack feature causes stray error messages](http://code.google.com/p/vim/source/detail?r=da17c7de616e3829e4f59923ffa138a067928d9e)
- [7.4.245 : crash for "vim -u NONE -N  -c '&&'"](http://code.google.com/p/vim/source/detail?r=80421d934ebde183ce545ab8d9eb3a4c2065c169)
- [7.4.246 : configure message for detecting smack are out of sequence](http://code.google.com/p/vim/source/detail?r=2cbac5608217cdce725a0f04d397d4466a3322fe)
- [7.4.247 : NUL and NL mixed up when giving input to system()](http://code.google.com/p/vim/source/detail?r=76863b4b48a3684905527dcc688e5a8c4c5e23d9)
- [7.4.248 : cannot distinguish between NL and NUL in output of system()](http://code.google.com/p/vim/source/detail?r=e5f1f2ea0b4a4834791924880f78272ef52eb087)
- [7.4.249 : using setreg() with a list of numbers does not work](http://code.google.com/p/vim/source/detail?r=0b9a66ea49f435536745be0e0a6154be7b607249)
- [7.4.250 : some test files missing from distribution](http://code.google.com/p/vim/source/detail?r=a8f3f45896288bd7e0a27e0c28c3cc3457ccc507)
- [7.4.251 : crash when BufAdd autocommand wipes out the buffer](http://code.google.com/p/vim/source/detail?r=29eb4c2a33ac701bfcd4d2e2bed7864eba876e0e)
- [7.4.252 : critical error in GTK, removing timer twice](http://code.google.com/p/vim/source/detail?r=a43a8262e1ceffc393625c53a1ffed9e6bcaba62)
- [7.4.253 : crash when using external reference in syntax regexp ](http://code.google.com/p/vim/source/detail?r=4901a36479f200b2e6700ad91c26911d92deb886)
- [7.4.254 : smack support detection is incomplete](http://code.google.com/p/vim/source/detail?r=251acc686ca41e4bccb037ef44cd7b486774d580)
- [7.4.255 : configure check for smack doesn't work with all shells](http://code.google.com/p/vim/source/detail?r=5595506b985a198abae41ab0150ee50b8bf1686c)
- [7.4.256 : (after 7.4.248) systemlist() does not work properly](http://code.google.com/p/vim/source/detail?r=afb542ea210cb9fc5fa8c5359bb4814370024b80)
- [7.4.257 : compiler warning, possibly for mismatch in parameter name](http://code.google.com/p/vim/source/detail?r=17903ded5e9a9d49ca73b324657b944f2954d4fd)
- [7.4.258 : configure fails if $CC contains options](http://code.google.com/p/vim/source/detail?r=e8ffd1e6c8dc62c604d34e879791404bd15cab33)
- [7.4.259 : warning for misplaced "const"](http://code.google.com/p/vim/source/detail?r=e4cd5bb75029d2c1208f3e31ebde4e03b16e8123)
- [7.4.260 : no error for function name with colon or lowercase char](http://code.google.com/p/vim/source/detail?r=6bc874e4789a0f912b4fd6b23afecf19d80b1605)
- [7.4.261 : an interactive :s to replace a "\\n" with line break may fail](http://code.google.com/p/vim/source/detail?r=43c6cd07c8defd8505acbe479c6970764c08e6f9)
- [7.4.262 : duplicate code in regexec()](http://code.google.com/p/vim/source/detail?r=0ea551fa607dc443b97c2fba97dc0c9cb0bcf303)
- [7.4.263 : GCC 4.8 compiler warning for hiding a declaration](http://code.google.com/p/vim/source/detail?r=af1bb39774f41c28eabd24d80cffc775695bc124)
- [7.4.264 : (after 7.4.260) can't define function g:Foo()](http://code.google.com/p/vim/source/detail?r=00acac0af680c2d8c82db5258474b121a5908926)
- [7.4.265 : (after 7.4.260) can't call function with "g:" in an expression](http://code.google.com/p/vim/source/detail?r=8ec9d2196bee0c5108f2d2c196a660a7f4e5f29f)
- [7.4.266 : test 62 fails](http://code.google.com/p/vim/source/detail?r=8f84e906d454a95d3167678a745dde9de442b604)
- [7.4.267 : (after 7.4.178) '\[ mark is in the wrong position after "gq"](http://code.google.com/p/vim/source/detail?r=75f222d67cea335efbe0274de6340dba174c1e7e)
- [7.4.268 : exists() on a funcref for script-local function does not work](http://code.google.com/p/vim/source/detail?r=1a5ed2626b26a982e307a206572121a557adf709)
- [7.4.269 : CTRL-U in Insert mode does not work after using a cursor key](http://code.google.com/p/vim/source/detail?r=81c26975e8f9dc7435353581346542409403f296)
- [7.4.270 : comparing pointers instead of the string they point to](http://code.google.com/p/vim/source/detail?r=c519c446c5488bfd48c93a03efae4ae3e0c1f162)
- [7.4.271 : compiler warning on 64 bit windows](http://code.google.com/p/vim/source/detail?r=88b0571de4327ba5127a483493bd7d46e6a9850e)
- [7.4.272 : using just "$" does not cause an error message](http://code.google.com/p/vim/source/detail?r=00228400629e28384f7f52556c3c119ba0d0a44d)
- [7.4.273 : "make autoconf/reconfig" may run configure pointlessly](http://code.google.com/p/vim/source/detail?r=747afb6a6de437131d9674e6b2bcc73d74863a89)

## 新着スクリプト

- [apprentice.vim : A dark low-contrast colorscheme.](http://www.vim.org/scripts/script.php?script_id=4905)
- [Peggi : A Parsing framework for Parsing Expression Grammar](http://www.vim.org/scripts/script.php?script_id=4906)
- [angular.vim : Some niceties for the AngularJS framework](http://www.vim.org/scripts/script.php?script_id=4907)
- [visualHtml : script to preview in live html, php and css edition](http://www.vim.org/scripts/script.php?script_id=4908)
- [vim-sails : Vim navigation plugin for Sails framework. Heavily inspired by vim-rails.](http://www.vim.org/scripts/script.php?script_id=4909)
- [nc.vim : Syntax highlighting for NC (G-Code) files](http://www.vim.org/scripts/script.php?script_id=4910)
- [LineComplete : Insert mode completion of entire lines based on looser matching.](http://www.vim.org/scripts/script.php?script_id=4911)
- [AlphaComplete : Insert mode completion based on any sequence of alphabetic characters.](http://www.vim.org/scripts/script.php?script_id=4912)
- [ttfzoom : For gvim, provide font zooming with Ctrl-MouseWheel like others](http://www.vim.org/scripts/script.php?script_id=4913)
- [repeatable-motions.vim : Make most motions repeatable](http://www.vim.org/scripts/script.php?script_id=4914)
- [heroku.vim : Heroku toolbelt and hk wrapper](http://www.vim.org/scripts/script.php?script_id=4915)
- [vim-onoff : Turn Vim options on. Off.](http://www.vim.org/scripts/script.php?script_id=4916)
- [vim-lldb : lldb debugger integration for vim](http://www.vim.org/scripts/script.php?script_id=4917)
- [otaku : Inspired by Lucky Star](http://www.vim.org/scripts/script.php?script_id=4918)
- [Tabv : Easily open relevant groupings of files as a tab in Vim](http://www.vim.org/scripts/script.php?script_id=4919)
- [Spiffy Foldtext : easy foldtext via a format string, with added features](http://www.vim.org/scripts/script.php?script_id=4920)
- [Simple TODO : A simple plugin that can use VIM as TODO](http://www.vim.org/scripts/script.php?script_id=4921)
- [FindOccurrence : Extended mappings for :isearch, :ilist and :ijump.](http://www.vim.org/scripts/script.php?script_id=4922)
- [autoswap.vim : Switch to open editor window instead of asking what to do with swapfile](http://www.vim.org/scripts/script.php?script_id=4923)
- [FormatToWidth : Apply the gq command to the selected / count width.](http://www.vim.org/scripts/script.php?script_id=4924)
- [snare : Snippets Again, with Regular Expressions](http://www.vim.org/scripts/script.php?script_id=4925)
- [vim-clang : Yet another C/C++ code completion plugin based on Clang](http://www.vim.org/scripts/script.php?script_id=4926)
- [Autosave : Auto save new empty buffer to backup folder.](http://www.vim.org/scripts/script.php?script_id=4927)
- [clang&#x5f;pro.vim : use clang  to complete c/c++ ,and gtags make a proj to find define and callers](http://www.vim.org/scripts/script.php?script_id=4928)
- [unite-spell-suggest : A spelling suggestion source for unite.vim](http://www.vim.org/scripts/script.php?script_id=4929)
- [EightHeader : Easily create custom (fold)headers, foldtext, toc, etc.](http://www.vim.org/scripts/script.php?script_id=4930)
- [tslime.vim : Send commands to a tmux repl session via vim.](http://www.vim.org/scripts/script.php?script_id=4931)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (2906)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2716)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2370)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1413)
5. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (1386)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1058)
7. [Pydiction : Tab-complete your Python code](http://www.vim.org/scripts/script.php?script_id=850) (1029)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (985)
9. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (912)
10. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (889)

## vim-jp issues

Open : 185 (+2) | Closed : 369 (+10)

- [Issue #543 : 7.4.227 「options」コマンド実行後エラー発生](https://github.com/vim-jp/issues/issues/543)
- [Issue #544 : os&#x5f;unix.c:2818にて型修飾子が間違っている。](https://github.com/vim-jp/issues/issues/544)
- [Issue #545 : 日本語を入力するとドットレジスタの中身がヒドイことになってる](https://github.com/vim-jp/issues/issues/545)
- [Issue #546 : gvim on Linux で f の後に日本語を指定できない](https://github.com/vim-jp/issues/issues/546)
- [Issue #547 : Stop adding me to the org!](https://github.com/vim-jp/issues/issues/547)
- [Issue #548 : systemlist() の戻り値を参照するとSEGVする](https://github.com/vim-jp/issues/issues/548)
- [Issue #549 : set nocompatibleしてある環境で:sだけ実行するとSEGV](https://github.com/vim-jp/issues/issues/549)
- [Issue #550 : neobundleのインストール](https://github.com/vim-jp/issues/issues/550)
- [Issue #551 : tabedit での :mksession! が不正](https://github.com/vim-jp/issues/issues/551)
- [Issue #552 : Chrome からのファイル D&D で内部エラー E341](https://github.com/vim-jp/issues/issues/552)
- [Issue #553 : コマンドラインのマッピングについて](https://github.com/vim-jp/issues/issues/553)
- [Issue #554 : Startifyの使い方](https://github.com/vim-jp/issues/issues/554)

