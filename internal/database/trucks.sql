-- For number, 4 digits before the decimal point and a precision of 2 digits after it.
CREATE TABLE IF NOT EXISTS trucks(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    fuel_type VARCHAR(25),
    payload DECIMAL(4,2), 
    distance DECIMAL(7,2)
) CHARSET=utf8mb4;

INSERT INTO trucks(fuel_type, payload, distance) 
    VALUES ROW('Diesel', 44, 500), ROW('Gasoline', 19, 200), ROW('Electricity', 3.5, 100);