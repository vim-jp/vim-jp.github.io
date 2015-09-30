---
layout: post
category: vimmagazine
title: Vim Magazine 2 月号
---

## 話題

- [Vim Advent Calendar 2012](http://atnd.org/events/33746)

  Vim Advent Calendar 2012 まだやってます。


## 今月の新機能

- 7.3.816 : `sha256()` 関数が追加されました。
- 7.3.826 : `:version` の出力が見やすくなりました。
- 7.3.834 : if\_ruby が Ruby 2.0 に対応しました。
- 7.3.836 : Cygwin 版でクリップボードが使えるようになりました。

## リリース情報

- [7.3.798 : (after 7.3.791) MzScheme: circular list works wrong](https://github.com/vim/vim/commit/39d7d51bcdc0edd8eb3034a9d094511b7b75e9e1)
- [7.3.799 : the color column is not correct when entering a buffer](https://github.com/vim/vim/commit/9c2e6cc6edb939828032c9179fe05b57fa84a384)
- [7.3.800 : the " mark is not adjusted when inserting lines](https://github.com/vim/vim/commit/b6a76ffc4cc79163c88b734e0a803385321b3ca6)
- [7.3.801 : ":window set nu?" displays the cursor line](https://github.com/vim/vim/commit/530e7dfafd7e951ac36e39c87bf72ebdb4caa428)
- [7.3.802 : no error when setting 'isk' to a value ending in a comma](https://github.com/vim/vim/commit/309379fc989a4f3337c9cf3361ab6eb9a456c626)
- [7.3.803 : (after 7.3.792) ":s/a/b/gc" is a no-op when typing 'q'](https://github.com/vim/vim/commit/11cb6e6e137afb9e20d68e69bf38d458a96701fd)
- [7.3.804 : (after 7.3.799) compiler warning for tiny build](https://github.com/vim/vim/commit/fffcfea9845021269da36bbe8eee5641ebb144c4)
- [7.3.805 : Lua version 5.2 is not detected properly on Arch Linux](https://github.com/vim/vim/commit/07e1da6bbdf4a2a98c4645868225ffad6df5ae3b)
- [7.3.806 : compiler warnings in Perl code when using Visual studio 2012](https://github.com/vim/vim/commit/e9d47cd08b0272ec278f03bb1be21f087ce63af6)
- [7.3.807 : popup menu does not work properly in some situations](https://github.com/vim/vim/commit/51971b33988e590901b9f6ad14a5a36f276afd0b)
- [7.3.808 : Python threads still do not work properly](https://github.com/vim/vim/commit/76d711c3b5397b749a67d229150d3c1ff3f33add)
- [7.3.809 : the dosinst.c program has a buffer overflow](https://github.com/vim/vim/commit/181ace28ed41ba15078285194a9b9d6e52b46844)
- [7.3.810 : 'relativenumber is reset unexpectedly](https://github.com/vim/vim/commit/f4e5e86627abe926ad201ef7dfa64b01bbef7567)
- [7.3.811 : useless termresponse parsing for SGR mouse](https://github.com/vim/vim/commit/be9c6276665be04dbd5543467383e959cbff30c3)
- [7.3.812 : when 'indentexpr' moves the cursor "curswant" not restored](https://github.com/vim/vim/commit/8fe8d9ed064d1a7c41fc43a565513dc41fd3aaed)
- [7.3.813 : the CompleteDone event is not always triggered](https://github.com/vim/vim/commit/a3914327f73b3d994af6ecb81dace22fab7d8dc7)
- [7.3.814 : Win32 console: problem if 'enc' differs from current codepage](https://github.com/vim/vim/commit/ffeedec023a47a71fe361c4f0cdbee1649ca8c5f)
- [7.3.815 : building with Cygwin and Ruby doesn't work](https://github.com/vim/vim/commit/c2149ea7cd782588563f235e2d1315eaa353ab7b)
- [7.3.816 : can't compute a hash](https://github.com/vim/vim/commit/af9aeb9250ea657b1f8c06ef4bab21d56c111297)
- [7.3.817 : test 89 fails with tiny and small features](https://github.com/vim/vim/commit/caca92be8a8556ff26f8ba37a517b68848b93744)
- [7.3.818 : when test 40 fails it may leave files behind](https://github.com/vim/vim/commit/707060e2ee59e50096e1a96abf00e18293fb0569)
- [7.3.819 : (CORRECTED) compiling without +eval and Python isn't working](https://github.com/vim/vim/commit/b2c5a5ac79848b7f889a8b279eed40be3064feca)
- [7.3.820 : errors when building with small features and Lua, Perl or Ruby](https://github.com/vim/vim/commit/feeaa688ebfd95281529191718b9194187995f7f)
- [7.3.821 : build with OLE and Cygwin is broken](https://github.com/vim/vim/commit/5004e8846e3e7ab3888b57746c4d84719c70ccf4)
- [7.3.822 : (after 7.3.799) crash when accessing freed buffer](https://github.com/vim/vim/commit/6d47df7ca0fd1dd649e050fd3bddbf8b3dbe266d)
- [7.3.823 : (after 7.3.821) building with Cygwin: '-lsupc++' is not needed](https://github.com/vim/vim/commit/452ff5bdc780029912387819bb50c7c2031344f1)
- [7.3.824 : can redefine builtin functions](https://github.com/vim/vim/commit/2142e5d85112f37e68374c81612d6277c6aaaf0b)
- [7.3.825 : with Python errors are not always clear](https://github.com/vim/vim/commit/4d36987c254b820ebfaf98957c0ae5182db0508e)
- [7.3.826 : list of features in :version output is hard to read](https://github.com/vim/vim/commit/445f3037ea1640683433df89ba383f2e4c79329a)
- [7.3.827 : (after 7.3.825) Python tests fail](https://github.com/vim/vim/commit/995a8cd18190c0b922796a8ee865e95bd1dbd916)
- [7.3.828 : mappings are not aware of wildmenu mode](https://github.com/vim/vim/commit/8738fc1be888d73fb4fe302cf085c0bcf4f52e91)
- [7.3.829 : 'showmatch' shows wrong match and may scroll the screen](https://github.com/vim/vim/commit/187d3acb7fbb82266802db9dc3c659f8b725fbd0)
- [7.3.830 : :mksession confuses bytes and characters when placing cursor](https://github.com/vim/vim/commit/558ddad8e836e0c6597c27eccc35275e7b4c4e6a)
- [7.3.831 : clumsy to handle the situation that a variable does not exist](https://github.com/vim/vim/commit/63dbda1caa88042c88d46ad139c99def2e8b80ca)
- [7.3.832 : compiler warning](https://github.com/vim/vim/commit/5c9626301b93621869bbc0662cf71ccb2f853c51)
- [7.3.833 : In the terminal scroll wheel always scrolls the active window](https://github.com/vim/vim/commit/40cf4b44ed6c7fae7d35b0652911cf7c97d9f2a2)
- [7.3.834 : Ruby 2.0 has a few API changes](https://github.com/vim/vim/commit/886ed691d26dd0253608eee174fecc24130e4de7)
- [7.3.835 : "xxd -i" fails on an empty file](https://github.com/vim/vim/commit/53eb37e9f9dfba51092f72553d79d3bb677a336b)
- [7.3.836 : clipboard does not work on Win32 when compiled with Cygwin](https://github.com/vim/vim/commit/693e40c2cdb65f19febde5664633ac465931cc63)
- [7.3.837 : empty lines in :version output when 'columns' is 320](https://github.com/vim/vim/commit/f13f45d59b1a6ab6681ee3953501f319e6d0128b)
- [7.3.838 : (after 7.3.830) insufficient testing for mksession](https://github.com/vim/vim/commit/fdf447b2863495170ed20b7f99adfe92e4eae3c0)
- [7.3.839 : some files missing in the list of distributed files](https://github.com/vim/vim/commit/e468e266c45806c782f387ef631a99bfb2575576)
- [7.3.841 : parse error in "cond ? one : two" expression with a subscript](https://github.com/vim/vim/commit/e17c2603355669a84813f454875d865f38172ff8)
- [7.3.842 : compiler warning](https://github.com/vim/vim/commit/988232fabba5d5dc154fcb962243a36740f02b12)
- [7.3.843 : (after 7.3.841) missing test file changes](https://github.com/vim/vim/commit/b8f8461d94e153c591eff126f508b25ce5c30639)

## 新着スクリプト

- [sunny : Color scheme for \_\_terminal\_\_ vim](http://www.vim.org/scripts/script.php?script_id=4419)
- [Fold License : A vim plugin to hide the annoyng license at the top of each file](http://www.vim.org/scripts/script.php?script_id=4420)
- [texlog.vim : Syntax highlighting and folding for TeX/LaTeX log files](http://www.vim.org/scripts/script.php?script_id=4421)
- [msbuild.vba : Vim plugin for MSBuild project support](http://www.vim.org/scripts/script.php?script_id=4422)
- [fat-finger : Simple vim plugin for common misspellings and typos](http://www.vim.org/scripts/script.php?script_id=4423)
- [usefold : indent based fold marker creation](http://www.vim.org/scripts/script.php?script_id=4424)
- [phonzia : A dark color scheme](http://www.vim.org/scripts/script.php?script_id=4425)
- [find-complete : Find files on path with file-name auto-complete functionality](http://www.vim.org/scripts/script.php?script_id=4426)
- [gf-ext : Extend "gf" to support opening files in external programs](http://www.vim.org/scripts/script.php?script_id=4427)
- [JavaScript Libraries Syntax : Syntax for popular JavasSript libs](http://www.vim.org/scripts/script.php?script_id=4428)
- [ Mini Year Todo & Gantt Planner for VIM  : Draw a todo list and its Gantt chart over the year ](http://www.vim.org/scripts/script.php?script_id=4429)
- [vimogen : Zero-configuration Vim Plugin Manager](http://www.vim.org/scripts/script.php?script_id=4430)
- [tcd.vim : Tab-specific directories](http://www.vim.org/scripts/script.php?script_id=4431)
- [zoomwintab.vim : zoomwintab.vim is a simple zoom window plugin](http://www.vim.org/scripts/script.php?script_id=4432)
- [ingo-library : Vimscript library of common functions.](http://www.vim.org/scripts/script.php?script_id=4433)
- [erroneous : The right way to handle errors](http://www.vim.org/scripts/script.php?script_id=4434)
- [tla.vim : TLA+ specification language](http://www.vim.org/scripts/script.php?script_id=4435)
- [psearch.vim : Quick access to search results](http://www.vim.org/scripts/script.php?script_id=4436)
- [redif.vim : Syntax highlighting for ReDIF format used by RePEc](http://www.vim.org/scripts/script.php?script_id=4437)
- [rythm.vim : Rythm template engine plugin](http://www.vim.org/scripts/script.php?script_id=4438)
- [Dev\_Delight : This Color-Scheme provides a very Elegant Display with Minimal Eye-Strain](http://www.vim.org/scripts/script.php?script_id=4439)
- [Flake8-vim : Automatic syntax checker for Python, bundled with pep8, pyflakes and mccabe.](http://www.vim.org/scripts/script.php?script_id=4440)
- [QuickFixCurrentNumber : Locate the quickfix item at the cursor position.](http://www.vim.org/scripts/script.php?script_id=4449)
- [ag.vim : Use ag, the\_silver\_searcher (better than ack, which is better than grep)](http://www.vim.org/scripts/script.php?script_id=4450)
- [pry : Pry/Vim Integration](http://www.vim.org/scripts/script.php?script_id=4451)
- [vim-javascript : New maintained version of vim-javascript](http://www.vim.org/scripts/script.php?script_id=4452)
- [play.vim : Play! framework plugin for Vim](http://www.vim.org/scripts/script.php?script_id=4453)
- [vim-vimlparser : VimL parser](http://www.vim.org/scripts/script.php?script_id=4454)
- [rbenv.vim : Minimal rbenv support](http://www.vim.org/scripts/script.php?script_id=4455)
- [vgdb : visual gdb in vim](http://www.vim.org/scripts/script.php?script_id=4456)
- [convert\_case.vim : Convert Case ](http://www.vim.org/scripts/script.php?script_id=4457)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (677)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (570)
3. [pathogen.vim : Easy manipulation of 'runtimepath', 'path', 'tags', etc](http://www.vim.org/scripts/script.php?script_id=2332) (373)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (312)
5. [Vim-R-plugin : Plugin to work with R](http://www.vim.org/scripts/script.php?script_id=2628) (306)
6. [c.vim : C/C++ IDE --  Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (273)
7. [snipMate : TextMate-style snippets for Vim](http://www.vim.org/scripts/script.php?script_id=2540) (263)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (253)
9. [bufexplorer.zip : Buffer Explorer / Browser](http://www.vim.org/scripts/script.php?script_id=42) (241)
10. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (224)

## vim-jp issues

Open : 145 (+13) | Closed : 176 (+7)

- [Issue #302 : 64bit版Vimから一部の64bit版のPythonモジュールを呼び出せない？](https://github.com/vim-jp/issues/issues/302)
- [Issue #303 : +kaoriya版の :Tutorial 実行時にSyntaxErrorが発生する](https://github.com/vim-jp/issues/issues/303)
- [Issue #304 : :simalt ~\<CR>でウインドウメニューを表示させた後、\<Esc>を押すと入力待ちになる](https://github.com/vim-jp/issues/issues/304)
- [Issue #305 : Ctrl-x modeのメッセージを出さずにポップアップメニューのみ表示したい](https://github.com/vim-jp/issues/issues/305)
- [Issue #306 : VimのDLL化 for Windows](https://github.com/vim-jp/issues/issues/306)
- [Issue #307 : ポップアップメニューが出ていないのにCtrl-x modeに入っている時がある](https://github.com/vim-jp/issues/issues/307)
- [Issue #308 : ウィンドウ位置調整がオカシイらしい](https://github.com/vim-jp/issues/issues/308)
- [Issue #309 : Windows で gvim -f が効かない](https://github.com/vim-jp/issues/issues/309)
- [Issue #310 : completeopt=previewを設定していると候補選択時に画面がチラつく](https://github.com/vim-jp/issues/issues/310)
- [Issue #311 : mklink すると .vimrc が読み込まれない?](https://github.com/vim-jp/issues/issues/311)
- [Issue #312 : 日本語メニューのスクリプトでグローバル変数が定義されている](https://github.com/vim-jp/issues/issues/312)
- [Issue #313 : 先行先読否定がCP932で正しく動作していない](https://github.com/vim-jp/issues/issues/313)
- [Issue #314 : Make\_cyg.mak を Make\_ming.mak と統合したい](https://github.com/vim-jp/issues/issues/314)
- [Issue #315 : if\_ruby が ruby2.0 でビルド出来ない。](https://github.com/vim-jp/issues/issues/315)
- [Issue #316 : toml-vim 作って](https://github.com/vim-jp/issues/issues/316)
- [Issue #317 : scala シンタックス](https://github.com/vim-jp/issues/issues/317)
- [Issue #318 : INSTALLpc.txtの改行コードを\<CR>\<LF>に](https://github.com/vim-jp/issues/issues/318)
- [Issue #319 : CygwinでビルドしたVimが make test を通らない](https://github.com/vim-jp/issues/issues/319)
- [Issue #320 : マルチバイト文字を含む行選択時の r がおかしい](https://github.com/vim-jp/issues/issues/320)
- [Issue #321 : クリップボードが文字化け](https://github.com/vim-jp/issues/issues/321)

