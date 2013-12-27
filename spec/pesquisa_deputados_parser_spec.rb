# encoding: utf-8

require 'spec_helper'

describe PesquisaDeputadosParser do
  before do
    stub_request(:get, 'http://www2.camara.leg.br/deputados/pesquisa')
      .to_return(status: 200,
                 body: File.read('spec/fixtures/pesquisa.html'),
                 headers: { content_type: 'text/html' })
  end

  describe '#deputados' do
    it 'returns 5 register' do
      expect(subject.deputados.count).to eq(5)
    end

    it 'returns names and id of deputados' do
      expect(subject.deputados).to eq([
        { id: '141463', nome_parlamentar: 'ABELARDO CAMARINHA' },
        { id: '74354', nome_parlamentar: 'ZENALDO COUTINHO' },
        { id: '73933', nome_parlamentar: 'ZEQUINHA MARINHO' },
        { id: '74145', nome_parlamentar: 'ZEZÉU RIBEIRO' },
        { id: '160625', nome_parlamentar: 'ZOINHO' }
      ])
    end
  end

  describe 'complete_info_url' do
    it 'returns the base complete info url' do
      url = 'http://www.camara.gov.br/internet/Deputado/dep_Detalhe.asp'
      expect(subject.complete_info_url).to eq(url)
    end

    it 'returns the complete info url to deputado' do
      deputado_id = '232'
      url = 'http://www.camara.gov.br/internet/Deputado'
      url << "/dep_Detalhe.asp?id=#{deputado_id}"
      expect(subject.complete_info_url(deputado_id)).to eq(url)
    end
  end

  describe 'bio_url' do
    it 'returns the base complete info url' do
      url = 'http://www2.camara.leg.br/deputados/pesquisa'
      url << '/layouts_deputados_biografia'
      expect(subject.bio_url).to eq(url)
    end

    it 'returns the complete info url to deputado' do
      deputado_id = '232'
      url = 'http://www2.camara.leg.br/deputados/pesquisa'
      url << "/layouts_deputados_biografia?pk=#{deputado_id}"
      expect(subject.bio_url(deputado_id)).to eq(url)
    end
  end

  # describe 'video_url' do
  #   it 'returns the video url' do
  #     expect(subject.video_url).to eq('http://www2.camara.leg.br
  #     /atividade-legislativa/webcamara/resultadoDep')
  #   end

  #   it 'returns the video url of deputado' do
  #     nome_parlamentar = 'TIRIRICA'
  #     expect(subject.video_url(nome_parlamentar)).to eq("http://
  #       www2.camara.leg.br/atividade-legislativa/webcamara
  #       /resultadoDep?dep=#{nome_parlamentar}")
  #   end
  # end
end
