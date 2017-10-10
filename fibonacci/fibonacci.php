<?php
error_reporting(0);
ini_set("memory_limit", 0);

function fibonacci($x, $y) {
	return $x + $y;
}

$x = $y = 0;

for ($j = 0; $j < 1000000; $j++) {
	$numbers = [];
	for ($i = 0; $i < 10; $i++) {
		if ($i <= 1) {
			$numbers[] = fibonacci($x, $i);
		} else {
			$x = $numbers[$i-2];
			$y = $numbers[$i-1];
			
			$numbers[] = fibonacci($x, $y);
		}
	}
}