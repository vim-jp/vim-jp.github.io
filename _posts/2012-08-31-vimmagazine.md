---
layout: post
category: vimmagazine
title: Vim Magazine 8 月号
---

# Vim Magazine 8 月号

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

- [7.3.619 : when executing a shell command Vim may become slow to respond](http://code.google.com/p/vim/source/detail?r=27ecf0c87bd20140d9e85d4fd581332e0916191e)
- [7.3.620 : building with recent Ruby on Win32 doesn't work](http://code.google.com/p/vim/source/detail?r=a6ebae140d89f35876628cde44f0e7ee5a0ae1d2)
- [7.3.621 : compiler warnings on 64 bit windows](http://code.google.com/p/vim/source/detail?r=ac13ea2b098d98e62408ec2d88026f690f68f940)
- [7.3.622 : XPM library for Win32 can't be found](http://code.google.com/p/vim/source/detail?r=52247eb615118965bb6a949307e9abfd378cd088)
- [7.3.623 : Perl 5.14 commands crash Vim on MS-Windows](http://code.google.com/p/vim/source/detail?r=a46a101bfddd4fcf30759f3847fa1e18546907f4)
- [7.3.624 : when cancelling input() it returns the third argument](http://code.google.com/p/vim/source/detail?r=06749e351b1c030acf50cf8c469f8973aff4a96a)
- [7.3.625 : "gn" does not handle zero-width matches correctly](http://code.google.com/p/vim/source/detail?r=f43ffd820a46bf47b3e55b56d3e5334302383757)
- [7.3.626 : Python interface doesn't build with Python 2.4 or older](http://code.google.com/p/vim/source/detail?r=956b1ed19282518e0287891d4e7b2b8a3f850402)
- [7.3.627 : expression not evaluated when using "n" flag with ":s"](http://code.google.com/p/vim/source/detail?r=dc65e6429d2c1b79719455d7303f3f95c5897965)
- [7.3.628 : ":open!" results in a confusing error message](http://code.google.com/p/vim/source/detail?r=1ed93878d6e224d46a7c125a82e52338b9bf0f6b)
- [7.3.629 : there is no way to make 'shiftwidth' follow 'tabstop'](http://code.google.com/p/vim/source/detail?r=7de7ef01288db1abc944293a28aac13ddd87879e)
- [7.3.630 : "|" does not behave correctly when 'virtualedit' is set](http://code.google.com/p/vim/source/detail?r=3416ac23d0758528919aaffbe02cecc6420b8aac)
- [7.3.631 : cannot complete user names](http://code.google.com/p/vim/source/detail?r=99f076ca8d846d553aefd7e626ca938bcf93f259)
- [7.3.632 : cannot select beyond 222 columns with the mouse in xterm](http://code.google.com/p/vim/source/detail?r=44038a9777aaf0e6502e0eb57b8f7ccf31798ec4)
- [7.3.633 : selection remains highlighted after selecting another text](http://code.google.com/p/vim/source/detail?r=ad5fbf3596e0a8ad15c260c5b529efc5ac689e13)
- [7.3.634 : Month/Day format for undo is confusing](http://code.google.com/p/vim/source/detail?r=b7eea24095073d55ac3c59664466b78a0525bebc)
- [7.3.635 : system call during startup sets 'lines' to a wrong value](http://code.google.com/p/vim/source/detail?r=f7f68f83fc335379bc9e3371c5118b8f81c8d5bc)
- [7.3.636 : (after 7.3.625) "gn" fails for some zero-width matches](http://code.google.com/p/vim/source/detail?r=616bc1ad4f12be0ef37fe67d2f6ed7535f20e515)
- [7.3.637 : cannot catch error caused by a foldopen when there is no fold](http://code.google.com/p/vim/source/detail?r=3b62d8f36cdf709d4e0575d63486ebe57ce7150c)
- [7.3.638 : unecessary redraw of the previous character](http://code.google.com/p/vim/source/detail?r=c0256c4bd91ed7b806aab002e4a45e440dd00343)
- [7.3.639 : it's not easy to build Vim on Windows with XPM support](http://code.google.com/p/vim/source/detail?r=1ec385a8faf4002fa300caf370ef2ac380d9702a)
- [7.3.640 : binary files for 7.3.639](http://code.google.com/p/vim/source/detail?r=c2c3577021ed4e354853d330e5fc60691d8750c5)
- [7.3.641 : ":mkview" uses ":normal" instead of ":normal!" for folds](http://code.google.com/p/vim/source/detail?r=c068389057c9a89fa351678b5f5776278971dabe)
- [7.3.642 : segfault with specific autocommands](http://code.google.com/p/vim/source/detail?r=79ac6744237ce5076a6d2e2b21198dfddc2cc9dd)
- [7.3.643 : MS-Windows: 'lines' is wrong when starting up maximized](http://code.google.com/p/vim/source/detail?r=fdac34e3afa5f0dea26dba5e96932e3b9a0be50a)
- [7.3.644 : dead code for BeOS GUI](http://code.google.com/p/vim/source/detail?r=fb7c028dcac7555fee78096813e3846d157cc91b)
- [7.3.645 : no tests for patch 7.3.625 and 7.3.637](http://code.google.com/p/vim/source/detail?r=1810ee914648ef3e16a8cf866299a32c4a9a00cb)
- [7.3.646 : undo file unusable after reloading a buffer](http://code.google.com/p/vim/source/detail?r=e70485d3f81df4d72cb40d6fa4ad4218cf1e8392)

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
- [pybreak : Fast toggling of set_trace statement in python files, for easier debugging](http://www.vim.org/scripts/script.php?script_id=4184)
- [PrevInsertComplete : Recall and insert mode completion for previously inserted text.](http://www.vim.org/scripts/script.php?script_id=4185)
- [dwm.vim : Tiled Window Management for Vim](http://www.vim.org/scripts/script.php?script_id=4186)
- [cmd.vim : Color scheme for the windows XP full screen console](http://www.vim.org/scripts/script.php?script_id=4187)
- [HybridText : A blend of useful document syntax for text files](http://www.vim.org/scripts/script.php?script_id=4188)
- [Close Location List Before Quit : Simple closes the location list before quit](http://www.vim.org/scripts/script.php?script_id=4189)
- [easymenu.vim : Data-driven, hierarchical Menu Creation](http://www.vim.org/scripts/script.php?script_id=4190)
- [GrepHere : List occurrences in the current buffer in the quickfix window.](http://www.vim.org/scripts/script.php?script_id=4191)
- [html_annoyance : HTML syntax file, removes annoying underlining from \<a> whitespace](http://www.vim.org/scripts/script.php?script_id=4192)
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
- [Issue #240 : windowsでユーザ名補完が出来ない](https://github.com/vim-jp/issues/issues/240)
- [Issue #241 : Windowsのシンボリックリンクを正しく扱えるようにしたい](https://github.com/vim-jp/issues/issues/241)
- [Issue #242 : e_noroomの翻訳を修正する](https://github.com/vim-jp/issues/issues/242)
- [Issue #243 : input()の complete引数="command" で、コマンドの引数を補完できない](https://github.com/vim-jp/issues/issues/243)
- [Issue #244 : ACLが保存・復帰されない (Cygwinの属性の保存)](https://github.com/vim-jp/issues/issues/244)
- [Issue #245 : getbufvar(), getwinvar(), gettabvar()にデフォルト値を指定できるようにしたい](https://github.com/vim-jp/issues/issues/245)

