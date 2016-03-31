---
layout: post
category: vimmagazine
title: Vim Magazine 6 月号
---

## 話題

- [momonga.vim #5](http://connpass.com/event/7231/)

  Vimをします（かなり）。 7月26日(土) 東京

## 今月の新機能

- 7.4.330 : `matchaddpos()` 関数が追加されました。
- 7.4.338 : `'breakindent'` オプションが追加されました。`'wrap'` で折り返したときにインデントするようにします。

## リリース情報

- [7.4.317 : crash when starting gvim, issue 230](https://github.com/vim/vim/commit/06e4a6df0a31138b50b81b6931d6f39f207cd456)
- [7.4.318 : fg/bg color settings ignored when checking syntax attributes](https://github.com/vim/vim/commit/e3a8bab6304117b5c341171299472e1a00e70c85)
- [7.4.319 : crash when putting zero bytes on the clipboard](https://github.com/vim/vim/commit/980e58f7b324980d40ce690506897d138e5e2b70)
- [7.4.320 : possible crash when an BufLeave autocommand deletes the buffer](https://github.com/vim/vim/commit/3be8585661f8b0a9b94e9bc1db6ebfa097c3270f)
- [7.4.321 : can't build with strawberry perl 5.20 + mingw-w64-4.9.0](https://github.com/vim/vim/commit/0e6c5ef7ef4f4a1eb16b5eebc97e3bf139d6276a)
- [7.4.322 : using "msgfmt" is hard coded, cannot use "gmsgfmt"](https://github.com/vim/vim/commit/e639eb44e9ca063f682f064095aa98fc781f5df1)
- [7.4.323 : substitute() with zero width pattern breaks multi-byte char](https://github.com/vim/vim/commit/8e7048ca4b699bbdb8452c0f8c642bfd66cc7bf4)
- [7.4.324 : in Ex mode, cyrillic characters are not handled](https://github.com/vim/vim/commit/2d54ec92d42a659aca4ec9e16521cd32767a7faa)
- [7.4.325 : incorrect display when starting GUI and setting 'laststatus'](https://github.com/vim/vim/commit/0ae36a5c8542d51dd0d70035e997d96e4a20dcbb)
- [7.4.326 : can't build Tiny version](https://github.com/vim/vim/commit/f0224c9f085849f4da87faa878279a641055fe83)
- [7.4.327 : when 'verbose' is set may get E724 repeatedly](https://github.com/vim/vim/commit/8502c704badd5f93875e56d887a6a1e5810d131f)
- [7.4.328 : selection of inner block is inconsistent](https://github.com/vim/vim/commit/7a54a90744529ab4d3291fa5cdf7656184b8469c)
- [7.4.329 : window not scrolled if moving cursor and go to another window](https://github.com/vim/vim/commit/ec1561cac59006213dd5405d164a94dc7d002806)
- [7.4.330 : using regexp pattern to show a position match can be slow](https://github.com/vim/vim/commit/b3414595c713ca161f082776f89417faddec7d2d)
- [7.4.331 : relative numbering not updated after a linewise yank](https://github.com/vim/vim/commit/7ec83432b7752a1c1956e967afa40611d40c179b)
- [7.4.332 : GTK: ugly gaps when a sign icon doesn't fit exactly](https://github.com/vim/vim/commit/58cbc914eacad98d70eff79292e77ecb24107c8f)
- [7.4.333 : compiler warning for unused function](https://github.com/vim/vim/commit/de993ea629c532686c3d888246609d2b0c1880b7)
- [7.4.334 : (after 7.4.330) unitialized variables, causing some problems](https://github.com/vim/vim/commit/deae0f25663bdcd30b7fd8eb440ab0d34ee7ed45)
- [7.4.335 : no digraph for the new rouble sign](https://github.com/vim/vim/commit/e8d1f20cbd425e33e4fcc09d46f98e5c9e6a833e)
- [7.4.336 : setting 'history' to a big value causes out-of-memory errors](https://github.com/vim/vim/commit/78159bbf9e241f162b1243853f45d08a3d81f1be)
- [7.4.337 : can't execute command line if there was an error before](https://github.com/vim/vim/commit/15a35c4f4a1670dd6ca228068a451f78d2bf75e0)
- [7.4.338 : cannot wrap lines taking indent into account](https://github.com/vim/vim/commit/597a422416f37f8e22ed8f561667d6bab8814958)
- [7.4.339 : local function is available globally](https://github.com/vim/vim/commit/8dc907d7d30c9f60f14f197e99f2f62001dd09f0)
- [7.4.340 : error from sed about illegal bytes when installing Vim](https://github.com/vim/vim/commit/72e8f0bcc1311aaec8949ddd9c0d0515c03ca3e0)
- [7.4.341 : sort() doesn't handle numbers well](https://github.com/vim/vim/commit/e8a3492548468c9d055adeb9f8d352e2c4482b19)
- [7.4.342 : clang gives warnings](https://github.com/vim/vim/commit/db5ffaab5a4755981ee7eff2f044b533ce3042a9)
- [7.4.343 : matchdelete() does not always update the right lines](https://github.com/vim/vim/commit/41d75239863d4b7a4142fc68cfa5d3a74a18d935)
- [7.4.344 : small issues related to matchaddpos()](https://github.com/vim/vim/commit/b6da44ae82f93d036ffb2ba37a009119688d9971)
- [7.4.345 : (after 7.4.338) indent is not updated when deleting indent](https://github.com/vim/vim/commit/a40aa760a1f49defccf2f8090d9f5ddf463e3303)
- [7.4.346 : (after 7.4.338) indent is not updated when changing 'briopt'](https://github.com/vim/vim/commit/9d7a592c8007d67627eac34edac14996926a99f8)
- [7.4.347 : test55 fails on some systems](https://github.com/vim/vim/commit/b7d929a86f4c3c4d127b37728b26f3a4632b8474)

## 新着スクリプト

- [caniuse.vim : Quickly search caniuse.com from vim](http://www.vim.org/scripts/script.php?script_id=4951)
- [dubs-red : A reddish version of the dubs scheme](http://www.vim.org/scripts/script.php?script_id=4952)
- [dubs-green : dubs-scheme gone green](http://www.vim.org/scripts/script.php?script_id=4953)
- [purpledrop : dark purple backdrop nice for coding](http://www.vim.org/scripts/script.php?script_id=4954)
- [Merginal : Fugitive extention to manage and merge Git branches](http://www.vim.org/scripts/script.php?script_id=4955)
- [vimbed : A plugin for embedding vim](http://www.vim.org/scripts/script.php?script_id=4956)
- [zeavim : Execute the documentation browser Zeal from Vim](http://www.vim.org/scripts/script.php?script_id=4957)
- [AdvancedSorters : Sorting of certain areas or by special needs.](http://www.vim.org/scripts/script.php?script_id=4958)
- [vimf90+ :   A fortran-90+ plugin for easy editing](http://www.vim.org/scripts/script.php?script_id=4959)
- [Kerbal Space Program Syntax : Simple syntax highlighting for Kerbal Space Program saves and .craft files.](http://www.vim.org/scripts/script.php?script_id=4960)
- [SmartColumnColor : vim plugin for alert when we reach a defined numbers of characters in a line](http://www.vim.org/scripts/script.php?script_id=4961)
- [FileBeagle : A VINE-spired (Vim Is Not Emacs) file system explorer](http://www.vim.org/scripts/script.php?script_id=4962)
- [OCamlJump : Allow to jum between let and in keyword.](http://www.vim.org/scripts/script.php?script_id=4963)
- [arnoldc.vim : Syntax highlighting for ArnoldC language](http://www.vim.org/scripts/script.php?script_id=4965)
- [Nefertiti : Dark-background Vim Color Scheme with Dynamic Brightness Control (GUI-only) ](http://www.vim.org/scripts/script.php?script_id=4967)
- [CWDPattern : Automatically changes CWD depending on patterns in file path of current buffer](http://www.vim.org/scripts/script.php?script_id=4968)
- [Fist of Vim : Super simple and awesome gist support for vim](http://www.vim.org/scripts/script.php?script_id=4969)
- [rdark-terminal2.vim : Modified rdark-terminal color scheme to enhance visibility](http://www.vim.org/scripts/script.php?script_id=4970)
- [SwapText : Mappings to exchange text with the previously deleted text. ](http://www.vim.org/scripts/script.php?script_id=4971)
- [vim-fanfou : Fanfou client for VIM](http://www.vim.org/scripts/script.php?script_id=4972)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (8771)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2033)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1892)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1233)
5. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (1186)
6. [boshiamy input method keymap file : This is the keymap file for Boshiamy chinese input method](http://www.vim.org/scripts/script.php?script_id=4393) (845)
7. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (842)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (823)
9. [github theme : A gvim colorscheme based on github's syntax highlighting](http://www.vim.org/scripts/script.php?script_id=2855) (660)
10. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (643)

## vim-jp issues

Open : 193 (+2) | Closed : 393 (+14)

- [Issue #571 : vim-users.jp の過去記事掲載について](https://github.com/vim-jp/issues/issues/571)
- [Issue #572 : substitute()での '\\\<\\|\\>'で文字化け](https://github.com/vim-jp/issues/issues/572)
- [Issue #573 : cmd.exe 版で「ﾎﾞ」が入力出来ない。](https://github.com/vim-jp/issues/issues/573)
- [Issue #574 : マルチバイトでのワードの区切り](https://github.com/vim-jp/issues/issues/574)
- [Issue #575 : Perl 5.20 on windows のビルドに失敗する](https://github.com/vim-jp/issues/issues/575)
- [Issue #576 : quickfix を開いた状態でコマンドラインウインドウを開閉すると特定のウインドウが小さくなる](https://github.com/vim-jp/issues/issues/576)
- [Issue #577 : cursor() でカーソルが動かない](https://github.com/vim-jp/issues/issues/577)
- [Issue #578 : Vim起動時にreplace modeになる](https://github.com/vim-jp/issues/issues/578)
- [Issue #579 : set history=100000000と設定するとコマンドラインモードに移行する度に&#x60;E342&#x60;が表示されVimを終了出来ない](https://github.com/vim-jp/issues/issues/579)
- [Issue #580 : matchaddpos()によるハイライトがvim -u NONEで一部残る](https://github.com/vim-jp/issues/issues/580)
- [Issue #581 : $ や h や w で画面右スクロールするとハングする](https://github.com/vim-jp/issues/issues/581)
- [Issue #582 : :options のローカライズ](https://github.com/vim-jp/issues/issues/582)
- [Issue #583 : Update syntax/vim.vim](https://github.com/vim-jp/issues/issues/583)
- [Issue #584 : html用のautoload関数が呼ばれなくてエラー](https://github.com/vim-jp/issues/issues/584)
- [Issue #585 : 例外が投げられる際に Vim script がうまく処理されない](https://github.com/vim-jp/issues/issues/585)
- [Issue #586 : 全角と半角で matchaddpos() の動きが違う](https://github.com/vim-jp/issues/issues/586)

