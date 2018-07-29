---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 07 月号

---

## 話題

*   [Asakusa.vim #2](https://asakusavim.connpass.com/event/93282/) が 2018-07-18 (水) に秋葉原で開催され、Vim使いたちの貴重な知見が共有されました
*   [VimConf 2018](https://vimconf.org/2018/) 関連
    *   VimConf 2018 が 2018-11-24 に秋葉原のアキバPLAZAで開催されます
    *   引き続き、[VimConf 2018の企業スポンサー公募](https://vim-jp.org/blog/2018/06/07/VimConf2018-sponsor-ja.html) 中です
    *   引き続き、[VimConf 2018のスピーカーの公募](https://vim-jp.org/blog/2018/06/28/VimConf2018-cfp-ja.html) 中です
*   [RailsDM 2018](https://techplay.jp/event/679666) というRuby on Railsのカンファレンスで、でVimのC実装に関する発表が行われました。
    *   [スライド](https://speakerdeck.com/ujihisa/vim-railsdm-2018-07)と[動画](https://www.youtube.com/watch?v=12CCKbGg1EA)
*   [meguro.vim #11](https://megurovim.connpass.com/event/94477/) が 2018-08-11 (土) に開催予定です。

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.1.0142: 組込み termcap の xterm, vt320 エントリにキーパッドのエスケープシーケンスが追加されました
*   8.1.0182: Unicode 11 に対応しました
*   8.1.0184: `wincol` と `winrow` の項目が `getwininfo()` の返り値に追加されました (関連: 8.1.0187)
*   8.1.0198: `'redrawtime'` で設定した時間を過ぎてシンタックスハイライトが無効になる場合はメッセージが出るようになりました
*   8.1.0214: `autochdir` が `has()` 関数と `:version` の出力結果で確認できるようになりました
*   8.1.0218: `matchadd()`, `matchaddpos()` 関数に現在のウィンドウ以外のウィンドウ番号やウィンドウIDを指定できる `window` の項目が追加されました
*   8.1.0225: インサートモードで `CTRL-O` コマンドを実行した際、`mode()` 関数が `niI`, `niR`, `niV` 等の文字列を返すようになりました

*   ランタイムの更新
    *   reStructuredText の ftplugin スクリプトで reStructuredText 標準ルールの設定を行うようになりました (`setlocal expandtab shiftwidth=3 softtabstop=3 tabstop=8`)
        *   これを抑制するには `let g:rst_style = 0` と vimrc に書く必要があります
    *   reStructuredText の ftplugin スクリプトで独自の折畳みを行うようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
