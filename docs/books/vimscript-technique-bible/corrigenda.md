---
layout: docs
title: Vim scriptテクニックバイブル ～Vim使いの魔法の杖 正誤表
---

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
