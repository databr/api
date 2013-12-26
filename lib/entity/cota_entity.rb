class CotaEntity < BaseEntity
  protected
    def the_model
      Cota
    end

    def type
      @type ||= :money
    end

    def attributes_for(model)
      attributes = {}
      attributes[:id] = "cota:#{model.id}"
      attributes[:type] = type
      attributes[:published_at] = model.data_emissao
      attributes[:verb] = "Gastei"
      attributes[:object] = model.descricao
      attributes[:subject] = {name: @deputado.nome_parlamentar, image: @deputado.image_url}
      attributes[:location] = {title: model.beneficiario, url: "#"}
      attributes[:value] = model.valor_documento.to_f
      attributes
    end

    def get_total(cotas)
      cotas.map(&:valor_liquido).reduce(&:+)
    end

    def group_by(type, data)
      case type
      when :month
        data.group_by{|c| c.data_emissao.beginning_of_month }
      when :year
        data.group_by{|c| c.data_emissao.beginning_of_year }
      end
    end
end
