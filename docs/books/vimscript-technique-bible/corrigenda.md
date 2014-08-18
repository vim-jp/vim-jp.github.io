---
layout: docs
title: Vim scriptテクニックバイブル ～Vim使いの魔法の杖 正誤表
---

## 正誤表

<table>
<tr>
  <th>ページ、行数</th>
  <th>誤</th>
  <th>正</th>
  <th>解説/言い訳</th>
</tr>
{% for item in site.data.books.vimscript_technique_bible.corrigenda %}
<tr>
  <td><nobr>{{item.page}}ページ</nobr>{% if item.line %} <nobr>{{item.line}}行目{% endif %}</nobr></td>
  <td><nobr>{{item.erratum}}</nobr></td>
  <td><nobr>{{item.correction}}</nobr></td>
  <td>{{item.excuse}}</td>
</tr>
{% endfor %}
</table>

## <a name="normal-index"><a>索引について

[@yrntrlmnmntさん](https://twitter.com/yrntrlmnmnt) が [普通の索引の印刷データ](http://vim-jp.org/docs/books/vimscript-technique-bible/corrigenda.html#normal-index) を作って公開してくれました。

<blockquote class="twitter-tweet" lang="ja"><p>Ｖｉｍ　ｓｃｒｉｐｔテクニックバイブルを読んでプラグイン（紙）を作ってみた&#10;※アルファベット順の索引 <a href="http://t.co/3t4f3kIzoc">pic.twitter.com/3t4f3kIzoc</a></p>&mdash; あなたの天然記念物／アナログ師（自称） (@yrntrlmnmnt) <a href="https://twitter.com/yrntrlmnmnt/statuses/499879383675326465">2014, 8月 14</a></blockquote>

必要な方は [google docs](https://docs.google.com/file/d/0B95Afy1gzPwjbHlKVldJNHVSd0k/edit) で印刷するか、XLSをダウンロードして印刷してご利用ください。

## <a name="normal-index"><a>Vim界隈(Vimスク本を含む)における「プラグイン」や「スクリプトファイル」などの用語についての補足文

[@yrntrlmnmntさん](https://twitter.com/yrntrlmnmnt) より、「Chapter5 Vim スクリプトを公開しよう」の各種用語が分かりにくいという指摘を頂いたので、執筆者の一人である [@kaoriyaさん](https://twitter.com/kaoriya) が [補足の文章](https://gist.github.com/koron/d2bc314e4087ca1a9a9e) を公開してくれました。

用語が分かりにくいと思った場合は参照してください。

## 関連ツイート

<!-- Twitter widget: "Vim script" OR Vimスク本 OR #Vimスク本 -->
<a class="twitter-timeline" data-dnt="true" href="https://twitter.com/search?q=%22Vim+script%22+OR+Vim%E3%82%B9%E3%82%AF%E6%9C%AC+OR+%23Vim%E3%82%B9%E3%82%AF%E6%9C%AC" data-widget-id="497924456598097920">"Vim script" OR Vimスク本 OR #Vimスク本 に関するツイート</a>
