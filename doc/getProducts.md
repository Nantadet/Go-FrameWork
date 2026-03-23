# GetProducts

## รายละเอียด

API นี้ใช้สำหรับดึงรายการสินค้าทั้งหมดในระบบ

เหมาะสำหรับ:

- แสดงรายการสินค้าในหน้า dashboard
- ใช้โหลดข้อมูลทั้งหมดก่อนเลือกแก้ไขหรือลบ
- ใช้ตรวจสอบข้อมูลสินค้าที่มีอยู่ในฐานข้อมูล

## Endpoint

```http
GET /api/product
```

## Method

- `GET`

## Path

- `/api/product`

## Request Parameters

ไม่มี request body และไม่มี path parameter

## Request Example

```bash
curl --location --request GET 'http://localhost:3000/api/product'
```

## Response Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| id | Integer | ID ของสินค้า |
| name | String | ชื่อสินค้า |
| price | Integer | ราคาสินค้า |

## Response Example

```json
[
  {
    "id": 1,
    "name": "Keyboard",
    "price": 890
  },
  {
    "id": 2,
    "name": "Mouse",
    "price": 390
  }
]
```

## Error Example

```json
{
  "error": "failed to fetch products"
}
```

## Notes

- ถ้าไม่มีข้อมูลสินค้า API จะคืน array ว่าง `[]`
- endpoint นี้ดึงข้อมูลจากฐานข้อมูลทั้งหมดโดยไม่มี pagination ในเวอร์ชันปัจจุบัน