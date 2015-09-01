---
layout: post
category: vimmagazine
title: Vim Magazine 11 月号
---

## 話題

- [Vim Advent Calendar 2014](http://qiita.com/advent-calendar/2014/vim)

  25日で終わる(笑) の Vim Advent Calendar 2014 です。

## 今月の新機能

- 7.4.502 : `'langnoremap'` オプションが追加されました。
- 7.4.503 : `writefile()` に 'a' (append) フラグが追加されました。

## リリース情報

- [7.4.489 : cursor movement wrong when 'lbr' with a number column](https://github.com/vim/vim/commit/fe3c41009887331dc9fee2bedf457728117be544)
- [7.4.490 : cannot specify the buffer to use for "do" and "dp"](https://github.com/vim/vim/commit/6a64365c952895ecc2219b303d9bf64dabf8f2e7)
- [7.4.491 : negative "topline" value for winrestview() has display errors](https://github.com/vim/vim/commit/b851a96d5c201291d9e955afaf55979b3e410f9f)
- [7.4.492 : Ctrl-O in Insert mode moves right after insert comment leader](https://github.com/vim/vim/commit/2f31e399780f434c9635a1da969020824440e03b)
- [7.4.493 : a TextChanged autocommand is triggered when saving a file](https://github.com/vim/vim/commit/086329d3f6b1182b6817433a1c498e1e746d2a3b)
- [7.4.494 : cursor shape is wrong after a CompleteDone autocommand](https://github.com/vim/vim/commit/eda73607a73fe5ff07eb52981df19c5bae129f70)
- [7.4.495 : XPM isn't used correctly in the Cygwin Makefile](https://github.com/vim/vim/commit/666578b53964cc34f2e9fc65c89612f063a5f3c5)
- [7.4.496 : many lines are both in Make&#x5f;cyg.mak and Make&#x5f;ming.mak](https://github.com/vim/vim/commit/4f7e821f26019c14f4470deb0867c919548d5cd5)
- [7.4.497 : NFA engine is very slow with some regexp patterns](https://github.com/vim/vim/commit/fda3729a064d6466cec6ac83dd1bfcc437ea4cc9)
- [7.4.498 : (after 7.4.497) typo in DOS makefile](https://github.com/vim/vim/commit/9c7d98d7c2b81e6e0b4d45389c56fc0e922a5bfe)
- [7.4.499 : substitute() can be slow with long strings](https://github.com/vim/vim/commit/e90c853ff6dcef63bff80be0e7379a80c276af8f)
- [7.4.500 : test 72 still fails once in a while](https://github.com/vim/vim/commit/72971a26119ce781161fb2a281fcaf27880a909e)
- [7.4.501 : (after 7.4.497) typo in file pattern](https://github.com/vim/vim/commit/c3940c76e8248ea7f618b3f1716d754c8e981f35)
- [7.4.502 : language mapping also applies to mapped characters](https://github.com/vim/vim/commit/4391cf98ec3b94f33dfd053cab25ed56c787bea9)
- [7.4.503 : cannot append a list of lines to a file](https://github.com/vim/vim/commit/6b2e938f1307f840165d9049d743161b01af811c)
- [7.4.504 : MS-Windows installer restricts install path to end in "vim"](https://github.com/vim/vim/commit/0434482e5aa5a46b7da27f745c22ec36d35f2c7e)
- [7.4.505 : problem on MS-Windows file name longer than MAX&#x5f;PATH bytes](https://github.com/vim/vim/commit/f9e6c3b8b00ea562e0b9ddf96f85c365cdf41627)
- [7.4.506 : MS-Windows: Cannot open a file with 259 characters](https://github.com/vim/vim/commit/374bf024d23d02064fa3da94d44af71a61588783)
- [7.4.507 : (after 7.4.496) building with MingW and Perl may fail](https://github.com/vim/vim/commit/345326aa3d7edfece7be57d690aa667482a961af)
- [7.4.508 : when generating ja.sjis.po, header is not correctly adjusted](https://github.com/vim/vim/commit/983c4e93dc560b00c5c7b752a3a7c1c452359c38)
- [7.4.509 : users are not aware their encryption is weak](https://github.com/vim/vim/commit/3a0c90898338c04bde314bede9f116f299eaddef)
- [7.4.510 : "-fwrapv" argument breaks use of cproto](https://github.com/vim/vim/commit/b103138b1e9ff2330901a0bc06378d2e5ed4f9ae)
- [7.4.511 : generating proto for if&#x5f;ruby.c uses type not defined elsewhere](https://github.com/vim/vim/commit/3e9a1614966926a6067c977921d5aa270eba5da7)
- [7.4.512 : cannot generate prototypes for Win32 files and VMS](https://github.com/vim/vim/commit/b8e0bdbda377d0a3c3e24bea62503b9f95f8249a)
- [7.4.513 : crash because reference count is wrong for list from getreg()](https://github.com/vim/vim/commit/42d84f8aed604e32e198a5fe31b3b7ab149ea63e)
- [7.4.514 : (after 7.4.492) memory access error](https://github.com/vim/vim/commit/ef6875be7a0b4dde9589e69b0a0229c4b976c45b)
- [7.4.515 : in a help buffer the global 'foldmethod' is used](https://github.com/vim/vim/commit/54fb43813443554c685306b92166218c53cdb18f)
- [7.4.516 : issue 253: completing function name with a # does not work](https://github.com/vim/vim/commit/8a349ff944dc01ca0e52f10144d8660aab430c9c)
- [7.4.517 : cursor may not end up in the right place on a wrapping line](https://github.com/vim/vim/commit/a365091dbb94e0daea5ea58d1facb1102868cca5)
- [7.4.518 : using status line height in width computations](https://github.com/vim/vim/commit/df46f6f07f90c0014352ff05aecb8caf61e7cfa0)
- [7.4.519 : (after 7.4.497) crash when using syntax highlighting](https://github.com/vim/vim/commit/dffa5b8ecfa2091495b77ab487c07792646d0633)
- [7.4.520 : Sun PCK locale is not recognzed](https://github.com/vim/vim/commit/57bc4633b24be82697d969c865bb018cfa47099f)
- [7.4.521 : when using "vep" a mark is moved to the next line](https://github.com/vim/vim/commit/06e7ce1a013535b11d5454630cc1770ef21ea461)
- [7.4.522 : specifying wrong buffer size for GetLongPathName()](https://github.com/vim/vim/commit/4c7b2f50fd6590c9b98d896978515edf0bf142bb)
- [7.4.523 : copy/paste no longer works is X11 server is restarted](https://github.com/vim/vim/commit/b1e265045cc872f2e1cf43e6a98bc3e3770a289d)
- [7.4.524 : when using ":ownsyntax" spell checking is messed up](https://github.com/vim/vim/commit/2683c8e7f7d8eb6222653d55849a9c556de27f45)
- [7.4.525 : map() leaks memory when there is an error in the expression](https://github.com/vim/vim/commit/b738c9a4421e6f8e799937fb18c13c7c6eb6c1a6)
- [7.4.526 : matchstr() fails on long text](https://github.com/vim/vim/commit/66a3e797a285a546348d25d3b7cba315b299a3de)
- [7.4.527 : still confusing regexp failure and NFA&#x5f;TOO&#x5f;EXPENSIVE](https://github.com/vim/vim/commit/8c731505b094bdad100484373387b4c4658300d5)
- [7.4.528 : crash when using matchadd()](https://github.com/vim/vim/commit/cbdf0a0b4adc19fb443bb7df542578c3e76970ce)
- [7.4.529 : no test for what 7.4.517 fixes](https://github.com/vim/vim/commit/baf0344ed7031369a0f355beb062eb4c34e35261)
- [7.4.530 : many commands take count/range that is not using line numbers](https://github.com/vim/vim/commit/b96a7f325c9047d6a65424d40e493d3e0eff26b5)
- [7.4.531 : comments about parsing an Ex command are wrong](https://github.com/vim/vim/commit/1c40a66badd83e9a93116c006097033725c9eb7d)
- [7.4.532 : when using 'incsearch' "2/pattern/e" highlights first match](https://github.com/vim/vim/commit/a3dfccc6e9a780570166b00c0e40c4b9a149e253)
- [7.4.533 : ":hardcopy" leaks memory in case of errors](https://github.com/vim/vim/commit/0a383969378a328ffa363eee0d9a4c3d3c8257fe)
- [7.4.534 : warnings when compiling if&#x5f;ruby.c](https://github.com/vim/vim/commit/901407491382198e1457f3c315b2aedc1d24bda9)
- [7.4.535 : (after 7.4.530) can't build with tiny features](https://github.com/vim/vim/commit/f240e181d720c8a056ad57dfa160b0e8c1371c3a)
- [7.4.536 : test 63 fails when using a black&white terminal](https://github.com/vim/vim/commit/f1e30ccb822c56693881da199c7efbe6c9c0dca1)
- [7.4.537 : value of v:hlsearch reflects an internal variable](https://github.com/vim/vim/commit/76440e2efe9e3582e9943d01183ec8b2f3e4eb4b)

## 新着スクリプト

- [VIntSearch : Vim Integrated Search - Awesome file search plugin using ctags and grep](http://www.vim.org/scripts/script.php?script_id=5047)
- [RootIgnore : Set 'wildignore' from git repo root](http://www.vim.org/scripts/script.php?script_id=5048)
- [ab-prefix : Conditionally expand abbreviations on a matching prefix.](http://www.vim.org/scripts/script.php?script_id=5049)
- [numbers.vim : plugin for intelligently toggling line numbers.](http://www.vim.org/scripts/script.php?script_id=5050)
- [directionalWindowResizer : Adjust window size with Ctrl +jklh](http://www.vim.org/scripts/script.php?script_id=5051)
- [shot-f : highlights the characters where the cursor can move directly (by f,F,t,T)](http://www.vim.org/scripts/script.php?script_id=5052)
- [rengbang : sequencial numbering with pattern](http://www.vim.org/scripts/script.php?script_id=5053)
- [atom-dark : Dark theme inspired by Atom's default dark theme](http://www.vim.org/scripts/script.php?script_id=5054)
- [vim-www : Define your favourite websites and search engines and open them from vim](http://www.vim.org/scripts/script.php?script_id=5055)
- [css&#x5f;color : CSS colour preview while editing](http://www.vim.org/scripts/script.php?script_id=5056)
- [buftabline : Use the tabline to render buffer tabs](http://www.vim.org/scripts/script.php?script_id=5057)
- [emoji&#x5f;complete : A plug-in for insert completion of emoji.](http://www.vim.org/scripts/script.php?script_id=5058)
- [vim-asterisk : &#x2a;-improved](http://www.vim.org/scripts/script.php?script_id=5059)
- [jspc.vim : JavaScript Parameter Complete](http://www.vim.org/scripts/script.php?script_id=5060)
- [yoda.vim : A VIM plug-in for completion using Clang.](http://www.vim.org/scripts/script.php?script_id=5061)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (5519)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2171)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1963)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (997)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (952)
6. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (779)
7. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (698)
8. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (692)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (676)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (654)

## vim-jp issues

Open : 226 (+0) | Closed : 443 (+12)

- [Issue #658 : writefile() で append したい。](https://github.com/vim-jp/issues/issues/658)
- [Issue #659 : 特定の条件で関数内の関数が関数として認識されない](https://github.com/vim-jp/issues/issues/659)
- [Issue #660 : :unmap が {rhs} に設定されているものを消去してしまう](https://github.com/vim-jp/issues/issues/660)
- [Issue #661 : カーソルキーの左、以降しかドットレジスタに記録されない](https://github.com/vim-jp/issues/issues/661)
- [Issue #662 : Vim の補完で入力に一致した候補が出てくる](https://github.com/vim-jp/issues/issues/662)
- [Issue #663 : GVim で \<Left> にマッピングを定義すると日本語入力がおかしくなる](https://github.com/vim-jp/issues/issues/663)
- [Issue #664 : jquery-latest.min.js を開くと落ちる](https://github.com/vim-jp/issues/issues/664)
- [Issue #665 : 折り返し時、textwidthの計算にインデントやshowbreakの文字数が含まれる](https://github.com/vim-jp/issues/issues/665)
- [Issue #666 : ｷﾘ番ｹﾞｯﾄ!](https://github.com/vim-jp/issues/issues/666)
- [Issue #667 : C の #define 文で 0 が現れるとハイライトが壊れる](https://github.com/vim-jp/issues/issues/667)
- [Issue #668 : 7.4.497 以降でregexengine=0の時にSEGVが起こる](https://github.com/vim-jp/issues/issues/668)
- [Issue #669 : syntax/netrw.vimにscriptencoding指定されていない。](https://github.com/vim-jp/issues/issues/669)

