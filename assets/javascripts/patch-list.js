$(function() {
  'use strict';

  $.getJSON("https://api.github.com/repos/vim/vim/commits?callback=?", function(items) {
    $.each(items.data, function(index) {
      if (index == 10) {
        return false;
      }
      var token = this.commit.message.split(/\n\n/, 2);
      var desc = $('<div/>').text(token[1] || '').html().replace(/\n(?=Solution:)/, '<br/>');
      $('<li/>').append(
        $('<a/>').attr({'href': this.html_url, 'target': '_blank'}).text(token[0].replace(/^patch (\d+\.\d+\.\d+).*$/, 'Patch $1'))
      ).append($('<br/>')).append($('<span/>').html(desc)).appendTo('#patches-list');
    });
  });
});
