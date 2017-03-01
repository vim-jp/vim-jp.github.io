---
layout: vimmagazine
category: vimmagazine
title: Vim Magazine 2017 年 2 月号

---

## 話題

* 2/11 に [Yokohama.vim #9](https://yokohamavim.connpass.com/event/47519/) が開催されました。
* 2/25 に [Meguro.vim #2](https://megurovim.connpass.com/event/49507/) が開催されました。

## 今月の新機能

* 8.0.0283: `mode()` が置換モードおよびインサートモード中の2種類の補完状態 "c" と "x" を返すようになりました。
* 8.0.0302: 端末キーコード (`'t_kl'` 等)が、 `:let` で設定できるようになりました。
* 8.0.0309: `json_decode()` が空のキーを受け付けるようになりました。
* 8.0.0369: いくつかのオプションが features に依存して定義されていなかったものが、常に定義されるようになりました。 (関連patch: 8.0.0372)

## Vimに関する脆弱性
* 8.0.0322 (related: 8.0.0376) <https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-5953>
* 8.0.0377 <https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-6349>
* 8.0.0378 <https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-6350>
