### Countdown Desk Calendar
My parents used to get Countdown Desk Calendar but it is no longer being published :sadpanda:

This lil number generates a letters round _or_ numbers round and a conundrum for each day of the year.  

The puzzles are placed into a template and sent to the numbers specified in `NUMBERS` env variable using [Twilio](https://www.twilio.com/).

The puzzle is sent daily at 06:39.

### Implementation
The puzzles currently sent are a letter puzzle and a conundrum. 

The letters puzzle selects a random word from the set, buffs it with random letters and shuffles the letters.

The conunndrum selects a random word from the set.

The "randomness" of word selection is in fact pseudo-random. `randn` is seeded with the year date of the puzzle and therefore the previous days solution can be re-generated simply by using `time.Now().Add(-1*time.Hour*24)` to seed. 

### todo
- finish implementation of numbers round
- change from numbers to whatsapp groups?
