<?php

// Abstraction: Shape
abstract class Shape
{
    protected $drawingTool;

    public function __construct(DrawingTool $drawingTool)
    {
        $this->drawingTool = $drawingTool;
    }

    abstract public function draw();
}

// Refined Abstraction: Circle
class Circle extends Shape
{
    public function draw()
    {
        return $this->drawingTool->drawCircle();
    }
}

// Refined Abstraction: Square
class Square extends Shape
{
    public function draw()
    {
        return $this->drawingTool->drawSquare();
    }
}

// Implementation: DrawingTool
interface DrawingTool
{
    public function drawCircle();
    public function drawSquare();
}

// Concrete Implementation: Pen
class Pen implements DrawingTool
{
    public function drawCircle()
    {
        return "Drawing a circle with a pen";
    }

    public function drawSquare()
    {
        return "Drawing a square with a pen";
    }
}

// Concrete Implementation: Brush
class Brush implements DrawingTool
{
    public function drawCircle()
    {
        return "Drawing a circle with a brush";
    }

    public function drawSquare()
    {
        return "Drawing a square with a brush";
    }
}

// Client code
$pen = new Pen();
$brush = new Brush();

$circleWithPen = new Circle($pen);
$squareWithBrush = new Square($brush);

echo $circleWithPen->draw() . "\n"; // Output: Drawing a circle with a pen
echo $squareWithBrush->draw() . "\n"; // Output: Drawing a square with a brush