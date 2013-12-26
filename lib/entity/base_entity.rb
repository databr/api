class BaseEntity
  def initialize(deputado)
    @deputado = deputado
  end

  def models
    @models ||= the_model.to_feed(@deputado)
  end

  def results
    by_year = group_by :year, models
    @results = []
    by_year.keys.sort.reverse.each do |_key|
      data = {year: _key.year}
      year_data = by_year[_key]
      year_total = get_total(year_data) if respond_to?(:get_total)
      data[:total] = year_total unless year_total.nil?
      data[:month] = []
      by_month = group_by :month, year_data
      by_month.keys.sort.reverse.each do |_month_key|
        data_month = {month: _month_key.month}
        month_data = by_month[_month_key]
        month_total = get_total(month_data) if respond_to?(:get_total)
        data_month[:total] = month_total unless month_total.nil?
        data_month[:data] = month_data.map{|model| attributes_for(model) }
        data[:month] << data_month
      end
      @results << data
    end
    @results
  end
end
