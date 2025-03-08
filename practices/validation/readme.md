1. register request

curl -X POST http://localhost:8002/register \
-H "Content-Type: application/json" \
-d '{
"username": "jd",
"email": "invalid-email",
"age": 17,
"password": "short"
}'

In this practice, only provide the validation error response

{"error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'min' tag\nKey: 'User.Email' Error:
Field validation for 'Email' failed on the 'email' tag\nKey: 'User.Age' Error:Field validation for 'Age' failed on the '
gte' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'min' tag","success":false}%

2. order request

curl -X POST http://localhost:8002/orders \
-H "Content-Type: application/json" \
-d '{
"customer_id": "",
"items": [
{
"product_id": "",
"name": "",
"quantity": 0,
"price": -10
},
{
"quantity": 11
}
],
"total_amount": 0,
"status": "invalid_status",
"notes": "'"$(printf "%0.sA" {1..600})"'"
}'

In this practice, only provide the validation error response

order response:

{"success":false,"error":"Key: 'Order.CustomerID' Error:Field validation for 'CustomerID' failed on the 'required'
tag\nKey: 'Order.Items[0].ProductID' Error:Field validation for 'ProductID' failed on the 'required' tag\nKey: '
Order.Items[0].Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Order.Items[0].Quantity'
Error:Field validation for 'Quantity' failed on the 'required' tag\nKey: 'Order.Items[0].Price' Error:Field validation
for 'Price' failed on the 'gt' tag\nKey: 'Order.Items[1].ProductID' Error:Field validation for 'ProductID' failed on
the 'required' tag\nKey: 'Order.Items[1].Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: '
Order.Items[1].Quantity' Error:Field validation for 'Quantity' failed on the 'lte' tag\nKey: 'Order.Items[1].Price'
Error:Field validation for 'Price' failed on the 'required' tag\nKey: 'Order.TotalAmount' Error:Field validation for '
TotalAmount' failed on the 'required' tag\nKey: 'Order.Status' Error:Field validation for 'Status' failed on the 'oneof'
tag\nKey: 'Order.Notes' Error:Field validation for 'Notes' failed on the 'max' tag"}%  


