---
layout: post
category: vimmagazine
title: Vim Magazine 1 月号
---

## 話題

- [Nagoya.vim #3](http://nagoyavim.connpass.com/event/11007/)

  Vim プラグインで遊ぶ会  2月7日(土) 名古屋

## 今月の新機能

- 7.4.566: :argdo、:bufdo、:windo、:tabdo で範囲指定がサポートされました。
- 7.4.601: feedkeys() に i フラグが追加されました。
- 7.4.605: # レジスタが書き込み可能になりました。

## リリース情報

- [7.4.561 : Ex range handling is wrong for buffer-local user commands](https://github.com/vim/vim/commit/abe382cd2ac032827d7ef286f3c19b21e85f30d4)
- [7.4.562 : segfault with wide screen and error in 'rulerformat'](https://github.com/vim/vim/commit/0027c218e9ba174e2884e9361ab9183c4d003dd7)
- [7.4.563 : no test for replacing on a tab in Virtual replace mode](https://github.com/vim/vim/commit/684d409f63eb4892e6f5d28acb85ac84213e2bab)
- [7.4.564 : FEAT&#x5f;OSFILETYPE is used even though it's never defined](https://github.com/vim/vim/commit/49a6ed8a8a5fa471bf99f8d068e7cb994c2af163)
- [7.4.565 : some ranges are not checked to be valid](https://github.com/vim/vim/commit/3ffc79a4a82918430940bfaa18f4da058fdbd0ca)
- [7.4.566 : :argdo, :bufdo, :windo and :tabdo don't take a range](https://github.com/vim/vim/commit/a162bc555e74c1a4f1b7b49871e292e05141b054)
- [7.4.567 : non-ascii vertical separater characters are always redrawn](https://github.com/vim/vim/commit/c60c4f6e06e222e90b2277c09fdaaac20ac9edf8)
- [7.4.568 : giving an error for ":0wincmd w" is a problem for some plugins](https://github.com/vim/vim/commit/8be6388b7649d9378cd1ba1627a4b0aed61b86e7)
- [7.4.569 : having CTRL-C interrupt does not check the mapping mode](https://github.com/vim/vim/commit/651863c94a882a97aec7968fc87a638ff78e56ff)
- [7.4.570 : building with dynamic library does not work for Ruby 2.2.0](https://github.com/vim/vim/commit/0c7485fdbb5023731b006572cc224cbf52e34288)
- [7.4.571 : can't build with tiny features](https://github.com/vim/vim/commit/435797304adc544562934f3ea1572fdc81f6dd92)
- [7.4.572 : address type of :wincmd depends on the argument](https://github.com/vim/vim/commit/84c8e5ab9cfb6826e880e8ae062ff3347bb6909e)
- [7.4.573 : mapping CTRL-C in Visual mode doesn't work](https://github.com/vim/vim/commit/5000869712f799d9ca25c0e45dc21d332edae5f4)
- [7.4.574 : no error for eval('$')](https://github.com/vim/vim/commit/615b9978dd32345ccb9f9d2ae0fc2b451ab62ea4)
- [7.4.575 : Unicode character properties are outdated](https://github.com/vim/vim/commit/ea67672534ae00cfd560c056340ed4da001e95f9)
- [7.4.576 : redrawing problem with 'relativenumber' and 'linebreak'](https://github.com/vim/vim/commit/ba3f58e29660d4f7501cb0e46162a399c20020c6)
- [7.4.577 : matching with virtual column has overhead on very long lines](https://github.com/vim/vim/commit/a20bcad15c826a8c43a0f66a6b438bb5a14c8196)
- [7.4.578 : after "$" in an empty line getcurpos() returns negative number](https://github.com/vim/vim/commit/084abaec061f830e2d389c416499d013e47fb0db)
- [7.4.579 : wrong cursor positioning when 'linebreak' set and lines wrap](https://github.com/vim/vim/commit/d574ea2f80b271aff4bac7dc7fa83b292f0ac486)
- [7.4.580 : ":52wincmd v" still gives an invalid range error](https://github.com/vim/vim/commit/164f3269b13080ab39495440d614efa29bd81fe5)
- [7.4.581 : compiler warnings for unitinialized variables](https://github.com/vim/vim/commit/7cbc703fdc7be3b5c0c83afe8362885b920aa132)
- [7.4.582 : (after 7.4.577) can't match "%>80v" properly](https://github.com/vim/vim/commit/ef795d1a5bb39a7cc74e7e36075a47afbfa75901)
- [7.4.583 : with tiny features test 16 may fail](https://github.com/vim/vim/commit/ff13610d7b40bee4ae58c3ce3e3b4ca7012a6c4f)
- [7.4.584 : with tiny features test&#x5f;command&#x5f;count may fail](https://github.com/vim/vim/commit/ff8a75dab738a2d1e81202b5dd1c03fc45d714f2)
- [7.4.585 : range for :bdelete does not work](https://github.com/vim/vim/commit/e823183254e79af4d384abd99d5bde9b013c4ec5)
- [7.4.586 : parallel building of documentation html files is not reliable](https://github.com/vim/vim/commit/6bf7a5933f50db03455104a018f5de1efd82b7fb)
- [7.4.587 : conceal does not work properly with 'linebreak'](https://github.com/vim/vim/commit/6a6028cf4b0a2c48e34f657d478adee814d471c5)
- [7.4.588 : ":0argedit foo" puts the new argument in the second place](https://github.com/vim/vim/commit/4bde00c59575f7547ab59946aea527c5be067662)
- [7.4.589 : MS-Windows console: Vim can't handle greek utf-8 characters](https://github.com/vim/vim/commit/1ec4dd41de276260cd1d934b8033804f2d2f0435)
- [7.4.590 : using ctrl&#x5f;x&#x5f;mode as if it contains flags](https://github.com/vim/vim/commit/eea478f885ab2fb4185053f1093d05b83bbdf5e1)
- [7.4.591 : test&#x5f;listlbr&#x5f;utf8 fails when conceal feature is not available](https://github.com/vim/vim/commit/df82661c9cf24a8947584197c756a88add5406ca)
- [7.4.592 : buffer with 'buftype' "nofile" is cleared on re-edit](https://github.com/vim/vim/commit/df5caa08f62fe67f11e48771f4a5125ebc5a69b3)
- [7.4.593 : crash when searching for "x\\{0,90000}"](https://github.com/vim/vim/commit/e0ad365498399c1bd34dd6361b3f7dc38e84e4ca)
- [7.4.594 : using block delete with 'breakindent' does not work properly](https://github.com/vim/vim/commit/1dc92334eb5e235af338f3b27277157839fa8dcc)
- [7.4.595 : the test&#x5f;command&#x5f;count test fails when using Japanese](https://github.com/vim/vim/commit/322a24af83a322b616da01969b37395df9ab3aa3)
- [7.4.596 : (after 7.4.592) tiny build doesn't compile](https://github.com/vim/vim/commit/932d88662f0d6636586aa3a321502609ebad73b5)
- [7.4.597 : cannot change the result of systemlist()](https://github.com/vim/vim/commit/9a492d456d15fe637f197fe872b76c164428fd3e)
- [7.4.598 : ":tabdo windo echo 'hi'" causes "&#x2a; register not to be changed](https://github.com/vim/vim/commit/5c27fd100aa5ec587e9d6b2a4a751172b890c837)
- [7.4.599 : out-of-memory error for complicated search pattern](https://github.com/vim/vim/commit/aaf3047de2123ae87811e947ddbaee97593fad57)
- [7.4.600 : memory wasted in struct because of aligning](https://github.com/vim/vim/commit/0cd040b96a14363de6928b4bc0a10d628ab40a52)
- [7.4.601 : it is not possible to have feedkeys() insert characters](https://github.com/vim/vim/commit/0a988df731605d692892f18317c26e6ab24d3d82)
- [7.4.602 : ":set" does not accept hex numbers as documented](https://github.com/vim/vim/commit/18400e6430755c97eb7e31525ab4744502603731)
- [7.4.603 : 'foldcolumn' may fill the whole window](https://github.com/vim/vim/commit/1c93429c48c59e22aa8e9cd76f102063799db938)
- [7.4.604 : running tests changes viminfo](https://github.com/vim/vim/commit/6bf7c523ad54f625547ef2cc6009c28c0dd337ed)
- [7.4.605 : the # register cannot be restored after jumping around](https://github.com/vim/vim/commit/3b3a9498d1eab3c28c524cce115160528a9a9297)
- [7.4.606 : may crash when using a small window](https://github.com/vim/vim/commit/7cc8ec47208f784418e40d1ede49c4aeb95b0b53)
- [7.4.607 : (after 7.4.598) compiler warnings for unused variables](https://github.com/vim/vim/commit/73a156bf362caed3b42f834e90e4409cf4d1df9b)
- [7.4.608 : (after 7.4.598) test&#x5f;eval fails without the clipboard feature](https://github.com/vim/vim/commit/4ac163ae5f137af236931e778660cf7878d70c25)

## 新着スクリプト

- [vim-abaqus : VIM filetype plugin for Abaqus](http://www.vim.org/scripts/script.php?script_id=5077)
- [permut.vim : Plugin for swapping text columns separated by arbitrary character.](http://www.vim.org/scripts/script.php?script_id=5078)
- [Darcula Color Scheme : Dark color scheme that is easy on the eyes](http://www.vim.org/scripts/script.php?script_id=5079)
- [MacDict : Search your system dictionary from within Vim](http://www.vim.org/scripts/script.php?script_id=5080)
- [Vim SQL Workbench : Provides sql access from vim to any DBMS](http://www.vim.org/scripts/script.php?script_id=5081)
- [vital-string-interpolation : String interpolation library for vim](http://www.vim.org/scripts/script.php?script_id=5082)
- [yowish : A yellowish dark colorscheme](http://www.vim.org/scripts/script.php?script_id=5083)
- [Falmouth : Falmouth is a pastel color scheme that is easy on the eyes.](http://www.vim.org/scripts/script.php?script_id=5084)
- [Boston : a key mapping plugin for commonly used coding shortcuts](http://www.vim.org/scripts/script.php?script_id=5085)
- [DidYouMean : Ask for the right file to open](http://www.vim.org/scripts/script.php?script_id=5086)
- [vim-pandoc-bundle : pandoc markdown support for vim + utilities](http://www.vim.org/scripts/script.php?script_id=5088)
- [vim-fetch : Make Vim handle line and column numbers in file names with a minimum of fuss](http://www.vim.org/scripts/script.php?script_id=5089)
- [vim-lastplace : Intelligently reopen files where you left off. But not for commit messages.](http://www.vim.org/scripts/script.php?script_id=5090)
- [perline-diff.vim : this line, that line, then diff comes up](http://www.vim.org/scripts/script.php?script_id=5091)
- [screensaver.vim : Screensavers for Vim](http://www.vim.org/scripts/script.php?script_id=5092)
- [readdir : A very minimal directory browser](http://www.vim.org/scripts/script.php?script_id=5093)
- [partial : Easily extract partials from your code](http://www.vim.org/scripts/script.php?script_id=5094)
- [codekit.vim : CodeKit 2 support for Vim](http://www.vim.org/scripts/script.php?script_id=5095)
- [Space Ident Highlighter : Plugin to mark vertically identation made by spaces](http://www.vim.org/scripts/script.php?script_id=5096)
- [vim-less : LESS css syntax updated to v2.3.0 - 2015](http://www.vim.org/scripts/script.php?script_id=5097)
- [cscope&#x5f;dynamic : quick automatic dynamic cscope updates](http://www.vim.org/scripts/script.php?script_id=5098)
- [vim-stay : Make Vim persist editing state without fuss](http://www.vim.org/scripts/script.php?script_id=5099)
- [vim-cursorword : Underlines the word under the cursor](http://www.vim.org/scripts/script.php?script_id=5100)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (7340)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1864)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1653)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (990)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (757)
6. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (684)
7. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (654)
8. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (634)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (581)
10. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (557)

## vim-jp issues

Open : 217 (-3) | Closed : 477 (+16)

- [Issue #682 : complete() を用いて補完を行うと \<C-l> で補完候補を選択してしまう](https://github.com/vim-jp/issues/issues/682)
- [Issue #683 : complete()による先頭マッチしない補完を連続で行った場合、ドットレジスタの値がおかしくなる](https://github.com/vim-jp/issues/issues/683)
- [Issue #684 : VisualモードでYとyyの挙動が違う](https://github.com/vim-jp/issues/issues/684)
- [Issue #685 : complete()による補完時に\<C-h>を押すと行補完になる](https://github.com/vim-jp/issues/issues/685)
- [Issue #686 : Ruby-2.2.0 で +ruby/dyn のビルドに失敗する](https://github.com/vim-jp/issues/issues/686)
- [Issue #687 : eval('$') は 0 でなくエラーになるべきではないか](https://github.com/vim-jp/issues/issues/687)
- [Issue #688 : range('a', 'z') で \['a', 'b', 'c', .... 'z'\] になってほしい](https://github.com/vim-jp/issues/issues/688)
- [Issue #689 : &#x60;set breakindent&#x60;かつ折り返し行の先頭が半角空白の場合、行をまたいだ矩形選択での削除でずれて削除される](https://github.com/vim-jp/issues/issues/689)
- [Issue #690 : システムエンコードが英語以外の環境でmake testを実施するとtest&#x5f;command&#x5f;countが失敗する](https://github.com/vim-jp/issues/issues/690)
- [Issue #691 : &#x60;set breakindent&#x60; 時にインデントのある行とない行を矩形選択するとハイライトがおかしい](https://github.com/vim-jp/issues/issues/691)
- [Issue #692 : complete() で補完を行うと自動折り返しが無効になる](https://github.com/vim-jp/issues/issues/692)
- [Issue #693 : 日本語と英語の境界の扱いが、モーションのwordとインサートモードでの &#x60;\<C-w>&#x60; で異なる](https://github.com/vim-jp/issues/issues/693)
- [Issue #694 : status/helpwanted の追加](https://github.com/vim-jp/issues/issues/694)

