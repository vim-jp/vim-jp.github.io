---
layout: post
category: vimmagazine
title: Vim Magazine 10 月号
---

# Vim Magazine 10 月号

## 話題

- [ujihisa.vim#3](http://vim-jp.org/ujihisa.vim-3/)
  An International Vim Conference 東京 11月3日(土)

- [Vim勉強会#11](http://cotocoto.jp/event/71098)
  Vim勉強会#11 兵庫 11月11日(日)

## 今月の新機能

- 7.3.688: Python 3.3がサポートされました。
- 7.3.693: 'softtabstop' を -1 に設定することで、'shiftwidth' と同じ幅、という指定ができるようになりました。
- 7.3.694: `shiftwidth()` 関数が追加されました。'shiftwidth' が 0 のときは 'tabstop' の幅と同じになりますが、実際に使われる値をとるのを簡単にするため。

## リリース情報

- [7.3.673 : "gN" while 'selection' is "exclusive" misses one character](http://code.google.com/p/vim/source/detail?r=ca1c025079b11c98363e11576d903af756a593fc)
- [7.3.674 : can't compile with Lua/dyn on Cygwin](http://code.google.com/p/vim/source/detail?r=4da6003f0f0eee0f3908f8627af6d13a94c10cc6)
- [7.3.675 : accessing uninitialized memory with a very long file name](http://code.google.com/p/vim/source/detail?r=9963f7a9b41b555a0660288dd94df44b6dee7c2e)
- [7.3.676 : Ruby compilation on Windows 32 bit doesn't work](http://code.google.com/p/vim/source/detail?r=5309996ddcedd7e2dcc8c46624348405290b8940)
- [7.3.677 : buf\_spname() is used inconsistently](http://code.google.com/p/vim/source/detail?r=8115f449a574bc146f75449161f1757aaa5bbb66)
- [7.3.678 : Ruby .so name may not be correct](http://code.google.com/p/vim/source/detail?r=de253f606048a2862a338190b071e5e3c7de7a41)
- [7.3.679 : Ruby detection uses Config, newer Ruby versions use RbConfig](http://code.google.com/p/vim/source/detail?r=dd0f25ede9a8322b5b638961835cc4743cd479a2)
- [7.3.680 : some files missing in the list of distributed files](http://code.google.com/p/vim/source/detail?r=52acec46b75ecf24d044982d4cc067ef7890e7d6)
- [7.3.681 : list of distributed files picks up backup files](http://code.google.com/p/vim/source/detail?r=a2fe24ec6a0d06d0af286b9fee9e307c33f27cca)
- [7.3.682 : (after 7.3.677) compiler complains about incompatible types](http://code.google.com/p/vim/source/detail?r=8eb5b61161b4c56b378ec829e360204102d22db4)
- [7.3.683 : ":python" may crash when vimbindeval() returns None](http://code.google.com/p/vim/source/detail?r=0f5ee262963528ddccfa1f54b83696f939d21b53)
- [7.3.684 : "make test" does not delete lua.vim](http://code.google.com/p/vim/source/detail?r=1cc37b8cfaa1c670e442e2f9444feeb1820228bc)
- [7.3.685 : no test for what patch 7.3.673 fixes](http://code.google.com/p/vim/source/detail?r=34aa94a421b7a03a71013ddbb92937dbaf83b1aa)
- [7.3.686 : cannot use CTRL-\\ e mapping when entering an expression](http://code.google.com/p/vim/source/detail?r=ba9c9d0c4ff4cf56c3c5d2dec2da71decbbe9f6a)
- [7.3.687 : test 16 fails when $DISPLAY is not set](http://code.google.com/p/vim/source/detail?r=0545dab1517f8ce2b8319c02572e4bb357567034)
- [7.3.688 : Python 3.3 is not supported](http://code.google.com/p/vim/source/detail?r=2696da84d37093da0dea419878e4241385eb515b)
- [7.3.689 : MzScheme and Lua may use a NULL string](http://code.google.com/p/vim/source/detail?r=3c072c1cb87372f3ebd345463abbc451e80143e2)
- [7.3.690 : crash with directory name equal to maximum path length](http://code.google.com/p/vim/source/detail?r=caa0ca9ad06cb515b7e6d2e4410e03802b58e167)
- [7.3.691 : state specific to the Python thread is discarded](http://code.google.com/p/vim/source/detail?r=7f10daa706bb6ed0c39f1d2905cd78368a98d19f)
- [7.3.692 : can't build GTK version with GTK 2.0](http://code.google.com/p/vim/source/detail?r=337a4368fd2b68b9822b94e595d77b4d6773bfa4)
- [7.3.693 : can't make 'softtabstop' follow 'shiftwidth'](http://code.google.com/p/vim/source/detail?r=fd96b3cc88ed71d4c4d24c6ac67c9d3a68ae6e6b)
- [7.3.694 : 'shiftwidth' is not so easy to use in indent files](http://code.google.com/p/vim/source/detail?r=8235ccf121bb241f8105d0ba6e1c92e9f8b04555)
- [7.3.695 : balloon cannot show multi-byte text](http://code.google.com/p/vim/source/detail?r=cd6c420e31d6843fa1231e1cf7a3d9eff0ad4e07)
- [7.3.696 : message about added spell language can be wrong](http://code.google.com/p/vim/source/detail?r=447c688b82eb4a5eb7cd0e16521094b77160fd0f)
- [7.3.697 : leaking resources when setting GUI font](http://code.google.com/p/vim/source/detail?r=7c72b5f7ada1f4a028315aa50f8739e0e83c2e74)
- [7.3.698 : Python 3 does not preserve state beween commands](http://code.google.com/p/vim/source/detail?r=b5b892472ecb1a10cce1709e41700b264b6485d2)
- [7.3.699 : manually set 'ttymouse' is overruled by automatic detection](http://code.google.com/p/vim/source/detail?r=4ffb6f9b58e0b51923fe6a71e6af158b6e3f7864)
- [7.3.700 : cannot detect URXVT and SGR mouse support](http://code.google.com/p/vim/source/detail?r=c656847932d33b5d68394ae1ea3153dbdba4d3ff)
- [7.3.701 : MS-Windows: Crash with stack overflow when setting 'encoding'](http://code.google.com/p/vim/source/detail?r=48af8656094539f8f71dc51e0a6aa14d48ddffd0)
- [7.3.702 : nmake from VS6 service pack 6 is not recognized](http://code.google.com/p/vim/source/detail?r=916b5c11a13e9978c85f9742546e70d614116883)
- [7.3.703 : when 'undofile' is reset the hash is computed unnecessarily](http://code.google.com/p/vim/source/detail?r=f69321485951a3c9bd0e9e7484a1901b94cae726)
- [7.3.704 : repeating "cgn" does not always work correctly](http://code.google.com/p/vim/source/detail?r=be8ad9c9b973a4d87b9bcfee2b6a8de1f6d3e071)
- [7.3.705 : mouse features are not in alphabetical order](http://code.google.com/p/vim/source/detail?r=0aa74f2f50a8333fd5544f5d4b383132be41704c)
- [7.3.706 : (after 7.3.697) can't build Motif version](http://code.google.com/p/vim/source/detail?r=d0ae72c2f5f94e2483ab0979eec652b1289679a8)
- [7.3.707 : (after 7.3.701) library name with non-latin characters fails](http://code.google.com/p/vim/source/detail?r=d24d309c365fc5c6a9324a6183d7b927d9f5132f)
- [7.3.708 : filler lines above the first line may be hidden](http://code.google.com/p/vim/source/detail?r=64427849c1582299f2af389c6225eead529f0329)
- [7.3.709 : compiler warning for unused argument](http://code.google.com/p/vim/source/detail?r=fb7dec372910b61478abdd7f79871724153ebfdd)
- [7.3.710 : (after 7.3.704) "fn" is broken](http://code.google.com/p/vim/source/detail?r=906eb5c0ef129b6876fa9c6c5a13e44807f7a503)
- [7.3.711 : (after 7.3.688) vim.current.buffer is not available](http://code.google.com/p/vim/source/detail?r=2a7b2943be22c27edb3c0334423c1b17275f18fb)
- [7.3.712 : nmake from VS2010 SP1 is not recognized](http://code.google.com/p/vim/source/detail?r=557ef119999ea687ccb84eddca295755c24989da)

## 新着スクリプト

- [butane.vim : Light up your buffers.](http://www.vim.org/scripts/script.php?script_id=4245)
- [longline.vim : Avoid long lines.](http://www.vim.org/scripts/script.php?script_id=4246)
- [tabdiff : A different way to diff.](http://www.vim.org/scripts/script.php?script_id=4247)
- [PatternComplete : Insert mode completion for matches of queried / last search pattern.](http://www.vim.org/scripts/script.php?script_id=4248)
- [write.vim : Get your author on.](http://www.vim.org/scripts/script.php?script_id=4249)
- [CommandAlias : Make aliases to vim command](http://www.vim.org/scripts/script.php?script_id=4250)
- [GuessIndent : Indentation guessing plugin](http://www.vim.org/scripts/script.php?script_id=4251)
- [easytree.vim : easytree.vim - is a simple tree file manager for vim](http://www.vim.org/scripts/script.php?script_id=4252)
- [TagManager : Project aware incremental tags manager](http://www.vim.org/scripts/script.php?script_id=4253)
- [plsql.vim : Syntax for Oracle PL/SQL language with folding feature](http://www.vim.org/scripts/script.php?script_id=4254)
- [xqmarklogic : vim file type plugin to run an xquery against a MarkLogic server](http://www.vim.org/scripts/script.php?script_id=4255)
- [Ozzy : Quick file launcher](http://www.vim.org/scripts/script.php?script_id=4256)
- [djangojump : jump between template, view, url  in django project](http://www.vim.org/scripts/script.php?script_id=4257)
- [sha3.vim : sha3 function](http://www.vim.org/scripts/script.php?script_id=4258)
- [vim-herokudoc-theme : Color scheme inspired Heroku documentation colors](http://www.vim.org/scripts/script.php?script_id=4260)
- [tumblr.vim : Syntax highlighting for Tumblr themes.](http://www.vim.org/scripts/script.php?script_id=4261)
- [vim-xdc-syntax : Syntax highlighter for Xilinx Design Constraints](http://www.vim.org/scripts/script.php?script_id=4262)
- [autotags : vim plugin for easy ctags and cscope handling in a separate directory ](http://www.vim.org/scripts/script.php?script_id=4263)
- [VimChatBot : A simple self-teaching chat bot game for Vim](http://www.vim.org/scripts/script.php?script_id=4264)
- [MotionComplete : Insert mode completion for chunks covered by queried {motion} or text object.](http://www.vim.org/scripts/script.php?script_id=4265)
- [BracketComplete : Insert mode completion for text inside various brackets.](http://www.vim.org/scripts/script.php?script_id=4266)
- [LineEndComplete : Insert mode completion that completes the rest of the line.](http://www.vim.org/scripts/script.php?script_id=4267)
- [ctagloader : The script will traverse path tree and load ctag found on the way.](http://www.vim.org/scripts/script.php?script_id=4268)
- [rvm.vim : Switch Ruby versions from inside Vim](http://www.vim.org/scripts/script.php?script_id=4269)
- [vim-eco : Eco support for Vim](http://www.vim.org/scripts/script.php?script_id=4270)
- [CY : Fast input method for Chinese (穿越输入法)](http://www.vim.org/scripts/script.php?script_id=4271)
- [Unicode RST Tables : Allows to create and edit restructuredText tables easily (Unicode and Python 3).](http://www.vim.org/scripts/script.php?script_id=4272)
- [autochmodx.vim : If a current file looks like an executable script, make it executable by chmod](http://www.vim.org/scripts/script.php?script_id=4273)
- [very monochrome grey theme : Grey color sheme](http://www.vim.org/scripts/script.php?script_id=4274)
- [sqlhana.vim : sqlhana.vim : Syntax file for SAP HANA](http://www.vim.org/scripts/script.php?script_id=4275)
- [SnippetCompleteSnipMate : Integrate snipMate snippets into SnippetComplete plugin.](http://www.vim.org/scripts/script.php?script_id=4276)
- [Sunset : Automatically set background on local sunrise/sunset time.](http://www.vim.org/scripts/script.php?script_id=4277)
- [Makeshift : Selects the right 'makeprg' for your build system](http://www.vim.org/scripts/script.php?script_id=4278)
- [vim-cmake-project : CMake project management in VIM](http://www.vim.org/scripts/script.php?script_id=4279)
- [bundler.vim : Lightweight support for Ruby's Bundler](http://www.vim.org/scripts/script.php?script_id=4280)
- [yars : Yet Another reStructuredText Syntax](http://www.vim.org/scripts/script.php?script_id=4281)
- [breakindent\_beta.vim : Improves indentation of wrapped lines by updating showbreak for current line.](http://www.vim.org/scripts/script.php?script_id=4282)
- [ctrlp-modified.vim : Easily open locally modified files in your git-versioned projects.](http://www.vim.org/scripts/script.php?script_id=4283)
- [windows\_remember\_size.vim : Each window remembers what size you give it, and restores it when refocused.](http://www.vim.org/scripts/script.php?script_id=4284)
- [toggle\_maximize.vim : Keybinds to temporarily maximize the current window, squashing the rest.](http://www.vim.org/scripts/script.php?script_id=4285)
- [navigation\_enhancer.vim : When moving between windows, uses history not cursor position to select target.](http://www.vim.org/scripts/script.php?script_id=4286)
- [highlight\_word\_under\_cursor.vim : Temporary highlighting of all occurrences of the currently focused word.](http://www.vim.org/scripts/script.php?script_id=4287)
- [highlight\_line\_after\_jump.vim : Briefly highlights the cursor line whenever the cursor jumps vertically.](http://www.vim.org/scripts/script.php?script_id=4288)
- [blinking\_statusline.vim : Flashes the status line of the focused window whenever you change window.](http://www.vim.org/scripts/script.php?script_id=4289)
- [back\_to\_recent\_buffer.vim : Goes back to the previous buffer (rather than C-O previous position).](http://www.vim.org/scripts/script.php?script_id=4290)
- [last\_edit\_marker.vim : Keybind to quickly return to the last code you edited.](http://www.vim.org/scripts/script.php?script_id=4291)
- [py-coverage : Show python coverage results](http://www.vim.org/scripts/script.php?script_id=4292)
- [STL Syntax : Improved C++11 STL syntax highlighting ](http://www.vim.org/scripts/script.php?script_id=4293)
- [git log : A simple Git Log and History diff tool](http://www.vim.org/scripts/script.php?script_id=4294)
- [smartgrep : Grep word in source file excluding comment](http://www.vim.org/scripts/script.php?script_id=4295)
- [ViewPort : Make a buffer from part of a file](http://www.vim.org/scripts/script.php?script_id=4296)
- [detailed : Show subtler color variations than your basic palette-shifting color scheme (cur](http://www.vim.org/scripts/script.php?script_id=4297)
- [vimside : Vimside - Vim Scala Ide](http://www.vim.org/scripts/script.php?script_id=4298)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (3044)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (2593)
3. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (1443)
4. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (1409)
5. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (1402)
6. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (1148)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (1115)
8. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (1083)
9. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (1080)
10. [minibufexpl.vim : Elegant buffer explorer - takes very little screen space](http://www.vim.org/scripts/script.php?script_id=159) (1059)

## vim-jp issues

Open : 129 (-1) | Closed : 137 (+6)

- [Issue #262 : DirectWriteで描画したい](https://github.com/vim-jp/issues/issues/262)
- [Issue #263 : 要望：候補を選択した時のautocmdがほしい](https://github.com/vim-jp/issues/issues/263)
- [Issue #264 : if\_python のおかしな挙動](https://github.com/vim-jp/issues/issues/264)
- [Issue #265 : autofmt を vim-jp に置きたい](https://github.com/vim-jp/issues/issues/265)
- [Issue #266 : deb を配布したい](https://github.com/vim-jp/issues/issues/266)

