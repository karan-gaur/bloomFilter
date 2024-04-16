with open('keywords1.txt', 'r') as file:
    contents = file.read().split(",")
    for word in contents:
        if len(word) > 0:
            print(word)
