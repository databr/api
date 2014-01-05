# encoding: utf-8
#
class CotaOnFeedEntity < BaseOnFeedEntity
  protected
    def model_class
      Cota
    end

    def type
      @type ||= :money
    end

    def verb
      'Gastei'
    end

    def attributes_for(model)
      attributes = _attributes
      attributes[:id] = "cota-#{model.id}"
      attributes[:published_at] = model.data_emissao
      attributes[:object] = { name: model.descricao }
      attributes[:location] = { title: model.beneficiario, url: '#' }
      attributes[:value] = model.valor_documento.to_f
      attributes
    end

    def get_total(cotas)
      cotas.map(&:valor_liquido).reduce(&:+)
    end

    def group_by(type, data)
      case type
      when :month
        data.group_by { |c| c.data_emissao.beginning_of_month }
      when :year
        data.group_by { |c| c.data_emissao.beginning_of_year }
      end
    end
end
