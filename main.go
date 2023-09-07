package main

import (
	"fmt"
)

// * ฟังก์ชันที่รับ pointer เป็นพารามิเตอร์ เพื่อแก้ไขค่าที่ตัวแปรนั้นชี้ไป
func modifyValue(p *int) {
	*p = *p * 2 //* คูณค่าที่ p ชี้ไปด้วย 2
}

func main() {
	x := 5
	fmt.Println("ค่าเริ่มต้นของ x:", x)

	//* ส่ง address ของ x ไปยังฟังก์ชัน modifyValue
	modifyValue(&x)

	fmt.Println("ค่าหลังจากการแก้ไข:", x)
}
