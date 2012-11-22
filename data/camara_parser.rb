class CamaraParser
  def initialize
    @agent = Mechanize.new
    @agent.pluggable_parser.default = Mechanize::Page
    @parser = @agent.get @url
  end
end
