CREATE TABLE IF NOT EXISTS `product_images` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    image_url VARCHAR(255),
    image_description VARCHAR(255)
);
