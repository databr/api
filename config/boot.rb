require "rubygems"
env = ENV['RACK_ENV'] ? ENV['RACK_ENV'] : 'development'
require 'bundler'
Bundler.require(:default, env)

require 'active_record'

require 'logger'

class ApiLogger
  def initialize
    @loggers = []
    @loggers << Logger.new("app.log")
    @loggers << Logger.new(STDOUT) if stdout?
    set_formatter
  end

  def fatal(message)
    log(Logger::FATAL, message)
  end

  def error(message)
    log(Logger::ERROR, message)
  end

  def info(message)
    log(Logger::INFO, message)
  end

  def debug(message)
    log(Logger::DEBUG, message)
  end

  def warn(message)
    log(Logger::WARN, message)
  end

  def stdout?
    ENV['RACK_ENV'] == 'development'
  end

  def debug?
    ENV['DEBUG'] == 'true' or ENV['RACK_ENV'] == 'development'
  end

  def level=(level)
    level = Logger::DEBUG if debug?
    @loggers.each do |logger|
      logger.level = level
    end
  end

  private

  def log(type, message)
    @loggers.each do |logger|
      logger.add(type){ message }
    end
  end

  def set_formatter
    @loggers.each do |logger|
      logger.formatter = Formatter.new
    end
  end
end

class Formatter
  SEVERITY_TO_COLOR_MAP   = {'DEBUG'=>'0;37', 'INFO'=>'32', 'WARN'=>'33', 'ERROR'=>'31', 'FATAL'=>'31', 'UNKNOWN'=>'37'}

  def call(severity, time, progname, msg)
    formatted_severity = sprintf("%-5s","#{severity}")

    formatted_time = time.strftime("%Y-%m-%d %H:%M:%S.") << time.usec.to_s[0..2].rjust(3)
    color = SEVERITY_TO_COLOR_MAP[severity]

    "#{formatted_time}\033[0m [\033[#{color}m#{formatted_severity}\033[0m] #{msg.to_s.strip} (pid:#{$$})\n"
  end
end

LOGGER = ApiLogger.new


if ENV['USE_PARSER'] == 'true'
  require './data/camara_parser'
  Dir[File.join("data/*.rb")].each {|f| require File.absolute_path(f) }
  Dir[File.join("lib/services/*.rb")].each {|f| require File.absolute_path(f) }
end

require './lib/entity/base_entity.rb'
Dir[File.join("lib/models/*.rb")].each {|f| require File.absolute_path(f) }
Dir[File.join("lib/entity/*.rb")].each {|f| require File.absolute_path(f) }

require 'aggregator'

database_config_file = File.read('config/database.yml')

if env == 'production'
  require 'erubis'
  e = Erubis::Eruby.new(database_config_file)
  database_config_file = e.result(binding())
end

database_config = YAML.load(database_config_file)[env]
ActiveRecord::Base.include_root_in_json = false
ActiveRecord::Base.logger = LOGGER
ActiveRecord::Base.establish_connection(database_config)

