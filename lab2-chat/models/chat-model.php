<?php
class Messages {
	protected $conn;
	
	public function __construct($conn) {
		$this->conn = $conn;
	}
	
	public function getMessages() {
		$sql = 'select * from message';
		$stmt = $this->conn->prepare($sql);
		return $stmt->fetchAll();
	}
}
?>