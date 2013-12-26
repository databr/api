# enconding: utf-8
class PropositionEntity < BaseEntity
  protected
    def the_model
      Proposition
    end

    def type
      @type ||= :project
    end

    def attributes_for(model)
      attributes = {}
      attributes[:id] = "project:#{model.id}"
      attributes[:type] = type
      attributes[:published_at] = model.presentations_at
      attributes[:verb] = "Aprensentei"
      attributes[:object] = model.name
      attributes[:content] = model.body
      attributes[:subject] = {name: @deputado.nome_parlamentar, image: @deputado.image_url}
      attributes[:location] = {title: "Camara", url: "#"}
      attributes
    end

    def group_by(type, data)
      case type
      when :month
        data.group_by{|c| c.presentations_at.beginning_of_month }
      when :year
        data.group_by{|c| c.presentations_at.beginning_of_year }
      end
    end
end

