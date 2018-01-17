$(function() {
  'use strict';

  $.getJSON("https://api.github.com/repos/vim/vim/commits?callback=?", function(items) {
    $.each(items.data, function() {
      var token = this.commit.message.replace(/\n\n+/g, '\n').split(/\n/, 2);
      $('<li/>').append(
        $('<a/>').attr({'href': this.html_url, 'target': '_blank'}).text(token[0])
      ).append($('<br/>')).append($('<span/>').html(token[1] || '')).appendTo('#patches-list');
    });
  });
});
