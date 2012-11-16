# encoding: utf-8

require 'spec_helper'

describe PesquisaDeputadosParser do
  let(:parser) { PesquisaDeputadosParser.new }

  describe '#deputados' do
    it 'returns 5 register' do
      expect(parser.deputados.count).to eq(5)
    end

    it 'returns names and id of deputados' do
      expect(parser.deputados).to eq([
        {id: '141463', name: 'ABELARDO CAMARINHA'},
        {id: '74354', name: 'ZENALDO COUTINHO'},
        {id: '73933', name: 'ZEQUINHA MARINHO'},
        {id: '74145', name: 'ZEZÉU RIBEIRO'},
        {id: '160625', name: 'ZOINHO'}
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
end
