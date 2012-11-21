# encoding: utf-8

require 'spec_helper'

describe PesquisaDeputadosParser do
  let(:parser) { PesquisaDeputadosParser.new }

  before do
    stub_request(:get, 'http://www2.camara.leg.br/deputados/pesquisa').
      to_return(:status => 200, :body => File.read('spec/fixtures/pesquisa.html'), :headers => {:'Content-Type' => 'text/html'})
  end

  describe '#deputados' do
    it 'returns 5 register' do
      expect(parser.deputados.count).to eq(5)
    end

    it 'returns names and id of deputados' do
      expect(parser.deputados).to eq([
        {id: '141463', nome_parlamentar: 'ABELARDO CAMARINHA'},
        {id: '74354', nome_parlamentar: 'ZENALDO COUTINHO'},
        {id: '73933', nome_parlamentar: 'ZEQUINHA MARINHO'},
        {id: '74145', nome_parlamentar: 'ZEZÉU RIBEIRO'},
        {id: '160625', nome_parlamentar: 'ZOINHO'}
      ])
    end
  end

  describe 'complete_info_url' do
    it 'returns the base complete info url' do
      expect(parser.complete_info_url).to eq('http://www.camara.gov.br/internet/Deputado/dep_Detalhe.asp')
    end

    it 'returns the complete info url to deputado' do
      deputado_id = "232"
      expect(parser.complete_info_url(deputado_id)).to eq("http://www.camara.gov.br/internet/Deputado/dep_Detalhe.asp?id=#{deputado_id}")
    end
  end

  describe 'bio_url' do
    it 'returns the base complete info url' do
      expect(parser.bio_url).to eq('http://www2.camara.leg.br/deputados/pesquisa/layouts_deputados_biografia')
    end

    it 'returns the complete info url to deputado' do
      deputado_id = "232"
      expect(parser.bio_url(deputado_id)).to eq("http://www2.camara.leg.br/deputados/pesquisa/layouts_deputados_biografia?pk=#{deputado_id}")
    end
  end

  describe 'video_url' do
    it 'returns the video url' do
      expect(parser.video_url).to eq('http://www2.camara.leg.br/atividade-legislativa/webcamara/resultadoDep')
    end

    it 'returns the video url of deputado' do
      nome_parlamentar = 'TIRIRICA'
      expect(parser.video_url(nome_parlamentar)).to eq("http://www2.camara.leg.br/atividade-legislativa/webcamara/resultadoDep?dep=#{nome_parlamentar}")
    end
  end
end
