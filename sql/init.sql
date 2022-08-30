CREATE TABLE m_product_category
(
    id            char(10) PRIMARY KEY,
    category_name varchar(100) not null
);
CREATE TABLE m_product
(
    id           char(10) PRIMARY KEY,
    product_name varchar(100) not null,
    category_id  char(10),
    constraint fk_product_category foreign key (category_id) references m_product_category (id)
);
