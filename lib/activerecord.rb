require "rubygems"
require "bundler/setup"
require 'active_record'

Dir[File.join("lib/models/*.rb")].each {|f| require File.absolute_path(f) }

env = ENV['RACK_ENV'] ? ENV['RACK_ENV'] : 'development'

database_config_file = File.read('config/database.yml')

if env == 'production'
  e = Erubis::Eruby.new(database_config_file)
  database_config_file = e.result(binding())
end

database_config = YAML.load(database_config_file)[env]

ActiveRecord::Base.establish_connection(database_config)

