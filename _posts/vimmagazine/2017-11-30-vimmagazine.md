---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 11 月号

---

## 話題

*   [VimConf 2017](http://vim-jp.org/blog/2017/08/04/vimconf2017-venue-and-date-ja.html) が 11/04 (土) に開催されました
*   [Fablic.vim #4](https://fablicvim.connpass.com/event/71673/) が 12/04 (月) に東京で開催予定です
*   [Meguro.vim #7](https://megurovim.connpass.com/event/72311/) が 12/16 (土) に東京で開催予定です
*   [Yokohama.vim #10](https://yokohamavim.connpass.com/event/73077/) が 12/10 (日) に神奈川で開催予定です

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1278: GUI ウィンドウのスクロールバー表示時にウィンドウのサイズを保持する `'guioptions'` オプションの `k` フラグが追加されました
    *   これによりスクロールバーが表示される時に GUI のウィンドウの最大化状態を解除されてしまう問題が修正されました
*   8.0.1279: メニューの初期化処理が高速化され、起動時間が速くなりました
*   8.0.1280: if_python における `None` が Vim script の `v:none` に変換されるようになりました
*   8.0.1281: Vim script で実行されるファイルタイプの検知処理が高速化され、起動時間が速くなりました (関連パッチ: 8.0.1282, 8.0.1284)
*   8.0.1287: ウィンドウローカルのディレクトリを保持する `'viewoptions'` オプションの `curdir` が追加されました
*   8.0.1295: UNIX 系の端末版 Vim でも `v:servername` が常に生成される `--enable-autoservername` オプションが configure スクリプトに追加されました (関連パッチ: 8.0.1297)
*   8.0.1299: Bracketed paste mode が端末ウィンドウで動くようになりました
*   8.0.1305: `writefile()` の第3引数のフラグに `fsync(2)` を実行するか否かのフラグ `s`, `S` が追加されました (関連パッチ: 8.0.1335)
    *   `s` が指定された場合は実行し、`S` が指定された場合は実行せず、どちらも指定されなかった場合は `'fsync'` オプションに従います
    *   今までは常に `fsync(2)` を実行していませんでした
*   8.0.1309: `'balloonevalterm'` オプションが追加され、Unix 系の GUI 環境から起動された端末版 Vim でも ballooneval 機能が使用できるようになりました (関連パッチ: 8.0.1323, 8.0.1324)
*   8.0.1318: `balloon_split()` 関数が追加され、`balloon_show()` 関数が端末版 Vim では List を受け付けるようになりました (関連パッチ: 8.0.1346)
*   8.0.1330: Vim の端末機能から実行されたジョブ内で Windows の場合も `VIM_SERVERNAME` 環境変数が定義されるようになりました
*   8.0.1332: `qfSeparator` ハイライトグループが追加されました
*   8.0.1336: xim 機能付きでコンパイル済みでなくとも `'imactivatefunc'` オプションが使えるようになりました (関連パッチ: 8.0.1338)
*   8.0.1343: Windows で色付き絵文字を表示できるようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
