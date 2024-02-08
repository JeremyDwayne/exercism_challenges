=begin
Write your code for the 'Reverse String' exercise in this file. Make the tests in
`reverse_string_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/reverse-string` directory.
=end

class Reverser
  def self.reverse(text)
    return text if text.empty?

    n = text.length
    (0...(n/2)).each do |i|
      temp = text[i]
      text[i] = text[n - 1 -i]
      text[n - 1 - i] = temp
    end
    text
  end
end
