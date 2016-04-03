---
layout: post
category: vimmagazine
title: Vim Magazine 10 月号
---

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

- [7.3.673 : "gN" while 'selection' is "exclusive" misses one character](https://github.com/vim/vim/commit/718f007499a5d3f0ff9c0a645780133131ab2b2e)
- [7.3.674 : can't compile with Lua/dyn on Cygwin](https://github.com/vim/vim/commit/1e91f2606b2380e6e443e7389f7648d86fb8c654)
- [7.3.675 : accessing uninitialized memory with a very long file name](https://github.com/vim/vim/commit/ec3cfeb19988dbeb4cfb27e735f3905149f79169)
- [7.3.676 : Ruby compilation on Windows 32 bit doesn't work](https://github.com/vim/vim/commit/b213da0b5742c6615837436ef4f75a416a534f06)
- [7.3.677 : buf\_spname() is used inconsistently](https://github.com/vim/vim/commit/e1704bada4554f4edc95ac6be50144dd7b95332b)
- [7.3.678 : Ruby .so name may not be correct](https://github.com/vim/vim/commit/6f3d6b48a67a87ce7561929b6d1c493d61df4bb5)
- [7.3.679 : Ruby detection uses Config, newer Ruby versions use RbConfig](https://github.com/vim/vim/commit/8139889a5f458b857f0b3d9d5888ca6951a77f5a)
- [7.3.680 : some files missing in the list of distributed files](https://github.com/vim/vim/commit/14f5827940ebd819a77048761a28c0b208addaee)
- [7.3.681 : list of distributed files picks up backup files](https://github.com/vim/vim/commit/7d11244c99bd92526429c491aea433a16f2dbd25)
- [7.3.682 : (after 7.3.677) compiler complains about incompatible types](https://github.com/vim/vim/commit/4ccb265bd41a2048edbe4a42b1d9ca92adcc9bb9)
- [7.3.683 : ":python" may crash when vimbindeval() returns None](https://github.com/vim/vim/commit/d1f13fd5978df5fd4b2e0fbc61e671edec958e15)
- [7.3.684 : "make test" does not delete lua.vim](https://github.com/vim/vim/commit/b35c338f96dab732fa112d7719f10191c7bb34f0)
- [7.3.685 : no test for what patch 7.3.673 fixes](https://github.com/vim/vim/commit/0c17875ef31e7e8b3f16e1b2e8da263c9a420b8b)
- [7.3.686 : cannot use CTRL-\\ e mapping when entering an expression](https://github.com/vim/vim/commit/b73568149f2f41a84b9b701ae512adbd79b2e93b)
- [7.3.687 : test 16 fails when $DISPLAY is not set](https://github.com/vim/vim/commit/172ddb60c2401cd24c80a096709226e37592111f)
- [7.3.688 : Python 3.3 is not supported](https://github.com/vim/vim/commit/7bc4f93c07a1e8b17e7d91f814b9801bd611d569)
- [7.3.689 : MzScheme and Lua may use a NULL string](https://github.com/vim/vim/commit/d04da7cb4afe77de05e4e08e7e86e504be6759e0)
- [7.3.690 : crash with directory name equal to maximum path length](https://github.com/vim/vim/commit/dac7569612d4088084ef81dfc32c3bb86aec090f)
- [7.3.691 : state specific to the Python thread is discarded](https://github.com/vim/vim/commit/b88adbf844612281d3ed11166bbf03368f7ed704)
- [7.3.692 : can't build GTK version with GTK 2.0](https://github.com/vim/vim/commit/205f9f5e2d4b25a94072644bbcdd2c8b20ad7b80)
- [7.3.693 : can't make 'softtabstop' follow 'shiftwidth'](https://github.com/vim/vim/commit/9f340fa57b91db9c04307c99cd4475f197d7a5c8)
- [7.3.694 : 'shiftwidth' is not so easy to use in indent files](https://github.com/vim/vim/commit/2d17fa3ebd378097e3c0feab813a2c60fc9bf474)
- [7.3.695 : balloon cannot show multi-byte text](https://github.com/vim/vim/commit/b6101cf77f3e9e8cc33a94a0c01db3bcde96eba9)
- [7.3.696 : message about added spell language can be wrong](https://github.com/vim/vim/commit/f3408e76fb500d56557489cd533d19d116fc260b)
- [7.3.697 : leaking resources when setting GUI font](https://github.com/vim/vim/commit/c4b98fbb3ea881621d051f84d67609f6789880bc)
- [7.3.698 : Python 3 does not preserve state beween commands](https://github.com/vim/vim/commit/003d14a26771bf7d2c05013cd2fd9d2622f373b9)
- [7.3.699 : manually set 'ttymouse' is overruled by automatic detection](https://github.com/vim/vim/commit/bffa06ddb8d25081dea48274aff2cb65ebef6ced)
- [7.3.700 : cannot detect URXVT and SGR mouse support](https://github.com/vim/vim/commit/cfb807026f678e2e99c98baa2b2b7e74b82b0fc7)
- [7.3.701 : MS-Windows: Crash with stack overflow when setting 'encoding'](https://github.com/vim/vim/commit/faca84059a6ef728e1326becc2f82345a59b50c7)
- [7.3.702 : nmake from VS6 service pack 6 is not recognized](https://github.com/vim/vim/commit/67a4f6cfb4fe0f8e8cc576a4f4196573ccc5e3fb)
- [7.3.703 : when 'undofile' is reset the hash is computed unnecessarily](https://github.com/vim/vim/commit/e8d8fd2add98a95e6c776deb54e2a2c59fddf1c7)
- [7.3.704 : repeating "cgn" does not always work correctly](https://github.com/vim/vim/commit/f00dc2627bc753560b82ae861c90fdd79dc454b4)
- [7.3.705 : mouse features are not in alphabetical order](https://github.com/vim/vim/commit/707cfb82dce17018da95ef0dc5466082258260f5)
- [7.3.706 : (after 7.3.697) can't build Motif version](https://github.com/vim/vim/commit/a9a2d8f4d0d8bfd10fdcf468e6614e196d39ffd0)
- [7.3.707 : (after 7.3.701) library name with non-latin characters fails](https://github.com/vim/vim/commit/17aa8cc255d313ac33eca4f2727ed2a6b8bdd8ef)
- [7.3.708 : filler lines above the first line may be hidden](https://github.com/vim/vim/commit/a80888d296f97ed28a4debc5019a7bbe86cd9121)
- [7.3.709 : compiler warning for unused argument](https://github.com/vim/vim/commit/95e09eaaadbc9a1b76bed8af0ad4a378010187c1)
- [7.3.710 : (after 7.3.704) "fn" is broken](https://github.com/vim/vim/commit/12a753a5ca03fe8beddc06201104e785baccb471)
- [7.3.711 : (after 7.3.688) vim.current.buffer is not available](https://github.com/vim/vim/commit/9c9cbf13511522c05e8cc8d038a858608a158533)
- [7.3.712 : nmake from VS2010 SP1 is not recognized](https://github.com/vim/vim/commit/e3a2246227cfc1c9a4415d2c3f8a71622a67eac3)

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

