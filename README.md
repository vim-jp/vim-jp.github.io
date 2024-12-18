# vim-jp.github.io

Source of <https://vim-jp.org/>

  * [EditSite](https://github.com/vim-jp/vim-jp.github.io/wiki/EditSite) サイト編集の手順 (ローカル環境での確認方法)

<a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="クリエイティブ・コモンズ・ライセンス" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br />この 作品 は <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">クリエイティブ・コモンズ 表示 4.0 国際 ライセンス</a>の下に提供されています。

## ローカルでビルド・確認する方法

### 必要なもの

* Docker
* (オプション) Python 3: ローカルHTTPサーバーをPython 3で動かす

### ビルド

以下のコマンドを実行すると `_site/` ディレクトリにHTML他が出力される。

```console
$ ./_scripts/docker_jekyll jekyll build
```

以下のようにするとインクリメンタルビルドが有効になり、
2回目以降のビルド時間が短縮される。

```console
$ ./_scripts/docker_jekyll jekyll build -I
```

### ローカルサーバー

以下のコマンドで `jekyll serve` を実行し、ビルドした上でローカルHTTPサーバーが
起動し、手元のWebブラウザで内容を確認できる。

```console
$ ./_scripts/docker_jekyll_serve
```

これは4000番のポートでコンテンツを供給する。
Webブラウザで <http://127.0.0.1:4000/> を開くことで内容を確認できる。

また以下のように同コマンドへオプションを指定すると、インクリメンタルビルドと
変更監視&自動再ビルド機能が有効になる。

```console
$ ./_scripts/docker_jekyll_serve -wI
```

Windowsなどのシステムでは変更監視が機能しない場合がある。
そのようなケースではPython 3でローカルサーバーを立て、
ビルドは手動でやってしまったほうが手っ取り早い場合があるかもしれない。
そのやり方の詳しくは後述。

ローカルサーバーだけが必要で、変更監視や自動更新が不要な場合は次のように
起動しても良いかもしれない。

```console
$ ./_scripts/docker_jekyll_serve --skip-initial-build --no-watch
```

#### Pythonを使用したローカルサーバー

以下のコマンドでPythonによるローカルHTTPサーバーを起動できる

```console
$ ./_scripts/python_http_server
```

これは8000番のポートでコンテンツを供給するので、
Webブラウザで <http://127.0.0.1:8000/> を開くことで内容を確認できる。
