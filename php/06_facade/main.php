<?php

class CPU {
    public function start() {
        echo "\nCPU - staring\n";
    }

    public function execute() {
        echo "\nCPU - executing\n";
    }
}

class Memory {
    public function load() {
        echo "\nMemory - loading\n";
    }
}

class Disk {
    public function read() {
        echo "\nDisk - reading\n";
    }
}

class ComputerFacade {
    private $cpuClass, $memoryClass, $diskClass;

    public function __construct() {
        $this->cpuClass = new CPU();
        $this->memoryClass = new Memory();
        $this->diskClass = new Disk();
    }

    public function start() {
        $this->cpuClass->start();
        $this->memoryClass->load();
        $this->diskClass->read();
        $this->cpuClass->execute();
    }
}

$facade = new ComputerFacade();

$facade->start();

?>