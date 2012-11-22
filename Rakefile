require "rubygems"
require "bundler"
Bundler.require

StandaloneMigrations::Tasks.load_tasks

namespace :data do
  task :enviroment do
    ENV['USE_PARSER'] = 'true'
    require './config/boot'
  end

  desc "get deputados"
  task :deputados => :enviroment do
    Deputado.save_from_pesquisa_parser
    Deputado.save_from_deputado_xml_parser
  end
end
