# encoding: utf-8
require 'digest/md5'

class VideoEntity < BaseEntity

  protected

    def type
      @type ||= :video
    end

    def attributes_for(model)
      deputado = model.deputado
      attributes[:id] = "video:#{model.id}"
      attributes[:type] = type
      attributes[:published_at] = model.event.starts_at
      attributes[:verb] = 'Postei'
      attributes[:object] = "http://localhost:4003/video/#{Digest::MD5.hexdigest(model.video_url.gsub("&d=1", ""))}.mp4"
      attributes[:video_url] = "http://localhost:4003/video/#{Digest::MD5.hexdigest(model.video_url.gsub("&d=1", ""))}.mp4"
      attributes[:video_url_raw] = model.video_url
      attributes[:subject] = { name: deputado.nome_parlamentar, image: deputado.image_url, url: "/#{deputado.uri}" }
      attributes[:location] = { title: model.event.local, url: '#' }
      attributes[:in] = { name: model.event.title, url: '#' }
      attributes
    end

    def get_total(cotas)
      nil
    end

    def group_by(type, data)
      case type
      when :month
        data.group_by { |v| v.event.starts_at.beginning_of_month }
      when :year
        data.group_by { |v| v.event.starts_at.beginning_of_year }
      end
    end
end
