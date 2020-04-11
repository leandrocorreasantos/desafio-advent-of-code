<?php
$andar = 0;

$arquivo = file('input.txt');

foreach($arquivo as $linha){
	$caracteres = str_split($linha);
	foreach($caracteres as $char){
		if($char == '('){
			$andar = $andar + 1;
		}else if($char == ')'){
			$andar = $andar - 1;
		}
	}
}

echo "Resultado: $andar \n";
?>
