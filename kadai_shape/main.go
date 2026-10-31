package main

import (
	"errors"
	"fmt"
	"math"
)

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Field=%s, Value=%.2f", e.Field, e.Value)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func validateShape(s Shape) error {
	if r, ok := s.(Rectangle); ok {
		if r.Height <= 0 {
			return &ValidationError{"高さは正の数を指定してください。", "Rectangle.Height", r.Height}
		}
		if r.Width <= 0 {
			return &ValidationError{"幅は正の数を指定してください。", "Rectangle.Width", r.Width}
		}
	}

	if c, ok := s.(Circle); ok {
		if c.Radius <= 0 {
			return &ValidationError{"半径は正の数を指定してください。", "Circle.Radius", c.Radius}
		}
	}
	return nil
}

func main() {
	// 処理対象の図形スライスを作成
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}

	fmt.Println("--- 図形のバリデーション ---")
	var calcShapes []Shape
	for i, s := range shapes {
		validated := validateShape(s)
		var validateResult string
		var validationError *ValidationError
		if errors.As(validated, &validationError) {
			validateResult = fmt.Sprintf("NG：[バリデーションエラー] %s：%s（Value：%.2f）\n [詳細] %v", validationError.Field, validationError.Message, validationError.Value, validationError.Error())
		} else {
			validateResult = "バリデーションOK"
			calcShapes = append(calcShapes, s)
		}

		switch s.(type) {
		case Rectangle:
			fmt.Printf("#%d：main.Rectangle => %s\n", i+1, validateResult)
		case Circle:
			fmt.Printf("#%d：main.Circle => %s\n", i+1, validateResult)
		}

	}

	fmt.Println("\n--- 検証済み図形の一括計算 ---")
	for _, cs := range calcShapes {
		switch s := cs.(type) {
		case Rectangle:
			fmt.Printf("図形：main.Rectangle｜面積：%.2f ｜周長：%.2f\n", s.Area(), s.Perimeter())
		case Circle:
			fmt.Printf("図形：main.Circle｜面積：%.2f ｜周長：%.2f\n", s.Area(), s.Perimeter())
		}
	}
}
