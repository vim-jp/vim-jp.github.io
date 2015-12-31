---
layout: post
category: vimmagazine
title: Vim Magazine 12 月号
---

## 今月の新機能

- 7.4.944: `assert_equal()`, `assert_false()`, `assert_true()` 関数が追加されました。
- 7.4.951: `sort()` 関数に 'N' フラグが追加されました。

## リリース情報

- [7.4.944 : writing tests for Vim script is hard](https://github.com/vim/vim/commit/43345546ae63710441f066648b8485fb545b3801)
- [7.4.945 : (after 7.4.944) new style testing is incomplete](https://github.com/vim/vim/commit/683fa185a4b4ed7595e5942901548b8239ed5cdb)
- [7.4.946 : (after 7.4.945) missing changes in source file](https://github.com/vim/vim/commit/bbfbaf9741deebb9f1ed790885bd571c4cbce17a)
- [7.4.947 : test&#x5f;listchars fails with MingW](https://github.com/vim/vim/commit/5311c02f25eed8f34e8a80becb98e86264f371c3)
- [7.4.948 : can't build when the insert&#x5f;expand feature is disabled](https://github.com/vim/vim/commit/20ad69ccfb60ef718bd26387ef0e5424461a643e)
- [7.4.949 : highlighting wrong for 'colorcolumn' with a fullwidth sign](https://github.com/vim/vim/commit/32a214e78df0120f92fe049eab1385c60f0cdb0b)
- [7.4.950 : v:errors is not initialized](https://github.com/vim/vim/commit/4649ded2877508fe343cbcf6f7e7fd277be0aab3)
- [7.4.951 : sorting number strings does not work as expected](https://github.com/vim/vim/commit/b00da1d6d1655cb6e415f84ecc3be5ff3b790811)
- [7.4.952 : 'lispwords' is tested in the old way](https://github.com/vim/vim/commit/6cd1345307440491580e5e86cb82c54ee9a46baa)
- [7.4.953 : a test may create the .res file with a wrong name](https://github.com/vim/vim/commit/de0ad40cb3c1bc691a754698ed16a5b6cdb4086b)
- [7.4.954 : when using Lua there may be a crash (issue #468)](https://github.com/vim/vim/commit/f609dcf8c1094f6fc95f4fc36321a1fb08a7110c)
- [7.4.955 : Vim doesn't recognize .pl6 and .pod6 files](https://github.com/vim/vim/commit/3b8fcd945c5f0ee104eaabcf969fb6f973e79c77)
- [7.4.956 : a few more file name extensions not recognized](https://github.com/vim/vim/commit/f882d9f89dbe24ab1ba4bc88529bef28242fd2ed)
- [7.4.957 : test&#x5f;tagcase fails when using another language than English](https://github.com/vim/vim/commit/bc96c29ffc753daef302d20322d1e3d560094f44)
- [7.4.958 : Vim checks if the directory "$TMPDIR" exists](https://github.com/vim/vim/commit/e1a61991d9b6fd5f65636d17583f93118268cda5)
- [7.4.959 : when setting 'term' the clipboard ownership is lost](https://github.com/vim/vim/commit/0379d01c52e7930ccfc9133f229fba54a2024a42)
- [7.4.960 : detecting every version of nmake is clumsy](https://github.com/vim/vim/commit/90f5d0a5c3bbfeefcbc4d6eac59cf225ec714b28)
- [7.4.961 : test107 fails in some circunstances](https://github.com/vim/vim/commit/9dc2ce398bb3456cc8f590ef0260459798b34d2a)
- [7.4.962 : cannot run the tests with gvim and individual new tests](https://github.com/vim/vim/commit/3f12a2421bda43a4e48c822541b75f72ee11125a)
- [7.4.963 : test&#x5f;listlbr&#x5f;utf8 sometimes fails](https://github.com/vim/vim/commit/1c57fe8b9450eb29c3e42a94527d4b7514f853e2)
- [7.4.964 : test 87 doesn't work in a shadow directory](https://github.com/vim/vim/commit/91376b63877c113fe9a3fff2c1b04bf9504f447f)
- [7.4.965 : on FreeBSD /dev/fd/ files are special](https://github.com/vim/vim/commit/941aea2b975623a0c8bc24b140881ef0032a8bb8)
- [7.4.966 : configure doesn't work with a space in a path](https://github.com/vim/vim/commit/49222bee65228c7b5994b33c1568394c3cbf4583)
- [7.4.967 : cross compilation on MS-windows doesn't work well](https://github.com/vim/vim/commit/6b90351786eb0915336b576cc930300bf5c9ac63)
- [7.4.968 : test86 and test87 are flaky in Appveyor](https://github.com/vim/vim/commit/b65c749ac5a8a990d53493e3b9677142b1b9e4ce)
- [7.4.969 : compiler warnings on Windowx x64 build](https://github.com/vim/vim/commit/35be4534c029148a89ccc41e8e465d793e7ed7c2)
- [7.4.970 : rare crash in getvcol()](https://github.com/vim/vim/commit/9ec021a2b0dd35ba744a8e2a9430a643c85b922a)
- [7.4.971 : the asin() function can't be used](https://github.com/vim/vim/commit/099fdde0f073315b7f2700786ae533d23a556348)
- [7.4.972 : memory leak when there is an error in setting an option](https://github.com/vim/vim/commit/a98849670674264de699d7ab22ae4b9b32e78f4a)
- [7.4.973 : pasting on the command line line inserts \<CR> characters](https://github.com/vim/vim/commit/6f62fed349bf829da2adb02619dc9acba13c8ab6)
- [7.4.974 : when using :diffsplit the cursor jumps to the first line](https://github.com/vim/vim/commit/f29a82dcd0914c76f595d475ddac4517371fab2b)
- [7.4.975 : using ":sort" on a very big file may cause corrupted text](https://github.com/vim/vim/commit/75e3ad019933f4879137775549261bf51985ab7d)
- [7.4.976 : when compiling for MSYS2 the Win32 clipboard is not enabled](https://github.com/vim/vim/commit/8def26a0f5f5535e9af64e715cb80845fc8ec322)
- [7.4.977 : problem with 'linebreak' set and using "space" in 'listchars'](https://github.com/vim/vim/commit/9bc01ebb957d2b30d57bd30d7aee6f1df2a336b0)
- [7.4.978 : test&#x5f;cdo fails when using another language than English](https://github.com/vim/vim/commit/c42b9c670ea621d4dac0f216e011a6db576c5136)
- [7.4.979 : when crypt key changes blocks read from disk are not decrypted](https://github.com/vim/vim/commit/4a8c2cfc56b9affc36934aa0f20d8cfd2b1511c8)
- [7.4.980 : tests for :cdo, :ldo, etc. are outdated](https://github.com/vim/vim/commit/57d7971b5f1621071176eea81cdb0d1fc50c925d)
- [7.4.981 : an error in a test script goes unnoticed](https://github.com/vim/vim/commit/4686b323e4bc0f466500b018959f6c8965f010f9)
- [7.4.982 : keeping the list of tests updated is a hassle](https://github.com/vim/vim/commit/7b6156f4cd4027b664a916ba546e9b05d4c49e11)
- [7.4.983 : executing one test after "make testclean" doesn't work](https://github.com/vim/vim/commit/a60824308cd9bc192c5d38fc16cccfcf652b40f6)
- [7.4.984 : searchpos() always starts searching in the first column](https://github.com/vim/vim/commit/ad4d8a192abf44b89371af87d70b971cd654b799)
- [7.4.985 : can't build with Ruby 2.3.0](https://github.com/vim/vim/commit/f2f6d297966ec0e357640b71a238e51afcaba6cc)
- [7.4.986 : test49 doesn't work on MS-Windows, test70 is listed twice](https://github.com/vim/vim/commit/da9888a3f0118ce1ce5acbdcf4720602c2de2a3b)
- [7.4.987 : (after 7.4.985) can't build with Ruby 1.9.2](https://github.com/vim/vim/commit/0d27f64f7188efef99062a3c5694027c12401670)
- [7.4.988 : (after 7.4.982) default test target is test49.out](https://github.com/vim/vim/commit/40bbceee2213a6fa8fdc1d3f3920d61fb5370803)
- [7.4.989 : leaking memory when hash&#x5f;add() fails; Coverity error 99126](https://github.com/vim/vim/commit/0107f5ba87ca9427500d0fc42ec80a1f3fca9fdb)
- [7.4.990 : test 86 fails on AppVeyor](https://github.com/vim/vim/commit/e5c5f0c66c9491aca013f30da6e4f730a7ba7db6)
- [7.4.991 : when running new style tests the output is not visible](https://github.com/vim/vim/commit/096c8bb40d51b22a4b1d761baf7bb79fb9e55a28)
- [7.4.992 : makefiles for MS-Windows in src/po are outdated](https://github.com/vim/vim/commit/013806229a1e15480592f6bc8453130685ec750b)
- [7.4.993 : test 87 is flaky on AppVeyor](https://github.com/vim/vim/commit/52f6ae1366b34fc5771595c0bd17c779a7f6f544)
- [7.4.994 : new style tests are not run on MS-Windows](https://github.com/vim/vim/commit/4e5a31c8b3e259605f4d8543aaae68578cf9b0d7)
- [7.4.995 : gdk&#x5f;pixbuf&#x5f;new&#x5f;from&#x5f;inline() is deprecated](https://github.com/vim/vim/commit/36e294c00c784b9ddd05a4fdbea2e331ab2b1ca8)
- [7.4.996 : new files missing from distro, outdated PC build instructions](https://github.com/vim/vim/commit/e292d80bede5cb0b9b1ca95176ad6c3fbaae2e0a)
- [7.4.997 : "make shadow" was sometimes broken](https://github.com/vim/vim/commit/4c7bb12c82914307e6bbb73d95cfb3ba7189813a)
- [7.4.998 : running tests in shadow directory fails; test 49 fails](https://github.com/vim/vim/commit/f9c8bd2137b045f9a64d63eefcf022b4726b1419)
- [7.4.999 : "make shadow" creates a broken link](https://github.com/vim/vim/commit/b8cb643eab0e84d6a41f5884c7e41736218425fb)
- [7.4.1000 : test 49 is slow and doesn't work on MS-Windows](https://github.com/vim/vim/commit/c06624661a3aa6642304c06db9cebe553a4cab17)
- [7.4.1001 : (after 7.4.1000) change in makefile missing](https://github.com/vim/vim/commit/f49e240c2def978247fa457aa105bb3024413f7d)
- [7.4.1002 : cannot run an individual test on MS-Windows](https://github.com/vim/vim/commit/604619784c7f9007a883c123231d080598bd49f5)
- [7.4.1003 : Travis could check a few more things](https://github.com/vim/vim/commit/7b5f0a15bce11754c47f849b2ddd68ba0909afac)
- [7.4.1004 : warnings when using Makefile without auto/config.mk](https://github.com/vim/vim/commit/08b7bae91adb79d30d4c923fd758e2f7cecd33ef)

## 新着スクリプト

- [vim-cloud-buffer : Save your buffers remotely ](http://www.vim.org/scripts/script.php?script_id=5277)
- [vim-uci : Syntax, ftdetect and ftplugin for OpenWRT's UCI format](http://www.vim.org/scripts/script.php?script_id=5278)
- [vim-stata : Vim Plugin for Running Selected Do-File in Stata](http://www.vim.org/scripts/script.php?script_id=5279)
- [snpsMan2VimSyntax : Perl script to parse Synopsys man pages to create a vim syntax file](http://www.vim.org/scripts/script.php?script_id=5280)
- [Tabstuff : Tabstuff is a plugin to sulplies various tabrelated stuff](http://www.vim.org/scripts/script.php?script_id=5281)
- [vim-do : Run shell commands asynchronously and show the output in Vim when they finish](http://www.vim.org/scripts/script.php?script_id=5282)
- [PAM-CRASH Syntax Highlighting : Syntax Highlighting for ESI VPS (Pam-Crash)](http://www.vim.org/scripts/script.php?script_id=5283)
- [vim-dtsm : Commands for TypeScript dtsm](http://www.vim.org/scripts/script.php?script_id=5284)
- [vim-babel : Wrapper around babel.js (ES2015, React, ...)](http://www.vim.org/scripts/script.php?script_id=5285)
- [vim-unite-cscope : Use cscope within vim](http://www.vim.org/scripts/script.php?script_id=5286)
- [Esper programming language syntax : Basic syntax highlighting](http://www.vim.org/scripts/script.php?script_id=5287)
- [vim-rubyformat : A simple ruby source code auto reformatter/beautifier for vim](http://www.vim.org/scripts/script.php?script_id=5288)
- [tgf-vim : This is a PureFun-Tiger Language syntax and indent script.](http://www.vim.org/scripts/script.php?script_id=5289)
- [partialnumber.vim : yet another line number option](http://www.vim.org/scripts/script.php?script_id=5290)
- [vim-uncommenter : Automatically deletes all comments in a programming source file.](http://www.vim.org/scripts/script.php?script_id=5291)
- [vim-gtest : Vim plugin to quickly select and asyncronously run GoogleTest](http://www.vim.org/scripts/script.php?script_id=5292)
- [vim-operator-flashy : Highlight yanked area](http://www.vim.org/scripts/script.php?script_id=5293)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1616)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1528)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (638)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (598)
5. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (565)
6. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (515)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (505)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (472)
9. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (464)
10. [Pydiction : Tab-complete your Python code](http://www.vim.org/scripts/script.php?script_id=850) (460)

## vim-jp issues

Open : 242 (+4) | Closed : 560 (+7)

- [Issue #792 : assert&#x5f;equal/assert&#x5f;true/assert&#x5f;false の第三引数 msg が出力されない。](https://github.com/vim-jp/issues/issues/792)
- [Issue #793 : 物理的なカーソル位置が欲しい](https://github.com/vim-jp/issues/issues/793)
- [Issue #794 : 補完の絞り込みチラつき改善](https://github.com/vim-jp/issues/issues/794)
- [Issue #795 : v:errors が初期化されていない](https://github.com/vim-jp/issues/issues/795)
- [Issue #796 : set listcharsのspacesに何か設定して set list linebreak するとlinebreakされない且つカーソル位置がズレる](https://github.com/vim-jp/issues/issues/796)
- [Issue #797 : オレオレsyntax/vim.vim generatorを自動実行して最新ファイルをWeb上に置いておきたい](https://github.com/vim-jp/issues/issues/797)
- [Issue #798 : 長い行で 'set number' すると行数表示やカーソル位置がおかしくなることがある](https://github.com/vim-jp/issues/issues/798)
- [Issue #799 : +conceal機能がright-to-leftに対応していない](https://github.com/vim-jp/issues/issues/799)
- [Issue #800 : vim-airline causes crash on macvim ](https://github.com/vim-jp/issues/issues/800)
- [Issue #801 : Kaoriya版VimでNetrwからファイルに移動すると、終了できなくなる](https://github.com/vim-jp/issues/issues/801)
- [Issue #802 : Ruby 2.3.0 で +ruby/dyn のビルドに失敗する](https://github.com/vim-jp/issues/issues/802)

