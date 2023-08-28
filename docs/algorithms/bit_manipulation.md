# Bit manipulation

There are 4 operators
- `|` - If any bit is 1, then the result will be 1. Otherwise, the result is 0.
- `&` - If all bits are 1, then the result will be 1. Otherwise, the result is 0.
- `^` - If the number of 1 bits is odd, then the result will be 1. Otherwise, the result is 0.
- `<<`, `>>` - Left and right shifts - Shifts move all bits over one place in the respective direction. Shifts correspond to multipling or dividing a number by 2. (Left shift = multiply by 2, right shift = floor division by 2)

Sometimes, we want to focus on specific bits rather than every bit in a number.
To do this, we can use something called a bitmask, also just called a mask.
A mask with one bit flipped can be used to check individual bits of other numbers.
For example, let's say that we were concerned only about the 2nd bit (0 indexed from the right, so the bit that represents the decimal number 4) of a number x.

### More about XOR:
- An easy way to think about XOR: if you do `x XOR 0`, then nothing changes with `x`. If you do `x XOR 1`, then `x` is flipped. It works both ways, so `1 XOR x` flips x, and `0 xor X` doesn't.
- If you xor a value with another value twice, nothing changes `X XOR 12345 XOR 12345 = X`, because all the bits that need to be flipped, will be flipped twice and become the same thgey were. This [problem](https://leetcode.com/problems/single-number/) can be solved with this rule.

### Hiding data in unused bits
You can hide data in the bits you don't use. E.g. if a problem says that your input is only positive, you can use minus to mark that element. When using that element again, you can use it's `ABS` value. 
Another example: if your input is very small, you can just use the lastes bits to hide data in them.
