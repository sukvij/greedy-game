CREATE TYPE campaign_status AS ENUM('active', 'passive');
CREATE TABLE campaigns (
    id BIGSERIAL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    cid VARCHAR(255) PRIMARY KEY NOT NULL,
    campaign_name VARCHAR(255),
    img VARCHAR(255),
    cta VARCHAR(255),
    status campaign_status NOT NULL,
);


CREATE TABLE targeting_rules (
    cid VARCHAR(255) primary key NOT NULL,
    rules JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);


-- // add foreign key to campaign also