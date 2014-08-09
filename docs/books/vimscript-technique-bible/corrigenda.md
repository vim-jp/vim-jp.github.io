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

## 関連ツイート

<!-- Twitter widget: "Vim script" OR Vimスク本 OR #Vimスク本 -->
<a class="twitter-timeline" data-dnt="true" href="https://twitter.com/search?q=%22Vim+script%22+OR+Vim%E3%82%B9%E3%82%AF%E6%9C%AC+OR+%23Vim%E3%82%B9%E3%82%AF%E6%9C%AC" data-widget-id="497924456598097920">"Vim script" OR Vimスク本 OR #Vimスク本 に関するツイート</a>
