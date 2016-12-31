---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 12 月号

---
2016年最後の発行となりました。
ちょっと新しい動きの乏しい月となりましたが、
Vim 8.0も落ち着いてきたということでしょう。

2016年中はVimとvim-jp、加えてVim Magazineへのご愛顧ご助力をありがとうございました。
2017年も引き続きよろしくお願いいたします。

## 話題

*   [Meguro.vim #1](https://megurovim.connpass.com/event/46044/) が 12/18(日) に目黒で開催されました
*   2016 年のアドベントカレンダーが終了しました
    *   [Vim](http://qiita.com/advent-calendar/2016/vim)
    *   [Vim (その2)](http://qiita.com/advent-calendar/2016/vim2)
    *   [Vim 8.0](http://qiita.com/advent-calendar/2016/vim8)

## 今月の新機能

*   8.0.0107: `ch_open()` のオプションに `"drop"` が追加されました。
*   8.0.0117: 並列 make (例 : `make -j 4`) に対応しました。
*   ランタイム更新 2016/12/01: `:compiler` に `ghc` が追加されました。

## リリース情報

- [8.0.0107 : when reading channel output in timer messages may be missing](https://github.com/vim/vim/commit/958dc6923d341390531888058495569d73c356c3)
- [8.0.0108 : (after 8.0.0107) the channel "drop" option is not tested](https://github.com/vim/vim/commit/65e08ee1d26aa7bf341ac0e0400839d696d1ab64)
- [8.0.0109 : still checking if memcmp() exists, should be everywhere now](https://github.com/vim/vim/commit/b129a447f3b580d4c941869672b0557c52c37e4d)
- [8.0.0110 : drop command doesn't use existing window](https://github.com/vim/vim/commit/5a030a540f4157d5c9905e3564282c92b4dcec9a)
- [8.0.0111 : the :history command is not tested](https://github.com/vim/vim/commit/eebd84eb94ed7f59a06a52cb4863563642f58899)
- [8.0.0112 : tests 92 and 93 are old style](https://github.com/vim/vim/commit/eca626fcdb73d480660c78b9f84cc043fa561922)
- [8.0.0113 : MS-Windows: dialog for saving changes on the wrong monitor](https://github.com/vim/vim/commit/87f3d202a90bd2d08a7afd55b3486b10bef858bb)
- [8.0.0114 : coding style not optimal](https://github.com/vim/vim/commit/b04a98f6c3cca14bf055934b0a793f4dc376858b)
- [8.0.0115 : when building with Cygwin libwinpthread isn't found](https://github.com/vim/vim/commit/e3af763d5e6b90a9b5d5706920e669fd8f0b6c77)
- [8.0.0116 : using CTRl-\] in English help language from 'helplang' is used](https://github.com/vim/vim/commit/6dbf66aa3e2197ce41f2b1cc7602bb9c15840548)
- [8.0.0117 : parallel make fails](https://github.com/vim/vim/commit/327054df45faf5390e7392708f58eb49e9f323d4)
- [8.0.0118 : "make proto" adds extra function prototype](https://github.com/vim/vim/commit/5162822914372fc916a93f85848c0c82209e7cec)
- [8.0.0119 : no test for using CTRL-R on the command line](https://github.com/vim/vim/commit/21efc3633edb58809c5dd89b025d34d7002e731c)
- [8.0.0120 : channel test is still flaky on OS X](https://github.com/vim/vim/commit/5643db84c6a9f15d14492cefd52647623aa2ac7c)
- [8.0.0121 : setting 'cursorline' changes the curswant column](https://github.com/vim/vim/commit/a2477fd3490c1166522631eee53c57d34321086a)
- [8.0.0122 : channel test is still flaky on OS X](https://github.com/vim/vim/commit/3fad98e8af247af8ebc49730646282a71ccdd47a)
- [8.0.0123 : modern Sun compilers define "&#x5f;&#x5f;sun" instead of "sun"](https://github.com/vim/vim/commit/a899e6ecc4523c7e411eaf6fbaa4197d70f6f39e)
- [8.0.0124 : internal error for assert&#x5f;inrange(1, 1)](https://github.com/vim/vim/commit/3421566376b5723213af502bd3c2b9debe025ef1)
- [8.0.0125 : not enough testing for entering Ex commands](https://github.com/vim/vim/commit/eaaa9bbda6ec0a8589a9b23720f95bffe01dc267)
- [8.0.0126 : display problem with 'foldcolumn' and a wide character](https://github.com/vim/vim/commit/6270660611a151c5d0f614a5f0248ccdc80ed971)
- [8.0.0127 : cancelling completion still formats text](https://github.com/vim/vim/commit/73fd4988866c3adc15b5d093efdf5e8cf70d093d)
- [8.0.0128 : (after 8.0.0126) display test fails on MS-Windows](https://github.com/vim/vim/commit/7089237885218eb8a19805bc2b75481c4efcd6ba)
- [8.0.0129 : parallel make still doesn't work](https://github.com/vim/vim/commit/0df3c7f2a05c2a99f2fb2747ae46bd6594052997)
- [8.0.0130 : configure uses "ushort" while the Vim code doesn't](https://github.com/vim/vim/commit/63de19e805a7df2b52ec0e705b6a668ecd8e1b64)
- [8.0.0131 : not enough test coverage for syntax commands](https://github.com/vim/vim/commit/73b484c4da00011317dc68ada4f5dfc6515ad263)
- [8.0.0132 : (after 8.0.0131) test fails because of using :finish](https://github.com/vim/vim/commit/4c8980b717f73042f1d625ee255fa74eddb989ba)
- [8.0.0133 : "2;'(" causes ml&#x5f;get errors in an empty buffer](https://github.com/vim/vim/commit/fe38b494fff56cd9b2fcaeef26a8fd7b6557d69c)
- [8.0.0134 : null pointer access reported by UBsan](https://github.com/vim/vim/commit/c4bfedabe057c05f09a455a5851089e177fa9c00)

## 新着スクリプト

- [minimalist.vim : A Material Colorscheme Darker for Vim](http://www.vim.org/scripts/script.php?script_id=5490)
- [vim-base64 : Vim plugin to encode/decode base64 strings](http://www.vim.org/scripts/script.php?script_id=5491)
- [vim-loggly-search : A vim plug-in to easily search in loggly](http://www.vim.org/scripts/script.php?script_id=5492)
- [far.vim : Find And Replace Vim plugin](http://www.vim.org/scripts/script.php?script_id=5493)
- [vtags : verdi like, verilog code signal trace and show topo script  ](http://www.vim.org/scripts/script.php?script_id=5494)
- [ExpandBackspace : Make backspace eat white space to the last tabstop](http://www.vim.org/scripts/script.php?script_id=5495)
- [coverage.vim : Show code coverage as signs](http://www.vim.org/scripts/script.php?script_id=5496)
- [Grammalecte : Vim plugin for the Grammalecte French grammar checker](http://www.vim.org/scripts/script.php?script_id=5497)
- [Broduo Color Scheme : A dark color scheme for Vim.](http://www.vim.org/scripts/script.php?script_id=5498)
- [gen&#x5f;tags.vim : A simple vim plugin that generate/maintain ctags and gtags database.](http://www.vim.org/scripts/script.php?script_id=5499)
- [vim-paragraph : Vim plain text utilities](http://www.vim.org/scripts/script.php?script_id=5500)
- [rpgle.vim : Free (7.1) syntax highlight, folds, indent, etc for RPG/ILE](http://www.vim.org/scripts/script.php?script_id=5501)
- [vim-workspace : The Vim Workspace Manager.](http://www.vim.org/scripts/script.php?script_id=5502)
- [yarn.vim : Yarn for Vim.](http://www.vim.org/scripts/script.php?script_id=5503)
- [Sprint : Async file running](http://www.vim.org/scripts/script.php?script_id=5504)
- [vmux : vim/neovim session handler within tmux](http://www.vim.org/scripts/script.php?script_id=5505)
- [ipynb&#x5f;notedown.vim : plugin for editing jupyter notebook (ipynb) files through notedown](http://www.vim.org/scripts/script.php?script_id=5506)
- [Arcadia : A Monochromatic colorscheme ](http://www.vim.org/scripts/script.php?script_id=5507)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1327)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1190)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (582)
4. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (506)
5. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (498)
6. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (423)
7. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (417)
8. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (415)
9. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (392)
10. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (354)

## vim-jp issues

Open : 248 (-1) | Closed : 748 (+6)

- [Issue #992 : if&#x5f;pythでvim.buffersした結果がvim8とneovimで違う](https://github.com/vim-jp/issues/issues/992)
- [Issue #993 : getbufinfo() で得られる lnum の意味を知りたい](https://github.com/vim-jp/issues/issues/993)
- [Issue #994 : 画面描画が遅くなる](https://github.com/vim-jp/issues/issues/994)
- [Issue #995 : BufWinLeave発生時にquitするとSEGVする](https://github.com/vim-jp/issues/issues/995)
- [Issue #996 : 'isident', 'iskeyword', 'isfname' に含まれない文字にマッチする正規表現パターンを構成したい](https://github.com/vim-jp/issues/issues/996)

