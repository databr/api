require "rubygems"
require "bundler/setup"
require 'active_record'

Dir[File.join("lib/models/*.rb")].each {|f| require File.absolute_path(f) }


ActiveRecord::Base.establish_connection(YAML.load(File.read(File.join('db','config.yml')))[ENV['ENV'] ? ENV['ENV'] : 'development'])
