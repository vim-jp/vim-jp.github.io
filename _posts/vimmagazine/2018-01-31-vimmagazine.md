---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 01 月号

---
## 話題

*   [vital.vim開発者会議2018-02](https://fablicvim.connpass.com/event/74842/)が 2018-02-15 (木) に開催されます
*   [Meguro.vim #8](https://megurovim.connpass.com/event/76881/)が 2018-02-17 (土) に開催されます
*   [Osaka.vim #12](https://osaka-vim.connpass.com/event/77504/)が 2018-02-17 (土) に開催されます

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1436: `has()` で確認可能な機能 `python_compiled`, `python3_compiled`, `python_dynamic`, `python3_dynamic` が追加されました
*   8.0.1445: コマンドラインが編集された時に実行される autocommand イベント `CmdlineChanged` が追加されました

*   ランタイムの更新
    *   [NSIS script のファイルタイププラグイン](https://github.com/k-takata/vim-nsis) のメンテナが変更され、プラグインが更新されました
    *   autodoc のシンタックスが追加されました
    *   Pike の プラグインのメンテナが変更され、更新されました
    *   CMOD のファイルタイププラグインが追加されました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.1429 : crash when calling term&#x5f;start() with empty argument](https://github.com/vim/vim/commit/ede35bbbd05306097bf3f4d603f2248ed1f4c5f1)
- [8.0.1430 : crash when term&#x5f;start() fails](https://github.com/vim/vim/commit/4aad53c36998c77ee119c6a46d002460876c2cbb)
- [8.0.1431 : MS-Windows: vimtutor fails if %TMP% has special chars](https://github.com/vim/vim/commit/0cbcd949e15ad95171e5b33881d3a30f17073dda)
- [8.0.1432 : after ":copen" can't get the window-ID of the quickfix window](https://github.com/vim/vim/commit/2ec364e94dbc080ccdf6c5dfc6f1653b5b7ded64)
- [8.0.1433 : illegal memory access after undo](https://github.com/vim/vim/commit/95dbcbea6d85a5b79d9617ab3863458fdf0217a0)
- [8.0.1434 : GTK: :promtfind does not put focus on text input](https://github.com/vim/vim/commit/d7823d5b7c32f73ca720373ea9c16b1b47f086df)
- [8.0.1435 : memory leak in test&#x5f;arabic](https://github.com/vim/vim/commit/501383236d203bacad758e82e47a07b877a3b63b)
- [8.0.1436 : not enough information about what Python version may work](https://github.com/vim/vim/commit/84b242c369a22b581c43de9de0152f0baedd71ab)
- [8.0.1437 : pkg-config doesn't work with cross compiling](https://github.com/vim/vim/commit/d6d304298a6b9842164a53e7be386d55d18ce79f)
- [8.0.1438 : filetype detection test not updated for change](https://github.com/vim/vim/commit/0479e910c423d71e2b96bc721feffad5808e767a)
- [8.0.1439 : if cscope fails a search Vim may hang](https://github.com/vim/vim/commit/1274d33493efb6250470a37b9f4432bb31e87d64)
- [8.0.1440 : terminal window: some vterm responses are delayed](https://github.com/vim/vim/commit/b50773c6df0bc2c9c2ab1afecc78083abc606de0)
- [8.0.1441 : using ":undo 0" leaves undo in wrong state](https://github.com/vim/vim/commit/ce46d934af35d0f774be7f996001db03cf0b894a)
- [8.0.1442 : (after 8.0.1439) using pointer before it is set](https://github.com/vim/vim/commit/a172b63ab8661019dba61285a738c8b6b55a33aa)
- [8.0.1443 : (after 8.0.1441) compiler gives uninitialized var warning](https://github.com/vim/vim/commit/059fd01021779ee369c1e55557275f6c349fda9e)
- [8.0.1444 : missing -D&#x5f;FILE&#x5f;OFFSET&#x5f;BITS=64 may cause problems](https://github.com/vim/vim/commit/ec0557f08b2660118eaedb94471e5ab0f87cf2a3)
- [8.0.1445 : cannot act on edits in the command line](https://github.com/vim/vim/commit/153b704e20f9c269450a7d3ea8cafcf942579ab7)
- [8.0.1446 : acessing freed memory after window command in auto command](https://github.com/vim/vim/commit/6f361c991221e96d5068c77b854967d997b1529b)
- [8.0.1447 : still too many old style tests](https://github.com/vim/vim/commit/cada78975eebc47f9b12de1a471639b5afd9ad2f)
- [8.0.1448 : segfault with exception inside :rubyfile command](https://github.com/vim/vim/commit/37badc898b8d167e11553b6d05908ffd35928a6e)
- [8.0.1449 : slow redrawing with DirectX](https://github.com/vim/vim/commit/a338adcf222b6a24e26ea5ae6a2ad27f914acb38)
- [8.0.1450 : GUI: endless loop when stopping cursor blinking](https://github.com/vim/vim/commit/1dd45fb4f3371f0256653b2186c8b4b3d26b3f41)
- [8.0.1451 : difficult to set the python home directories properly](https://github.com/vim/vim/commit/94073167e3aa8cbe18380e93a2fc8e8165438cc8)

## 新着スクリプト

- [pike.vim : New and improved Pike syntax colouring](https://vim.sourceforge.io/scripts/script.php?script_id=5640)
- [autodoc.vim : Autodoc syntax colouring](https://vim.sourceforge.io/scripts/script.php?script_id=5641)
- [twiggy : Git branch management](https://vim.sourceforge.io/scripts/script.php?script_id=5643)
- [If I Only... : An only slightly smarter :only](https://vim.sourceforge.io/scripts/script.php?script_id=5644)
- [cmod.vim : Cmod syntax colouring (Pike development)](https://vim.sourceforge.io/scripts/script.php?script_id=5645)
- [Pickachu : Color, date, and file picker for Vim using Zenity](https://vim.sourceforge.io/scripts/script.php?script_id=5646)
- [jbuilder.vim : Jbuilder syntax hilighting](https://vim.sourceforge.io/scripts/script.php?script_id=5647)
- [textfilter : Several useful text filters.](https://vim.sourceforge.io/scripts/script.php?script_id=5648)
- [Gruvbox 8 : Simplified and optimized Gruvbox colorscheme](https://vim.sourceforge.io/scripts/script.php?script_id=5649)
- [Realcolors : Utilities for tweaking and creating 24-bit true color schemes for Vim at runtime](https://vim.sourceforge.io/scripts/script.php?script_id=5650)
- [qt-support.vim : Simple Qt support for Vim](https://vim.sourceforge.io/scripts/script.php?script_id=5651)
- [vim-fetlang : Vim language support for Fetlang, the fetish-themed programming language](https://vim.sourceforge.io/scripts/script.php?script_id=5652)
- [vim-auf : Format Your Lines And Be Reminded By Coloring](https://vim.sourceforge.io/scripts/script.php?script_id=5653)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://vim.sourceforge.io/scripts/script.php?script_id=273) (1031)
2. [wombat256.vim : Wombat for 256 color xterms](https://vim.sourceforge.io/scripts/script.php?script_id=2465) (781)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](https://vim.sourceforge.io/scripts/script.php?script_id=1658) (713)
4. [OmniCppComplete : C/C++ omni-completion with ctags database](https://vim.sourceforge.io/scripts/script.php?script_id=1520) (368)
5. [python.vim : Enhanced version of the python syntax highlighting script](https://vim.sourceforge.io/scripts/script.php?script_id=790) (365)
6. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](https://vim.sourceforge.io/scripts/script.php?script_id=2945) (299)
7. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](https://vim.sourceforge.io/scripts/script.php?script_id=213) (289)
8. [nginx.vim : initial version](https://vim.sourceforge.io/scripts/script.php?script_id=1886) (286)
9. [AutoComplPop : Automatically opens popup menu for completions](https://vim.sourceforge.io/scripts/script.php?script_id=1879) (269)
10. [yaml.vim : Syntax coloring and functions for YAML](https://vim.sourceforge.io/scripts/script.php?script_id=739) (262)

## vim-jp issues

Open : 259 (+4) | Closed : 887 (+5)

- [Issue #1138 : :rubyfile が必ずSEGVする](https://github.com/vim-jp/issues/issues/1138)
- [Issue #1139 : スワップファイルが壊れる](https://github.com/vim-jp/issues/issues/1139)
- [Issue #1140 : Vim の :terminal で非カノニカルモード (min=1) の read() がブロックする](https://github.com/vim-jp/issues/issues/1140)
- [Issue #1141 : エスケープ文字を入れた時に DFA エンジンと NFA エンジンで結果が異なる](https://github.com/vim-jp/issues/issues/1141)
- [Issue #1142 : 'undo 1' 後に redo すると元に戻るが 'undo 0' 後だと戻らない](https://github.com/vim-jp/issues/issues/1142)
- [Issue #1143 : ~/.vim/pack が symlink だとプラグイン中の一部ファイルが読み込まれない](https://github.com/vim-jp/issues/issues/1143)
- [Issue #1144 : histdel関数を実行後、vimを再起動するとhistoryが復元している](https://github.com/vim-jp/issues/issues/1144)
- [Issue #1145 : ハイライトグループ "CursorIM" が反映されない](https://github.com/vim-jp/issues/issues/1145)
- [Issue #1146 : 矩形選択で全角文字への置換が出来ない](https://github.com/vim-jp/issues/issues/1146)
