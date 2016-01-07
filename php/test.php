Hey this is some content above the code
<?php 
$name = 'Luc';
$fullName = $name . ' Johnson';

class Person {
	protected $name;
	
	public function __construct($n) {
		$this->name = $n;
	}
	
	public function getName() {
		return $this->name;
	}
}

function foo($bar) {
	echo "Hey this is the foo fighting function\n";
}

echo "Hello, {$fullName}s\n";
foo(NULL);
?>
And this is some content below
