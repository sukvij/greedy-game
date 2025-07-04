CREATE TABLE campaigns (
    cid VARCHAR(255) NOT NULL,
    campaign_name VARCHAR(255),
    img VARCHAR(255),
    cta VARCHAR(255),
    status ENUM('ACTIVE', 'PASSIVE') NOT NULL,
    PRIMARY KEY (cid)
);