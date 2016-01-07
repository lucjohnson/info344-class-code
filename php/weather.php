<?php
$ipUrl = "http://ip-api.com/json";
$ip = json_decode(file_get_contents($ipUrl), true);
echo '<pre>';
print_r($ip);

$url = "http://api.openweathermap.org/data/2.5/weather?q=Seattle,us&units=imperial&appid=2de143494c0b295cca9337e1e96b00e0";
$results = json_decode(file_get_contents($url), true);
print_r($results);
foreach($results as $result) {
	print_r($result);
}
echo '</pre>';
?>