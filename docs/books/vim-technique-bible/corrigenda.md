---
layout: docs
title: Vimテクニックバイブル 作業効率をカイゼンする150の技 正誤表
---
# 1-2 project.vimを使いやすくする
* P.26 「インストール・設定」を差し替えます。<br />

> 誤：本Tipsを利用するには、リスト1の設定をVimエディタの設定ファイルに追加してください。<br />
> 正：本Tipsを利用するには、設定ファイルディレクトリにafter/pluginディレクトリを作成し、project.vimという名前でリスト1の設定ファイルを用意してください。<br />

<hr />

# 1-14 Vimプラグインを管理する
* P.51 pathogen.vim との違い<br />

> 誤：特定のディレクトリを 'runtimepath' に追加するしたり削除したりできます。<br />
> 正：特定のディレクトリを 'runtimepath' に追加したり削除したりできます。<br />

<hr />

# 2-2 vimfilerを用いてファイル操作を行う

※：正誤表というより、補足情報です。

<https://github.com/Shougo/vimfiler>

現在、私のvimfilerリポジトリにて、vimfiler Ver.3が開発中です。
vimfiler Ver.3はp.358にもあるように、unite.vim Ver.3に依存する形で再設計をされています。
よって、unite.vim Ver.3がインストールされていないと動作しません。注意してください。

さらに、次のキーバインドが変更されています。

* P.58 表1 <br />

マッピング$は削除されています。ゴミ箱の汎用的な実装ができなかったためです。
ゴミ箱が欲しい場合は、外部コマンドを利用してください。

マッピングl(エル)は「カーソル位置のファイルを開く／カーソル位置のディレクトリに移動する」となっていますが、
正確には、「vimfilerの関連付けでカーソル位置のファイルを開く／カーソル位置のディレクトリに移動する」です。
「カーソル位置のファイルをVimで開く(:editで開く)」のは、マッピングeとなります。

* P.58 表2 <br />

マッピングdはゴミ箱へ移動するのではなく、ファイルの強制削除となります。これにより、マッピングDは削除されています。

* P.59 補足　unite.vimのブックマーク情報と連携する<br />

「`<C-p>/<C-n>`でヒストリを辿ることができます」とありますが、vimfiler Ver.3では、この機能に相当するのはデフォルトで`<C-j>`となります。
`<C-p>/<C-n>`のマッピングは削除されています。
履歴移動はuniteインタフェース化され、より使いやすくなっています。

<hr />

# 2-7 ファイルの文字コードを変換する
* P.67 文字不正時の処理を指定する<br />

> 誤：ファイルの中にエンコーディングで期待しない文字が1つでも含まれる場合、Vimはそのファイルをスキップしてしまいます。<br />
> 正：ファイルの中にエンコーディングで期待しない文字が1つでも含まれる場合、Vimはそのファイルの変換をスキップしてしまいます。<br />

<hr />

# 2-9 最後の編集位置にカーソルを自動的に移動させる

> この設定をしていて、Gitのコミットメッセージ入力時にカーソルが移動する場合、git config --global core.editor "vim -i NONE"とする必要がある。<br />

<hr />

# 3-3 ビジュアルモードで選択されたテキストのみを対象に置換処理を実行する
* P.88 3行目

> 誤： \@!は否定後先読み<br />
> 正： \@!は否定先読み<br />

<hr />

# 3-5 検索パターンにマッチしないワードを検索する
* P.92 リスト6

> 誤： <br />
>      :%s/&\(amp;\)\@!/\&/g
>      :%s/&\(lt;\)\@!/\&/g
>      :%s/&\(gt;\)\@!/\&/g
>      :%s/&\(quot;\)\@!/\&/g
> 正： <br />
>      :%s/&\(amp;\|lt;\|gt;\|quot;\)\@!/\&/g

<hr />

# 4-12 スクリーンサイズを切り替える
* P.134 表1 1段目のコマンド例<br />

> 誤：`FullScreen 4`<br />
> 正：`ScreenMode 4`<br />

<hr />

# 5-14 テキストオブジェクトを括弧やダブルクォートで囲む
* P.179, P.180, P.181 リスト1, リスト3, リスト5の名称<br />

> 誤：コマント<br />
> 正：コマンド<br />

*実際には「コマント」になっているのでは無く、「ト」と「 ゙」(0x3099)がutf-8の合成文字になっており、おそらくDTP時に撥ねられたと思われる。*

<hr />

# 5-14 テキストオブジェクトを括弧やダブルクォートで囲む
* P.180 ビジュアルモードで括弧で囲むテキストを選択する
    * <span style="color:red">最新版のsurround.vimではビジュアルモード時の *s* が廃止されました。代わりに *S* を利用します。</span>

### ビジュアルモードで括弧で囲むテキストを選択する
ビジュアルモードで括弧で囲む範囲を選択してから、 *S* コマンドを使うとテキストを括弧で囲めます。
*S* の後に入力したテキストで選択範囲を囲めるので、見た目にもわかりやすく、最初はこちらの方法の方が使いやすいと思います。
リスト1例文に対して、ビジュアルモードで範囲を選択、 *S&#60;strong>* とコマンドを入力してみてください。
選択範囲が *&#60;strong>* タグで囲まれ、前のコマンド例の結果と同じくリスト2の結果になります。
ビジュアルモードで選択してから実行する *S* コマンドのフォーマットを以下に、使用例を表2に示します。

    S{surround}
    
    # {surround}には、囲みとして使いたいオブジェクト、括弧、ダブルクォーテーション、タグなどを指定します

### 表2 Sコマンドの使用例
<table border=1><thead><tr>
<th>コマンド</th><th>説明</th>
</tr></thead>
<tbody>
<tr><td>S"</td><td>選択箇所をダブルクォーテーションで囲む</td></tr>
<tr><td>S{</td><td>選択箇所を括弧で囲む</td></tr>
<tr><td>S&lt;strong&gt;</td><td>選択箇所を&lt;strong&gt;タグで囲む</td></tr>
</tbody></table>

<hr />

# 5-15 テキストオブジェクト機能を拡張する
リスト2が抜けていて、リスト1→リスト3→リスト4という並びになっています。
内容には問題はありません。

<hr />

# 6-5 APIドキュメントを参照する
* P.194
  * alc ソースは最新版では削除されました。詳しくは [こちら](http://d.hatena.ne.jp/thinca/20120512/1336758552)。
  * 以下のソースが追加されました。
    * rfc
    * webdict [紹介](http://d.hatena.ne.jp/thinca/20120512/1336761779)

* P.195 コマンドライン補完

> 誤：:Refコマンドはコマンドラインを提供しています。<br />
> 正：:Refコマンドはコマンドライン補完を提供しています。<br />

<hr />

# 6-8 errormarker.vimでエラー行を目立たせる
* P.202

> 誤：リスト1
>     if has('win32') || has('win64')
>       let g:errormarker_erroricon = expand('~/.vim/signs/err.bmp')
>       let g:errormarker_warningicon = expand('~/.vim/signs/warn.bmp')
>     else
>       let g:errormarker_erroricon = expand('~/.vim/signs/err.bmp')
>       let g:errormarker_warningicon = expand('~/.vim/signs/err.png')
>     endif
>
> 正：リスト1
>     if has('win32') || has('win64')
>       let g:errormarker_erroricon = expand('~/.vim/signs/err.bmp')
>       let g:errormarker_warningicon = expand('~/.vim/signs/warn.bmp')
>     else
>       let g:errormarker_erroricon = expand('~/.vim/signs/err.png')
>       let g:errormarker_warningicon = expand('~/.vim/signs/err.png')
>     endif
>
> else側がbmpになっており、本文の解説と矛盾している。
>

* P.203

> 誤：それが面倒な場合、fpluginや<br />
> 正：それが面倒な場合、ftpluginや<br />

> 誤：使用例：例えば、Rubyのスクリプトで:makeによりエラーチェックを走らせた場合<br />
> 正：使用例：例えば、Perlのスクリプトで:makeによりエラーチェックを走らせた場合<br />

> スクリーンショットが明らかにPerlスクリプトです。Rubyにもprint文はあるので完全な間違いではありませんが……。<br />

<hr />

# 6-11 vimshellを用いてVimからシェルのコマンドを実行する
※：正誤表というより、補足情報です。

* P.211

「vimshellはvimprocというプラグインに依存しているため、それだけでは動作しません。下記の
URLで配布されているvimprocも同時にインストールします。」
とありますが、vimshell Ver.9.1からは補完機能がneocomplcacheに依存するようになりました。
neocomplcacheがインストールされていない環境下では警告が表示されます。
自動補完を使用する・しないに関わらず、neocomplcacheをインストールすることを推奨します。

* P.211

vimprocのコンパイル方法として、このページに記述されているMakefile名はvimproc ver.6.1以下でのみ有効です。
vimproc ver.7.0からはMakefileの名前が異なっています。
正確なMakefile名については、":help vimproc-install"に書いてあるものを参照してください。

* P.213

「vimshellは実装上の制限により、iexeコマンドを用いないと起動したコマンドに入力を送ることができません」
とありますが、この制限は最新版であるvimshell Ver.9で撤廃されました。
ただしvimshell Ver.9を使用するためには、vimproc Ver.6.0以上をインストールし、DLLをコンパイルしておく必要があります。
iexeコマンドはもちろん残っていますが、iexeコマンドを用いる必要はありません。
ただし、Windows環境ではコマンドに対して適切なオプションを与える必要があります。
iexeではそれを自動的に処理しているため、Windows環境はiexeを用いる方法がオススメです。

<hr />

# 6-12 vimshellを用いてインタプリタと通信する
※：正誤表というより、補足情報です。

* P.215

「vimshellは普通のシェルとは異なり、コマンドにキーボードからの入力を与えることができません」
とありますが、この制限は最新版であるvimshell Ver.9で撤廃されました。
ただしvimshell Ver.9を使用するためには、vimproc Ver.6以上をインストールし、DLLをコンパイルしておく必要があります。

* P.217

「現在のvimshellは残念ながらvimshellバッファでコマンドに入力を送信することができません。
それはかなり不便なように思えます。確かに不便なので、筆者も改善をしたいと考えています。」
とありますが、この機能は最新版であるvimshell Ver.9で実装されました:-)

<hr />

# 8-16 Vimエディタで標準入力をフィルター し、標準出力に結果を出力する

* P.303
※：間違いではありませんが、こちらのURLの方が望ましいだろう、ということで。

> 誤：https://github.com/jakobi/script-archive/tree/master/vim/ <br />
> 正：<https://raw.github.com/jakobi/script-archive/master/vim/pipe.vim> <br />

<hr />

# 9-5 ソースコード中のtrue・falseをトグルで置き換える
* P.314 リスト2 CTRL-Cにtoggle.vimプラグインの機能をマッピングする

> 誤：vmap &lt;C-C&gt; &gt;Plug&gt;ToggleV<br />
> 正：vmap &lt;C-C&gt; &lt;Plug&gt;ToggleV<br />

<hr />

# 9-8 zencodingをより活用する
* P.321

> 誤：リスト1
> 
>     let g:user_zen_settings = {
>     \  'lang' : 'ja',
>     \  'indentation' : '\t',
>     \  'html' : {
>     \    'indentation' : ' ',
>     \  },
>     \  'css' : {
>     \    'filters' : 'fc',
>     \  },
>     \ },
>     \}
> 
> 正：リスト1
> 
>     let g:user_zen_settings = {
>     \  'lang' : 'ja',
>     \  'indentation' : '\t',
>     \  'html' : {
>     \    'indentation' : ' ',
>     \  },
>     \  'css' : {
>     \    'filters' : 'fc',
>     \  },
>     \}
> 
> 中括弧が一つ多い

* P.323

> 誤：リスト3
> 
>     let g:user_zen_settings = {
>     \  'javascript' : {
>     \    'snippets' : {
>     \      'jq' : "$(function() {\n\t${cursor}${child}\n});",
>     \      'jq:each' : "$.each(${cursor}, function(index, item)\n\t${child}\n});",
>     \      'fn' : "(function() {\n\t${cursor}\n})();",
>     \      'tm' : "setTimeout(function() {\n\t${cursor}\n}, 100);",
>     \    },
>     \  },
>     \}
>     
> 正：リスト3
> 
>     let g:user_zen_settings = {
>     \  'javascript' : {
>     \    'snippets' : {
>     \      'jq' : "$(function() {\n\t${cursor}${child}\n});",
>     \      'jq:each' : "$.each(${cursor}, function(index, item) {\n\t${child}\n});",
>     \      'fn' : "(function() {\n\t${cursor}\n})();",
>     \      'tm' : "setTimeout(function() {\n\t${cursor}\n}, 100);",
>     \    },
>     \  },
>     \}
>     
> jq:eachに関数の中括弧開始が無い
>

<hr />

# 9-10 neocomplcacheを使用して自動補完する

* P.329

> クイックマッチ
> g:neocomplcache_enable_quick_matchオプションをオンにした場合、-記号を入力すると、
> 補完候補の横に英数字が出てくるので、それを入力することによって、簡単に候補を選択できます。
>
> neocomplcacheのクイックマッチ機能はneocomplcache Ver.6.2で廃止されました。
>
> さらに追記：ただし、neocomplcache最新版とunite.vimの最新版をインストールしている場合、
> &lt;Plug&gt;(neocomplcache_start_unite_quick_match)のマッピングで同等のことができます。

<hr />

# 9-11 neocomplcacheを使用して自動補完する

* P.333

> g:neocomplcache_plugin_disable
>
> 追記：neocomplcache Ver.7以降では、この変数はg:neocomplcache_source_disableという名前になっています。

* P.334

> g:neocomplcache_plugin_completion_length
>
> 追記：neocomplcache Ver.7以降では、この変数はg:neocomplcache_source_completion_lengthという名前になっています。

<hr />

# 9-13 neocomplcacheを使用してスニペット補完を行う
* P.340

> Vimのスニペット展開のためのプラグインとしては、snipMateが有名ですが、neocomplcache
もスニペット機能を内蔵しています。
>
> 追記：Ver.7以降ではスニペット補完するsourceが標準添付でなくなりました。

> neocomplcacheのスニペット機能は標準添付のsnippets_complete sourceにより実現さ
れています。neocomplcacheをインストールさえすれば、特に他のプラグインをインストールする必要が
ありません。
>
> 追記：neocomplcache Ver.7以降では、スニペット機能が標準添付ではありません。
> 別配布(<https://github.com/Shougo/neocomplcache-snippets-complete>)となっているため、各自でインストールする必要があります。  
> 再追記：最新版では、neocomplcache-snippets-complete は neosnippet となりました。
> neocomplcache-snippets-complete はまだリポジトリが残っていますが、使用すると警告メッセージが出ます。
> neosnippet (<https://github.com/Shougo/neosnippet>)に移行することを推奨します。
> ちなみに、neosnippet では neocomplcache に対する依存を排除しています。

# 10-5 unite.vimを本格的にカスタマイズする

* P.360 表2 :Unite系コマンドのオプション一覧

> 誤：-bufer-name<br />
> 正：-buffer-name<br />

# 10-6 unite.vimのbuffer-name-optionを知る

* P.362

> 本文中では「buffer-name-option」となっていますが、これは古い仕様です。
>unite.vim Ver.3.1以降では「buffer-name-option」は「profile」と改名されました。
>buffer-nameではなく、profile-nameを指定することで設定を行います。
>もちろん、profile-nameを設定していないときはprofile-nameにbuffer-nameが使用されるため、互換性があります。
>詳しい仕様については、ヘルプを参照してください。

>unite.vim Ver.3.1以降においても、互換性のために
>ここで紹介されている「unite#set\_buffer\_option()」も残されてはいますが非推奨です。
>代わりに「unite#set\_profile()」を使用してください。引数の仕様は同じです。

* P.363 :UniteResumeとバッファ名

>unite.vim Ver.3.0以降では、「files@1」のように、buffer-nameにpost fixが付くようになりました。
>同じbuffer-nameで複数のuniteバッファを開けるようにするためです。
>':UniteResume'の引数には、このpost fix付きの名前を指定しなくてはいけません。注意してください。
>':Unite resume'からresumeするときは問題ありません。

* P.364 特定のバッファ名でのみ有効なキーマッピング

>例ではbuffer-nameで分岐させていますが、前述の通りbuffer-nameには自動的にpost fixが付いてしまうので、
>unite.vim Ver.3以降ではうまく動作しません。
>新しいunite.vimでは
>if l:unite.buffer_name ==# 'grep@1'
>とするか、
>if l:unite.profile_name ==# 'grep'
>とします。ただし、profile_nameが実装されているのはunite.vim Ver.3.1以降です。
