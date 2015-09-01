---
layout: post
category: vimmagazine
title: Vim Magazine 9 月号
---

## 話題

- [momonga.vim #2](http://partake.in/events/55ef537f-07fb-4e36-8f27-666f7925a4d7)

  みんなでvimの便利を便利していきます 10月13日(日) 東京

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 まだやってます。

## リリース情報

- [7.4.017 : ":help !!" does not find the "!!" tag in the help file](https://github.com/vim/vim/commit/966d71c852e1126883846817d3a16590268298f5)
- [7.4.018 : when completing item becomes unselected](https://github.com/vim/vim/commit/f0d9e005375409a9f3a887a1a946afb9a45cdd34)
- [7.4.019 : file name completion fails with Chinese characters](https://github.com/vim/vim/commit/00b764aa4115271aa553c8443adc6cec6859e69e)
- [7.4.020 : NFA engine matches too much with \\@>](https://github.com/vim/vim/commit/78c93e4b6a87720a50ac14c40aa077ed6b86145b)
- [7.4.021 : NFA regexp: Using \\ze may result in wrong end](https://github.com/vim/vim/commit/9be44818921a9255b3a26bf5566f1939bf38804e)
- [7.4.022 : deadlock while exiting, because of allocating memory](https://github.com/vim/vim/commit/bec9c20884290cc64e63d4eec531c0a2fbdcd3f6)
- [7.4.023 : MS-Windows: warning for 64 bit type](https://github.com/vim/vim/commit/3b7ac49e2672263840b5854402e7d24a7f3a8a23)
- [7.4.024 : current user cannot always use the undo file he created](https://github.com/vim/vim/commit/3b262393419ddf73e0bce95e0d896c46a95236d9)
- [7.4.025 : reading before start of a string](https://github.com/vim/vim/commit/0300e465aa9b034455f2c98d9996d5a3b04e9900)
- [7.4.026 : clang warning for int shift overflow](https://github.com/vim/vim/commit/95235e64d8329b8c0fbd9311d98626afe86ad911)
- [7.4.027 : valgrind error when using CTRL-X CTRL-F at start of the line](https://github.com/vim/vim/commit/dd4073480011fe1cc58408a33963154fcb41673d)
- [7.4.028 : equivalence classes are not working for multi-byte characters](https://github.com/vim/vim/commit/e6a2fa65732eb068b4d10f65e0a3f7353d3dffdc)
- [7.4.029 : an error in a pattern is reported twice](https://github.com/vim/vim/commit/917789fdf702b18edd9788adf5a60123b9a31ce8)
- [7.4.030 : the -mno-cygwin argument is no longer supported by Cygwin](https://github.com/vim/vim/commit/9952c102dd66e14101b1abbe6510dde973f96945)
- [7.4.031 : ":diffoff!" resets options even when 'diff' is not set](https://github.com/vim/vim/commit/00462ffb7d0fc6448eb5d10681ef4e1f812c7310)
- [7.4.032 : NFA engine does not match the NUL character](https://github.com/vim/vim/commit/595cad2ea1c77c40a34a240c71290fdef6aabd3b)
- [7.4.033 : if terminal has 20 lines test 92 and 93 overwrite input file](https://github.com/vim/vim/commit/6848c8b561e6c35a3e646c99a82e03c4aa5def92)
- [7.4.034 : using "p" in Visual block mode only changes the first line](https://github.com/vim/vim/commit/ec11aef914d3f1570c40e9795c145fffaf94edd4)
- [7.4.035 : MS-Windows: mouse pointer flickers when going to Normal mode](https://github.com/vim/vim/commit/268b55b5fbe68ca24020391c3c3055f71cff7277)
- [7.4.036 : NFA engine does not capture group correctly when using \\@>](https://github.com/vim/vim/commit/699c12076de7d8811045cd2b98df78d7108b68ed)
- [7.4.037 : "\\ze" in a sub-pattern does set the end of the match](https://github.com/vim/vim/commit/f211884fa151a3c39b1a85a115d6d88ce85bbe54)
- [7.4.038 : Using "zw" and "zg" when 'spell' is off give confusing error](https://github.com/vim/vim/commit/134bf07ca0e28addeeb67edc4fceeba00388d7fc)
- [7.4.039 : MS-Windows: \< MSCV10 can't handle symlinks to a directory](https://github.com/vim/vim/commit/2ee95f7d13566ba393ed70bd4401e7164b0cc2f9)
- [7.4.040 : error on exit when variable holds reference to a script scope](https://github.com/vim/vim/commit/cca74134564a99f638a583d79d2cfd3ca393673d)
- [7.4.041 : Visual selection does not remain after being copied over](https://github.com/vim/vim/commit/033d8888e8d6648b5540dc3dda85aba3c5a192de)
- [7.4.042 : after ":setlocal" 'spell'/'spellang' :spelldump doesn't work](https://github.com/vim/vim/commit/7a18fdc8a4d3c1408f54074ff9268735a277093a)
- [7.4.043 : VMS can't handle long function names](https://github.com/vim/vim/commit/386dcde30d677c26a596cf7479f8dd9ee498978b)
- [7.4.044 : can't build with old MSVC](https://github.com/vim/vim/commit/8962fda1dacf374d78bd02ca487800c2519c57b2)
- [7.4.045 : substitute() has a problem with pattern starting with "\\ze".](https://github.com/vim/vim/commit/ceb84af027ac782f0b9de0e66f321258358e391a)

## 新着スクリプト

- [FindYamlKey : Search for nested keys in a YAML document](http://www.vim.org/scripts/script.php?script_id=4712)
- [reco : backup and recovery solution for Vim](http://www.vim.org/scripts/script.php?script_id=4713)
- [KillBrackets : Removes matching brackets/parens. Very handy in LISP/Scheme](http://www.vim.org/scripts/script.php?script_id=4714)
- [Hardcore Mode : A mode that helps get rid of the habit of using evil cursor keys, mouse, and GUI](http://www.vim.org/scripts/script.php?script_id=4715)
- [files.vim : List files for partial search using 'ls' command](http://www.vim.org/scripts/script.php?script_id=4716)
- [svnim : svn plugin](http://www.vim.org/scripts/script.php?script_id=4717)
- [SuperSub : Enhanced Substitute Command](http://www.vim.org/scripts/script.php?script_id=4718)
- [yankitute : Regex powered yank+substitute](http://www.vim.org/scripts/script.php?script_id=4719)
- [vim-django-i18n : transform text to django i18n format](http://www.vim.org/scripts/script.php?script_id=4720)
- [vim-polyglot : A collection of language packs for Vim.](http://www.vim.org/scripts/script.php?script_id=4721)
- [DuplicateWrite :  Cascade the writing of a file to another location.](http://www.vim.org/scripts/script.php?script_id=4722)
- [fontdetect : Detect installed fonts](http://www.vim.org/scripts/script.php?script_id=4723)
- [Ultimate Test Utility : A straightforward, fully documented unit testing utility.](http://www.vim.org/scripts/script.php?script_id=4724)
- [vim-emoji : Emoji in Vim](http://www.vim.org/scripts/script.php?script_id=4725)
- [Extended-FT : Extended functionality for 'f' and 't' commands](http://www.vim.org/scripts/script.php?script_id=4726)
- [stringliteral : C++ multi-line string literals from common text](http://www.vim.org/scripts/script.php?script_id=4727)
- [Tags Updater : Simple tags automatic updater](http://www.vim.org/scripts/script.php?script_id=4728)
- [VerticalHelp : Launch help in vertical window for wide-screen displays](http://www.vim.org/scripts/script.php?script_id=4729)
- [pastek.vim : Syntax for Pastek files](http://www.vim.org/scripts/script.php?script_id=4730)
- [cmake.vim : Incorporates CMake support to Vim.](http://www.vim.org/scripts/script.php?script_id=4731)
- [integrake : Abuse Rake to do project chores around Vim](http://www.vim.org/scripts/script.php?script_id=4732)
- [Sol : Simple colorscheme created for fun and day usage](http://www.vim.org/scripts/script.php?script_id=4733)
- [united-front : Autoshare yank registers across vim instances.](http://www.vim.org/scripts/script.php?script_id=4734)

## 月間ダウンロードランキング

1. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2037)
2. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2019)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1312)
4. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1209)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (1033)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (894)
7. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (742)
8. [github theme : A gvim colorscheme based on github's syntax highlighting](http://www.vim.org/scripts/script.php?script_id=2855) (661)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (641)
10. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (617)

## vim-jp issues

Open : 174 (-1) | Closed : 297 (+7)

- [Issue #466 : undoを実行したときのメッセージが表示されない場合がある](https://github.com/vim-jp/issues/issues/466)
- [Issue #467 : input()で何も入力なしで\<CR>叩いた場合と\<Esc>した場合を区別したい](https://github.com/vim-jp/issues/issues/467)
- [Issue #468 : Windows でシンボリックリンクのディレクトリに :cd できない](https://github.com/vim-jp/issues/issues/468)
- [Issue #469 : normal コマンドから実行したスクリプトで input() が動作しない](https://github.com/vim-jp/issues/issues/469)
- [Issue #470 : :\[range\]source はどうだろうか?](https://github.com/vim-jp/issues/issues/470)
- [Issue #471 : compatible について解説したほうが良いか?](https://github.com/vim-jp/issues/issues/471)

