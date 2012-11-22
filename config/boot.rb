require "rubygems"
env = ENV['RACK_ENV'] ? ENV['RACK_ENV'] : 'development'
require 'bundler'
Bundler.require(:default, env)

require 'active_record'

if ENV['USE_PARSER'] == 'true'
  require './data/camara_parser'
  Dir[File.join("data/*.rb")].each {|f| require File.absolute_path(f) }
end

Dir[File.join("lib/models/*.rb")].each {|f| require File.absolute_path(f) }


database_config_file = File.read('config/database.yml')

if env == 'production'
  require 'erubis'
  e = Erubis::Eruby.new(database_config_file)
  database_config_file = e.result(binding())
end

database_config = YAML.load(database_config_file)[env]

ActiveRecord::Base.establish_connection(database_config)

