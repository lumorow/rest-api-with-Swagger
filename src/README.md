# Instruction
## Enter 'make' - create candy client and candy server
## Enter 'make server_run' - server run
## Another terminal 'make client_run' - client run
### Your test client should support flags:
'-k' (accepts two-letter abbreviation for the candy type: "CE - 10 cents", "AA - 15 cents", "NT - 17 cents", "DE - 21 cents" or "YR - 23 cents"),
'-c' (count of candy to buy) and
'-m' (amount of money you "gave to machine").
So, the "buying request" should look like this:
 ____________
### Example ./build/candy-client -k AA -c 1 -m 20

< Thank you! >
------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
Your change is 15
