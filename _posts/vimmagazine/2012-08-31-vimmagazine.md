---
layout: post
category: vimmagazine
title: Vim Magazine 8 月号
---

## 話題

- [vimrc読書会](http://vim-jp.org/reading-vimrc/)

  毎週土曜日夜23時(JST)にオンラインのチャットルーム [LingrのVim部屋](http://lingr.com/room/vim) にて vimrc 読書会が開催されています。お気軽にご参加ください。

- [TokyoVim#10](http://partake.in/events/c719e827-b312-4b71-a0b8-862f153d8045)

  Vimに関する何かをやる会 東京 9月1日


## 今月の新機能

  - 7.3.629

    `'shiftwidth'` オプションに `0` を設定することで `'tabstop'` の値を使
    うように設定できるようになりました。

  - 7.3.631

    `:edit ~user<Tab>` でユーザー名の補完ができるようになりました。
    `:command` に `-complete=user` が追加されました。

  - 7.3.632

    SGR形式の拡張マウストラッキングのサポートが追加されました。
    対応端末で 222 以上の座標を扱えるようになります。


## リリース情報

- [7.3.619 : when executing a shell command Vim may become slow to respond](https://github.com/vim/vim/commit/e4195c5d84027a3137ad4bc0ba5b6bd14b8a3a04)
- [7.3.620 : building with recent Ruby on Win32 doesn't work](https://github.com/vim/vim/commit/da3cb831e938d672361995d1dec2de3dae72ee5b)
- [7.3.621 : compiler warnings on 64 bit windows](https://github.com/vim/vim/commit/09168a77e4d191081ee7eefd98c13a91f197e896)
- [7.3.622 : XPM library for Win32 can't be found](https://github.com/vim/vim/commit/e66822e3d1aa4eac1cb4b5a47a4a9c4a28581021)
- [7.3.623 : Perl 5.14 commands crash Vim on MS-Windows](https://github.com/vim/vim/commit/c271c48e8043ec388487b7489fb4692a154791d8)
- [7.3.624 : when cancelling input() it returns the third argument](https://github.com/vim/vim/commit/04b2751899886a18399286b92ef097ec8a74e911)
- [7.3.625 : "gn" does not handle zero-width matches correctly](https://github.com/vim/vim/commit/ba6ba36fa2ca4a6877e1cca9c62d48bd37188b8e)
- [7.3.626 : Python interface doesn't build with Python 2.4 or older](https://github.com/vim/vim/commit/c1a995db49854947a94e60be918c03a643a4292c)
- [7.3.627 : expression not evaluated when using "n" flag with ":s"](https://github.com/vim/vim/commit/07e31c571a13855cf67c69f3a676038510141699)
- [7.3.628 : ":open!" results in a confusing error message](https://github.com/vim/vim/commit/b02612b6411659b6385aea44fe8abb5a2e16d4d7)
- [7.3.629 : there is no way to make 'shiftwidth' follow 'tabstop'](https://github.com/vim/vim/commit/14f247414733b9dafb50e59320f75e8b56ac6ec6)
- [7.3.630 : "|" does not behave correctly when 'virtualedit' is set](https://github.com/vim/vim/commit/2dac213ac6896bffbbfdd9afbdb0ad0e8b032532)
- [7.3.631 : cannot complete user names](https://github.com/vim/vim/commit/2430586629dfccd637a23c45793fc182d2cb2b87)
- [7.3.632 : cannot select beyond 222 columns with the mouse in xterm](https://github.com/vim/vim/commit/2b9578f0f8cdba144e996273bca586b384e33d90)
- [7.3.633 : selection remains highlighted after selecting another text](https://github.com/vim/vim/commit/dd82d699c8cb3e60b2da968fb40622dc497f71ca)
- [7.3.634 : Month/Day format for undo is confusing](https://github.com/vim/vim/commit/b9ce83e135331d74f09ec094528243f33c3ada85)
- [7.3.635 : system call during startup sets 'lines' to a wrong value](https://github.com/vim/vim/commit/4d93dc2699845009e0f263ed10196309c1175c2e)
- [7.3.636 : (after 7.3.625) "gn" fails for some zero-width matches](https://github.com/vim/vim/commit/dde0efed85c0f78048409ef6deada0c6e8d0be3c)
- [7.3.637 : cannot catch error caused by a foldopen when there is no fold](https://github.com/vim/vim/commit/00b8ae0d3d90f80fc817857dca26359b175b18ed)
- [7.3.638 : unecessary redraw of the previous character](https://github.com/vim/vim/commit/451cf63751673e5cc7ecfbd112325217d9a955b5)
- [7.3.639 : it's not easy to build Vim on Windows with XPM support](https://github.com/vim/vim/commit/c6c1d8b80236305f7e7b053be8dea63688fa648e)
- [7.3.640 : binary files for 7.3.639](https://github.com/vim/vim/commit/18a1b9bd5fa5925a2dc4b702b37ea0217fcb6ea3)
- [7.3.641 : ":mkview" uses ":normal" instead of ":normal!" for folds](https://github.com/vim/vim/commit/002a4edc5b66c90245ca3eebc564635ed0af1ff1)
- [7.3.642 : segfault with specific autocommands](https://github.com/vim/vim/commit/8f913993666b154b233ab9923845fa165d9b3260)
- [7.3.643 : MS-Windows: 'lines' is wrong when starting up maximized](https://github.com/vim/vim/commit/3633dc012919238e66c006b1037e2a1fece2b18d)
- [7.3.644 : dead code for BeOS GUI](https://github.com/vim/vim/commit/0551410a14d8fdfa657392c0d6b3cc8dd643752d)
- [7.3.645 : no tests for patch 7.3.625 and 7.3.637](https://github.com/vim/vim/commit/dd7d84647521723a0c95b72752beccdb7ce2cd02)
- [7.3.646 : undo file unusable after reloading a buffer](https://github.com/vim/vim/commit/465748e41179169f9b432deaae9cabd0e55bb9e3)

## 新着スクリプト

- [stl : C++ STL Syntax File](http://www.vim.org/scripts/script.php?script_id=4153)
- [allegro : Allegro gaming library syntax file](http://www.vim.org/scripts/script.php?script_id=4154)
- [AlignFromCursor : Perform :left / :right only for the text part right of the cursor.](http://www.vim.org/scripts/script.php?script_id=4155)
- [LogViewer : Comfortable examination of multiple parallel logfiles.](http://www.vim.org/scripts/script.php?script_id=4156)
- [pb.vim : plugin for interacting with Mac OS X pastebuffer](http://www.vim.org/scripts/script.php?script_id=4157)
- [bwm color scheme : Color scheme](http://www.vim.org/scripts/script.php?script_id=4158)
- [phpvim : PHP syntax file updated for PHP 5.4](http://www.vim.org/scripts/script.php?script_id=4159)
- [Drupal : A very simple plugin to help develop with Drupal](http://www.vim.org/scripts/script.php?script_id=4160)
- [flexddpl.vim : Basic highlighting for FlexDDPL policies](http://www.vim.org/scripts/script.php?script_id=4161)
- [Simple CSS align (SCSSA) : Align selected CSS properties](http://www.vim.org/scripts/script.php?script_id=4162)
- [michael.vim : A dark background color scheme with (mostly) matte colors](http://www.vim.org/scripts/script.php?script_id=4163)
- [hybrid.vim : A dark colourscheme combining Jellybeans, Solarized and Tomorrow-Night](http://www.vim.org/scripts/script.php?script_id=4164)
- [directoryBrowser : Directory listing and file operations in vim](http://www.vim.org/scripts/script.php?script_id=4165)
- [locator : Shows where is your cursor, including folds, classes, functions, and more.](http://www.vim.org/scripts/script.php?script_id=4166)
- [Area search : Search for area containing keywords](http://www.vim.org/scripts/script.php?script_id=4167)
- [SyntaxRange : Define a different filetype syntax on regions of a buffer.](http://www.vim.org/scripts/script.php?script_id=4168)
- [VPSIONKI DATA : Portable Vim code to handle database as CSV file (eg. for contacts)](http://www.vim.org/scripts/script.php?script_id=4169)
- [Vdebug : Multi-language DBGP debugger client (PHP, Python, Perl, Ruby, etc.)](http://www.vim.org/scripts/script.php?script_id=4170)
- [sideways.vim : Move function arguments (and other delimited-by-something items) left and right](http://www.vim.org/scripts/script.php?script_id=4171)
- [switch.vim : Switch segments of text with predefined replacements](http://www.vim.org/scripts/script.php?script_id=4172)
- [GrepCommands : Perform :vimgrep over arguments, buffers, windows.](http://www.vim.org/scripts/script.php?script_id=4173)
- [rgbr : a simple color scheme for R](http://www.vim.org/scripts/script.php?script_id=4174)
- [rslate : a modified slate theme for R](http://www.vim.org/scripts/script.php?script_id=4175)
- [Rainbow Parentheses Improved : help you read complex code by showing diff level of parentheses in diff color !!](http://www.vim.org/scripts/script.php?script_id=4176)
- [undotree.vim : Display your undo history in a graph.](http://www.vim.org/scripts/script.php?script_id=4177)
- [CursorLineCurrentWindow : Only highlight the screen line of the cursor in the currently active window.](http://www.vim.org/scripts/script.php?script_id=4178)
- [X Marks the Spot : Easy mappings for creating and navigating through lower-case marks](http://www.vim.org/scripts/script.php?script_id=4179)
- [vim-dict : The Dict client for Vim](http://www.vim.org/scripts/script.php?script_id=4180)
- [omgrofl.vim : syntax file for the Omgrofl language](http://www.vim.org/scripts/script.php?script_id=4181)
- [shellasync.vim : shellasync.vim plugin for asynchronously executing shell commands in vim](http://www.vim.org/scripts/script.php?script_id=4182)
- [vim-g : Quick Google lookup directly from Vim](http://www.vim.org/scripts/script.php?script_id=4183)
- [pybreak : Fast toggling of set\_trace statement in python files, for easier debugging](http://www.vim.org/scripts/script.php?script_id=4184)
- [PrevInsertComplete : Recall and insert mode completion for previously inserted text.](http://www.vim.org/scripts/script.php?script_id=4185)
- [dwm.vim : Tiled Window Management for Vim](http://www.vim.org/scripts/script.php?script_id=4186)
- [cmd.vim : Color scheme for the windows XP full screen console](http://www.vim.org/scripts/script.php?script_id=4187)
- [HybridText : A blend of useful document syntax for text files](http://www.vim.org/scripts/script.php?script_id=4188)
- [Close Location List Before Quit : Simple closes the location list before quit](http://www.vim.org/scripts/script.php?script_id=4189)
- [easymenu.vim : Data-driven, hierarchical Menu Creation](http://www.vim.org/scripts/script.php?script_id=4190)
- [GrepHere : List occurrences in the current buffer in the quickfix window.](http://www.vim.org/scripts/script.php?script_id=4191)
- [html\_annoyance : HTML syntax file, removes annoying underlining from \<a> whitespace](http://www.vim.org/scripts/script.php?script_id=4192)
- [UtilityFunctions : Data structures and utility functions](http://www.vim.org/scripts/script.php?script_id=4193)
- [FileTags : Easily locate previously-used files, directories and URLs](http://www.vim.org/scripts/script.php?script_id=4194)
- [harlequin : Dark, high contrast colorscheme inspired by molokai and badwolf](http://www.vim.org/scripts/script.php?script_id=4195)
- [GrepTasks : Grep for tasks and TODO markers.](http://www.vim.org/scripts/script.php?script_id=4196)
- [numsign.vim :  jump between signs  according sign ID or Line Number](http://www.vim.org/scripts/script.php?script_id=4197)
- [asyncfinder.vim : asyncfinder.vim - simple asynchronous fuzzy file finder for vim](http://www.vim.org/scripts/script.php?script_id=4198)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (3285)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2776)
3. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (1540)
4. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (1407)
5. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1304)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1191)
7. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1163)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (1157)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (1155)
10. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (1070)

## vim-jp issues

Open : 129 (+5) | Closed : 116 (+9)

- [Issue #232 : Winでのマルチスレッドバグの可能性を検証する](https://github.com/vim-jp/issues/issues/232)
- [Issue #233 : Vim 7.3.604からinput()のキャンセルの挙動がおかしい](https://github.com/vim-jp/issues/issues/233)
- [Issue #234 : windows 2000 をサポートすべきか](https://github.com/vim-jp/issues/issues/234)
- [Issue #235 : :mksession で落ちる](https://github.com/vim-jp/issues/issues/235)
- [Issue #236 : :mksession で落ちる](https://github.com/vim-jp/issues/issues/236)
- [Issue #237 : 'errorbells'がおかしい？](https://github.com/vim-jp/issues/issues/237)
- [Issue #238 : Windows環境下にて「sign unplace * file=C:\...\hoge.vim」が動かない。](https://github.com/vim-jp/issues/issues/238)
- [Issue #239 : rで日本語に書き換えができない](https://github.com/vim-jp/issues/issues/239)
- [Issue #240 : windowsでユーザー名補完が出来ない](https://github.com/vim-jp/issues/issues/240)
- [Issue #241 : Windowsのシンボリックリンクを正しく扱えるようにしたい](https://github.com/vim-jp/issues/issues/241)
- [Issue #242 : e\_noroomの翻訳を修正する](https://github.com/vim-jp/issues/issues/242)
- [Issue #243 : input()の complete引数="command" で、コマンドの引数を補完できない](https://github.com/vim-jp/issues/issues/243)
- [Issue #244 : ACLが保存・復帰されない (Cygwinの属性の保存)](https://github.com/vim-jp/issues/issues/244)
- [Issue #245 : getbufvar(), getwinvar(), gettabvar()にデフォルト値を指定できるようにしたい](https://github.com/vim-jp/issues/issues/245)

