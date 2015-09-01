---
layout: post
category: vimmagazine
title: Vim Magazine 12 月号
---

## 今月の新機能

- 7.4.542: `:command` コマンドに範囲指定の種類を指定する `-addr` 引数が追加されました。

## リリース情報

- [7.4.538 : tests fail with small features plus Python](https://github.com/vim/vim/commit/0b105416066c95ade3604ec2139d8367d3c6e74e)
- [7.4.539 : (after 7.4.530) crash when computing buffer count](https://github.com/vim/vim/commit/4d84d9325f9537427f58c5ab0c1ce9a68be9d697)
- [7.4.540 : (after 7.4.539) cannot build with tiny and small features](https://github.com/vim/vim/commit/f8a447c6dc51e55d38d72c7a893c7483de112cdb)
- [7.4.541 : crash when doing a range assign](https://github.com/vim/vim/commit/b2a851fee41a5a1faddcb862aef2beca69dc063a)
- [7.4.542 : cannot specify the type of range for a user command](https://github.com/vim/vim/commit/f1d6ccf2f9c8a8ae2c0ec4577946397c103ead2b)
- [7.4.543 : (after 7.4.232) "1,3s/\\n//" joins two lines instead of three](https://github.com/vim/vim/commit/cc2b9d5dc08cefa0342a25ece71b21d4b4b32e00)
- [7.4.544 : warnings for unused arguments with a combination of features](https://github.com/vim/vim/commit/d96c52c71b85d848f23f58dfb8463279a39bfe84)
- [7.4.545 : highlighting for multi-line matches is not correct](https://github.com/vim/vim/commit/aff5c3a5304d85b53b00fd67418d53056b663ba9)
- [7.4.546 : repeated use of vim&#x5f;snprintf() with a number](https://github.com/vim/vim/commit/2f1e51a43c5e0b7560396898bdb7384550f05481)
- [7.4.547 : "vit" does not select multi-byte character at the end correctly](https://github.com/vim/vim/commit/8340dd9b2919299bd5802227fcdf2df706592d41)
- [7.4.548 : compilation fails with native version of MinGW-w64](https://github.com/vim/vim/commit/2d78cd7756ac5c666e85222aea54d45b04b949a4)
- [7.4.549 : function name not recognized correctly when inside a function](https://github.com/vim/vim/commit/ef92390cb7586b8c46bfb06664e25e5e9cfbeb9c)
- [7.4.550 : curs&#x5f;rows() function is always called with 2nd argument false](https://github.com/vim/vim/commit/3f9be97015e2ba4cb6d6654117745d912ec59287)
- [7.4.551 : "ygn" may yank too much; Issue 295](https://github.com/vim/vim/commit/b12db9f92ea74039bb0b4110add4232f0bf338ea)
- [7.4.552 : langmap applies to Insert mode expression mappings](https://github.com/vim/vim/commit/c68c72ea7bbc92ce7cc2c746ce6408dc6cbe06cd)
- [7.4.553 : various small issues](https://github.com/vim/vim/commit/caad4f0a0b428f3dc4e9e395ad0b1cddb38d3bcd)
- [7.4.554 : missing part of patch 7.4.519](https://github.com/vim/vim/commit/6f2dd9e75eab236d1bcc63b17d02f554111f87ec)
- [7.4.555 : test&#x5f;close&#x5f;count may fail for some combination of features](https://github.com/vim/vim/commit/36ff08479f0398e09298f5742dcdfa0f8e070a32)
- [7.4.556 : failed commands in Python interface not handled correctly](https://github.com/vim/vim/commit/ae38d05271f9f25b6ad39b10e5f41d9656a7f1ab)
- [7.4.557 : one more small issue](https://github.com/vim/vim/commit/140e995ed530770408a49f69f61d68f95b07e327)
- [7.4.558 : when the X server restarts Vim may get stuck](https://github.com/vim/vim/commit/527a6785c73306ff72d83cbc2d1ea14f37aca142)
- [7.4.559 : block append inside tab is wrong when virtualedit is set](https://github.com/vim/vim/commit/fc3f23bedfc848cd77f97c62bf4f39a01d46994e)
- [7.4.560 : memory leak using :wviminfo; Issue 296](https://github.com/vim/vim/commit/e88b0033f67b45472006578798794dfd58413a9f)

## 新着スクリプト

- [argwrap.vim : Wrap and unwrap function arguments, lists and dictionaries in Vim](http://www.vim.org/scripts/script.php?script_id=5062)
- [Casetrate : Modify the case of identifiers](http://www.vim.org/scripts/script.php?script_id=5063)
- [FileSync : FileSync offers file synchronization functionality.](http://www.vim.org/scripts/script.php?script_id=5064)
- [filestyle : Check and make (experimental) the opened buffer to correspond Vim settings](http://www.vim.org/scripts/script.php?script_id=5065)
- [Escape Search : escape special characters in your search expressions](http://www.vim.org/scripts/script.php?script_id=5066)
- [vim-packer : basic vim/packer integration](http://www.vim.org/scripts/script.php?script_id=5067)
- [vim-terraform : basic vim/terraform integration](http://www.vim.org/scripts/script.php?script_id=5068)
- [Vim-Swoop : Grep and replace easily in multiple files being context aware !](http://www.vim.org/scripts/script.php?script_id=5069)
- [vim-tab : plugin that maintain each tab different working directory](http://www.vim.org/scripts/script.php?script_id=5070)
- [Spacegray.vim : An elegant dark colorscheme for Vim](http://www.vim.org/scripts/script.php?script_id=5071)
- [vim-nose : Plugin wrapping nosetests command to be executed inside Vim](http://www.vim.org/scripts/script.php?script_id=5072)
- [EntryComplete : Insert mode completion based on lines in designated files or buffers.](http://www.vim.org/scripts/script.php?script_id=5073)
- [grimoire.vim : Access Clojure Grimoire documentation from Vim](http://www.vim.org/scripts/script.php?script_id=5074)
- [cf5-compile.vim : Compile files right from vim.](http://www.vim.org/scripts/script.php?script_id=5075)
- [tex-syntax : Vim TeX syntax file](http://www.vim.org/scripts/script.php?script_id=5076)

## 月間ダウンロードランキング

1. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (5841)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1992)
3. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1743)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (956)
5. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (743)
6. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (711)
7. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (671)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (662)
9. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (656)
10. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (586)

## vim-jp issues

Open : 220 (-6) | Closed : 461 (+18)

- [Issue #670 : cursorline の背景色が通常シンタックスの背景色に負けてしまう。](https://github.com/vim-jp/issues/issues/670)
- [Issue #671 : :h col() が実際の挙動と異なる](https://github.com/vim-jp/issues/issues/671)
- [Issue #672 : unix 用の ctermfg/ctermbg を cmd.exe 版に食わせるとサイケデリックになる](https://github.com/vim-jp/issues/issues/672)
- [Issue #673 : 複数行に渡る検索の hlsearch のハイライトが正しくない](https://github.com/vim-jp/issues/issues/673)
- [Issue #674 : mingw64 でビルドに失敗する](https://github.com/vim-jp/issues/issues/674)
- [Issue #675 : Windows で長いファイル名のファイルがあるフォルダに対して、ワイルドカードを使うと落ちることがある](https://github.com/vim-jp/issues/issues/675)
- [Issue #676 : &#x60;call setreg('"', \[\])&#x60; で E341](https://github.com/vim-jp/issues/issues/676)
- [Issue #677 : mapgroup 欲しい？](https://github.com/vim-jp/issues/issues/677)
- [Issue #678 : xlose/backlog の導入](https://github.com/vim-jp/issues/issues/678)
- [Issue #679 : pathとincludeを設定して、suffixesaddを設定してincludeexprを設定したのに](https://github.com/vim-jp/issues/issues/679)
- [Issue #680 : cmd.exe版で「β」が入力出来ない。](https://github.com/vim-jp/issues/issues/680)
- [Issue #681 : \[WIP\]Windows のコンソールでも「森鷗外」を表示したい。](https://github.com/vim-jp/issues/issues/681)

