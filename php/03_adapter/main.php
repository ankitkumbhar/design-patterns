<?php

interface stripeInterface {
    public function ProcessPayment();
}

class StripeGateway implements stripeInterface {
    public function ProcessPayment() {
        echo "\nStripeGateway - ProcessPayment\n";
    }
}

interface payPalInterface {
    public function MakePayment();
}

class PayPalGateway implements payPalInterface {
    public function MakePayment() {
        echo "\nPayPayGateway - MakePayment\n";
    }
}

class StripeAdapter implements payPalInterface {
    private $stripeObj;

    public function __construct($obj) {
        $this->stripeObj = $obj;
    }

    public function MakePayment() {
        $this->stripeObj->ProcessPayment();
    }
}

$obj = new StripeAdapter(new StripeGateway());
$obj->MakePayment();

$obj2 = new PayPalGateway();
$obj2->MakePayment();

?>