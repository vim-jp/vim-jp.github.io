---
layout: post
category: vimmagazine
title: Vim Magazine 3 月号
---

## 話題

- Vim のリポジトリの Github への移行が発表されました。

  [github への移行準備](http://vim-jp.org/blog/2015/03/25/preparations-for-moving-to-github.html)

## 今月の新機能

- 7.4.652: xxd に -e (switch to little-endian hexdump) と -o (offset) が追加されました。

- 7.4.654: glob() と globpath() に 'alllinks' 引数が追加されました。

- 7.4.668: glob2regpat() 関数が追加されました。

## リリース情報

- [7.4.649 : compiler complains about ignoring return value of fwrite()](http://code.google.com/p/vim/source/detail?r=8bee881c3ca5d72008268db5dde8e2940226f6bd)
- [7.4.650 : configure check may fail because the dl library is not used](http://code.google.com/p/vim/source/detail?r=7026d6f3367fca47d664a1b1ef1b302f5c0edc45)
- [7.4.651 : (after 7.4.582) can't match "%>80v" for multi-byte characters](http://code.google.com/p/vim/source/detail?r=2c7f279d419caab2b6329e0415322ebf4fa3f312)
- [7.4.652 : xxd lacks a few features](http://code.google.com/p/vim/source/detail?r=6c94cedccda4569141ddf71005a4087b8df833dc)
- [7.4.653 : Insert mode completion may have CTRL-L work like CTRL-P](http://code.google.com/p/vim/source/detail?r=41d50eba0c0f472b9d4399e01989679ad893ed2f)
- [7.4.654 : glob()/ globpath() cannot include links to non-existing files](http://code.google.com/p/vim/source/detail?r=6f1e4b7942c328173b1cb2b322314e028669a866)
- [7.4.655 : text deleted by "dit" depends on indent of closing tag](http://code.google.com/p/vim/source/detail?r=950e24f26ef8cc460c66238efce9104eaf424f79)
- [7.4.656 : (after 7.4.654) missing changes for glob() in one file](http://code.google.com/p/vim/source/detail?r=056809de0b298e7f32a76ace0575088a639f3b24)
- [7.4.657 : (after 7.4.656) compiler warnings for pointer mismatch](http://code.google.com/p/vim/source/detail?r=9fda0c52db0b0aa8819370d6061c8af61a8a13a8)
- [7.4.658 : 'formatexpr' is evaluated too often](http://code.google.com/p/vim/source/detail?r=9f74286c4404346bb69d10575eb5ac82cb635add)
- [7.4.659 : when 'ruler' is set the preferred column is reset (Issue 339)](http://code.google.com/p/vim/source/detail?r=5a12a68613670ddb7d3be0c693738400c61ef24b)
- [7.4.660 : using freed memory if g:colors&#x5f;name is changed](http://code.google.com/p/vim/source/detail?r=476a12a96bb1e5ccc71b1c1010d529deca271617)
- [7.4.661 : using "0 CTRL-D" in Insert mode may have CursorHoldI interfere](http://code.google.com/p/vim/source/detail?r=21719d38d2ff265c7f55230df8529780280f2cf9)
- [7.4.662 : selecting text object can be wrong when 'M' is in the 'cpo'](http://code.google.com/p/vim/source/detail?r=24e0b6dd52e1b259559f142e8c04aa0a74128146)
- [7.4.663 : when using netbeans a buffer is not found in another tab](http://code.google.com/p/vim/source/detail?r=77172f5bac73b69838da5b9161ef5574cf3b7189)
- [7.4.664 : effect of 'numberwidth' being set not immediately visible](http://code.google.com/p/vim/source/detail?r=27d9a4d8c1689123939b6c3114e8e60c9656f507)
- [7.4.665 : 'linebreak' does not work properly with multi-byte characters](http://code.google.com/p/vim/source/detail?r=80abb7288d39c917b875f6788246356f8f1da645)
- [7.4.666 : there is a chance that Vim may lock up](http://code.google.com/p/vim/source/detail?r=781104cb0d069389b2db31a7553118f44e4f24c1)
- [7.4.667 : 'colorcolumn' isn't drawn in closed fold like 'cursorcolumn'](http://code.google.com/p/vim/source/detail?r=fd495ad521dc55f76636b7dde74f59739512fa45)
- [7.4.668 : can't use a glob pattern as a regexp pattern](http://code.google.com/p/vim/source/detail?r=c52bb45c675242c167626ad303fbc8fb8ecaf844)
- [7.4.669 : when netbeans is active the sign column always shows up](http://code.google.com/p/vim/source/detail?r=97cc4ee3e095ce4e7178c41a39bac864ba6d2c7d)
- [7.4.670 : using 'cindent' for Javascript is less than perfect](http://code.google.com/p/vim/source/detail?r=c4cb95fb4c585d81adc794421f16f62a30f0fc30)
- [7.4.671 : (after 7.4.665) warning for shadowing a variable](http://code.google.com/p/vim/source/detail?r=8faab3792a924c58b592170980fc74244d7cf402)
- [7.4.672 : shell command completion does not see local directories](http://code.google.com/p/vim/source/detail?r=6529590f6c436594c3ba177c61efd9b42b3cbc08)
- [7.4.673 : first syntax entry gets wrong sequence number zero](http://code.google.com/p/vim/source/detail?r=67e7a819c8115a149d2e780fa1fa7d9fa96a6c02)
- [7.4.674 : (after 7.4.672) missing changes in one file](http://code.google.com/p/vim/source/detail?r=f26c1d294156c092290595897694627867809199)
- [7.4.675 : when FileReadPost moves cursor inside line it gets moved back](http://code.google.com/p/vim/source/detail?r=d879db7c2f56d8c16d997c4b37cee38036d7a721)
- [7.4.676 : on Mac, configure can't handle non-default Python framework](http://code.google.com/p/vim/source/detail?r=692bc2c52002dcc3a6c0e809057d46792e70db29)
- [7.4.677 : (after 7.4.676) configure fails when using python-config-dir](http://code.google.com/p/vim/source/detail?r=c8ff693e9ebc074137d0566868561b13dc11fe0c)
- [7.4.678 : when using --remote the directory may end up being wrong](http://code.google.com/p/vim/source/detail?r=b225e8fed8cc0e320527922e6abf1c7b8c6b83b6)
- [7.4.679 : color values greater than 255 cause problems on MS-Windows](http://code.google.com/p/vim/source/detail?r=43a07fa781552a64fe3e9be1bb504f8bac146a24)
- [7.4.680 : CTRL-W in Insert mode does not work well for multi-byte chars](http://code.google.com/p/vim/source/detail?r=c939d19bd86edd083bd690e2b6e1315877b098d7)
- [7.4.681 : MS-Windows: with minimized Vim window height is incorrect](http://code.google.com/p/vim/source/detail?r=dc96dd80aef84bbdf0537fbf9c9f5ab63a7ab649)
- [7.4.682 : search and match highlighting replace cursorline highlighting](http://code.google.com/p/vim/source/detail?r=f242cf1f98975b4840e5f4644c3c3e84d08d96eb)
- [7.4.683 : typo in the vimtutor command](http://code.google.com/p/vim/source/detail?r=6f55637de261d7e407502bd4d62557ec64aad2f9)

## 新着スクリプト

- [RedirCommands : plugin for redirecting command output](http://www.vim.org/scripts/script.php?script_id=5122)
- [codeforces.vim : Vim plugin for CodeForces](http://www.vim.org/scripts/script.php?script_id=5123)
- [Vinfo : Read Info documentation files](http://www.vim.org/scripts/script.php?script_id=5124)
- [love-launch : Run LOVE](http://www.vim.org/scripts/script.php?script_id=5125)
- [vim-snoopy : snoop around the script-local things](http://www.vim.org/scripts/script.php?script_id=5126)
- [vim-ranger : A vim plugin which uses ranger as explorer](http://www.vim.org/scripts/script.php?script_id=5127)
- [vim-quick-radon : vim-quick-radon is a Vim plugin that applies radon to your current file.](http://www.vim.org/scripts/script.php?script_id=5128)
- [textobj-functioncall : The textobject plugin to select a region that calls function.](http://www.vim.org/scripts/script.php?script_id=5129)
- [vim-tlv-mode : Transaction-Level Verilog support for Vim](http://www.vim.org/scripts/script.php?script_id=5130)
- [VimPyServer : A Python server to remote control Vim without X-windows or vim-server](http://www.vim.org/scripts/script.php?script_id=5131)
- [vim-mail : A plugin to ease mail redaction from vim](http://www.vim.org/scripts/script.php?script_id=5132)
- [vim-compile : Yet another compile plugin for vim](http://www.vim.org/scripts/script.php?script_id=5133)
- [todo-txt.vim : Vim plugin for Todo.txt](http://www.vim.org/scripts/script.php?script_id=5134)
- [ColumnTags : Navigate your tags in column-view](http://www.vim.org/scripts/script.php?script_id=5135)
- [plugconf : A VIM plugin for managing configurations of other plugins.](http://www.vim.org/scripts/script.php?script_id=5136)
- [vim.ergonomic : use vim more ergonomically](http://www.vim.org/scripts/script.php?script_id=5137)
- [vim-markdown-wiki : Links manipulation and navigation in markdown pages](http://www.vim.org/scripts/script.php?script_id=5138)
- [vim-scrum-markdown : manage a scrum dashboard inside vim](http://www.vim.org/scripts/script.php?script_id=5139)
- [vim-http-client : Make HTTP requests from Vim without wrestling the command line!](http://www.vim.org/scripts/script.php?script_id=5140)
- [vim-uncrustify : Helper to call uncrustify in vim.](http://www.vim.org/scripts/script.php?script_id=5141)
- [unmaze.vim : A vim plugin that shows the C/C++ preprocessor conditions that need to be met fo](http://www.vim.org/scripts/script.php?script_id=5142)
- [Quicker&#x5f;file : 2-step search plugin for quicker open](http://www.vim.org/scripts/script.php?script_id=5143)
- [gulp-vim : Simple gulp wrapper for Vim](http://www.vim.org/scripts/script.php?script_id=5144)
- [spss syntax highlighting file : syntax highlighting file based on spss version 22](http://www.vim.org/scripts/script.php?script_id=5145)
- [AtelierSulphurpool : Sulphurpool is a purplish colorscheme](http://www.vim.org/scripts/script.php?script_id=5146)
- [ShowSpaces : Highlight indentation characters if they contain spaces, even mixed with tabs](http://www.vim.org/scripts/script.php?script_id=5148)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (6646)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1963)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1770)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1116)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (798)
6. [vim-orgmode : Text outlining and task management for Vim based on Emacs' Org-Mode](http://www.vim.org/scripts/script.php?script_id=3642) (773)
7. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (760)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (680)
9. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (620)
10. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (602)

## vim-jp issues

Open : 226 (+7) | Closed : 495 (+9)

- [Issue #706 : 'cursorlinenr' オプションの追加](https://github.com/vim-jp/issues/issues/706)
- [Issue #707 : unix端末版vimで、シェルの呼び出しのタイミングでtcap-queryの応答が返るとエコーバックする](https://github.com/vim-jp/issues/issues/707)
- [Issue #708 : showmatch for ruby is too slow](https://github.com/vim-jp/issues/issues/708)
- [Issue #709 : Patch 7.4.651 のテストが無い](https://github.com/vim-jp/issues/issues/709)
- [Issue #710 : search.c findmatchlimit()でflagsがFM&#x5f;SKIPCOMMの場合の処理が未実装になっている](https://github.com/vim-jp/issues/issues/710)
- [Issue #711 : cpoptions+=M時のテキストオブジェクト選択の挙動が微妙](https://github.com/vim-jp/issues/issues/711)
- [Issue #712 : vimでカーソルの表示位置がずれる](https://github.com/vim-jp/issues/issues/712)
- [Issue #713 : ordered-dict の実装](https://github.com/vim-jp/issues/issues/713)
- [Issue #714 : set shiftwidth=0 の場合に対応していないインデントファイルが存在する](https://github.com/vim-jp/issues/issues/714)
- [Issue #715 : WindowsのgVimでホイールスクロールするとゴミが残る](https://github.com/vim-jp/issues/issues/715)
- [Issue #716 : 'breakindent'オプションを使用時に長い行でカーソルの位置がずれる](https://github.com/vim-jp/issues/issues/716)
- [Issue #717 : exコマンド :qflistdo と :loclistdo の追加](https://github.com/vim-jp/issues/issues/717)
- [Issue #718 : 単純なVim操作で&#x60;E315: ml&#x5f;get: invalid lnum: 2&#x60;が発生する。](https://github.com/vim-jp/issues/issues/718)
- [Issue #719 : 通常操作で cursorline の残像が残る](https://github.com/vim-jp/issues/issues/719)
- [Issue #720 : hg updateでsrc/auto/configureが更新された後のmakeでconfigure設定が引き継がれない](https://github.com/vim-jp/issues/issues/720)
- [Issue #721 : vim.exe が良く分からないけど落ちる](https://github.com/vim-jp/issues/issues/721)

