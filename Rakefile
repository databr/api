require "rubygems"
require "bundler"
Bundler.require

StandaloneMigrations::Tasks.load_tasks

namespace :data do
  desc "get deputados"
  task :deputados do

    Dir[File.join("data/*.rb")].each {|f| require File.absolute_path(f) }
    require './lib/activerecord'

    Deputado.save_from_pesquisa_parser
  end
end
