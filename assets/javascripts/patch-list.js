$(function() {
  'use strict';

  $.getJSON("//vim-jp.herokuapp.com/patches/json?callback=?", function(items) {
    $.each(items, function() {
      var desc = $('<div/>').html(this.description).text().replace(/^[^\n]*\n\n/, '').replace(/\n(?=Solution:)/, '<br/>')
      $('<li/>').append(
        $('<a/>').attr({'href': this.link, 'target': '_blank'}).text("Patch " + this.title)
      ).append($('<br/>')).append($('<span/>').html(desc)).appendTo('#patches-list');
    });
  });
});
