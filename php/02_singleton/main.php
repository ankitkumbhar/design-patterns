<?php

class Singleton {
    private static $instance;

    public static function new($objectName) {
        
        if (self::$instance == NULL) {
            echo "\nCreating object for : $objectName\n";
            
            self::$instance = new Singleton;
        }

        return self::$instance;
    }

    public function display() {
        echo "\nThis is singleton class method\n";
    }
}

$obj1 = Singleton::new("firstObject");
$obj1->display();

$obj2 = Singleton::new("secondObject");
$obj2->display();

?>