#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

require 'yaml'

root = File::expand_path("#{File::dirname(__FILE__)}/..")

commit_hashes = ["7bf9665fd74fe399874bd783c508f37da161026", "0d83ffed5358bf8b3b07d81347f753813ad9a1c5"]
diff = ""
commit_hashes.each do |hash|
    diff += `git diff #{hash}^..#{hash}`
end

entries = []
while diff.size != 0
# _posts/2011-09-22-how_to_write_patches.md b/_posts/2011-09-22-how_to_write_patches.md
  if /^diff --git a\/(.+?\.md)/ =~ diff
    path = $+
    rest = $'
    if /^diff --git/ =~ rest
      body = $`
    else
      body = rest
    end
    diff = rest

    if body =~ /-category: (.+)$/
      old_category = $+.chomp
    else
      old_category = nil
    end
    if body =~ /\+category: (.+)$/
      new_category = $+.chomp
    else
      new_category = nil
    end

    header = nil
    open("#{root}/#{path}", "r") do |f|
      entity = f.read
      if /^---(.+?)^---/m =~ entity
        header = YAML.load($+)
      else
        puts "header not found: #{path}"
      end
    end

    entries << {:path => path, :old_category => old_category, :new_category => new_category, :header => header}
  else
    diff = ""
  end
end

entries.each do |entry|
  if entry[:old_category].nil?
    entry[:header].delete("category")
  else
    entry[:header]["category"] = entry[:old_category]
  end
  entry[:header]["layout"] = "redirect"
  entry[:header]["title"] = "redirect to " + entry[:header]["title"]
  new_path = entry[:path].gsub(/(\d{4}-\d{2}-\d{2}-)(.+)$/, '\1redirect_to_\2')
  entry[:header]["redirect-to"] = entry[:path].gsub(/.+(\d{4})-(\d{2})-(\d{2})-(.+)\.md$/, "/#{entry[:new_category]}/\\1/\\2/\\3/\\4.html")
  entry[:header]["permalink"] = entry[:path].gsub(/.+(\d{4})-(\d{2})-(\d{2})-(.+)\.md$/, "/#{entry[:old_category].nil? ? "" : "#{entry[:old_category]}/"}\\1/\\2/\\3/\\4.html")
  open("#{root}/#{new_path}", "w") do |f|
    f.write("---\n")
    entry[:header].each do |key, value|
      f.write("#{key}: #{value}\n")
    end
    f.write("---\n")
  end
end

