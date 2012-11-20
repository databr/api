require "rubygems"
require "bundler/setup"
require 'active_record'

Dir[File.join("lib/models/*.rb")].each {|f| require File.absolute_path(f) }


ActiveRecord::Base.establish_connection(YAML.load(File.read(File.join('config','database.yml')))[ENV['RACK_ENV'] ? ENV['RACK_ENV'] : 'development'])
