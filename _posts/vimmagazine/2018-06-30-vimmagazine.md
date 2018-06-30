---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 06 月号

---

## 話題

*   [asakusa.vim #1](https://asakusavim.connpass.com/event/90854/) が 2018-06-21 (木) に開催され、Vim使いたちの貴重な知見が共有されました
*   [meguro.vim #10](https://megurovim.connpass.com/event/89330/) が 2018-06-23 (土) に開催されました
*   VimConf 2018 関連
    *   VimConf 2018 が 2018-11-24 に秋葉原のアキバPLAZAで開催されます
    *   [VimConf 2018のスポンサー公募](https://vim-jp.org/blog/2018/06/07/VimConf2018-sponsor-ja.html) が始まっています
    *   [VimConf 2018のスピーカーの公募](https://vim-jp.org/blog/2018/06/28/VimConf2018-cfp-ja.html) が始まっています

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.1.0011: `maparg()` 関数と `mapcheck()` 関数が rhs が `<Nop>` であるマッピングに対して `"<Nop>"` という文字列を返すようになりました (今までは rhs が `<Nop>` あるいは存在しないマッピング両方に対して空文字列を返していました)
*   8.1.0027: プロンプトバッファ (`set buftype=prompt`) が追加されました (関連: 8.1.0028, 8.1.0032, 8.1.0035-0036, 8.1.0042, 8.1.0064, 8.1.0069-0071)
    *   関数 `prompt_addtext()`, `prompt_setprompt()`, `prompt_setcallback()` が追加されました
*   8.1.0037: 任意のバッファに行を追記できる関数 `appendbufline()` が追加されました
*   8.1.0039: 任意のバッファの行を削除できる関数 `deletebufline()` が追加されました
*   8.1.0047: `:unlet $VAR` にコマンドライン補完が追加されました
*   8.1.0049: `:terminal` を実行した時にシェルから参照可能な環境変数 `VIM_TERMINAL` が追加されました
*   8.1.0091: 指定した pid に `SIGTRAP` シグナルを送る `debugbreak()` 関数が追加されました (Windows でのみ使用可能です)
*   8.1.0105: 可変タブ機能 (`+vartabs`) が追加されました
    *   `'varsofttabstop'`, `'vartabstop'` に設定されたタブ幅のリストが `'softtabstop'`, `'tabstop'` の代わりに使用されます
*   8.1.0110: `'shortmess'` オプションが `F` フラグを含んでいても `:file` コマンドの出力が省略されないようになりました

*   ランタイムの更新
    *   PHP のインデントスクリプトに `g:PHP_noArrowMatching`, `b:PHP_noArrowMatching` 変数が追加されました (`:help PHP_noArrowMatching`)。1 にするとメソッド呼び出しの `->` の位置を揃えないようになります

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。

## リリース情報

- [8.1.0001 : the netrw plugin does not work](https://github.com/vim/vim/commit/00e45563384de01949b9c4df9155424604887b33)
- [8.1.0002 : :stopinsert changes the message position](https://github.com/vim/vim/commit/2abad54cede7c33f090fbde19a9a3eb3dbf78e53)
- [8.1.0003 : the :compiler command is not tested](https://github.com/vim/vim/commit/10561fe0a37efd611ccb1517b37de12d2e37af0f)
- [8.1.0004 : test for :compiler command sometimes fails](https://github.com/vim/vim/commit/54651f74948650bc68d6ac5b20376f2f985e80c5)
- [8.1.0005 : test for :compiler command fails on MS-Windows](https://github.com/vim/vim/commit/d19b234154476401a6ad8e690f9705419095f44e)
- [8.1.0006 : syn&#x5f;id2cterm&#x5f;bg() may be undefined](https://github.com/vim/vim/commit/a772baf85afac44bb8eb20ff808637ee6eedec54)
- [8.1.0007 : no test for "o" and "O" in Visual block mode](https://github.com/vim/vim/commit/2e94976abd1cd6b94db38d4f2a1cfd71808b8100)
- [8.1.0008 : no test for strwidth()](https://github.com/vim/vim/commit/42ab17b8e32352210c4e273a4a4161a287d2c159)
- [8.1.0009 : tabpages insufficiently tested](https://github.com/vim/vim/commit/dbe8869c41a3b3309207a434c17db295d358e873)
- [8.1.0010 : efm&#x5f;to&#x5f;regpat() is too long](https://github.com/vim/vim/commit/6bff719f7e472e918c60aa336de03e799b806c4f)
- [8.1.0011 : maparg() and mapcheck() confuse empty and non-existing](https://github.com/vim/vim/commit/f88a5bc10232cc3fac92dba4e8455f4c14311f8e)
- [8.1.0012 : misplaced #endif](https://github.com/vim/vim/commit/1f131aee0ed4383009c0a67cc15581cc4c0dcad7)
- [8.1.0013 : using freed memory when changing terminal cursor color](https://github.com/vim/vim/commit/4f7fd56b423245786f223a370c6fbfc12e3f37eb)
- [8.1.0014 : qf&#x5f;init&#x5f;ext() is too long](https://github.com/vim/vim/commit/6053f2d29a979ffed1fe01b0a2f28e23750530e9)
- [8.1.0015 : cursor color wrong when closing a terminal window](https://github.com/vim/vim/commit/05af9a419198245f0810301ac9a3d59a9432ef21)
- [8.1.0016 : possible crash in term&#x5f;wait()](https://github.com/vim/vim/commit/d45aa55d42211eb2f42e0a04ff77537a5df51c40)
- [8.1.0017 : shell command completion has duplicates](https://github.com/vim/vim/commit/62fe66f251263715968442e237742d9d3dfd5fa1)
- [8.1.0018 : using "gn" may select wrong text when wrapping](https://github.com/vim/vim/commit/bdb657924d73c98b0ab28411749571e893b699a9)
- [8.1.0019 : error when defining a Lambda with index of a function result](https://github.com/vim/vim/commit/b4518563c73460150344a57879bf5b22cb8b1c77)
- [8.1.0020 : cannot tell whether a register is executing or recording](https://github.com/vim/vim/commit/0b6d911e5de1a1c10a23d4c2ee1b0275c474a2dd)
- [8.1.0021 : clang warns for undefined behavior](https://github.com/vim/vim/commit/d315cf551f1d15609c4d7cf724e471de55f5cdac)
- [8.1.0022 : repeating put from expression register fails](https://github.com/vim/vim/commit/833093bfb0e4a7f89b5adc66babcfa8ac09cfda9)
- [8.1.0023 : gcc 8.1 warns for use of strncpy()](https://github.com/vim/vim/commit/22c105640cdef25d98390f973d2ed53dd9039f6e)
- [8.1.0024 : % command not testded on #ifdef and comment](https://github.com/vim/vim/commit/6b69e5c646faecf2a686bfe71dc17a2a1c6d344f)
- [8.1.0025 : no test for the undofile() function](https://github.com/vim/vim/commit/e5fa11186fde4a19e505eba403d3af8c61d11304)
- [8.1.0026 : terminal test fails with very tall terminal](https://github.com/vim/vim/commit/33c5e9fa7af935c61a8aac461b9664c501003440)
- [8.1.0027 : difficult to make a plugin that feeds a line to a job](https://github.com/vim/vim/commit/f273245f6433d5d43a5671306b520a3230c35787)
- [8.1.0028 : prompt buffer test fails on MS-Windows](https://github.com/vim/vim/commit/1149382d21402474c771862d082a541e23beb108)
- [8.1.0029 : terminal test fails on MS-Windows when "wc" exists](https://github.com/vim/vim/commit/1580f755c91d12a3e63d60eec8733f8b79debc31)
- [8.1.0030 : stoping Vim running in a terminal may not work](https://github.com/vim/vim/commit/3339d3dad93a44eac76c71cfa48c86a20d3dddcc)
- [8.1.0031 : terminal test aucmd&#x5f;on&#x5f;close is flaky](https://github.com/vim/vim/commit/c8523e2e6cd072d86a9412f465aa9eef53f5675e)
- [8.1.0032 : BS in prompt buffer starts new line](https://github.com/vim/vim/commit/6b810d92a9cd9378ab46ea0db07079cb789f9faa)
- [8.1.0033 : keys to stop Vim in terminal are wrong](https://github.com/vim/vim/commit/acb9effecc9f54b93398a44cc0ec40e77978e094)
- [8.1.0034 : cursor not restored with ":edit #"](https://github.com/vim/vim/commit/adb8fbec4f4059d214fe6acf2485ffd35e803450)
- [8.1.0035 : not easy to switch between prompt buffer and other windows](https://github.com/vim/vim/commit/6d41c78e353b630bc1a72cbff9160311d2a81e8c)
- [8.1.0036 : not restoring Insert mode if leaving prompt buffer with mouse](https://github.com/vim/vim/commit/891e1fd894720d0b99a9daefa41e8181844f819a)
- [8.1.0037 : cannot easily append lines to another buffer](https://github.com/vim/vim/commit/ca851593a660f08aba5c134f90c238d4a3e983e6)
- [8.1.0038 : popup test causes Vim to exit](https://github.com/vim/vim/commit/680c99b2efc7e2a9e6b3d611133f43fb7f8c6811)
- [8.1.0039 : cannot easily delete lines in another buffer](https://github.com/vim/vim/commit/d79a26219d7161e9211fd144f0e874aa5f6d251e)
- [8.1.0040 : warnings from 64-bit compiler](https://github.com/vim/vim/commit/e31e256ba1769a3a3ed7840d5cc9a01ab058b8bc)
- [8.1.0041 : attribute "width" missing from python window attribute list](https://github.com/vim/vim/commit/e950f9992b291a07e4fb41cb561156f5b382cd5d)
- [8.1.0042 : if omni completion opens a window Insert mode is stopped](https://github.com/vim/vim/commit/f98b845dd185dfadfa7a622a42452bfa6809d4e0)
- [8.1.0043 : ++bad argument of :edit does not work properly](https://github.com/vim/vim/commit/7580849df9d6c7c515f5ed784019336d8a8ec0f0)
- [8.1.0044 : if a test function exists Vim this may go unnoticed](https://github.com/vim/vim/commit/8903676d3d4c3fcb37d2a55c67370f30806b1626)
- [8.1.0045 : popup test isn't run completely](https://github.com/vim/vim/commit/ae0f30b761eb62e1b6bfc83fb4a6d1a47bf48320)
- [8.1.0046 : loading a session file fails if 'winheight' is big](https://github.com/vim/vim/commit/1c3c10492a291270fa89b3c8df11828792f927d3)
- [8.1.0047 : no completion for :unlet $VAR](https://github.com/vim/vim/commit/19834010889fc5bfa0f88b3ba83133dae6c0a35d)
- [8.1.0048 : vim&#x5f;str2nr() does not handle numbers close to the maximum](https://github.com/vim/vim/commit/07ccf7ce7fb948fd4d080b817e9fbaea9e721dab)
- [8.1.0049 : shell cannot tell running in a terminal window](https://github.com/vim/vim/commit/d7a137fb0d980545dd567bee9c24cf7b9c3a2eae)
- [8.1.0050 : $VIM&#x5f;TERMINAL is also set when not in a terminal window](https://github.com/vim/vim/commit/493359eb3b10377d5c3524e91d911809b8ac7a76)
- [8.1.0051 : MS-Windows: missing #endif](https://github.com/vim/vim/commit/79c6b51e403d4ea755b6637b01d1426a06ca9f9a)
- [8.1.0052 : when mapping to \<Nop> times out the next mapping is skipped](https://github.com/vim/vim/commit/83f4cbd973731872b633d6ba0caf850fb708d70c)
- [8.1.0053 : first argument of 'completefunc' has inconsistent type](https://github.com/vim/vim/commit/ffa9684150f5441e84d492e7184ef73587bd6c6c)
- [8.1.0054 : compiler warning for using %ld for "long long"](https://github.com/vim/vim/commit/5ecdf96422af3171cf948865bf057c2a61f042bc)
- [8.1.0055 : complete test has wrong order of arguments](https://github.com/vim/vim/commit/52d3aaeea4c85406f46e2051729e4f007f83b35e)
- [8.1.0056 : crash when using :hardcopy with illegal byte](https://github.com/vim/vim/commit/43dee181f596c81b99e200b6cdfeb02ecfed42c8)
- [8.1.0057 : popup menu displayed wrong when using autocmd](https://github.com/vim/vim/commit/6ba3ec1bace67513a352326864cebc16b3c5bc56)
- [8.1.0058 : display problem with margins and scrolling](https://github.com/vim/vim/commit/bfa4246768e28335ed9b98f83019ea58b480158e)
- [8.1.0059 : displayed digraph for "ga" wrong with 'encoding' "cp1251"](https://github.com/vim/vim/commit/bc5020aa4d7ef4aea88395eff858f74fc881eab9)
- [8.1.0060 : crash when autocommands delete the current buffer](https://github.com/vim/vim/commit/600323b4ef51a58a8e800d8ca469383a3c911db7)
- [8.1.0061 : window title is wrong after resetting and setting 'title'](https://github.com/vim/vim/commit/84a9308511871d9ff94c91a1c6badb92300ded98)
- [8.1.0062 : popup menu broken if a callback changes the window layout](https://github.com/vim/vim/commit/491ac28d5f91505519c623ebc1a9ab08834bf367)
- [8.1.0063 : Mac: NSStringPboardType is deprecated](https://github.com/vim/vim/commit/d595a1910c5672808e44afa028e253e47f03651f)
- [8.1.0064 : typing CTRL-W in a prompt buffer shows mode "-- --"](https://github.com/vim/vim/commit/942b4541a2d8e8df8369ab70e112dbbbe0c7c0aa)
- [8.1.0065 : (after 8.1.0062) balloon displayed at the wrong position](https://github.com/vim/vim/commit/0e6e179f55b85b9e82e74d3f993928d52f0f4d43)
- [8.1.0066 : nasty autocommand causes using freed memory](https://github.com/vim/vim/commit/c3ffc9b8d3015dc5280b297b4e3deb4f34944bd4)
- [8.1.0067 : syntax highlighting not working when re-entering a buffer](https://github.com/vim/vim/commit/a5616b0136cea2104a475d143a1685d71e9b2d3d)
- [8.1.0068 : nasty autocommands can still cause using freed memory](https://github.com/vim/vim/commit/2f82ca7d79148ae931bf28a747ede06ba8a65de8)
- [8.1.0069 : cannot handle pressing CTRL-C in a prompt buffer](https://github.com/vim/vim/commit/0e5979a6d491f68c4a8c86fab489016919329a6b)
- [8.1.0070 : missing part of the changes for prompt&#x5f;setinterrupt()](https://github.com/vim/vim/commit/222cd20e2662e7478cfe42b78cc4f1c153ca819d)
- [8.1.0071 : terminal debugger only works with the terminal feature](https://github.com/vim/vim/commit/b3307b5e7e7bd3962b0d5c61a94e638564c146b0)
- [8.1.0072 : use of 'termwinkey' is inconsistent](https://github.com/vim/vim/commit/dcdeaaf1505b96aa7a31ccd4becc92c32119bdda)
- [8.1.0073 : crash when autocommands call setloclist()](https://github.com/vim/vim/commit/0366c0161e988e32420d2f37111a60129684905b)
- [8.1.0074 : (after 8.1.0073) crash when running quickfix tests](https://github.com/vim/vim/commit/d6b01a2d3884ee0abe10aad161f584889a496e3f)
- [8.1.0075 : no Vim logo in README file](https://github.com/vim/vim/commit/252b7ee82384520e20df7728387a1a441c47c6e6)
- [8.1.0076 : command getting cleared with CTRL-W : in a terminal window](https://github.com/vim/vim/commit/0ce7413a8318ec0b01386c54ee09d3d94216cb15)
- [8.1.0077 : header of README file is not nice](https://github.com/vim/vim/commit/5c3670718bebacb3a9a54522cab2924a6bfbc3d4)
- [8.1.0078 : "..." used inconsistently in messages](https://github.com/vim/vim/commit/c166927a32fe5c054ad35deecff00aa12c629cf7)
- [8.1.0079 : superfluous space in messages](https://github.com/vim/vim/commit/8df6e5d4670891608e791244b0c2ec0db387f710)
- [8.1.0080 : can't see the breakpoint number in the terminal debugger](https://github.com/vim/vim/commit/de1a83147a28d3db8f9fca415ffc8fa04218eac2)
- [8.1.0081 : the terminal debugger doesn't adjust to changed 'background'](https://github.com/vim/vim/commit/f07f9e731eb97bbdbd1b0b3983750589e4f557a6)
- [8.1.0082 : in terminal window, typing : at more prompt, inserts ':'](https://github.com/vim/vim/commit/1d4754f96fa5bff1c349cdb71560c55675f50d03)
- [8.1.0083 : "is" and "as" have trouble with quoted punctuation](https://github.com/vim/vim/commit/8516071124dbb7ad7caa43cc98ae3c57ae093c9e)
- [8.1.0084 : user name completion does not work on MS-Windows](https://github.com/vim/vim/commit/828c3d70833a0689cc07581f2a67d06430675da5)
- [8.1.0085 : no test for completing user name and language](https://github.com/vim/vim/commit/5f8f2d378a4f6d7db12806f3e35ec6f7fc6bd1f3)
- [8.1.0086 : no tests for libcall() and libcallnr()](https://github.com/vim/vim/commit/1ceebb4efc455dc6c34e0cd2c2adbd00939f038b)
- [8.1.0087 : v:shell&#x5f;error is always zero when using terminal for "!cmd"](https://github.com/vim/vim/commit/f9c3883b11b33f0c548df5e949ba59fde74d3e7b)
- [8.1.0088 : terminal test for stdout and stderr is a bit flaky](https://github.com/vim/vim/commit/5319191a2a726c18a153ee53e9fac506340d16c7)
- [8.1.0089 : error when ending the terminal debugger](https://github.com/vim/vim/commit/a15b0a936d2013e99e9323da9b920a5d93e4fb5a)
- [8.1.0090 : "..." used inconsistently in a message](https://github.com/vim/vim/commit/9b0c5c23bd5260caef82a4f3dcc945c129857c52)
- [8.1.0091 : MS-Windows: Cannot interrupt gdb when program is running](https://github.com/vim/vim/commit/4551c0a9fcdbdef52836d4852686d54b5e47fdaf)
- [8.1.0092 : (after 8.1.0091) prompt buffer test fails](https://github.com/vim/vim/commit/71ef1ba5e996f34d3e0acbe1d89c4c6bfa5e98ba)
- [8.1.0093 : non-MS-Windows: Cannot interrupt gdb when program is running](https://github.com/vim/vim/commit/2ed890f1f810f977ec6a235efd8bf58adddcd0e7)
- [8.1.0094 : help text "usage:" is not capatalized](https://github.com/vim/vim/commit/32aaf5ae07af1d5befaf997688fde5d6637d43f8)
- [8.1.0095 : dialog for ":browse tabnew" says "new window"](https://github.com/vim/vim/commit/39902a06d92750c203d86c921b9d69995f949d97)
- [8.1.0096 : inconsistent use of the word autocommands](https://github.com/vim/vim/commit/8c55533c6f109db2a0fff69651887f9474eb09c6)
- [8.1.0097 : superfluous space before exclamation mark](https://github.com/vim/vim/commit/5efa0102de6ed6049fb19e1e83787e5b3b24b6a2)
- [8.1.0098 : segfault when pattern with \\z() is very slow](https://github.com/vim/vim/commit/bcf9442307075bac40d44328c8bf7ea21857b138)
- [8.1.0099 : exclamation mark in error message not needed](https://github.com/vim/vim/commit/3c867daaf09e8ac6ce4b9d43d6fbbfdd7689702d)
- [8.1.0100 : terminal debugger: error when setting a watch point](https://github.com/vim/vim/commit/6dccc962f3d019f94656e77309aa00af7bffa4c3)
- [8.1.0101 : no test for getcmdwintype()](https://github.com/vim/vim/commit/81612b7a7dc50a6c0da9f42fa48f1f576a4ad616)
- [8.1.0102 : cannot build without syntax highlighting](https://github.com/vim/vim/commit/1f8c469fe22adf7b42843c3abf118ed339e6da0e)
- [8.1.0103 : long version string cannot be translated](https://github.com/vim/vim/commit/35fb6fbf727c72eb5987a1556bbba26dca14cc48)
- [8.1.0104 : can't build without the +eval feature](https://github.com/vim/vim/commit/5ec7414a1c0512832f60c17437d6374cbf4b08e9)
- [8.1.0105 : all tab stops are the same](https://github.com/vim/vim/commit/04958cbaf25eea27eceedaa987adfb354ad5f7fd)
- [8.1.0106 : build fails when HAVE&#x5f;DATE&#x5f;TIME is undefined](https://github.com/vim/vim/commit/eee3e9466951e6619e7db51cd027bafdfbebb08d)
- [8.1.0107 : Python: getting buffer option clears message](https://github.com/vim/vim/commit/defe6424aee6201241b7cb231b62db4bbb9f4a9f)
- [8.1.0108 : no Danish translations](https://github.com/vim/vim/commit/a4a2934e59429e2ff68d478556ad8b2043fb0fda)
- [8.1.0109 : new po makefile missing from distribution](https://github.com/vim/vim/commit/7cb6eecd32b0b51bd0a25df968386d0b09379afa)
- [8.1.0110 : file name not displayed with ":file"](https://github.com/vim/vim/commit/fc0896093c3b3e753859a5f929921933e7a2e6cd)
- [8.1.0111 : .po files do not use recommended names](https://github.com/vim/vim/commit/d1d037e90192ab64e4cec00b9d008b33bc69d979)
- [8.1.0112 : no error when using bad arguments with searchpair()](https://github.com/vim/vim/commit/3dddb09c98825acefa6f2d94bb369b8e00d7b3e5)
- [8.1.0113 : compiler warning for unused variable](https://github.com/vim/vim/commit/829adb74605604ff0bec33862cde850efb294711)
- [8.1.0114 : confusing variable name](https://github.com/vim/vim/commit/0119a59ffdfb21cf1c0a56e7ed6105e875150163)
- [8.1.0115 : the matchparen plugin may throw an error](https://github.com/vim/vim/commit/b7a5ab112ab2267c54acdb6a326642afcd797bc4)
- [8.1.0116 : display problem with 'vartabstop' and 'linebreak'](https://github.com/vim/vim/commit/a87b72cc316e065d66dcbcf7ec1cde330adef3a3)
- [8.1.0117 : URL in install program still points to SourceForge](https://github.com/vim/vim/commit/bd87eb38c558521ff709651ac190ffd1b9dee057)
- [8.1.0118 : duplicate error message for put command](https://github.com/vim/vim/commit/f52f9ea8f5fb3df51a308c56f2bf66f735ef3ca7)
- [8.1.0119 : failing test goes unnoticed because messages is not written](https://github.com/vim/vim/commit/ec12d6490923fb5ff147ebf7e28d32f2b4977e3b)
- [8.1.0120 : buffer 'modified' set even when :sort has no changes](https://github.com/vim/vim/commit/dc9e955fb07f410d5d3e981ce18d895dd2847c85)
- [8.1.0121 : crash when using ballooneval related to 'vartabstop'](https://github.com/vim/vim/commit/ca4b613c692b927c9345dc9e3d6be5de35266be2)
- [8.1.0122 : translators don't always understand the maintainer message](https://github.com/vim/vim/commit/0c183198727b80fd7e52fb29cc1dcd08d79a281c)
- [8.1.0123 : MS-Windows: colors are wrong after setting 'notgc'](https://github.com/vim/vim/commit/b0eb14f9392a8d95045182c4e520530319ce44c2)
- [8.1.0124 : has('vcon') returns true even for non-win32 terminal](https://github.com/vim/vim/commit/d8b37a53bd29cab78c6997aa75207385213f23e2)
- [8.1.0125 : virtual edit replace with multi-byte fails at end of line](https://github.com/vim/vim/commit/630afe889a2a02b367ea8eaaa48e66ed81e77ff3)
- [8.1.0126 : various problems with 'vartabstop'](https://github.com/vim/vim/commit/307ac5c68e9e624ab713136d79f35bb73f780d2d)
- [8.1.0127 : build failure when disabling the session feature](https://github.com/vim/vim/commit/7c365fb190225d98af818daf195dc19921fe01ca)

## 新着スクリプト

- [vim-indexer : Vim Indexer plugin that provides async tags generation with project management.](https://www.vim.org/scripts/script.php?script_id=5699)
- [voodooChild.vim : Low contrast colorscheme for low light conditions.](https://www.vim.org/scripts/script.php?script_id=5700)
- [smullyan.vim : Light colorscheme for brightly lit conditions.](https://www.vim.org/scripts/script.php?script_id=5701)
- [j.vim : j syntax](https://www.vim.org/scripts/script.php?script_id=5702)
- [k.vim : syntax for the k language.](https://www.vim.org/scripts/script.php?script_id=5703)
- [git-blame.vim :  See Git Blame information in the status bar for the currently selected line.](https://www.vim.org/scripts/script.php?script_id=5704)
- [put&#x5f;blank&#x5f;lines.vim : Mapping targets to insert blank lines](https://www.vim.org/scripts/script.php?script_id=5705)
- [markdown&#x5f;autoformat.vim : Automatic formatting heuristics for Markdown](https://www.vim.org/scripts/script.php?script_id=5706)
- [redact&#x5f;pass.vim : Set options to edit pass(1) files securely](https://www.vim.org/scripts/script.php?script_id=5707)
- [vim-colors-nobold : Vim color schemes with "bold" attribute removed - except for text console.](https://www.vim.org/scripts/script.php?script_id=5708)
- [peachpuff256mod : peachpuff256mod is the terminal version of the vim builtin peachpuff color theme](https://www.vim.org/scripts/script.php?script_id=5709)
- [object.vim : A pythonic object oriented framework.](https://www.vim.org/scripts/script.php?script_id=5710)
- [paste&#x5f;open.vim : Mapping targets for opening new lines in paste mode](https://www.vim.org/scripts/script.php?script_id=5711)
- [diff&#x5f;prune.vim : Reverse changes in selected lines of a diff](https://www.vim.org/scripts/script.php?script_id=5712)
- [clear&#x5f;local&#x5f;maps.vim : Clear buffer-local leader mappings](https://www.vim.org/scripts/script.php?script_id=5713)
- [vim-json-format : format json text ](https://www.vim.org/scripts/script.php?script_id=5714)
- [make&#x5f;target.vim : Make target for Makefile recipe under cursor](https://www.vim.org/scripts/script.php?script_id=5715)
- [cmdwin&#x5f;ctrlc.vim : Tweak CTRL-C command line window behavior to be a bit more useful](https://www.vim.org/scripts/script.php?script_id=5716)

## 月間ダウンロードランキング

1. [VimCoder.jar : The TopCoder Arena editor plugin with Vim integration.](https://www.vim.org/scripts/script.php?script_id=3321) (4)
2. [AutoComplPop : Automatically opens popup menu for completions](https://www.vim.org/scripts/script.php?script_id=1879) (4)
3. [Windows PowerShell Syntax Plugin : Syntax coloring, indenting and filetype detection for Windows PowerShell](https://www.vim.org/scripts/script.php?script_id=1327) (3)
4. [matchit.zip : extended % matching for HTML, LaTeX, and many other languages](https://www.vim.org/scripts/script.php?script_id=39) (3)
5. [systemverilog.vim : Indent & syntax script for Verilog and SystemVerilog](https://www.vim.org/scripts/script.php?script_id=4743) (3)
6. [verilog&#x5f;systemverilog.vim : Extending Verilog syntax highlighting for SystemVerilog](https://www.vim.org/scripts/script.php?script_id=1586) (3)
7. [taglist.vim : Source code browser (supports C/C++, java, perl, python, tcl, sql, php, etc)](https://www.vim.org/scripts/script.php?script_id=273) (3)
8. [Align : Help folks to align text, eqns, declarations, tables, etc](https://www.vim.org/scripts/script.php?script_id=294) (3)
9. [fugitive.vim : A Git wrapper so awesome, it should be illegal](https://www.vim.org/scripts/script.php?script_id=2975) (3)
10. [SuperTab : Do all your insert-mode completion with Tab.](https://www.vim.org/scripts/script.php?script_id=1643) (2)

## vim-jp issues

Open : 260 (-1) | Closed : 914 (+5)

- [Issue #1171 : MacVimで標準日本語IMの「Windows風のキー操作」を有効にすると\<C-u>や\<C-o>が効かなくなる](https://github.com/vim-jp/issues/issues/1171)
- [Issue #1172 : 本家helpのtypo報告スレ (Vim 8.1)](https://github.com/vim-jp/issues/issues/1172)
- [Issue #1173 : 現在のモードに応じたハイライトグループがあったらどうだろうか?](https://github.com/vim-jp/issues/issues/1173)
- [Issue #1174 : &#x60;:verbose set xxx?&#x60; でファイル名だけでなく行番号も欲しい](https://github.com/vim-jp/issues/issues/1174)
