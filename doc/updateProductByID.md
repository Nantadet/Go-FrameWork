# UpdateProductByID

## รายละเอียด

API นี้ใช้สำหรับอัปเดตข้อมูลสินค้าตาม `id`

ระบบจะค้นหาสินค้าตาม `id` ก่อน ถ้าพบจึงจะอัปเดตชื่อและราคาใหม่ให้

## Endpoint

```http
PUT /api/product/:id
```

## Method

- `PUT`

## Path

- `/api/product/:id`

## Request Parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| id | Integer | Yes | ID ของสินค้าที่ต้องการอัปเดต |
| name | String | Yes | ชื่อสินค้าใหม่ |
| price | Integer | Yes | ราคาสินค้าใหม่ |

## Request Example

```bash
curl --location --request PUT 'http://localhost:3000/api/product/1' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Mechanical Keyboard",
  "price": 1290
}'
```

## Response Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| id | Integer | ID ของสินค้า |
| name | String | ชื่อสินค้าหลังอัปเดต |
| price | Integer | ราคาสินค้าหลังอัปเดต |

## Response Example

```json
{
  "id": 1,
  "name": "Mechanical Keyboard",
  "price": 1290
}
```

## Error Codes

| Code | Message | Description |
| --- | --- | --- |
| 400 | invalit request body. | `id` ไม่ถูกต้อง |
| 400 | Invalit request body | body ไม่ถูกต้อง |
| 400 | request Name | ส่ง `name` ว่างหลัง trim space |
| 400 | Product not found | ไม่พบสินค้าที่ต้องการอัปเดต |
| 500 | Error fetching school | เกิดข้อผิดพลาดตอนดึงข้อมูลเดิม |
| 400 | Error updating schools | เกิดข้อผิดพลาดตอนบันทึกข้อมูลใหม่ |

## Notes

- ข้อความ error ตอนนี้ยังอิงจากชื่อเก่าในโค้ดบางจุด เช่น `school` และ `invalit`
- endpoint นี้มีการ `TrimSpace` ให้กับ `name` ก่อนตรวจสอบค่าว่าง