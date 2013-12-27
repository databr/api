# encoding: utf-8
require 'harvestman'

class DeputadoVideoService
  def self.save_video_from_deputados_parser
    parser = DeputadosVideoParser.new
    parser.videos
  end
end
