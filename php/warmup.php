<?php
	$num = rand(1, 100);
	echo "your new random value is $num\n";
	
	echo "\nMONTHS\n";
	$months = cal_info(0);
	$months = $months['months'];
	foreach($months as $month) {
		echo $month . "\n";
	}
	
	echo "\nALPHABETIZED MONTHS\n";
	sort($months);
	foreach($months as $m) {
		echo $m . "\n";
	}
?>