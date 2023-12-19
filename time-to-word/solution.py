
words_dict = {
        1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
        6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
        11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
        15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen",
        19: "nineteen", 20: "twenty"
    }
    
# Define words for multiples of 10
tens_dict = {
    10: "ten", 20: "twenty"
}

# Function to convert minutes to words
def minutes_to_words(m):
    if m == 1:
        return "one minute"
    elif m == 15:
        return "quarter"
    elif m == 30:
        return "half"
    elif m == 45:
        return "quarter"
    elif m < 20:
        return words_dict[m] + " minutes"
    else:
        tens = m // 10 * 10
        ones = m % 10
        return tens_dict[tens] + " " + words_dict[ones] + " minutes"
            
def timeInWords(h, m):
    if m == 0:
        return words_dict[h] + " o' clock"
    elif m <= 30:
        return minutes_to_words(m) + " past " + words_dict[h]
    else:
        return minutes_to_words(60 - m) + " to " + words_dict[h + 1]


result = timeInWords(5,47)
print(result)