FirstLast ← parse ⊂ ⊢ , ⊢⇌ .
FirstLastDigits ← FirstLast ▽ ∊, "123456789" .
&fras "../../2023/01/input.txt" # read input
⊜□ ≠@\n .                       # partition into lines
⊐∵FirstLastDigits .             # find and parse first and last digits in string
/+ .                            # and sum them all
