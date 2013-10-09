class Aggregator
  def self.build(sets)
    @aggregate = {}
    sets.flatten.each do |data|
      @aggregate[data[:year]] ||= {}
      @aggregate[data[:year]][:total] = data[:total] if data[:total]
      @aggregate[data[:year]][:month] ||= {}
      data[:month].each do |month|
        @aggregate[data[:year]][:month][month[:month]] ||= {}
        @aggregate[data[:year]][:month][month[:month]][:total] = month[:total] if month[:total]
        @aggregate[data[:year]][:month][month[:month]][:data] ||= []
        @aggregate[data[:year]][:month][month[:month]][:data] << month[:data]
        @aggregate[data[:year]][:month][month[:month]][:data].flatten!
      end
    end
    @result = []
    @aggregate.keys.sort.reverse.each do |_key|
      data = {year: _key}
      year_data = @aggregate[_key]
      data[:total] = year_data[:total]
      data[:month] = []
      year_data[:month].keys.sort.reverse.each do |_month_key|
        month_data = year_data[:month][_month_key]
        month_data[:month] = _month_key
        data[:month] << month_data
      end
      @result << data
    end
    @result
  end
end
