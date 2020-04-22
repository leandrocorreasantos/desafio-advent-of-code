<?php
$andar = 0;
$subsolo = 0;
$posicao = 0;

$arquivo = file('input.txt');

foreach($arquivo as $linha){
	$caracteres = str_split($linha);
	foreach($caracteres as $char){
		$posicao++;

		if($char == '('){
			$andar = $andar + 1;
		}else if($char == ')'){
			$andar = $andar - 1;
		}

		if($subsolo == 0){
			if($andar == -1){
				$subsolo = $posicao;
			}
		}

	}
}

echo "Resultado parte 1: $andar \n";
echo "Resultado parte 2: $subsolo \n";
?>
