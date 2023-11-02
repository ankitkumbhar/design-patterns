<?php

class Product {
    private $name, $price, $description;

    public function setName($name) {
        $this->name = $name;
    }

    public function setPrice($price) {
        $this->price = $price;
    }

    public function setDescription($desc) {
        $this->description = $desc;
    }

    public function displayProduct() {
        echo "\nProduct details : \n";
        echo "\nName : " . $this->name . "\nPrice : " . $this->price . "\nDesc : " . $this->description . "\n";
    }
}

class ProductBuilder {
    private $product;

    public function __construct() {
        $this->product = new Product();
    }

    public function setName($name) {
        $this->product->setName($name);
        
        return $this;
    }

    public function setPrice($price) {
        $this->product->setPrice($price);
        
        return $this;
    }

    public function setDesc($desc) {
        $this->product->setDescription($desc);
        
        return $this;
    }

    public function build() {
        return $this->product;
    }
}

$obj = new ProductBuilder();
$obj->setName("Laptop")
    ->setPrice(140000)
    ->setDesc("Smaple description");

$product = $obj->build();

$product->displayProduct();

?>