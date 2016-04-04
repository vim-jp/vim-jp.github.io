---
layout: post
category: vimmagazine
title: Vim Magazine 6 月号
---

## 話題

- [TokyoVim#8](http://partake.in/events/538c8c05-dab3-4169-8c9c-73e8375d7087)

  Vimに関する何かをやる会 東京 7月7日

- [第五回 Sugamo.vim](http://atnd.org/event/sugamovim05)

  Vimについてｇｄｇｄ語ろうの会 東京 7月15日


## リリース情報

- [7.3.532 : compiler warning from Clang](https://github.com/vim/vim/commit/112f318551e342f19b6af18d1661f65138b41ba8)
- [7.3.533 : memory leak when writing undo file](https://github.com/vim/vim/commit/d2aed44c77f165a23c495e5836bc33d3133fea2e)
- [7.3.534 : (after 7.3.461) autoindent fails with InsertCharPre autocmd](https://github.com/vim/vim/commit/704984ac87979fd20b8ba732df3abd3197814f7f)
- [7.3.535 : many #ifdefs for MB\_MAXBYTES](https://github.com/vim/vim/commit/9a920d8c31bc4d267d8790ba0d83b5774316e08f)
- [7.3.536 : German sharp s is not seen as a word character](https://github.com/vim/vim/commit/88178de99f9cefb92450ff796e541926f3533cec)
- [7.3.537 : unecessary call to init\_spell\_chartab()](https://github.com/vim/vim/commit/24ee83b0a0613c635cb6ce1292891c7fa65510cc)
- [7.3.538 : 'efm' does not handle Tabs in pointer lines](https://github.com/vim/vim/commit/f13de07e492fc0a4a0af12434f3d282f9c989ed4)
- [7.3.539 : redraw multi-byte char on command line does not work properly](https://github.com/vim/vim/commit/64fdf5ceae75d85e139a0d3812e0544f5bf6cb56)
- [7.3.540 : cursor is left on the text instead of the command line](https://github.com/vim/vim/commit/bc256d91eadb8f32d6a5833c1878684d3c75fb2d)
- [7.3.541 : when joining lines comment leaders need to be removed manually](https://github.com/vim/vim/commit/8134039744ffa581f7c217df58131b709317c1c8)
- [7.3.542 : (after 7.3.506) function is sometimes unused](https://github.com/vim/vim/commit/08bc274e88b32848b592d34600d51c8784357f3e)
- [7.3.543 : the cursor is in the wrong line after using ":copen"](https://github.com/vim/vim/commit/ab984db2969c3f522eb4da694a954cbe6cec7797)
- [7.3.544 : no autocommand for :quit before deciding to exit](https://github.com/vim/vim/commit/3b53dfb3b0743af7d6ae381a766e1bb2018fd01e)
- [7.3.545 : autocommands may close a window that is already being closed](https://github.com/vim/vim/commit/362ce4804819f39d5a4a21923577f3ccc59c8ad5)
- [7.3.546 : weird line break](https://github.com/vim/vim/commit/3b393a0b537cd7057d3462ef16685b3957eb8e1a)
- [7.3.547 : (after 7.3.541) compiler warning for uninitialized variable](https://github.com/vim/vim/commit/802053f14ac4da63c027df2f1a8333c9a26da7dc)
- [7.3.548 : compiler warning on 64 bit Windows](https://github.com/vim/vim/commit/27ba0885493243945b16e14056c0a98a3ab14021)
- [7.3.549 : in 'cinoptions' "0s" is interpreted as one shiftwidth](https://github.com/vim/vim/commit/48d279215f08667ccfa243afe874ef93e228dd57)
- [7.3.550 : (after 7.3.541) with "j" in 'fo' a list leader is not removed](https://github.com/vim/vim/commit/e04a48f20413f3f926d26394fad6431795348af7)
- [7.3.551 : on :tablose a TabEnter autocommand is triggered too early](https://github.com/vim/vim/commit/a8596c47724b97822924f5ffe5d50476de31ff4b)
- [7.3.552 : inside comments formatting does not use the "2" flag in 'fo'](https://github.com/vim/vim/commit/bfe3bf806a8a4300289055643d13d19e2dbc8967)
- [7.3.553 : text displayed one cell off if 'listchars' contains "precedes"](https://github.com/vim/vim/commit/5641f38d41068e92c2ba15cb6359f8d80188f6cf)
- [7.3.554 : compiler warning for unused argument](https://github.com/vim/vim/commit/795ec43112af40e02ddc8fcc1d0a7800a33f2a6e)
- [7.3.555 : building on IBM z/OS fails](https://github.com/vim/vim/commit/77c193579b39bbeacd4ffa4a16b5f3cd00c39cee)
- [7.3.556 : compiler warnings on 64 bit Windows](https://github.com/vim/vim/commit/36105782d2dfe526a0ca361d1c6c06e6f7a9ca6f)
- [7.3.557 : crash when an autocommand wipes out a buffer when it is hidden](https://github.com/vim/vim/commit/9e931224db8f8843ddc778307984dbe37ebde47c)
- [7.3.558 : (after 7.3.552) memory access error](https://github.com/vim/vim/commit/dc7e85ee5d29bd65681ff61b37d8e8cd38c22330)
- [7.3.559 : home\_replace() does not work with 8.3 filename](https://github.com/vim/vim/commit/9158f9e42355bfa1d956cb75cdaef37b4ab80f35)
- [7.3.560 : get an error for a locked argument in extend()](https://github.com/vim/vim/commit/ed46560bf0673ab9347e9880bc6e26ce547abff1)
- [7.3.561 : refresh: always in a complete function breaks the "." command](https://github.com/vim/vim/commit/5e1a0a9a65b221d376a3ab85d06d918e651dd123)
- [7.3.562 : ":profdel" works when the +profile feature is disabled](https://github.com/vim/vim/commit/38bdbd6c6f39e5bffad22ec956fa1383e04fa1ae)
- [7.3.563 : (after 7.3.557) can't build with tiny features](https://github.com/vim/vim/commit/50a12b40783550fb8fd7a6283815b107863292ab)
- [7.3.564 : (after 7.3.559) warning for pointer conversion](https://github.com/vim/vim/commit/d12f811816989fe969401a8f3b7d286580653a23)
- [7.3.565 : can't generate proto file for Python 3](https://github.com/vim/vim/commit/394315603402fe12b5339d36d35c8e1a2796f404)
- [7.3.566 : (after 7.3.561) redo works incorrectly without refresh:always](https://github.com/vim/vim/commit/22189a4bd659c2b9f6766cd45bacd11915d05296)
- [7.3.567 : missing copyright notice](https://github.com/vim/vim/commit/002850940e84c399dcbf2082f08cae36cc44e7b1)
- [7.3.568 : bad indents for #ifdefs](https://github.com/vim/vim/commit/a7014df97532a4171276aa7e3b878e80e88e513c)
- [7.3.569 : evaluating Vim expression in Python is insufficient](https://github.com/vim/vim/commit/db91395312a02527ed973c8376d8e26e5b63ff53)
- [7.3.570 : ":vimgrep" does not obey 'wildignore'](https://github.com/vim/vim/commit/8f5c6f003ade1d481c4db0f2e2fffac209dce77c)
- [7.3.571 : duplicated condition](https://github.com/vim/vim/commit/f6a2b08c5444712ddb106547695cc4d59652ce7a)
- [7.3.572 : duplicate statement in if and else](https://github.com/vim/vim/commit/a53c60d33c917455857de96a44e109da7d305d8a)
- [7.3.573 : using array index before bounds checking](https://github.com/vim/vim/commit/88b1ba151ae7bb66dc92c85f1eb63fe7fcc03051)
- [7.3.574 : a CTRL-L character is not pasted on the search command line](https://github.com/vim/vim/commit/e79abddb2fbc306beb61204432c92f3739bafd2c)
- [7.3.575 : "ygt" tries to yank instead of giving an error](https://github.com/vim/vim/commit/89f940fcacf0b0996c05537034b37b8c6bb87cf5)
- [7.3.576 : formatting of lists inside comments is not right yet](https://github.com/vim/vim/commit/96b7ca5142d72d2c8e79d15f9c38f41d6657be40)
- [7.3.577 : size of memory does not fit in 32 bit unsigned](https://github.com/vim/vim/commit/11b73d668f6209acbe06988a804028c5c86799d2)
- [7.3.578 : misplaced declaration.](https://github.com/vim/vim/commit/f27839c1a9e7390bb850d768e1146014b221cc43)
- [7.3.579 : (after 7.3.569) can't compile with Python 2.5](https://github.com/vim/vim/commit/2afa3238a13fe3f4769fa777ada34ec0d1ea5548)
- [7.3.580 : warning on 64 bit MS-Windows](https://github.com/vim/vim/commit/3cd3e7ab17535f48969b7e5467d5bbdebbecebed)
- [7.3.581 : problems compiling with Python](https://github.com/vim/vim/commit/cc3e85f169ba44a6ac4b66c6fbb1774f7f7ae0db)
- [7.3.582 : missing pieces in test OK file](https://github.com/vim/vim/commit/49346f4155b01a3d73f575f49359dd108eda939f)

## 新着スクリプト

- [css3 mod : a css syntax script to support css3](http://www.vim.org/scripts/script.php?script_id=4089)
- [Difference mode color settings : plugin for setting colors in difference mode](http://www.vim.org/scripts/script.php?script_id=4091)
- [Vitra : Trac UI for vim](http://www.vim.org/scripts/script.php?script_id=4092)
- [shuffle.vim : Do the random shuffle of lines in vim's buffer](http://www.vim.org/scripts/script.php?script_id=4093)
- [tips.vim : Show vim tips everytime when you enter vim](http://www.vim.org/scripts/script.php?script_id=4094)
- [vim-django : Support for Django through VIM](http://www.vim.org/scripts/script.php?script_id=4095)
- [maven-ide : maven-ide plugin](http://www.vim.org/scripts/script.php?script_id=4096)
- [fisa-color-scheme : soft color scheme for terminals with 256 colors](http://www.vim.org/scripts/script.php?script_id=4097)
- [sharefix : Share quickfixes among commands and functions](http://www.vim.org/scripts/script.php?script_id=4098)
- [oh-là-là : High contrast dark color scheme with black background](http://www.vim.org/scripts/script.php?script_id=4099)
- [abc-vim : abc music notation syntax, 2.1 and up](http://www.vim.org/scripts/script.php?script_id=4100)
- [PEGJS.vim : syntax highlighting for PEG.js grammar files](http://www.vim.org/scripts/script.php?script_id=4101)
- [vim-brunch : Easy file navigation for brunch projects](http://www.vim.org/scripts/script.php?script_id=4102)
- [GdbFromVim : This pluging lets you debug your applications with gdb from vim,](http://www.vim.org/scripts/script.php?script_id=4103)
- [gdbmgr : vim interface to gdb](http://www.vim.org/scripts/script.php?script_id=4104)
- [pysource.vim : View python source easily](http://www.vim.org/scripts/script.php?script_id=4105)
- [fish-syntax : fish syntax based on sh](http://www.vim.org/scripts/script.php?script_id=4106)
- [autumnleaf\_modified.vim : a light, simple and beautiful color scheme for programming and good for your eye](http://www.vim.org/scripts/script.php?script_id=4109)
- [sml\_polyml.vim : Poly/ML integration](http://www.vim.org/scripts/script.php?script_id=4110)
- [log4j.vim : Adds syntax highlighting for log4j and log4php log files](http://www.vim.org/scripts/script.php?script_id=4111)
- [riv.vim : managing and writing reStructuredText Documents.](http://www.vim.org/scripts/script.php?script_id=4112)
- [hickop : simple dark colorscheme for vim and gvim](http://www.vim.org/scripts/script.php?script_id=4113)
- [LineMotion : a Vim Plug-in, move fast in linei of text.](http://www.vim.org/scripts/script.php?script_id=4114)
- [BufferPersist : Save certain buffers somewhere when quitting them.](http://www.vim.org/scripts/script.php?script_id=4115)
- [MessageRecall : Browse and re-insert previous (commit, status) messages.](http://www.vim.org/scripts/script.php?script_id=4116)
- [VcsMessageRecall : Browse and re-insert previous VCS commit messages.](http://www.vim.org/scripts/script.php?script_id=4117)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2847)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2427)
3. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (1345)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1265)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1173)
6. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (1074)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (1072)
8. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1065)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (1030)
10. [rails.vim : Ruby on Rails: easy file navigation, enhanced syntax highlighting, and more](http://www.vim.org/scripts/script.php?script_id=1567) (998)

## vim-jp issues

Open : 119 (+2) | Closed : 95 (+9)

- [Issue #204 : コマンドラインで日本語が化ける](https://github.com/vim-jp/issues/issues/204)
- [Issue #205 : MSVCなvimで入力した日本語が化ける](https://github.com/vim-jp/issues/issues/205)
- [Issue #206 : :write >>でファイルを追記するとき追記元と追記先のエンコーディングが異なると日本語(2バイト文字)が文字化けする](https://github.com/vim-jp/issues/issues/206)
- [Issue #207 : \`=...\` が 'wildignore' を参照する](https://github.com/vim-jp/issues/issues/207)
- [Issue #208 : 補完時にfindstart=0の場合にfindstart=1で返した位置からカーソル位置までがgetline(".")から消える。](https://github.com/vim-jp/issues/issues/208)
- [Issue #209 : vim-cpp を提案する上での配布形態について](https://github.com/vim-jp/issues/issues/209)
- [Issue #210 : マップされたキーのみでコマンドを実行すると履歴に残らない](https://github.com/vim-jp/issues/issues/210)
- [Issue #211 : set tags でワイルドカードを設定しPHP のオムニ補完を実行すると QuickFix ウインドウが開く](https://github.com/vim-jp/issues/issues/211)
- [Issue #212 : tags に /path/to/\*\*/tags; を使うとメモリ枯渇になる件の検討](https://github.com/vim-jp/issues/issues/212)
- [Issue #213 : Windowsで$TMPにスペースが含まれているとsystem()が一時ファイルを開けない](https://github.com/vim-jp/issues/issues/213)
- [Issue #214 : linuxでのmake testがパスしない](https://github.com/vim-jp/issues/issues/214)

