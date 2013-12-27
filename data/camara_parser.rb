# encoding: utf-8
class CamaraParser
  def initialize
    @agent = Mechanize.new
    @agent.pluggable_parser.default = Mechanize::Page
    puts "\e[35m> GET #{@url}\e[0m"
    @parser = @agent.get @url
  end
end
