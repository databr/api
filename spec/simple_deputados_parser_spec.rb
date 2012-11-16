# encoding: utf-8

require 'spec_helper'

describe SimpleDeputadosParser do
  let(:parser) { SimpleDeputadosParser.new }

  before do
    stub_request(:get, "http://www2.camara.leg.br/deputados/pesquisa").
      to_return(:status => 200, :body => File.read("spec/fixtures/pesquisa.html"), :headers => {:"Content-Type" => "text/html"})
  end

  describe '#deputados' do
    it "returns 5 register" do
      expect(parser.deputados.count).to eq(5)
    end

    it "returns names and id of deputados" do
      expect(parser.deputados).to eq([
        {id: "141463", name: "ABELARDO CAMARINHA"},
        {id: "74354", name: "ZENALDO COUTINHO"},
        {id: "73933", name: "ZEQUINHA MARINHO"},
        {id: "74145", name: "ZEZÉU RIBEIRO"},
        {id: "160625", name: "ZOINHO"}
      ])
    end
  end
end
