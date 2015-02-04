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

- [7.4.561 : Ex range handling is wrong for buffer-local user commands](http://code.google.com/p/vim/source/detail?r=82c6a3bddb23c7343fefc271b33960a20506ead5)
- [7.4.562 : segfault with wide screen and error in 'rulerformat'](http://code.google.com/p/vim/source/detail?r=b5df99582638f6f6cad794b7049453d7ff2c70b2)
- [7.4.563 : no test for replacing on a tab in Virtual replace mode](http://code.google.com/p/vim/source/detail?r=9dc52b404c1b3f32d36df2841c4c8eddea6e4407)
- [7.4.564 : FEAT&#x5f;OSFILETYPE is used even though it's never defined](http://code.google.com/p/vim/source/detail?r=9657929ee56251b52885cd0660941b0aab2af3ba)
- [7.4.565 : some ranges are not checked to be valid](http://code.google.com/p/vim/source/detail?r=22f95a018b13dc9d8a3246d990a1ea0b5b3a264d)
- [7.4.566 : :argdo, :bufdo, :windo and :tabdo don't take a range](http://code.google.com/p/vim/source/detail?r=a88d4dc02bf438ff9e6ee119be8188e342ea459e)
- [7.4.567 : non-ascii vertical separater characters are always redrawn](http://code.google.com/p/vim/source/detail?r=0b245c8dbd19755e08b5248110a537a74ddb92c4)
- [7.4.568 : giving an error for ":0wincmd w" is a problem for some plugins](http://code.google.com/p/vim/source/detail?r=747cbac7047bf8d003ff0672b25b60e0236fdb1b)
- [7.4.569 : having CTRL-C interrupt does not check the mapping mode](http://code.google.com/p/vim/source/detail?r=ccb6962585f1ea5e7f872d79e59565e7a95234db)
- [7.4.570 : building with dynamic library does not work for Ruby 2.2.0](http://code.google.com/p/vim/source/detail?r=816c584ff1740cd118d5a35fbea98ed28ed2e2e2)
- [7.4.571 : can't build with tiny features](http://code.google.com/p/vim/source/detail?r=97b2ff29ae3a213b5c72bd8ea3aa2442ef355e08)
- [7.4.572 : address type of :wincmd depends on the argument](http://code.google.com/p/vim/source/detail?r=cba15023c40371ed51ee0ed50fbdf23d6860f0ba)
- [7.4.573 : mapping CTRL-C in Visual mode doesn't work](http://code.google.com/p/vim/source/detail?r=8cfbc34ae4aa2b91b5b012b2e8dfe86327b449c9)
- [7.4.574 : no error for eval('$')](http://code.google.com/p/vim/source/detail?r=45ff9dd354eaff958c0ba646134e1f5cd6e9186f)
- [7.4.575 : Unicode character properties are outdated](http://code.google.com/p/vim/source/detail?r=fe0f33ac63af3c16a9920fe08bd3582aa0f02f50)
- [7.4.576 : redrawing problem with 'relativenumber' and 'linebreak'](http://code.google.com/p/vim/source/detail?r=749fc929da45d07efc5352dd8a0c00f8ed9f2163)
- [7.4.577 : matching with virtual column has overhead on very long lines](http://code.google.com/p/vim/source/detail?r=bcf9d3a6007fff06d0521339f93dc870e5231eaa)
- [7.4.578 : after "$" in an empty line getcurpos() returns negative number](http://code.google.com/p/vim/source/detail?r=b7b3afcae90e5b8e9f57ec3b172fed14f6912579)
- [7.4.579 : wrong cursor positioning when 'linebreak' set and lines wrap](http://code.google.com/p/vim/source/detail?r=134a7d55f9fa170c8c4591328eb550e70b756616)
- [7.4.580 : ":52wincmd v" still gives an invalid range error](http://code.google.com/p/vim/source/detail?r=a5a6f1467b1ea73697b3265d8f12805bd998c81e)
- [7.4.581 : compiler warnings for unitinialized variables](http://code.google.com/p/vim/source/detail?r=e1874a4524f8c97efed78e215fd5ff91f3569351)
- [7.4.582 : (after 7.4.577) can't match "%>80v" properly](http://code.google.com/p/vim/source/detail?r=dfd593d81818468ad60cad4fc66c561299f8d969)
- [7.4.583 : with tiny features test 16 may fail](http://code.google.com/p/vim/source/detail?r=f2978794e801bfe8a587865492302cc0a16783cc)
- [7.4.584 : with tiny features test&#x5f;command&#x5f;count may fail](http://code.google.com/p/vim/source/detail?r=79b3d83592c5ad63ce0c8d276d15f570e95dd941)
- [7.4.585 : range for :bdelete does not work](http://code.google.com/p/vim/source/detail?r=1d3139ae6ce67d69ce8bc221bfc67dc1b47e95d9)
- [7.4.586 : parallel building of documentation html files is not reliable](http://code.google.com/p/vim/source/detail?r=667cb2eb17a52d9ac28a1b6f686935814e769ecd)
- [7.4.587 : conceal does not work properly with 'linebreak'](http://code.google.com/p/vim/source/detail?r=70793fedd25a509568abbd4026565675e830b1ca)
- [7.4.588 : ":0argedit foo" puts the new argument in the second place](http://code.google.com/p/vim/source/detail?r=a03e143b93587cf29c91121acdf66f6e5dc6e470)
- [7.4.589 : MS-Windows console: Vim can't handle greek utf-8 characters](http://code.google.com/p/vim/source/detail?r=8ff2d1b611bfeb17b594262cbf9313f819b4483a)
- [7.4.590 : using ctrl&#x5f;x&#x5f;mode as if it contains flags](http://code.google.com/p/vim/source/detail?r=5b8d26b85437ffa91cd586fa138aba80bba897a9)
- [7.4.591 : test&#x5f;listlbr&#x5f;utf8 fails when conceal feature is not available](http://code.google.com/p/vim/source/detail?r=7b8e3be930f39af3cdc1fe969a6de16f2d91e5d2)
- [7.4.592 : buffer with 'buftype' "nofile" is cleared on re-edit](http://code.google.com/p/vim/source/detail?r=005a6b65086666ed56531e4c8a531a3ee0cea638)
- [7.4.593 : crash when searching for "x\\{0,90000}"](http://code.google.com/p/vim/source/detail?r=bdc8e71633e45ebe5f1b59e4328fcf22a5866cb7)
- [7.4.594 : using block delete with 'breakindent' does not work properly](http://code.google.com/p/vim/source/detail?r=f9d02ce2f745f75003ff570b2a596c755ccc86ba)
- [7.4.595 : the test&#x5f;command&#x5f;count test fails when using Japanese](http://code.google.com/p/vim/source/detail?r=d83436d119871584ddb012ecf0aa4d01030f65ee)
- [7.4.596 : (after 7.4.592) tiny build doesn't compile](http://code.google.com/p/vim/source/detail?r=6da912e32896c57227a0f297aa1d3312a653cf0a)
- [7.4.597 : cannot change the result of systemlist()](http://code.google.com/p/vim/source/detail?r=c124a8f34ed1cf510d384b384db6f34be5d61a3e)
- [7.4.598 : ":tabdo windo echo 'hi'" causes "&#x2a; register not to be changed](http://code.google.com/p/vim/source/detail?r=b01ca71f93b293ba93489bc2320c12caf61bf289)
- [7.4.599 : out-of-memory error for complicated search pattern](http://code.google.com/p/vim/source/detail?r=1ef8ce97fc40b59d7218d7300685d40df63a407c)
- [7.4.600 : memory wasted in struct because of aligning](http://code.google.com/p/vim/source/detail?r=436d6c9e57f2f6def73e86891cbf72a0328aa9bd)
- [7.4.601 : it is not possible to have feedkeys() insert characters](http://code.google.com/p/vim/source/detail?r=2561531decf15c1a02f62a1b96ff778e12a5dcbc)
- [7.4.602 : ":set" does not accept hex numbers as documented](http://code.google.com/p/vim/source/detail?r=4e31d9f7c896a346b2346fc4ff005b8b255172c9)
- [7.4.603 : 'foldcolumn' may fill the whole window](http://code.google.com/p/vim/source/detail?r=0c8738124a7b9d6d1d3e18124ad97310e9618c92)
- [7.4.604 : running tests changes viminfo](http://code.google.com/p/vim/source/detail?r=f625b6302d048df7d9f9e738b773f91dd4bbd58e)
- [7.4.605 : the # register cannot be restored after jumping around](http://code.google.com/p/vim/source/detail?r=c1ed973fb58faa650dcd4d25a98c669edb8ca0d4)
- [7.4.606 : may crash when using a small window](http://code.google.com/p/vim/source/detail?r=22f164bd7e8ea7de1e4c2a41020dde971112b862)
- [7.4.607 : (after 7.4.598) compiler warnings for unused variables](http://code.google.com/p/vim/source/detail?r=52b89a52ffd4f2a925b42be1aa8041bc0c0dff31)
- [7.4.608 : (after 7.4.598) test&#x5f;eval fails without the clipboard feature](http://code.google.com/p/vim/source/detail?r=6b86d256fd3360a5a0a8f89c3efeda4e961153c0)

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

