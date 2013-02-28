// vim:set ts=8 sts=2 sw=2 tw=0 et:
function escape(str) {
  return str.replace('&', '&amp;')
            .replace('<', '&lt;')
            .replace('>', '&gt;')
            .replace('"', '&quot;')
            .replace("'", '&#x27;');
}

google.load("visualization", "1");
$(function() {
  var query = new google.visualization.Query('https://spreadsheets.google.com/tq?key=0AvP2GcrHKTfudEdwWm85ajZtd0ltZjRUcVFPWHpzWHc&gid=0&pub=1');
  query.send(function(response) {
    if(response.isError()) {
      return;
    }
    var data = response.getDataTable();
    for (var row = 0; row < data.getNumberOfRows(); row++) {
      var name    = data.getFormattedValue(row, 1);
      var twitter = data.getFormattedValue(row, 2);
      var github  = data.getFormattedValue(row, 3);
      var website = data.getFormattedValue(row, 4);
      var description = data.getFormattedValue(row, 5);

      var items = [];

      items.push([ 'li', { 'class': 'name' }, name ]);
      if (twitter) {
        var twitter_url = 'https://twitter.com/' + twitter;
        var twitter_img = 'http://api.twitter.com/1/users/profile_image/'
            + twitter + '.png&size=mini';
        items.unshift([
          'li', { 'class': 'faceicon' },
          [ 'img', { 'src': twitter_img } ]
        ]);
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': twitter_url }, '@' + twitter ]
        ]);
      }

      items.push([ 'li', { 'class': 'clear' } ]);

      if (github) {
        var github_url = 'https://github.com/' + github;
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': github_url }, github_url ]
        ]);
      }

      if (website) {
        items.push([
          'li', { 'class': 'link' },
          [ 'a', { 'href': website }, website ]
        ]);
      }

      if (description) {
        /* FIXME: adapt to JsonML.
        var link = escape(description).replace(
            /((http|https):\/\/[\x21-\x7e]+)/gi, '<a href="$1">$1</a>');
         */
        items.push([ 'li', { 'class': 'desc' }, description ]);
      }

      $.jqml([
        'div', { 'class': 'person', },
        [ 'ul' ].concat(items)
      ]).appendTo('#vimmers-container');
    }
    $('#vimmers-container').masonry({
      isAnimated: true,
      isResizable: true
    });
  });
});
