# encoding: UTF-8
#
class DeputadoPropositionsService
  def self.save_propositions
    parser = DeputadoXMLParser.new
    parser.deputados.each do |deputado|
      begin
      parser = DeputadoProposicaoParser.new(OpenStruct.new(deputado), ENV['YEAR'])
      parser.propositions.each do |proposition_data|
        proposition = Proposition.where(proposition_id: proposition_data[:proposition_id])
        proposition.create! proposition_data unless proposition.first
      end
      rescue => e
        puts "Error: #{deputado}"
        `echo "#{deputado} - - - #{e}" >> proposition_err-#{parser.object_id}`
      end
    end
  end
end
