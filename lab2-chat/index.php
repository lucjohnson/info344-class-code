<?php
require_once 'connection.php';
require_once 'models/chat-model.php';

$conn = getConnection();
$model = new Messages();
$messages = $model->getMessages();
var_dump($messages);
?>
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Chattin'</title>
</head>
<body>
	<form action="" method="GET">
		<input type="textarea"></input>
		<input type="submit"></input>
	</form>
</body>
</html>