# CRC-32

Used in ethernet frames for error detection. The Frame Check Sequence (FCS) in the ethernet frame trailer is a CRC-32 checksum.

The Cyclic Redundancy Check algorithm as defined in IEEE 802.3-2022 Standard for Ethernet is below.

1) The first 32 bits of the frame are complemented. 
2) The n bits of the protected fields are then considered to be the coefficients of a polynomial M(x) of
degree n – 1. (The first bit of the Destination Address field corresponds to the x(n–1) term and the last
bit of the MAC Client Data field (or Pad field if present) corresponds to the x0 term.)
3) M(x) is multiplied by $x^{32}$ and divided by G(x), producing a remainder R(x) of degree <=31.
4) The coefficients of R(x) are considered to be a 32-bit sequence.
5) The bit sequence is complemented and the result is the CRC.

Note: I had some trouble wrapping my head around step 3, specifically the multiplication part. Multiply by $x^{32}$ because 32 is the degree of a 32 bit polynomial, so for a the polynomial represented by 1001 which is degree 3 multiply by 3.

## Worked example

given a 10 bit input and using a 4 bit polynomial, 
polynomial = $G(x) = x^3 + 1$ (1001)
input = 1010110110

1) 0101110110
2) $M(x) = 0x^9+x^8+0x^7+x^6+x^5+x^4+0x^3+x^2+x+0$
3) $M(x)* x^3 / G(x)= (x^{11}+x^9+x^8+x^7+x^5+x^4) / (x^3 + 1)$ 
$$ 
\begin{array}{r}
x^8+x^6+x^4-x^3+x^2+1\\
{x^3 + 1 }\overline{)x^{11}+x^9+x^8+x^7+x^5+x^4}  \\
-x^{11}-0x^{10}-0x^9-x^8 \\
\hline
x^9+x^7+x^5+x^4 \\
-x^9-0x^8-0x^7-x^6\\
\hline  
x^7-x^6+x^5+x^4\\
-x^7-0x^6-0x^5-x^4\\
\hline
-x^6+x^5\\
+x^6+0x^5+0x^4+x^3\\
\hline
x^5+x^3\\
-x^5-0x^4-0x^3-x^2\\
\hline
x^3-x^2\\
-x^3-0x^2-0x-1\\
\hline
-x^2-1\\
\end{array} \\

R =-x^2-1 
$$

This is implimented as a series of xors  
In the binary representation $x^{11}+x^9+x^8+x^7+x^5+x^4$ is 
0101110110000

$$
\begin{array}{r}
0101011101\\
{1001 }\overline{)0101110110000}\\
\oplus            0000000000000\\
\hline
                   101110110000\\
\oplus             100100000000\\
\hline
                    01010110000\\
\oplus              00000000000\\
\hline
                     1010110000\\
\oplus               1001000000\\
\hline
                      011110000\\
\oplus                000000000\\
\hline
                       11110000\\
\oplus                 10010000\\
\hline
                        1100000\\
\oplus                  1001000\\
\hline
                         101000\\
\oplus                   100100\\
\hline
                          01100\\
\oplus                    00000\\
\hline
                           1100\\
\oplus                     1001\\
\hline
                            101\\



\end{array} 
R=101
$$

4) Represented as binary coefficents $x^8+x^6+x^4-x^3+x^2+1$ is $0101011101$ and the remainder $R(x) = -x^2-1$ is $101$

5) $CRC = 010$ Which is then appended to the message