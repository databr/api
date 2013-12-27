# encoding: UTF-8
#
class DeputadoPropositionsService
  def self.save_propositions
    parser_xml = DeputadoXMLParser.new
    parser_xml.deputados.each do |deputado|
      begin
        parser = DeputadoProposicaoParser.new(OpenStruct.new(deputado), ENV['YEAR'])
        parser.propositions.each do |proposition_data|
          puts "\e[32m  * Finding Proposition #{proposition_data[:number]}(#{proposition_data[:proposition_id]}) from save propositions\e[0m"
          proposition = Proposition.where(proposition_id: proposition_data[:proposition_id])
          if proposition.first
            puts "\e[32m    * Updating Proposition #{proposition_data[:number]}(#{proposition_data[:proposition_id]}) from save propositions\e[0m"
            proposition.first.update_attributes proposition_data
          else
            puts "\e[32m    * Creating Proposition #{proposition_data[:number]}(#{proposition_data[:proposition_id]}) from save propositions\e[0m"
            proposition.create! proposition_data
          end
        end
      rescue
        puts "\e[31m    - Error from save propositions\e[0m"
        `echo "#{deputado} -  " >> proposition_err`
       next
      end
    end
  end
end
