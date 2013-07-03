#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

require 'json'
require 'open-uri'
require 'net/http'

root = File::expand_path("#{File::dirname(__FILE__)}/..")

vimmers = JSON.load open "#{root}/vimmers/vimmers.json"

env = ENV['http_proxy']
if env then
  uri = URI(env)
  proxy_host, proxy_port = uri.host, uri.port
else
  proxy_host, proxy_port = nil, nil
end
vimmers.each do |x|
  if !x['twitter'].empty?
    x['twitter_icon'] = Net::HTTP::Proxy(proxy_host, proxy_port).get_response(URI.parse("http://api.dan.co.jp/twicon/#{x['twitter']}"))["location"]
  end
  warn x['twitter_icon']
end

puts vimmers.to_json
