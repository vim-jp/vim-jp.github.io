$(function() {
  'use strict';

  $.getJSON("//vim-jp.herokuapp.com/patches/json?callback=?", function(items) {
    $.each(items, function() {
      $('<li/>').append(
        $('<a/>').attr({'href': this.link, 'target': '_blank'}).text("Patch " + this.title)
      ).append($('<br/>')).append($('<span/>').text(this.description)).appendTo('#patches-list');
    });
  });
});
