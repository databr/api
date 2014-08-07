# encoding: utf-8
class DeputiesJsonParser
  def initialize(file)
    @file = File.expand_path(file).to_s
    @parser = JSON.parse(File.read(@file))
  end

  def parser
    @parser
  end
end
