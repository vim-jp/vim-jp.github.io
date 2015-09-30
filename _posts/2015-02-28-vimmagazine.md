---
layout: post
category: vimmagazine
title: Vim Magazine 2 月号
---

## 話題

- [TokyoVim#24](http://tokyovim.connpass.com/event/12389/)

  Vimに関する何かをやる会です。 3月21日(土) 東京

## 今月の新機能

- 7.4.627: "xn" termcap 機能がサポートされました。

## リリース情報

- [7.4.609 : the garbage collector can run out of stack space](https://github.com/vim/vim/commit/2459a5ecaa43c8549ea53e9364253ff891676da5)
- [7.4.610 : some function headers may be missing from generated .pro files](https://github.com/vim/vim/commit/016154919e317a8ce2b7f2e33db2f39ca3a8edb9)
- [7.4.611 : (after 7.4.609) syntax error](https://github.com/vim/vim/commit/b641df49655bb34762b44aa8c6e03e5f464f73d6)
- [7.4.612 : test&#x5f;eval fails on Mac](https://github.com/vim/vim/commit/e08dd4e49e6d67686e3ba1322a641488ad67c711)
- [7.4.613 : the NFA engine does not implement the 'redrawtime' time limit](https://github.com/vim/vim/commit/70781ee4035b5fd5e3cbb3fe4c7646e19119f0a8)
- [7.4.614 : there is no test for what patch 7.4.601 fixes](https://github.com/vim/vim/commit/d7ce7a9ad2d9311f7ec3368eeddec8fce6d8e890)
- [7.4.615 : Vim hangs when freeing a lot of objects](https://github.com/vim/vim/commit/e71eea801ec779e2d19f4b486313b59de2137a8a)
- [7.4.616 : cannot insert a tab in front of a block](https://github.com/vim/vim/commit/f2c03d7301d35590a20cc43431950acc3a2f6036)
- [7.4.617 : wrong ":argdo" range does not cause an error](https://github.com/vim/vim/commit/c0a37b9db38e722b9e5ac3ac8b4c4af86bdf018c)
- [7.4.618 : (after 7.4.609) luaV&#x5f;setref() is missing a return statement](https://github.com/vim/vim/commit/9d49da8cbbd46d528be38ec3ac394a2ce9062093)
- [7.4.619 : (after 7.4.618) luaV&#x5f;setref() not returning the correct value](https://github.com/vim/vim/commit/b84634de30884f4f6a17b5872f4553ec746797c5)
- [7.4.620 : compiler warning for unitinialized variable](https://github.com/vim/vim/commit/934b13619318c5d0721797266b9a9060546e07a5)
- [7.4.621 : (after 7.4.619) returning 1 in the wrong function](https://github.com/vim/vim/commit/f554a3285e45d8c3f95f9b8f43bd0897185b95cb)
- [7.4.622 : compiler warning for unused argument](https://github.com/vim/vim/commit/168e04966c48b0d7d4726934ed2c39a9a0fec2af)
- [7.4.623 : crash with pattern: \\(\\)\\{80000}](https://github.com/vim/vim/commit/a1d2c58985584116d20fa5e132137d8ff1a535f7)
- [7.4.624 : may leak memory or crash when vim&#x5f;realloc() returns NULL](https://github.com/vim/vim/commit/9abd5c6507154eabdfe8256940a24f090db0f533)
- [7.4.625 : possible NULL pointer dereference](https://github.com/vim/vim/commit/24a6ff88bc8710c305deba405d00061dec909125)
- [7.4.626 : MSVC with W4 gives useless warnings](https://github.com/vim/vim/commit/eb2928595bca45cb37f0906c76896dce806b5543)
- [7.4.627 : the last screen cell is not updated](https://github.com/vim/vim/commit/494838a3fec2fb1abcac6256271810a762225bed)
- [7.4.628 : compiler warning for variable might be clobbered by longjmp](https://github.com/vim/vim/commit/8872ef1b4e1b304cc168ea4c8f68645498c155be)
- [7.4.629 : Coverity warning for Out-of-bounds read](https://github.com/vim/vim/commit/98f52500b6bca75539f50ea31cb333223141dfeb)
- [7.4.630 : redo is wrong for insert mode completion with autocommands](https://github.com/vim/vim/commit/c51b02d7d9c75c241eb6435d410422d0fa9147a1)
- [7.4.631 : the default conceal character is a dash instead of a space](https://github.com/vim/vim/commit/4a427106954857c6f84c0b0dc07725e302a433c6)
- [7.4.632 : (after 7.4.592) breaks the netrw plugin](https://github.com/vim/vim/commit/eaa330e8e913823f1514bc419f23b885d6028cb9)
- [7.4.633 : after 7.4.630 the problem persists](https://github.com/vim/vim/commit/be20f9fa5da001258a40a5f01e38befac4ebfe80)
- [7.4.634 : marks are not restored after redo + undo](https://github.com/vim/vim/commit/f65aad5554f1d1f972f0bd38059da70e3cdf4c87)
- [7.4.635 : fileformat set to "mac" if a file starts with a very long line](https://github.com/vim/vim/commit/05eb612ff3597fb7102f892bfd320f28b56d2fc6)
- [7.4.636 : a search with end offset gets stuck at end of file](https://github.com/vim/vim/commit/4653911950e8bedff444adc9a2dcbebd9b36dc07)
- [7.4.637 : buffer number for autocommand is wrong](https://github.com/vim/vim/commit/1e997827f12b38359fe8a4e7b033bbb6465b6668)
- [7.4.638 : can't build with Lua 5.3 on Windows](https://github.com/vim/vim/commit/46538ee46cb04ed47c546d92583d6eb023e51b32)
- [7.4.639 : combination of linebreak and conceal doesn't work well](https://github.com/vim/vim/commit/8fc6bc71266e342d339c851b3ee472357f917597)
- [7.4.640 : after joining lines in Insert mode undo does not work properly](https://github.com/vim/vim/commit/c3bbad085c3ec304b111dc95aed97fea4c38a177)
- [7.4.641 : the tabline menu was using ":999tabnew" which is now invalid](https://github.com/vim/vim/commit/dfd7691bb85b345bd86cf30945a66acf7c782920)
- [7.4.642 : when using "gf" escaped spaces are not handled](https://github.com/vim/vim/commit/d45c07ac7499358c5cb096cadb675ce74ae3eaf6)
- [7.4.643 : using the default file format for Mac files (Issue 77)](https://github.com/vim/vim/commit/c6b7217ff502b8fc28d6c861d25c1943e30973ad)
- [7.4.644 : Stratus VOS doesn't have sync()](https://github.com/vim/vim/commit/ba17ed6797421b377ec0cac7ec186d24bf5b3979)
- [7.4.645 : window count wrong when splitting window in BufAdd autocommand](https://github.com/vim/vim/commit/8da9bbfd02957b79edd595c8c7397453012510b0)
- [7.4.646 : ":bufdo" may start at a deleted buffer](https://github.com/vim/vim/commit/e25bb90b2eb0ceed2caa5202ca62003e367021ae)
- [7.4.647 : files differ after running the tests on MS-Windows](https://github.com/vim/vim/commit/13cf77ba1715803355029ed119e55a12f6865db1)
- [7.4.648 : (after 7.4.647) tests broken on MS-Windows](https://github.com/vim/vim/commit/bac203ea2d0e3b2a5e4eae4dbe48244e592926e2)

## 新着スクリプト

- [enabler : Load plugins/bundles as needed](http://www.vim.org/scripts/script.php?script_id=5101)
- [vital-snoop.vim : Snoop script-local functions](http://www.vim.org/scripts/script.php?script_id=5102)
- [haystack.vim : Fuzzy matching algorithm](http://www.vim.org/scripts/script.php?script_id=5103)
- [FastFold : Speed up Vim by updating folds only when just](http://www.vim.org/scripts/script.php?script_id=5104)
- [h2cppx : parse the c++ header file and generate c++ implement code ](http://www.vim.org/scripts/script.php?script_id=5105)
- [Indentwise : Indent-level based motion](http://www.vim.org/scripts/script.php?script_id=5106)
- [vim-conda : Conda environment manager](http://www.vim.org/scripts/script.php?script_id=5107)
- [vim-hackernews : Browse Hacker News in Vim](http://www.vim.org/scripts/script.php?script_id=5108)
- [delview : 'delview' command to delete mkview/loadview files](http://www.vim.org/scripts/script.php?script_id=5109)
- [pl65.vim : Syntax file for Noahsoft&#039s PL65 for Atari 8-bit computers](http://www.vim.org/scripts/script.php?script_id=5110)
- [glsl&#x5f;es&#x5f;100.vim : Syntax file for GLSL ES 100](http://www.vim.org/scripts/script.php?script_id=5111)
- [Perfect Dark : Dark & cold color scheme inspired by gedit](http://www.vim.org/scripts/script.php?script_id=5112)
- [sonicpi.vim : Syntax highlighting and autocomplete for Sonic Pi / Ruby files](http://www.vim.org/scripts/script.php?script_id=5113)
- [vim-webdevicons : vim web developer filetype font unicode icons for NERDTree and vim-airline](http://www.vim.org/scripts/script.php?script_id=5114)
- [vim-fixprotocol : parse a FIX protocol message into xml](http://www.vim.org/scripts/script.php?script_id=5115)
- [vim-consul : basic vim/consul integration](http://www.vim.org/scripts/script.php?script_id=5116)
- [vim-stardict : Looking up meaning of words inside Vim using SDCV dictionary](http://www.vim.org/scripts/script.php?script_id=5117)
- [AppendToClip : plugin for appending to the default register](http://www.vim.org/scripts/script.php?script_id=5118)
- [WholeLineColor : plugin for changing background color of lines](http://www.vim.org/scripts/script.php?script_id=5119)
- [SpecialLocationComplete : Insert mode completion for special custom patterns.](http://www.vim.org/scripts/script.php?script_id=5120)
- [EnhancedDiff : A Vim plugin for creating better diffs (sometimes)](http://www.vim.org/scripts/script.php?script_id=5121)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (5270)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1450)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1318)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (950)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (668)
6. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (607)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (580)
8. [vim-misc : Miscellaneous auto-load Vim scripts](http://www.vim.org/scripts/script.php?script_id=4597) (526)
9. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (430)
10. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (429)

## vim-jp issues

Open : 219 (+2) | Closed : 486 (+9)

- [Issue #695 : 空白を表示(list)しているとwrapが効かない](https://github.com/vim-jp/issues/issues/695)
- [Issue #696 : ヘルプの「'g' で始まるコマンド」の一覧に含まれていないコマンドがある](https://github.com/vim-jp/issues/issues/696)
- [Issue #697 : --term=mlterm オプション付きで起動した mlterm で Vim の背景が乱れる](https://github.com/vim-jp/issues/issues/697)
- [Issue #698 : 7.4.615 以降のGCの挙動を検証する](https://github.com/vim-jp/issues/issues/698)
- [Issue #699 : プログラマーがじゃぶじゃぶコード書きたく鳴るような射幸心を煽りまくるvimの説明文](https://github.com/vim-jp/issues/issues/699)
- [Issue #700 : VIM::evaluate の説明が古い](https://github.com/vim-jp/issues/issues/700)
- [Issue #701 : 'cinoptions'オプションに &#x60;J1&#x60; が指定されているとインデントがおかしくなる](https://github.com/vim-jp/issues/issues/701)
- [Issue #702 : cygwin上でvimコマンドを入力しても何も起こらない](https://github.com/vim-jp/issues/issues/702)
- [Issue #703 : mintty 上でも native vim.exe を動かせる様にしたい](https://github.com/vim-jp/issues/issues/703)
- [Issue #704 : マルチバイトを含む行に仮想桁のシンタックスハイライトがマッチしない](https://github.com/vim-jp/issues/issues/704)
- [Issue #705 : vimrun.exe がないと system() を実行できない](https://github.com/vim-jp/issues/issues/705)

