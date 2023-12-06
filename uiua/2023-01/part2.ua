FirstLast ← parse ⊂ ⊢ , ⊢⇌ .
FirstLastDigitsWithString ← FirstLast ∊, "123456789" .
&fras "../../2023/01/input.txt" # read input
⊜□ ≠@\n .                       # partition into lines

⊐∵FirstLastDigitsWithString .   # find and parse first and last digits in string, including words
/+ .                            # and sum them all
