class CotaEntity
  def initialize(cotas)
    @cotas = cotas
  end

  def results
    by_year = @cotas.group_by{|c| c.data_emissao.beginning_of_year }
    @results = []
    by_year.keys.sort.reverse.each do |_key|
      data = {year: _key.year}
      year_cotas = by_year[_key]
      year_total = get_total(year_cotas)
      data[:total] = year_total
      data[:month] = []
      by_month = year_cotas.group_by{|c| c.data_emissao.beginning_of_month }
      by_month.keys.sort.reverse.each do |_month_key|
        data_month = {month: _month_key.month}
        month_cotas = by_month[_month_key]
        month_total = get_total(month_cotas)
        data_month[:total] = month_total
        data_month[:data] = month_cotas.map(&:attributes)
        data[:month] << data_month
      end
      @results << data
    end
    @results
  end

  private
    def get_total(cotas)
      cotas.map(&:valor_liquido).reduce(&:+)
    end
end
