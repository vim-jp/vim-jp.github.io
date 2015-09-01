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
- 7.4.235 : `exepath()` 関数が追加されました。コマンドのフルパスを得ることができます。
- 7.4.236 : パッチバージョンが `has("patch-7.4.123")` で確認できるようになりました。
- 7.4.238 : smack (Linux のセキュリティモジュールの一つ) のサポートが追加されました。
- 7.4.241 : `submatch()` にリストを返すための第2引数が追加されました。
- 7.4.242 : `getreg()` にリストを返すための別の引数が追加されました。
- 7.4.243 : `setreg()` がリストを許容するようになりました。
- 7.4.247 : `system()` の input に任意でリストを使うようになりました。
- 7.4.248 : `systemlist()` が追加されました。

## リリース情報

- [7.4.230 : error when using ":options"](https://github.com/vim/vim/commit/e7a88a8d4ebd2419ad926578ddc0ee4bf43cf069)
- [7.4.231 : an error in ":options" is not caught by the tests](https://github.com/vim/vim/commit/efa304d760640e5c575bef97d689f7c360b20779)
- [7.4.232 : ":%s/\\n//" uses a lot of memory](https://github.com/vim/vim/commit/fd3fe98b78edb10f2d439b0c6d56e6d70eecb92d)
- [7.4.233 : escaping special chars in ":!cmd %" is inconsistant](https://github.com/vim/vim/commit/31b7d38611541030fe54898a7d8b6a714fa0fa59)
- [7.4.234 : can't get the command that was used to start Vim](https://github.com/vim/vim/commit/a1706c958e69086f5c9eb7d79779ed839441ff60)
- [7.4.235 : it is not easy to get the full path of a command](https://github.com/vim/vim/commit/c7f025536e9e5b7b95b55d09094febb627395d05)
- [7.4.236 : it's not that easy to check the Vim patch version](https://github.com/vim/vim/commit/7f3be402cecb458ac2a8d385bb7303d8b27f9af4)
- [7.4.237 : (after 7.4.236) has() not checking for specific patch](https://github.com/vim/vim/commit/6716d9af113a5c4bcc7fdf8fd24f3d633db386c7)
- [7.4.238 : Vim does not support the smack library](https://github.com/vim/vim/commit/5bd32f47ec5121a5485d180add1dacb73472e4b2)
- [7.4.239 : ":e +" does not position cursor at end of the file](https://github.com/vim/vim/commit/3e45159ceb465082c183e79e2547cb6cc408167e)
- [7.4.240 : ":tjump" shows "\\n" as "\\\\n"](https://github.com/vim/vim/commit/fe5aab63feb2b03656700d3738d46a19e99edde0)
- [7.4.241 : submatch() does not distinguish between a NL and a NUL](https://github.com/vim/vim/commit/41571769c9a236fd07b333a5eb98c461636b466c)
- [7.4.242 : getreg() does not distinguish between a NL and a NUL character](https://github.com/vim/vim/commit/b7cb42bc3878fcb62ed407f47f0a2cc960aa7c1e)
- [7.4.243 : cannot use setreg() to add text that includes a NUL](https://github.com/vim/vim/commit/5a50c2255c447838d08d3b4895a3be3a41cd8eda)
- [7.4.244 : (after 7.4.238) the smack feature causes stray error messages](https://github.com/vim/vim/commit/57a728d1df7451f5b2b6b1f933182b5af9513d83)
- [7.4.245 : crash for "vim -u NONE -N  -c '&&'"](https://github.com/vim/vim/commit/21e854e5ce08ad419441136b1698d535ff818a72)
- [7.4.246 : configure message for detecting smack are out of sequence](https://github.com/vim/vim/commit/4ed89cdf41cf3d509fd30e6f7c577a7c090679a9)
- [7.4.247 : NUL and NL mixed up when giving input to system()](https://github.com/vim/vim/commit/57ebe6e2f94edad6adc43246d98919e728095211)
- [7.4.248 : cannot distinguish between NL and NUL in output of system()](https://github.com/vim/vim/commit/39c29ed5118ab513554d1d51d6a98e65f32784ba)
- [7.4.249 : using setreg() with a list of numbers does not work](https://github.com/vim/vim/commit/7d647820ed7c747bbc5618366ce6dfcf0006398d)
- [7.4.250 : some test files missing from distribution](https://github.com/vim/vim/commit/e9a54227f93bd658988d853684aa24529af9414b)
- [7.4.251 : crash when BufAdd autocommand wipes out the buffer](https://github.com/vim/vim/commit/4c7ab1bb5722de662db04550b74256671f20c4a2)
- [7.4.252 : critical error in GTK, removing timer twice](https://github.com/vim/vim/commit/7bcdb7d1669b0a37cd86446dab0c66f2e1971fa7)
- [7.4.253 : crash when using external reference in syntax regexp ](https://github.com/vim/vim/commit/5a4e160ce423f56456227c3d7d8d8a069120b377)
- [7.4.254 : smack support detection is incomplete](https://github.com/vim/vim/commit/c09551ab796af1f8aa19338d99e49267860fcd9c)
- [7.4.255 : configure check for smack doesn't work with all shells](https://github.com/vim/vim/commit/e29b1feead084f19e862bd719f140a490e3ed7f2)
- [7.4.256 : (after 7.4.248) systemlist() does not work properly](https://github.com/vim/vim/commit/b21a29be56fb0e125d9f736bfdef8dde5a1ceae0)
- [7.4.257 : compiler warning, possibly for mismatch in parameter name](https://github.com/vim/vim/commit/0eac828ab0199a639dd3657558fe9ed1411f43cd)
- [7.4.258 : configure fails if $CC contains options](https://github.com/vim/vim/commit/c8836f702532b0bc3dd16972e6b504a7340e90e2)
- [7.4.259 : warning for misplaced "const"](https://github.com/vim/vim/commit/62f167f716beb8bfeaadb8ec506a257827f701a2)
- [7.4.260 : no error for function name with colon or lowercase char](https://github.com/vim/vim/commit/9bdfb0025cba78c7a917f7f9420fe00136918e1c)
- [7.4.261 : an interactive :s to replace a "\\n" with line break may fail](https://github.com/vim/vim/commit/93fc481b578b18c282e60d251455685602449526)
- [7.4.262 : duplicate code in regexec()](https://github.com/vim/vim/commit/2af78a14088a158208eab0853201b8ce07b2ac03)
- [7.4.263 : GCC 4.8 compiler warning for hiding a declaration](https://github.com/vim/vim/commit/163d0da508d1a206e6a6ee7e2a34fa0a583ae16c)
- [7.4.264 : (after 7.4.260) can't define function g:Foo()](https://github.com/vim/vim/commit/eccb7fc3158877d93194e6b7c0f7e542b4544137)
- [7.4.265 : (after 7.4.260) can't call function with "g:" in an expression](https://github.com/vim/vim/commit/a4f317df89e662a964197f2d586ac24cf801f14f)
- [7.4.266 : test 62 fails](https://github.com/vim/vim/commit/33e87789a740fadcae473c60a00de2a13c55a7d0)
- [7.4.267 : (after 7.4.178) '\[ mark is in the wrong position after "gq"](https://github.com/vim/vim/commit/d69bd9af3ccf5edd2138fb7abd68d35f7b84ef7e)
- [7.4.268 : exists() on a funcref for script-local function does not work](https://github.com/vim/vim/commit/355a95a079d3c8c64cf496df434741891d07985a)
- [7.4.269 : CTRL-U in Insert mode does not work after using a cursor key](https://github.com/vim/vim/commit/3d1956bcc9b2f606a8b6b03b3852a6eed9847348)
- [7.4.270 : comparing pointers instead of the string they point to](https://github.com/vim/vim/commit/1ff32c5c2cec3597a5695ca79825ab53441dbf06)
- [7.4.271 : compiler warning on 64 bit windows](https://github.com/vim/vim/commit/121f9bdde4a474ae729fd0b1e5fc9ad1ffcd8651)
- [7.4.272 : using just "$" does not cause an error message](https://github.com/vim/vim/commit/e512c8c049c2e5768c0ea86531093224e2919955)
- [7.4.273 : "make autoconf/reconfig" may run configure pointlessly](https://github.com/vim/vim/commit/b96c69d80e626f1f06ed8646c679968af1a1d923)

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

