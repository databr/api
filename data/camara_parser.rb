class CamaraParser
  def initialize
    @agent = Mechanize.new
    @parser = @agent.get @url
  end
end
