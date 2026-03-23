# DeleteProductByID

## รายละเอียด

API นี้ใช้สำหรับลบข้อมูลสินค้าตาม `id`

ระบบจะพยายามลบ record ที่ตรงกับ `id` แล้วคืนผลลัพธ์เป็นข้อความ JSON กลับมา

## Endpoint

```http
DELETE /api/product/:id
```

## Method

- `DELETE`

## Path

- `/api/product/:id`

## Request Parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| id | Integer | Yes | ID ของสินค้าที่ต้องการลบ |

## Request Example

```bash
curl --location --request DELETE 'http://localhost:3000/api/product/1'
```

## Response Example

```json
"delated"
```

## Error Codes

| Code | Message | Description |
| --- | --- | --- |
| 400 | invalid request body | `id` ไม่ถูกต้อง |
| 404 | product not found | ไม่พบข้อมูลสินค้า |
| 500 | failed to fetch product | เกิดข้อผิดพลาดระหว่างลบข้อมูล |

## Notes

- response success ปัจจุบันสะกดเป็น `delated` ตามโค้ดที่มีอยู่จริง
- ในเวอร์ชันนี้ยังไม่มี response object แบบ `{ "message": "deleted" }`