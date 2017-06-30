---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 6 月号

---

## 話題

*   [Meguro.vim #4](https://megurovim.connpass.com/event/57955/) が 07/15 に開催されます。既に定員は超えているようですが、キャンセルや増員がないとも限らないので注目です。
*   [Fabric.vim #2](https://fablicvim.connpass.com/event/60339/) が 07/20 に開催されます。こちらはまだ定員に余裕があります。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.0609: 終了方法の説明が改善されました
*   8.0.0612: パッケージディレクトリが `'runtimepath'` に追加されるタイミングが速まりました
*   8.0.0613: パッケージで ftdetect が使えるようになりました
*   8.0.0621: `:stag` コマンドが `'switchbuf'` オプションを尊重するようになりました
*   8.0.0625: `shellescale()` が第2引数に応じて、改行をエスケープするようになりました
*   8.0.0630: range が1行の時に限り `:g[lobal]` コマンドがネスト可能になりました
*   8.0.0634: `getqflist()` と `setqflist()` の引数 "nr" に最後の quickfix list を取得する "$" を指定できるようになりました
*   8.0.0641: `QuickFixLine` ハイライトグループが追加されました。これは quickfix window の選択されている行に適用されます
*   8.0.0647: `'redrawtime'` オプションがシンタックスハイライトにも適用されるようになりました
*   8.0.0652: Unicode 10 に対応しました
*   8.0.0654: `:endfunction` の後ろの文字を無視しないようになりました
*   8.0.0655: `:delfunction` コマンドに `!` を追加することで、関数を削除しやすく、存在しない関数を指定してもエラーにならなくなりました
*   8.0.0657: `getqflist()` と `setqflist()` の引数に `items` が指定できるようになりました
*   8.0.0661: urxvt でマウスが使えるようになりました。そのエスケープシーケンス `Esc[*M` と `Esc[*m` に対応しました
*   8.0.0664: tmux でマウスが使えるようになりました
*   8.0.0670: タイマーのコールバック関数の中で `input()` を使えるようになりました
*   8.0.0675: `'hlsearch'` の優先度が `'colorcolumn'` よりも高くなりました
*   8.0.0683: ビジュアルベルが実行された後の500ミリ秒間は次のビジュアルベルが実行されなくなりました
*   8.0.0688: `:resize` コマンドが `FileType` autocmd の最中でも利用できるようになりました
*   8.0.0689: 検索時、CTRL-L でコマンドラインにパターンを挿入する際に `~` がエスケープされるようになりました
*   ランタイムの更新
    *   `'shiftwidth'` に 0 を設定してもすべての indent ファイルが問題なく使えるようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.0.0607 : after :bwipe + :new bufref might still be valid](https://github.com/vim/vim/commit/45e5fd135da5710f24a1acc142692f120f8b0b78)
- [8.0.0608 : cannot manipulate other than the current quickfix list](https://github.com/vim/vim/commit/a3921f48c6b31a035c80fda49925dd3b42df0dec)
- [8.0.0609 : some people still don't know how to quit](https://github.com/vim/vim/commit/28a8193e3113f676f89fb6312b099d849df881d3)
- [8.0.0610 : the screen is redrawn when default 'background' is detected](https://github.com/vim/vim/commit/4b974d54434b8020771be7225da94648e9ec961a)
- [8.0.0611 : the screen is not redrawn after sending t&#x5f;u7](https://github.com/vim/vim/commit/976787d1f31451ca7a88e774a03e6c24ddc67876)
- [8.0.0612 : pack dirs are added to 'runtimepath' too late](https://github.com/vim/vim/commit/ce876aaa9a250a5a0d0e34b3a2625e51cf9bf5bb)
- [8.0.0613 : the conf filetype is used before ftdetect from packages](https://github.com/vim/vim/commit/3e54569b17683318e0cb6693ab0024c2ad1e3e8f)
- [8.0.0614 : float2nr() is not exactly right](https://github.com/vim/vim/commit/863e80b4451b5102b41bebf9ddca3a420de746fa)
- [8.0.0615 : using % with :hardcopy wrongly escapes spaces](https://github.com/vim/vim/commit/bf15b8d78b22661db8b19d662b62bb9a061cdd37)
- [8.0.0616 : not always setting 'background' correctly after :hi Normal](https://github.com/vim/vim/commit/1615b36b91b094263240d7b555283ddf33208f62)
- [8.0.0617 : (after 8.0.0615) hardcopy test hangs on MS-Windows](https://github.com/vim/vim/commit/763209c57bf50ae777f9c2929eeea01eff7ae6ee)
- [8.0.0618 : NFA regex engine handles \[0-z\] incorrectly](https://github.com/vim/vim/commit/bb7943b7920ef2f88cb9b6f46c34c7946c370819)
- [8.0.0619 : GUI gets stuck if timer uses feedkeys()](https://github.com/vim/vim/commit/9472eec83c3f9c191814dc81dd82498c10b1fc9c)
- [8.0.0620 : checking for HAVE&#x5f;GTK&#x5f;MULTIHEAD is not needed](https://github.com/vim/vim/commit/b463e8d999ec812d656876f313efbeaeed663b45)
- [8.0.0621 : :stag does not respect 'switchbuf'](https://github.com/vim/vim/commit/8ad80dea089ffeb1a845199c013e9bb4be1cd22e)
- [8.0.0622 : selecting quoted text fails with 'selection' "exclusive"](https://github.com/vim/vim/commit/c5e2b040b490c2f4dd50c945840bc176bfcccb29)
- [8.0.0623 : error for invalid regexp is not very informative](https://github.com/vim/vim/commit/966e58e413ffa88af8d748e697aa2999571fcd7b)
- [8.0.0624 : (after 8.0.0623) warning for unused variable in tiny build](https://github.com/vim/vim/commit/6c95fbc9ae64f3a7619070e830f0c35aa4f0ada9)
- [8.0.0625 : shellescape() always escapes a newline](https://github.com/vim/vim/commit/206155280def51160a9d81d983aed639015ffb44)
- [8.0.0626 : in the GUI the cursor may flicker](https://github.com/vim/vim/commit/e21d69eec1870a3f4732653aa8ee25d5da10128c)
- [8.0.0627 : "gn" selects only one character with 'nowrapscan'](https://github.com/vim/vim/commit/add8dce38de65a0c64e8f54d6bdcadb45a8de2cf)
- [8.0.0628 : cursor disappears after silent mapping](https://github.com/vim/vim/commit/f085f4266e07b36279c56d43fd0d73ed080046ae)
- [8.0.0629 : checking for ambigous width is not working](https://github.com/vim/vim/commit/6b1da3312e15c065b373c9ec2732f31a77cee61f)
- [8.0.0630 : it is not easy to work on lines without a match](https://github.com/vim/vim/commit/f84b122a99da75741ae686fabb6f81b8b4755998)
- [8.0.0631 : can't build with Perl 5.26](https://github.com/vim/vim/commit/fa4161cb0c266378012d3fde0eca56da8739528d)
- [8.0.0632 : the quotestar test is still a bit flaky](https://github.com/vim/vim/commit/1c13c0fe3e7b15750464ffbc39a4648aa5c639be)
- [8.0.0633 : the client-server test is still a bit flaky](https://github.com/vim/vim/commit/60964f68740b8abcbb2d3f0f3aeade21d1bacb22)
- [8.0.0634 : cannot easily get to the last quickfix list](https://github.com/vim/vim/commit/875feea6ce223462d55543735143d747dcaf4287)
- [8.0.0635 : when 'ignorecase' is set script detection is inaccurate](https://github.com/vim/vim/commit/3a429efb628a3925d13c3fe415e02a7ce117071f)
- [8.0.0636 : when reading the undo file fails may use uninitialized data](https://github.com/vim/vim/commit/56f2db562ddc6c69026d55360f0cfaacd8adc26a)
- [8.0.0637 : crash when using some version of GTK 3](https://github.com/vim/vim/commit/21b34b63b3f63dd54b338c4433fa7117dad9507b)
- [8.0.0638 : cannot build with new MSVC version](https://github.com/vim/vim/commit/d7383881c8e9f474a909b54b21abb0aa2b412ba2)
- [8.0.0639 : the cursor position is set when editing a new commit message](https://github.com/vim/vim/commit/9a48961d8bd7ffea14330b9b0181a6cdbe9288f7)
- [8.0.0640 : mismatch between help and actual message](https://github.com/vim/vim/commit/83064068eaabf75a7d235b0eec561dccbcb96b31)
- [8.0.0641 : cannot set a separate highlighting for the quickfix line](https://github.com/vim/vim/commit/2102035488e80ef6fd5038ed15d21672712ba0f6)
- [8.0.0642 : writefile() continues after detecting an error](https://github.com/vim/vim/commit/8cf91286ca46a501d24e4b7d631b193256782c88)
- [8.0.0643 : when a pattern search is slow Vim becomes unusable](https://github.com/vim/vim/commit/fbd0b0af6800f6ff89857863d6a07ea03f09ff6c)
- [8.0.0644 : the timeout for 'hlsearch' is not tested](https://github.com/vim/vim/commit/5b1affefd0e96154517ec6f71300086ae6d22d24)
- [8.0.0645 : no error for illegal back reference in NFA engine](https://github.com/vim/vim/commit/1ef9bbe215e13a273e74fccaddd8fc5a42c76b6e)
- [8.0.0646 : the hlsearch test fails on fast systems](https://github.com/vim/vim/commit/0946326580e6f034fe2c88d041407ea0fde980ab)
- [8.0.0647 : syntax highlighting can cause a freeze](https://github.com/vim/vim/commit/06f1ed2f78c5c03af95054fc3a8665df39dec362)
- [8.0.0648 : possible use of NULL pointer](https://github.com/vim/vim/commit/fadacf01d0dbcc7a96ef5eee0ad57956eeab04d7)
- [8.0.0649 : when opening a help file the filetype is set several times](https://github.com/vim/vim/commit/9049298f8d0bbc237b7c666c7ad6cdabe738e8fc)
- [8.0.0650 : for extra help files the filetype is set more than once](https://github.com/vim/vim/commit/868831f1224cfc17a4de21fb7386873254592308)
- [8.0.0651 : build failure without the auto command feature](https://github.com/vim/vim/commit/157069b04e64eff72d2b34a2f727dfb6e2190cf3)
- [8.0.0652 : unicode information is outdated](https://github.com/vim/vim/commit/383aa84c0d3456c12987ba7e12042dcffba0eac6)
- [8.0.0653 : the default highlight for QuickFixLine is not good](https://github.com/vim/vim/commit/c768a208ca8e0e0fec900c18d5d9a593357ad793)
- [8.0.0654 : no warning for text after :endfunction](https://github.com/vim/vim/commit/663bb2331626944cea156374858131fcd323b9e9)
- [8.0.0655 : not easy to make sure a function does not exist](https://github.com/vim/vim/commit/d6abcd154cdc6a8dd4b7c6ccad37617ea8a1b4aa)
- [8.0.0656 : cannot use ! after some user commands](https://github.com/vim/vim/commit/6f9a476b2f2f0bb5c50d703ec4fc535ffd5bfe8f)
- [8.0.0657 : cannot get and set quickfix list items](https://github.com/vim/vim/commit/6a8958db259d4444da6e6956e54a6513c1af8860)
- [8.0.0658 : spell test is old style](https://github.com/vim/vim/commit/d2c061d24c0534f1f1b92f3462ed6ae8fa848a9a)
- [8.0.0659 : no test for conceal mode](https://github.com/vim/vim/commit/4d785895d1f8b54cdd3fabd87446ca692f49e94e)
- [8.0.0660 : silent install on MS-Windows shows dialog](https://github.com/vim/vim/commit/a1fed064d1eca6ae2c526311ec5464175c6f19e2)
- [8.0.0661 : recognizing urxvt mouse codes does not work well](https://github.com/vim/vim/commit/a529ce068ba84d53343f3732b6a1ed4ad1fe3a68)
- [8.0.0662 : stray FIXME for fixed problem](https://github.com/vim/vim/commit/4670490673ed98502a09b74fbabe785b47e3d289)
- [8.0.0663 : unexpected error message only when 'verbose' is set](https://github.com/vim/vim/commit/f8be461d0284110b321be748fea206d4169b98bb)
- [8.0.0664 : mouse does not work in tmux](https://github.com/vim/vim/commit/6d006f9e950364898240b7a6358516bbf73fe458)
- [8.0.0665 : warning for uninitialized variable](https://github.com/vim/vim/commit/090209bfbd41af5d8c8e44f8366d290431ffbf18)
- [8.0.0666 : dead for loop](https://github.com/vim/vim/commit/5fe691240bff11e9618252486147f0156e875666)
- [8.0.0667 : memory access error when command follows :endfunc](https://github.com/vim/vim/commit/53564f7c1a2998d92568e07fff1f2a4c1cecb646)
- [8.0.0668 : nsis installer script does not work](https://github.com/vim/vim/commit/a1bd86e0f2056f796390bc0cd3aba5c89513d0d2)
- [8.0.0669 : CTRL-N at start of the buffer does not work correctly](https://github.com/vim/vim/commit/24a9e348aa88a6c60ae0cdf5c4a777d8c03b08ca)
- [8.0.0670 : can't use input() in a timer callback](https://github.com/vim/vim/commit/1e8e14552e0cc8881411eb8fbe39a654dae42554)
- [8.0.0671 : hang when typing CTRL-C in confirm() in timer](https://github.com/vim/vim/commit/4eb6531b03445b4d492bc52fea0b6dcd886583af)
- [8.0.0672 : third item of synconcealed() changes too often](https://github.com/vim/vim/commit/cc0750dc6e878394ab0fd922b7ea4280918ae406)
- [8.0.0673 : build failure without conceal feature](https://github.com/vim/vim/commit/ea20de81462a7e1a7f8fc1d53ad618efb0663c17)
- [8.0.0674 : cannot build with eval but without timers](https://github.com/vim/vim/commit/5d7be4f0fa04a9210a14fca69b4f9e34613b378f)
- [8.0.0675 : 'colorcolumn' has a higher priority than 'hlsearch'](https://github.com/vim/vim/commit/774e5a9673260b1b8b88463669213a96637f72e8)
- [8.0.0676 : crash when closing quickfix window in autocmd](https://github.com/vim/vim/commit/182a17b1e80b92826204d967808df0d30eb2ef27)
- [8.0.0677 : setting 'filetype' may switch buffers](https://github.com/vim/vim/commit/1814183b865059679f6ee526ec23fc575e536e66)
- [8.0.0678 : closing a window does not trigger resizing](https://github.com/vim/vim/commit/8eeeba8c025ff844e6514c4a60cec11bf1fc1b35)
- [8.0.0679 : using freed memory](https://github.com/vim/vim/commit/41cc038ff83498c589c7d25b3d2984145528eb92)
- [8.0.0680 : plugins in start packages are sourced twice](https://github.com/vim/vim/commit/07ecfa64a18609a986f21d6132d04ee8934f3200)
- [8.0.0681 : unnamed register only contains the last deleted text](https://github.com/vim/vim/commit/18d90b95c49d9ff1c635dd762864022aab8e71f1)
- [8.0.0682 : no test for synIDtrans()](https://github.com/vim/vim/commit/0b2eef24bcbe2c85c90bbde914a1782cbedc5c72)
- [8.0.0683 : visual bell flashes too quickly](https://github.com/vim/vim/commit/2e147caa14f622dfd1c1def8e07c113b9b85d4b2)
- [8.0.0684 : old style tests are not nice](https://github.com/vim/vim/commit/28b238225ae618f63cfe5d3d723120960a941da7)
- [8.0.0685 : when conversion fails written file may be truncated](https://github.com/vim/vim/commit/e6bf655bc4de1b7f4586e1f5c2fc4978141c3aa3)
- [8.0.0686 : extra redraw when using CTRL-L in second window](https://github.com/vim/vim/commit/9f5f7bf4d5f757ef885acf74ce03c25429a328aa)
- [8.0.0687 : minor issues related to quickfix](https://github.com/vim/vim/commit/86f100dc0922e83bead7bcd5fd2bb2abbf153f46)
- [8.0.0688 : cannot resize the window in a FileType autocommand](https://github.com/vim/vim/commit/9c4fefffb65a2ed9b4a5b0f1bde0da8f349470b5)
- [8.0.0689 : ~ character not escaped when extending search pattern](https://github.com/vim/vim/commit/a693d0584b9a7ccce98813dda3a6badb209904c7)
- [8.0.0690 : compiler warning on non-Unix system](https://github.com/vim/vim/commit/c2226845eb207bcd1a24e1afa941acbfecbc5170)
- [8.0.0691 : compiler warning without the linebreak feature](https://github.com/vim/vim/commit/a83fe75ca76f6f74daa1e5a620054b973b43379b)

## 新着スクリプト

- [golint-fixer.vim : A vim tool for fix goling warnings, which generate by vim-go or ale.](http://www.vim.org/scripts/script.php?script_id=5574)
- [WWDC17 : Colorful light color scheme inspired by WWDC17 page](http://www.vim.org/scripts/script.php?script_id=5575)
- [Absolute : Light Colorscheme](http://www.vim.org/scripts/script.php?script_id=5576)
- [CompleteParameter : complete parameter after select the completion.](http://www.vim.org/scripts/script.php?script_id=5577)
- [LiterateVimrc : Plugin for creating ](http://www.vim.org/scripts/script.php?script_id=5578)
- [change-case.vim : Operator to change case (camelCase, snake&#x5f;case, UPPER CASE etc.) of text](http://www.vim.org/scripts/script.php?script_id=5579)
- [ZFVimTxtHighlight : highlight plain text files, suit most cases](http://www.vim.org/scripts/script.php?script_id=5580)
- [ZFVimIndentMove : move cursor quickly accorrding to indent](http://www.vim.org/scripts/script.php?script_id=5581)
- [ZFVimFoldBlock : fold code block using regexp](http://www.vim.org/scripts/script.php?script_id=5582)
- [DeleteHiddenBuffers : Clean your file buffer when you work for long periods on a large project](http://www.vim.org/scripts/script.php?script_id=5583)
- [Vim PHP Wrapper : A plug-in to speed up PHP development](http://www.vim.org/scripts/script.php?script_id=5584)
- [stunter.vim : stunter.vim: run vim test](http://www.vim.org/scripts/script.php?script_id=5585)
- [factorus : Vim plugin for automated refactoring](http://www.vim.org/scripts/script.php?script_id=5586)

## 月間ダウンロードランキング

1. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](http://www.vim.org/scripts/script.php?script_id=273) (1078)
2. [The NERD tree : A tree explorer plugin for navigating the filesystem](http://www.vim.org/scripts/script.php?script_id=1658) (886)
3. [wombat256.vim : Wombat for 256 color xterms](http://www.vim.org/scripts/script.php?script_id=2465) (614)
4. [python.vim : Enhanced version of the python syntax highlighting script](http://www.vim.org/scripts/script.php?script_id=790) (498)
5. [AutomaticLaTeXPlugin : Background compilation, completion, bib serch, toc and other nice features.](http://www.vim.org/scripts/script.php?script_id=2945) (397)
6. [OmniCppComplete : C/C++ omni-completion with ctags database](http://www.vim.org/scripts/script.php?script_id=1520) (359)
7. [nginx.vim : initial version](http://www.vim.org/scripts/script.php?script_id=1886) (355)
8. [molokai : A port of the monokai scheme for TextMate](http://www.vim.org/scripts/script.php?script_id=2340) (354)
9. [c.vim : C/C++ IDE -- Write and run programs. Insert statements, idioms, comments etc.](http://www.vim.org/scripts/script.php?script_id=213) (344)
10. [winmanager : A windows style IDE for Vim 6.0](http://www.vim.org/scripts/script.php?script_id=95) (303)

## vim-jp issues

Open : 250 (-5) | Closed : 806 (+6)

- [Issue #1056 : autocmdで自動でcopenした時にBufReadPost quickfixが反応しない](https://github.com/vim-jp/issues/issues/1056)
