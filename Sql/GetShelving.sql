SELECT product_id, products.shelf_id, is_main, product_count,
       product_name, shelf_name
FROM products LEFT JOIN shelving
ON products.shelf_id = shelving.shelf_id
WHERE product_id = $1
