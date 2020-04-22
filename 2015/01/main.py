andar = 0
subsolo = 0
posicao = 0

with open('input.txt', 'r') as arquivo:
    for linha in arquivo:
        for char in linha:
            posicao = posicao + 1

            if char == '(':
                andar = andar + 1
            elif char == ')':
                andar = andar - 1

            if subsolo == 0:
                if andar == -1:
                    subsolo = posicao

    print('Resultado parte 1: {}'.format(andar))
    print('Resultado parte 2: {}'.format(subsolo))
