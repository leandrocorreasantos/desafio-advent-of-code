andar = 0

with open('input.txt', 'r') as arquivo:
    for linha in arquivo:
        for char in linha:
            if char == '(':
                andar = andar + 1
            elif char == ')':
                andar = andar - 1

    print('Resultado: {}'.format(andar))
