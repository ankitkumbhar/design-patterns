<?php

define("PaymentGatewayTypeStripe", 1);
define("PaymentGatewayTypePaypal", 2);

interface PaymentGateway {
    public function ProcessPayment();
}

class StripeGateway implements PaymentGateway {
    public function ProcessPayment() {
        echo "\nStripe Gateway : Process Payment\n";
    }
}

class PaypalGateway implements PaymentGateway {
    public function ProcessPayment() {
        echo "\nPaypal Gateway : Process Payment\n";
    }
}

class Factory {
    public function factoryMethod($gatewayType) {
        switch($gatewayType) {
            case 1:
                return new StripeGateway();
                break;
            case 2:
                return new PaypalGateway();
                break;
        }
    }
}

$factory = new Factory();

$obj = $factory->factoryMethod(PaymentGatewayTypeStripe);
$obj->ProcessPayment();

$obj = $factory->factoryMethod(PaymentGatewayTypePaypal);
$obj->ProcessPayment();

?>
