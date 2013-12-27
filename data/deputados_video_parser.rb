# encoding: utf-8
require File.expand_path('../pesquisa_deputados_parser', __FILE__)

class DeputadosVideoParser < PesquisaDeputadosParser
  URLS = {
    video:         'http://www2.camara.leg.br/atividade-legislativa/webcamara/resultadoDep',
    video_page:    'http://www2.camara.leg.br/atividade-legislativa/webcamara'
  }

  def videos
    Harvestman.crawl "#{URLS[:video]}?dep=*", deputados.map { |d| URI.escape(d[:nome_parlamentar]) }, :plain do
      css('#content .listaTransmissoes li') do
        date_raw, hour_raw = css('.timestamp').strip.split(' - ')
        date = date_raw.split('/')
        hour = hour_raw.split(':')

        datetime = Time.new date[2], date[1], date[0], hour[0], hour[1], 0, '-03:00'
        title = css('h4')
        local = css('.descricao .local')
        description = css('.descricao').strip.gsub(local, '')

        begin
          video_path = @document.css('.midia a').attribute('href').value
        rescue
          next
        end
        video_page_url = "#{DeputadosVideoParser::URLS[:video_page]}/#{URI.escape(video_path)}"

        _video_page_url = URI.parse(video_page_url)
        query = Hash[_video_page_url.query.split('&').map { |x| x.split('=') }]

        event = Event.where(session_id: query['codSessao']).first
        unless event
          event = Event.create session_id: query['codSessao'], starts_at: datetime, title: title, local: local, description: description
        end
        puts "Event: #{title} at #{local} #{datetime}"

        nome_parlamentar = URI.unescape query['dep']
        deputado = Deputado.find_by_nome_parlamentar nome_parlamentar

        agent = Mechanize.new
        agent.pluggable_parser.default = Mechanize::Page
        video_page_parser = agent.get video_page_url
        video_url = video_page_parser.search('#iconeDonwload').attribute('href').value

        event.videos.first_or_create video_url: video_url, deputado_id: deputado.id
        puts "Video: #{video_url}"
        puts '=='
      end
    end
  end
end
