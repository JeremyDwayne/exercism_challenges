=begin
Write your code for the 'Matrix' exercise in this file. Make the tests in
`matrix_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/matrix` directory.
=end
class Matrix
  attr_reader :matrix

  def initialize(text)
    @matrix = prep_text(text)
  end

  def row(row_num)
    matrix[row_num - 1] 
  end

  def column(col_num)
    output = []
    matrix.each do |row|
      output << row[col_num - 1]
    end
    output
  end
  
  private

  def prep_text(text)
    temp = text.split("\n")
    temp.map! { |x| x.split(' ').map!(&:to_i) }
  end
end
