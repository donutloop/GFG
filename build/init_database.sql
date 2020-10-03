CREATE DATABASE IF NOT EXISTS `product`;

USE `product`;

CREATE TABLE IF NOT EXISTS `seller` (
  `id_seller` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `phone` VARCHAR(100) NOT NULL,
  `uuid` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`id_seller`),
  UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  ROW_FORMAT = DYNAMIC;

CREATE TABLE IF NOT EXISTS `product`
(
  `id_product` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name`       VARCHAR(200)     NOT NULL,
  `brand`      VARCHAR(200)     NOT NULL,
  `stock`      INT(10) DEFAULT 0,
  `fk_seller`  INT(10) unsigned NOT NULL,
  `uuid`       VARCHAR(36)      NOT NULL,
  PRIMARY KEY (`id_product`),
  UNIQUE KEY `uuid` (`uuid`),
  CONSTRAINT fk_seller FOREIGN KEY (fk_seller) REFERENCES seller (id_seller)
) ENGINE = InnoDB
  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

INSERT INTO seller (id_seller, name, email, phone, uuid) VALUES
(1, 'Christene Maggio', 'christene.maggio@seller.com', '202-555-0143', UUID()),
(2, 'Owen Ringgold', 'owen.ringgold@seller.com', '202-555-0188', UUID()),
(3, 'Shani Marinello', 'shani.marinello@seller.com', '202-555-0138', UUID()),
(4, 'Lajuana Mooring', 'lajuana.mooring@seller.com', '202-555-0103', UUID()),
(5, 'Tatyana Moua', 'tatyana.moua@seller.com', '202-555-0178', UUID()),
(6, 'Chelsie Wurster', 'chelsie.wurster@seller.com', '202-555-0187', UUID()),
(7, 'Syble Coria', 'syble.coria@seller.com', '202-555-0132', UUID()),
(8, 'Cathleen Swick', 'cathleen.swick@seller.com', '202-555-0166', UUID()),
(9, 'Cathleen Wurster', 'cathleen.wurster@seller.com', '202-555-0166', UUID()),
(10, 'Tatyana Ringgold', 'tatyana.ringgold@seller.com', '202-555-0166', UUID()),
(11, 'Chelsie Maggio', 'chelsie.maggio@seller.com', '202-555-0166', UUID()),
(12, 'Lajuana Marinello', 'lajuana.marinellok@seller.com', '202-555-0166', UUID()),
(13, 'Owen Swick', 'owen.swick@seller.com', '202-555-0166', UUID()),
(14, 'Christene Ringgold', 'christene.ringgold@seller.com', '202-555-0166', UUID());

INSERT INTO product (id_product, name, brand, stock, fk_seller, uuid) VALUES
(1, 'Raja', 'RJ', 100, 1, UUID()),
(2, 'Pure Linen Plain Shirt', 'ShirtsCo', 44, 1, UUID()),
(3, 'Long Sleeve Nursing Tops', 'ShirtsCo', 12, 1, UUID()),
(4, 'Christian Dior Pure Poison Eau De Parfum Spray', 'Christian Dior', 1031, 2, UUID()),
(5, 'Speed TR Flexweave Shoes', 'Flexweave', 1300, 2, UUID()),
(6, 'Berlin New Shirt', 'ShirtsCo', 2100, 2, UUID()),
(7, '2 Pack Bikini Brief', 'ShirtsCo', 3200, 2, UUID()),
(8, 'Organic Moisturizing Lip Balm', 'Organic', 10, 2, UUID()),
(9, 'Combi Leather Sandals', 'ShoesCo', 1100, 3, UUID()),
(10, 'Sasha', 'Sasha', 120, 3, UUID()),
(11, 'Womens Bella Surfing Maxi Tank Dress', 'ShirtsCo', 2100, 4, UUID()),
(12, 'Plano Tee', 'TeeCo', 10, 4, UUID()),
(13, 'Solid Cross Back Bikini Top', 'ShirtsCo', 120, 3, UUID()),
(14, 'Ottana Sweater', 'Ottana', 3200, 1, UUID()),
(15, 'Brushed Herringbone Pant With Tape Detail In Loose Tapered Fit', 'Ottana', 100, 5, UUID()),
(16, 'Mina', 'Mina', 321, 6, UUID()),
(17, 'Claire Top', 'Claire', 1321, 7, UUID()),
(18, 'Calvin Klein Fully Delicious Sheer Plumping Lip Gloss', 'Calvin Klein', 132, 8, UUID()),
(19, 'Shara Shara White Stem Sleeping Mask', 'Shara Shara', 1021, 9, UUID()),
(20, 'Classy & Fabulous Coat', 'ShirtsCo', 1332, 11, UUID()),
(21, 'Black Cherry Dress', 'ShirtsCo', 11, 12, UUID()),
(22, 'Storm Jacket', 'ShirtsCo', 145, 14, UUID()),
(23, 'Sweet Daisy Top', 'ShirtsCo', 14400, 14, UUID());
