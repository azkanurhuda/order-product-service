CREATE TABLE orders (
                        id INT PRIMARY KEY,
                        customer_id INT,
                        product_id INT,
                        quantity INT NOT NULL,
                        total DECIMAL(10, 2) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        FOREIGN KEY (customer_id) REFERENCES costumers(id),
                        FOREIGN KEY (product_id) REFERENCES products(id)
)  engine = InnoDB;
