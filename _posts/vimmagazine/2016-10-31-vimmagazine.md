---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2016 年 10 月号

---

## 話題

*   VimConf 2016 関連
    *   タイムテーブルが決定しました
    *   参加者の募集が行われました

## 今月の新機能

*   8.0.0023: `gd` と `gD` がコメント内と文字列のマッチを無視するようになりました。
*   8.0.0034: `:messages` コマンドのサブコマンド `clear` を補完できるようになりました。
*   ランタイム更新 2016/10/31: 拡張子 .rs がファイルタイプ `rust` にアサインされました。

## リリース情報

- [8.0.0020 : the regexp engines are not reentrant](https://github.com/vim/vim/commit/6100d02aab7c8294b581cb299250eea164b50e9d)
- [8.0.0021 : GUI: cursor may be drawn on second half of double byte char](https://github.com/vim/vim/commit/84dbd494dca599ecff05b2c2279d402c12e6d197)
- [8.0.0022 : if channel is missing final NL remaining chars are dropped](https://github.com/vim/vim/commit/ec68a99464055029c01082762517e97245ddae0c)
- [8.0.0023 : "gd" and "gD" may find a match in a comment or string](https://github.com/vim/vim/commit/226630a030c0d41145e1109f09633360fc9c999d)
- [8.0.0024 : when netbeans channel closes, "DETACH" is put in output part](https://github.com/vim/vim/commit/8ddef48d1eade1911b946fdda8c73c80856e6273)
- [8.0.0025 : inconsistent use of spaces vs tabs in gd test](https://github.com/vim/vim/commit/936c48f8ca82a0257640c8c9d0792538f5a7e813)
- [8.0.0026 : error format with %W, %C and %Z does not work](https://github.com/vim/vim/commit/9b4579481892a62e7e002498b9eddaaf75bbda49)
- [8.0.0027 : channel is closed when there may still be something to read](https://github.com/vim/vim/commit/dc0ccaee68ca24d10050117fbec757ad33590a17)
- [8.0.0028 : superfluous semicolons](https://github.com/vim/vim/commit/a7c023ec88233a2db4c1f53121921e110ece8d43)
- [8.0.0029 : exceptions for old systems make MS-Windows code complicated](https://github.com/vim/vim/commit/cea912af725c54f4727a0565e31661f6b29c6bb1)
- [8.0.0030 : mouse mode is not automatically detected for tmux](https://github.com/vim/vim/commit/0ba407012c63064f03f1a5677677d4da423e5a73)
- [8.0.0031 : after ":bwipeout" 'fileformat' is not set to the default](https://github.com/vim/vim/commit/e8ef3a093453b73594e15462d4de50b011c8ba66)
- [8.0.0032 : tests may change the input file when something goes wrong](https://github.com/vim/vim/commit/3e8474dd50f64c998bb665ce852f584a58dede6b)
- [8.0.0033 : cannot use overlapping positions with matchaddpos()](https://github.com/vim/vim/commit/a6c27ee6db2c328e0ab0e6d143e2a295a0bb9c9a)
- [8.0.0034 : no completion for ":messages"](https://github.com/vim/vim/commit/9e507ca8a3e1535e62de4bd86374b0fcd18ef5b8)
- [8.0.0035 : order of matches for 'omnifunc' is messed up](https://github.com/vim/vim/commit/472e85970ee3a80abd824bef510df12e9cfe9e96)
- [8.0.0036 : detecting that a job has finished may take a while](https://github.com/vim/vim/commit/97792de2762cc79cc365a8a0b858f27753179577)
- [8.0.0037 : get E924 when switching tabs](https://github.com/vim/vim/commit/0a9046fbcb33770517ab0220b8100c4494bddab2)
- [8.0.0038 : OPEN&#x5f;CHR&#x5f;FILES not defined for FreeBSD/Debian mix](https://github.com/vim/vim/commit/ca291aec99b60fe81eaab36aa718e51421bb88d5)
- [8.0.0039 : when Vim 8 reads an old viminfo next time marks are not read](https://github.com/vim/vim/commit/156919f99afd1ac11d19d4270afbc1afb7245640)
- [8.0.0040 : (after 8.0.0033) matchaddpos() whole line highlighting broken](https://github.com/vim/vim/commit/8507747600bddfd6a68aed057840856bf5548e61)
- [8.0.0041 : get undo item with Insert mode completion without inserting](https://github.com/vim/vim/commit/869e35270ecffd9024958880cb03f6f0bb01ea93)
- [8.0.0042 : (after 8.0.0041) Insert mode completion is not saved for undo](https://github.com/vim/vim/commit/cbd3bd6cbed5baf418b037b17ad46e339ff59174)
- [8.0.0043 : (after 8.0.0041) Insert mode completion is not saved for undo](https://github.com/vim/vim/commit/9ec7fa82a2c4f0283862ed704c2940959e6130ee)
- [8.0.0044 : ml&#x5f;get error In diff mode](https://github.com/vim/vim/commit/025e3e0bafbc85cc4e365145af711edf99d0a90d)
- [8.0.0045 : calling job&#x5f;stop() right after job&#x5f;start() does not work](https://github.com/vim/vim/commit/bb09ceb95477ecc271854b3fdd8d2776eca66adf)
- [8.0.0046 : using NUL instead of NULL](https://github.com/vim/vim/commit/7756e7465d627ff9cd01e59625484a8c302ef853)
- [8.0.0047 : crash when using the preview window from an unnamed buffer](https://github.com/vim/vim/commit/50e5376926dc2ec4a26a7a16f8f0f3213c4afdf0)
- [8.0.0048 : on Windows job&#x5f;stop() stops cmd.exe, not processes it runs](https://github.com/vim/vim/commit/a350bab042605422304161df462b173ffa31dcb9)
- [8.0.0049 : concealing and match highlighting interfere](https://github.com/vim/vim/commit/2f97912800e86a296c001832bbbf2fc425f1e533)
- [8.0.0050 : an exiting job is detected with a large latency](https://github.com/vim/vim/commit/01688ad545ff0809ddad5c8fa6b149dc5d67312b)
- [8.0.0051 : (after 8.0.0048) job&#x5f;stop() change breaks tests](https://github.com/vim/vim/commit/14207f487c9e79a913256a41e3e9716b03b46955)
- [8.0.0052 : (after 8.0.0049) conceal test passes even without the bug fix](https://github.com/vim/vim/commit/35a1f59d635d9a655e1267c18f7cc757afd0d5b0)
- [8.0.0053 : (after 8.0.0047) no test for what 8.0.0047 fixes](https://github.com/vim/vim/commit/60ef3e81f4a54d9f7ee617d57021f0811ec8ada5)
- [8.0.0054 : (after 8.0.0051) Windows: job&#x5f;stop() does not stop processes](https://github.com/vim/vim/commit/fb63090b62801d718fe7e1f44407358404c08724)
- [8.0.0055 : minor comment and style deficiencies](https://github.com/vim/vim/commit/3a117e19e02bf29cfc5e398470dd7851ae3d6803)

## 新着スクリプト

- [lilycomment.vim : comment plugin for csharp](http://www.vim.org/scripts/script.php?script_id=5456)
- [vim-tail : Makes vim behave like tail -f](http://www.vim.org/scripts/script.php?script_id=5457)
- [vim&#x5f;comment&#x5f;toggle : Simple and lightweight comment toggle scripts](http://www.vim.org/scripts/script.php?script_id=5458)
- [syntax for verilog 2001 : set vim for verilog](http://www.vim.org/scripts/script.php?script_id=5459)
- [vim-markdown-toc : A plugin to generate table of contents for Markdown files.](http://www.vim.org/scripts/script.php?script_id=5460)
- [dutch&#x5f;peasants : A dark scheme of earth colours](http://www.vim.org/scripts/script.php?script_id=5461)
- [Despacio : Vim Colorscheme](http://www.vim.org/scripts/script.php?script_id=5462)
- [TabList for terminal vim : Get list of tab pages and the windows in them in a quickfix like window.](http://www.vim.org/scripts/script.php?script_id=5463)
- [vim-kramdown-tab : A vim plugin to resolve indent issue of content nested in Markdown lists, especi](http://www.vim.org/scripts/script.php?script_id=5464)
- [vim-jobs : Manage async jobs ](http://www.vim.org/scripts/script.php?script_id=5466)
- [glsl.vim : Utilities for WebGL/GLSL/Three.js developers.](http://www.vim.org/scripts/script.php?script_id=5468)
- [Subnet Search : Match all IP addresses withing a given IPv4 CIDR address block](http://www.vim.org/scripts/script.php?script_id=5469)
- [MUcomplete : Chained (fallback) completion that works the way you want!](http://www.vim.org/scripts/script.php?script_id=5471)
- [pace.vim : Measure the pace of typing (in Insert mode &c.)](http://www.vim.org/scripts/script.php?script_id=5472)
- [nordisk : A calm, fairly low contrast, dark color scheme](http://www.vim.org/scripts/script.php?script_id=5473)
- [auto&#x5f;mkdir2.vim : Automatically create directories](http://www.vim.org/scripts/script.php?script_id=5474)
- [Install a Plugin in Vim : symlinks don\\\\&#039t overwrite existing files, uninstall only deletes symlinks](http://www.vim.org/scripts/script.php?script_id=5475)
- [xdg&#x5f;open.vim : Run xdg-open from Vim; re-implements netrw's gx.](http://www.vim.org/scripts/script.php?script_id=5476)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1448)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (1347)
3. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (620)
4. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (544)
5. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (421)
6. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (420)
7. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (410)
8. [pathogen.vim : Poor man's package manager. Easy manipulation of 'runtimepath' et al](http://www.vim.org/scripts/script.php?script_id=2332) (410)
9. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (399)
10. [SuperTab : Do all your insert-mode completion with Tab.](http://www.vim.org/scripts/script.php?script_id=1643) (394)

## vim-jp issues

Open : 237 (+0) | Closed : 737 (+16)

- [Issue #959 : Win32版vim.exeをmintty上で実行すると文字化け](https://github.com/vim-jp/issues/issues/959)
- [Issue #960 : Windows版GVimでカーソル表示のゴミが表示される](https://github.com/vim-jp/issues/issues/960)
- [Issue #961 : Windows で job&#x5f;start にダブルクオート付き引数を渡すと正しくエスケープされない。](https://github.com/vim-jp/issues/issues/961)
- [Issue #962 : C言語プリプロセッサで defined が並ぶとシンタックスがおかしい](https://github.com/vim-jp/issues/issues/962)
- [Issue #963 : y コマンドの設定する編集末尾位置マーク  '\]  に一貫性がない](https://github.com/vim-jp/issues/issues/963)
- [Issue #964 : noautocmd で CursorMoved を一時的に無効にできない](https://github.com/vim-jp/issues/issues/964)
- [Issue #965 : getregtype('&#x5f;') が常に 'v' をかえす](https://github.com/vim-jp/issues/issues/965)
- [Issue #966 : matchaddpos() が一部重複する pos を描画しない](https://github.com/vim-jp/issues/issues/966)
- [Issue #967 : job&#x5f;start に環境変数が渡せない](https://github.com/vim-jp/issues/issues/967)
- [Issue #968 : Netrw のバグが多すぎる](https://github.com/vim-jp/issues/issues/968)
- [Issue #969 : job&#x5f;start() してすぐに job&#x5f;stop() すると job プロセスを終了できない場合がある](https://github.com/vim-jp/issues/issues/969)
- [Issue #970 : Test&#x5f;help&#x5f;complete が失敗する](https://github.com/vim-jp/issues/issues/970)
- [Issue #971 : matchaddpos() が行を描画しない](https://github.com/vim-jp/issues/issues/971)
- [Issue #972 : vim.exe の画面リストアで正しくリストア出来ていない事がある](https://github.com/vim-jp/issues/issues/972)
- [Issue #973 : Windows7 で job&#x5f;stop の kill が正しく動作しない](https://github.com/vim-jp/issues/issues/973)
- [Issue #974 : :hideの後にコメントが書けない](https://github.com/vim-jp/issues/issues/974)
