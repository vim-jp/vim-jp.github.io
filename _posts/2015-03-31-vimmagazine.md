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

- [7.4.649 : compiler complains about ignoring return value of fwrite()](https://github.com/vim/vim/commit/cf48767cd17130958a3076eed1872b6950947a0a)
- [7.4.650 : configure check may fail because the dl library is not used](https://github.com/vim/vim/commit/5d3fbf3bfea3b698f42a44c2732e58fa63f01446)
- [7.4.651 : (after 7.4.582) can't match "%>80v" for multi-byte characters](https://github.com/vim/vim/commit/4f36dc3bf7118a3fa5ccdae1a37963860dc71e6c)
- [7.4.652 : xxd lacks a few features](https://github.com/vim/vim/commit/4dcdf299100781766d09829b9b80d7fb1ad48f37)
- [7.4.653 : Insert mode completion may have CTRL-L work like CTRL-P](https://github.com/vim/vim/commit/e421450a5ef8c6327fcc9b90a5c3d6b26b274830)
- [7.4.654 : glob()/ globpath() cannot include links to non-existing files](https://github.com/vim/vim/commit/a245bc79b4c6b83a4b5b6cdb95c4d2165762a20b)
- [7.4.655 : text deleted by "dit" depends on indent of closing tag](https://github.com/vim/vim/commit/b6c2735c56f1541159e1ad95c3f17a52b3a94f1d)
- [7.4.656 : (after 7.4.654) missing changes for glob() in one file](https://github.com/vim/vim/commit/d8b77f7dc04e5721989df9c505b8568194261a39)
- [7.4.657 : (after 7.4.656) compiler warnings for pointer mismatch](https://github.com/vim/vim/commit/ab11a5914b04228249d429cbc2dd6c56ae6c8232)
- [7.4.658 : 'formatexpr' is evaluated too often](https://github.com/vim/vim/commit/0f8dd840fc6a614450db60ebe405d6201a2ecc3e)
- [7.4.659 : when 'ruler' is set the preferred column is reset (Issue 339)](https://github.com/vim/vim/commit/488eb26ec7fda1dc79f3d856d40da072332630b6)
- [7.4.660 : using freed memory if g:colors&#x5f;name is changed](https://github.com/vim/vim/commit/c7dc1f4a53fdb26b88c484003142c575655481b7)
- [7.4.661 : using "0 CTRL-D" in Insert mode may have CursorHoldI interfere](https://github.com/vim/vim/commit/6c5bdb751c0c77be7d03a6000134d1df367763d2)
- [7.4.662 : selecting text object can be wrong when 'M' is in the 'cpo'](https://github.com/vim/vim/commit/438b64ab71cd724129c4eec840be16c52602ebc8)
- [7.4.663 : when using netbeans a buffer is not found in another tab](https://github.com/vim/vim/commit/404c942fb85c390f4b73c1733088871b0b0d7964)
- [7.4.664 : effect of 'numberwidth' being set not immediately visible](https://github.com/vim/vim/commit/6b31467aa7f304a80343216c11a1588dff361229)
- [7.4.665 : 'linebreak' does not work properly with multi-byte characters](https://github.com/vim/vim/commit/76feaf1bfed170d2fc5c34e7e758885554ecc794)
- [7.4.666 : there is a chance that Vim may lock up](https://github.com/vim/vim/commit/b1fc2bf29c560ca60515950812e8672c6dca1a4c)
- [7.4.667 : 'colorcolumn' isn't drawn in closed fold like 'cursorcolumn'](https://github.com/vim/vim/commit/fbc25b20cbf29d6f9846e97875474f18e8d0c7be)
- [7.4.668 : can't use a glob pattern as a regexp pattern](https://github.com/vim/vim/commit/825e7ab423ce34a57f4819cc3745246246f38081)
- [7.4.669 : when netbeans is active the sign column always shows up](https://github.com/vim/vim/commit/3b7b83649cd9c486afd6fbdb33e4450616c4bd46)
- [7.4.670 : using 'cindent' for Javascript is less than perfect](https://github.com/vim/vim/commit/dcefba993456263100785f18d21d780cf531f2dc)
- [7.4.671 : (after 7.4.665) warning for shadowing a variable](https://github.com/vim/vim/commit/4df702999d14955255fcdfb820511767dcfec463)
- [7.4.672 : shell command completion does not see local directories](https://github.com/vim/vim/commit/b5971141dff0c69355fd64196fcc0d0d071d4c82)
- [7.4.673 : first syntax entry gets wrong sequence number zero](https://github.com/vim/vim/commit/e7154eb5c391aa39d5a87b47b9f7338f1a65354e)
- [7.4.674 : (after 7.4.672) missing changes in one file](https://github.com/vim/vim/commit/77b7710d4e0ec9c1f3d0d057d9fa335420273b37)
- [7.4.675 : when FileReadPost moves cursor inside line it gets moved back](https://github.com/vim/vim/commit/eab316bdf9494eb1e076dfc5c8ec7ae000a0560f)
- [7.4.676 : on Mac, configure can't handle non-default Python framework](https://github.com/vim/vim/commit/6c927557e2b0b2fbc2dfc91f604bf1117d0bf77a)
- [7.4.677 : (after 7.4.676) configure fails when using python-config-dir](https://github.com/vim/vim/commit/a161e26d558204311f6e76208fa4cd3b2dbbab0b)
- [7.4.678 : when using --remote the directory may end up being wrong](https://github.com/vim/vim/commit/f11ce667b06ffb7447fb55638a04b5fa5cfe29a0)
- [7.4.679 : color values greater than 255 cause problems on MS-Windows](https://github.com/vim/vim/commit/6383b92d8c272f75d1fe144f75e2e35ad7e11b04)
- [7.4.680 : CTRL-W in Insert mode does not work well for multi-byte chars](https://github.com/vim/vim/commit/310f2d59b2b20c642088feb5e6dfe323cc570923)
- [7.4.681 : MS-Windows: with minimized Vim window height is incorrect](https://github.com/vim/vim/commit/71371b15794941dbeb01e46827d6d501dc95acbd)
- [7.4.682 : search and match highlighting replace cursorline highlighting](https://github.com/vim/vim/commit/09deeb7c945d2677722df5b40959f97b646b6092)
- [7.4.683 : typo in the vimtutor command](https://github.com/vim/vim/commit/1ca2e361a869c58a775b5fabb2aaf4e2675ea7ba)

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

