require 'spec_helper'

describe CotaXMLParser do
  subject { CotaXMLParser.new(File.join(File.dirname(__FILE__), 'fixtures', 'cotas.xml')) }
  describe "#cotas" do
    it 'returns 5 register' do
      expect(subject.cotas.count).to eq(4)
      expect(subject.cotas[0].keys.count).to eq(21)
    end
  end
end
