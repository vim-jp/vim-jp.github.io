#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

require 'date'
require 'yaml'

root = File::expand_path("#{File::dirname(__FILE__)}/..")
conf = YAML.load_file("#{root}/_config.yml")

if !ARGV.empty? && conf['downloads'].has_key?(ARGV[0])
  rel = conf['downloads'][ARGV[0]]
  ver = rel['version'].gsub(/\s/, '')
  file = "_posts/#{Date.today.to_s}-#{ARGV[0]}-#{ver}.md"
  open("#{root}/#{file}", "wb") {|f|
    f.puts <<"EOS"
---
layout: post
title: #{rel['title']} #{rel['version']} リリース
---

#{rel['title']} #{rel['version']} がリリースされました。
-------------------------------------------------------

#{rel['info']}
EOS
  }
  puts <<"EOS"
Wrote: #{file}
You should do:
  git add #{file}'
  git commit
EOS
else
  puts "usage: #{File.basename(__FILE__)} {release-name}"
  puts "--------"
  puts conf['downloads'].keys.join("\n")
end
