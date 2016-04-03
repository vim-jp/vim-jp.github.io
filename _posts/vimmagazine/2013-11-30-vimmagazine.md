---
layout: post
category: vimmagazine
title: Vim Magazine 11 月号
---

## 話題

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 ついに完結。

- [Vim Advent Calendar 2013](http://atnd.org/events/45072)

  Vim Advent Calencar 2013 が始まりました。

- [TokyoVim#16](http://partake.in/events/c17663f3-cb0f-4bfe-8610-96177a8edf66)

  Vimに関する何かをやる会です。  12月7日(土) 東京


## 今月の新機能

  - 7.4.057: `byteidxcomp({expr}, {nr})` が追加されました。

  - 7.4.073: 'undolevels' オプションが global-local になりました。

  - 7.4.079: `v:hlsearch` が追加されました。

  - 7.4.083: `:keeppatterns` が追加されました。

  - 7.4.089: 'spelllang' オプションに "cjk" フラグが追加されました。スペルチェック有効時に日本語がスペルエラーとして表示されるのを防ぎます。


## リリース情報

- [7.4.053 : test75 has a wrong header](https://github.com/vim/vim/commit/b69a714bb75e25881eb3d4ce116894272535202b)
- [7.4.054 : reading past end of the 'stl' string](https://github.com/vim/vim/commit/1180822670e987ef204dfbf24e55fc09eed89650)
- [7.4.055 : Mac: Where macros are defined depends on the system](https://github.com/vim/vim/commit/2b1b53f5d9674c6fc41eab86178e3dcfd948d297)
- [7.4.056 : Mac: Compilation problem with OS X 10.9 Mavericks](https://github.com/vim/vim/commit/4cc95d1495706593d30ce2fae3fc592a2ffdaab8)
- [7.4.057 : byteidx() does not work for composing characters](https://github.com/vim/vim/commit/0ffbbf9ea871dd64c2be933d6718f0b34fe960c4)
- [7.4.058 : warnings on 64 bit Windows](https://github.com/vim/vim/commit/92c2db8be669b9e5038a5a9b8e02ca8856747d77)
- [7.4.059 : set&#x5f;last&#x5f;cursor() may encounter w&#x5f;buffer being NULL](https://github.com/vim/vim/commit/9db1293317abd1aca101bf2d35caef0354f66e8b)
- [7.4.060 : declaration has wrong return type for PyObject&#x5f;SetAttrString()](https://github.com/vim/vim/commit/0b400087d076ba1a476d753c019fbfd96d8ae58e)
- [7.4.061 : (after 7.4.056) configure check in the wrong place](https://github.com/vim/vim/commit/39766a7595461b89d1aebff72ac3853225e6a4c5)
- [7.4.062 : (after 7.4.061) wrong configure check for AvailabilityMacros.h](https://github.com/vim/vim/commit/18e5469b6e73dfbd85ac9831224cd4ec29cbe308)
- [7.4.063 : crash when using invalid key in Python dictionary](https://github.com/vim/vim/commit/ba2d7ffc4bd69175ddeffb77a7b459787baa0d92)
- [7.4.064 : in Visual block mode replacing with CR does not break lines](https://github.com/vim/vim/commit/d9820538bd6c83c51d24e3481cadf79ccaa9549e)
- [7.4.065 : the character typed at the hit-enter prompt is recorded twice](https://github.com/vim/vim/commit/332a2cadcd19a0ac69b9c41ed6f4add89128fdfe)
- [7.4.066 : MS-Windows: swap file name wrong if colon in file name](https://github.com/vim/vim/commit/69c3500f9630db9b194cb82167d0fccef9efb9a1)
- [7.4.067 : CTRL-\\ CTRL-O moves the cursor after inserting comment leader](https://github.com/vim/vim/commit/f332a656c3aa26ee95849341198eb730a150445d)
- [7.4.068 : cannot build Vim on Mac with non-Apple compilers](https://github.com/vim/vim/commit/0958e0fbe7307f0b46b8f692cbd097fbf93c90f6)
- [7.4.069 : cannot right shift comment lines starting with #](https://github.com/vim/vim/commit/6bcbcc59be58d0c3b3cd53ac105c6eb7d0b87f06)
- [7.4.070 : (after 7.4.069) can't compile with tiny features](https://github.com/vim/vim/commit/dce7c91de95ff565a18535390a9abd4fe0a4a3f1)
- [7.4.071 : (after 7.4.069) passing limits around too often](https://github.com/vim/vim/commit/84dbb62fca26784eb706f5309f9285464d1a2989)
- [7.4.072 : crash when using Insert mode completion](https://github.com/vim/vim/commit/4c1e626103397f8ac84ab7ca29cc74d0b3cc06e6)
- [7.4.073 : setting undolevels for one buffer changes undo in another](https://github.com/vim/vim/commit/f5a2fd880ae8f6225814209ab73783f65078a4d5)
- [7.4.074 : when undo'ing all and making a change, undo structure is wrong](https://github.com/vim/vim/commit/aa88732efd21347587b8b8ce0f37a4042e668f42)
- [7.4.075 : locally setting 'undolevels' is not tested](https://github.com/vim/vim/commit/c1c023ecfc19bc4ef3e24a43cbee3f4f54f5706b)
- [7.4.076 : "cgn" does not wrap around the end of the file](https://github.com/vim/vim/commit/fcea03ddd3dc54c8a0a786af5e75bbfbcf8fb3b7)
- [7.4.078 : MSVC 2013 is not supported](https://github.com/vim/vim/commit/e5878f4be9f0cb6719220fa17f113fa24b1f3039)
- [7.4.079 : a script cannot detect whether 'hlsearch' highlighting is on](https://github.com/vim/vim/commit/8050efa07d7a66e8e2e88253d079a106a7f2601a)
- [7.4.080 : (after 7.4.079) missing documentation for v:hlsearch](https://github.com/vim/vim/commit/817a8801d7d0edb0c2776c22ef0dd5ac2b2d73c8)
- [7.4.081 : (after 7.4.078) wrong logic when ANALYZE is "yes"](https://github.com/vim/vim/commit/815135e4082dce22a264cc1f176187221d5313cf)
- [7.4.082 : using "gf" in a changed buffer suggests adding "!"](https://github.com/vim/vim/commit/45d3b1454c1001001ef97746556d3142788a8383)
- [7.4.083 : hard to avoid adding a used pattern to the search history](https://github.com/vim/vim/commit/a939e434407b749adf4d50ea9f3f13b6a9abaf86)
- [7.4.084 : Python: interrupt not being properly discarded](https://github.com/vim/vim/commit/d6b8a5253b293b90a90af4320e9fd1c6e587ad2b)
- [7.4.085 : can't move cursor when inserting text in Visual block mode](https://github.com/vim/vim/commit/3f75e422be7dba58a6c6fc58dcc51b882bd4adaf)
- [7.4.086 : can't skip over expression when not evaluating for dict member](https://github.com/vim/vim/commit/0f8de8dfb2cc7b778457d9c84305453bd1ef3600)
- [7.4.087 : compiler warning on 64 bit Windows systems](https://github.com/vim/vim/commit/efe06f4dd8713e5a8dc8c537b0fdf89101a87c20)
- [7.4.088 : Asian characters are always marked as spell errors](https://github.com/vim/vim/commit/cc63c647eee82ca4f3a9c7b09c1bd86933e913d4)
- [7.4.089 : Vim doesn't set the security context on a renamed file](https://github.com/vim/vim/commit/0671de335f87c15aa989e99952098feea1285830)
- [7.4.090 : Win32: no completion if dir name contains an exclamation mark](https://github.com/vim/vim/commit/8f5610df7312023e2da5dc8f885e1b159a57dee2)
- [7.4.091 : (after 7.4.089) missing semicolon](https://github.com/vim/vim/commit/e874744e94e172d33edd91ccb0f5cdf5105b6be3)
- [7.4.092 : (after 7.4.088) can't build small version](https://github.com/vim/vim/commit/2593e037c9e9a00ebd86fa3f8a8011c39511c314)
- [7.4.093 : configure can't use LuaJIT on ubuntu 12.04](https://github.com/vim/vim/commit/9421bdc2cd7cdcd801bcde4d41ea0a2c60b66199)
- [7.4.094 : configure may not find that -lint is needed for gettext()](https://github.com/vim/vim/commit/49b6a57d19e46258c2b4c7fb71a5e78f6a7840d1)
- [7.4.095 : (after 7.4.093) regexp for LuaJIT version doesn't work on BSD](https://github.com/vim/vim/commit/49b10273f0dd1b10e94ea675b18088a27364edbd)
- [7.4.096 : can't change directory to an UNC path](https://github.com/vim/vim/commit/ffa2220950a8f2856f4d7042768888c27e51178a)
- [7.4.097 : unexpected behavior change related to 'virtualedit'](https://github.com/vim/vim/commit/1e42f7ac1624058f8097bc8d6406af21b600b8f5)
- [7.4.098 : error for line numbers out of range when using ":'\<,'>del"](https://github.com/vim/vim/commit/d07c6e1e82d4321da0fb11ff408aeaeb9fba1f68)
- [7.4.099 : append in blockwise Visual mode with "$" is wrong](https://github.com/vim/vim/commit/cafaa8a9502f64d5c23e51c1f89c5b322deb22fe)
- [7.4.100 : NFA regexp doesn't handle backreference correctly](https://github.com/vim/vim/commit/272fb5823627e2e088b8f39ea516fc8a9341572e)
- [7.4.101 : using \\1 in pattern goes one line too far](https://github.com/vim/vim/commit/438ee5bf603f3cee2c8ec924d081885c0086ba19)
- [7.4.102 : crash when interrupting "z="](https://github.com/vim/vim/commit/ace8d8ee2f5d6ddd6d47272609c75155ce533f97)
- [7.4.103 : dos installer escapes spaces in the diff command wrong](https://github.com/vim/vim/commit/c62a644eaf31ff1ea6e3d569d0fe8c9d4eb3d908)
- [7.4.104 : ":help s/\\&#x5f;" reports an internal error](https://github.com/vim/vim/commit/174a848b677c6cd66def00d36a2b25332ca49e38)
- [7.4.105 : completing a tag pattern may give an error for invalid pattern](https://github.com/vim/vim/commit/dff54ac7e85f14f91faeadbaff22b23b2e5efb84)
- [7.4.106 : can't build with Ruby using Cygwin](https://github.com/vim/vim/commit/e88a5f3a2c18426a68c4591d61bc612276878b46)
- [7.4.107 : Python try/catch doesn't catch Vim error in vim.eval()](https://github.com/vim/vim/commit/9fee7d4729af19e7ce4950ede8de358c5eeb3772)
- [7.4.108 : "zG" and "zW" leave temp files around on MS-Windows](https://github.com/vim/vim/commit/34b466edfe6acdf610899da18e81537a8e5d52c8)
- [7.4.109 : ColorScheme autocommand matches with the current buffer name](https://github.com/vim/vim/commit/b95186fd36bae9716181bfdd203fb27f410e3c4f)
- [7.4.110 : "gUgn" cannot be repeeated](https://github.com/vim/vim/commit/ba2d44f33863e115b1858ab572829f403ad21883)

## 新着スクリプト

- [vim-airline-tomato : Pomodoro Technique](http://www.vim.org/scripts/script.php?script_id=4756)
- [SlateDark : Dark vim color scheme inspired by Slate.](http://www.vim.org/scripts/script.php?script_id=4757)
- [Cosco : Colon and semicolon insertion bliss](http://www.vim.org/scripts/script.php?script_id=4758)
- [minimal&#x5f;gdb : lightweight vim->gdb broker. set and auto export breakpoints to gdb](http://www.vim.org/scripts/script.php?script_id=4759)
- [vim-gnote : keep your note in your gmail special box](http://www.vim.org/scripts/script.php?script_id=4760)
- [HowMuch : A Vim calculator for visual selections](http://www.vim.org/scripts/script.php?script_id=4761)
- [vim-sunflower : Switch colorscheme depending on whether it is day or night.](http://www.vim.org/scripts/script.php?script_id=4762)
- [cppnav : C/C++ File navigation - using omnicppcomplete](http://www.vim.org/scripts/script.php?script_id=4763)
- [NERV-ous : Evangelion-inspired dark colorscheme](http://www.vim.org/scripts/script.php?script_id=4764)
- [axiom.vim : Syntax files for Axiom computer algebra system files](http://www.vim.org/scripts/script.php?script_id=4765)
- [sovim : Goes up on directory structures and autosources Vimscript file.](http://www.vim.org/scripts/script.php?script_id=4766)
- [tabwins.vim : Create Tabs of symmetric or asymmetric Windows, empty or populated](http://www.vim.org/scripts/script.php?script_id=4767)
- [natbib.vim : Syntax support for the LaTeX package natbib](http://www.vim.org/scripts/script.php?script_id=4768)
- [airlineish : Custom theme for vim-airline (works perfect with dark background / UI)](http://www.vim.org/scripts/script.php?script_id=4769)
- [noerrmsg.vim : Hides the ErrorMsg while in insert mode](http://www.vim.org/scripts/script.php?script_id=4770)
- [maven errorformat : errorformat for maven3.10 compiler plugin2.5.1](http://www.vim.org/scripts/script.php?script_id=4771)
- [marslo.vim : A dark-background color scheme made by Marslo](http://www.vim.org/scripts/script.php?script_id=4772)
- [grishin-color-scheme : A koehler-similar scheme](http://www.vim.org/scripts/script.php?script_id=4773)
- [p4checkout : Automatically check out RO files from Perforce on edit](http://www.vim.org/scripts/script.php?script_id=4774)
- [tex-breqn.vim : LaTeX breqn package syntax highlighting support](http://www.vim.org/scripts/script.php?script_id=4775)
- [Projx : Light weight project explorer](http://www.vim.org/scripts/script.php?script_id=4776)
- [repeatableMapping : Set up mappings that can be repeated via repeat.vim.](http://www.vim.org/scripts/script.php?script_id=4777)
- [QFEnter : A vim plugin for intuitive file opening from Quickfix window.](http://www.vim.org/scripts/script.php?script_id=4778)
- [vim-slamhound : Slamhound integration for vim.](http://www.vim.org/scripts/script.php?script_id=4779)
- [bashizzle.vim : deadsimple snippet/templating using bash](http://www.vim.org/scripts/script.php?script_id=4780)
- [RPN : Vim RPN calculator](http://www.vim.org/scripts/script.php?script_id=4781)
- [v2ex : v2ex for vim.](http://www.vim.org/scripts/script.php?script_id=4782)
- [npm.vim : Run NPM commands in Vim.](http://www.vim.org/scripts/script.php?script_id=4783)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (2710)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2429)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1472)
4. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1310)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (967)
6. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (947)
7. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (916)
8. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (878)
9. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (866)
10. [tSkeleton : File Templates and Code Skeletons/Snippets](http://www.vim.org/scripts/script.php?script_id=1160) (807)

## vim-jp issues

Open : 177 (-5) | Closed : 314 (+16)

- [Issue #481 : autocmd.txtのautocmd一覧にTextChangedがない](https://github.com/vim-jp/issues/issues/481)
- [Issue #482 : 行末が揃ってないバッファで矩形選択で$Aして挿入すると余計な文字が挿入される](https://github.com/vim-jp/issues/issues/482)
- [Issue #483 : zG で作成された一時ファイルが終了時に削除されない](https://github.com/vim-jp/issues/issues/483)
- [Issue #484 : GUIタブページのD&Dによる並び変え機能追加の提案](https://github.com/vim-jp/issues/issues/484)
- [Issue #485 : clang 3.3でvimがビルドできない？](https://github.com/vim-jp/issues/issues/485)
- [Issue #486 : 複数の &#x2a;.py ファイルを開くと2番目以降のファイルで expandtab が効かない](https://github.com/vim-jp/issues/issues/486)
- [Issue #487 : 選択された補完候補のdictを得られるようにする](https://github.com/vim-jp/issues/issues/487)
- [Issue #488 : VV&#x5f;VERSIONの初期化位置がおかしい。](https://github.com/vim-jp/issues/issues/488)
- [Issue #489 : netrwで全角ではじまるディレクトリ名が扱えない](https://github.com/vim-jp/issues/issues/489)
- [Issue #490 : mbyte.c utf&#x5f;char2cells doublewidthに絵文字](https://github.com/vim-jp/issues/issues/490)
- [Issue #491 : mbyte.c utf&#x5f;char2cells ambiguousが不完全(?)](https://github.com/vim-jp/issues/issues/491)

