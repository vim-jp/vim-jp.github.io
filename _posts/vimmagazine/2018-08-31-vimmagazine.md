---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2018 年 08 月号

---

## 話題

*   [Asakusa.vim #3](https://asakusavim.connpass.com/event/95401/) が 2018-08-03 (金) に二子玉川で開催されました
*   [meguro.vim #11](https://megurovim.connpass.com/event/94477/) が 2018-08-11 (土) に開催されました
*   [VimConf 2018](https://vimconf.org/2018/) 関連
    *   VimConf 2018 が 2018-11-24 に秋葉原のアキバPLAZAで開催されます

## 今月の新機能

今月の新機能及びユーザーに影響のある変更は以下のとおりです。

*   8.1.0232: if_ruby 実行中にエラーが起きた場合、そのバックトレースが Vim のコマンドラインに表示されるようになりました
*   8.1.0251: `set backupdir=~/vimbackupdir//` のように指定すると、`~/vimbackupdir/` 以下にファイルのフルパスが連結されたファイル名で書き出されるようになりました
    *   例: 上記の設定を行った際、ファイル `/home/vim/vimrc` のバックアップファイルは `~/vimbackupdir/%home%vim%vimrc{backupext}` となります (`{backupext}` = `'backupext'` オプションの値)
    *   スワップファイルの場合は以前から可能です `set directory=~/swapdir//`
*   8.1.0253: タイトルスタックに対応した端末では、端末のタイトルが正しく復帰されるようになりました
    *   参考記事: [「Vimを使ってくれてありがとう」にさようなら](https://qiita.com/ttdoda/items/903e85f07d58018c851d)
*   8.1.0263: チャンネルのログにどのパート ("sock", "out", "err", "in") かの情報が出力されるようになりました

## Vimに関する脆弱性

特筆すべき脆弱性の報告はありませんでした。
