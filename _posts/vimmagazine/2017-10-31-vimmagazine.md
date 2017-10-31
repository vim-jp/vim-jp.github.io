---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 10 月号

---

## 話題

*   [Meguro.vim #6](https://megurovim.connpass.com/event/67608/) が 10/21 (土) に開催されました
*   [VimConf 2017](http://vim-jp.org/blog/2017/08/04/vimconf2017-venue-and-date-ja.html) が 11/04 (土) に開催予定です

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.0.1182: 'mzschemedll' オプションと 'mzschemegcdll' オプションが追加されました (関連パッチ: 8.0.1213)
*   8.0.1192: Windows で端末機能がデフォルトで有効になりました
*   8.0.1201: `yL` 等が 'scrolloff` の影響を受けないようになりました
*   8.0.1206: コマンドラインに入る時・抜ける時に実行される `CmdlineEnter` イベントと `CmdlineLeave` イベントが追加されました
*   8.0.1211: GTK と Windows でタブページをドラッグドロップで並べ替えられるようになりました
*   8.0.1232: Windows 版のインストーラで、Windows ライクなキーマッピングを行うかどうかの選択肢が出るようになりました
*   8.0.1234: Win32 GUI 版で合成文字や異体字が正しく表示されるようになりました
*   8.0.1236: `has('osxdarwin')` で `--enable-darwin` でコンパイルされたのかチェックできるようになりました
*   8.0.1238: 'incsearch' と 'hlsearch' がオンの場合、インクリメンタルサーチ中にも全ての候補をハイライトするようになりました
*   8.0.1239: `searchpair()` 関数の skip パラメーターに lambda を渡せるようになりました
*   8.0.1240: Windows において `term_start()` に環境変数が渡されるようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
