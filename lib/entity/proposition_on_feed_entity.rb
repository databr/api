# encoding: utf-8
class PropositionOnFeedEntity < BaseOnFeedEntity
  protected
    def model_class
      Proposition
    end

    def type
      @type ||= :project
    end

    def verb
      'Aprensentei'
    end

    def attributes_for(model)
      attributes = _attributes
      attributes[:id] = "proposition-#{model.proposition_id}"
      attributes[:published_at] = model.presentations_at
      attributes[:object] = { name: model.name, url: model.url }
      attributes[:content] = model.body
      attributes[:location] = { title: 'Camara', url: '#' }
      attributes
    end

    def group_by(type, data)
      case type
      when :month
        data.group_by { |c| c.presentations_at.beginning_of_month }
      when :year
        data.group_by { |c| c.presentations_at.beginning_of_year }
      end
    end
end
