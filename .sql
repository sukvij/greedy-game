CREATE database greedy-game;
CREATE TABLE campaigns (
    cid VARCHAR(255) primary key NOT NULL,
    campaign_name VARCHAR(255),
    img VARCHAR(255),
    cta VARCHAR(255),
    status VARCHAR(255) NOT NULL CHECK (status IN ('ACTIVE', 'PASSIVE'))
);


CREATE TABLE targeting_rules (
    cid VARCHAR(255) primary key NOT NULL,
    rules JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);


-- // add foreign key to campaign also
ALTER TABLE targeting_rules
ADD CONSTRAINT fk_campaign
FOREIGN KEY (cid)
REFERENCES campaigns (cid)
ON DELETE CASCADE;


-- select * from campaigns inner join targeting_rules on 
-- campaigns.cid = targeting_rules.cid 
-- where targeting_rules.rules ->'include_country' @> '["US"]'::jsonb


-- to include
-- countries := []string{"usa", "ca"}
-- countryJSON, _ := json.Marshal(countries)
-- db.Where("t.rules->'include_country' @> ?", string(countryJSON))








-- SELECT 
--     c.cid,
--     c.campaign_name,
--     c.img,
--     c.cta,
--     c.status,
--     t.rules
-- FROM targeting_rules t
-- INNER JOIN campaigns c ON t.cid = c.cid
-- WHERE 
--      (
--         (t.rules->'include_country' IS NOT NULL AND t.rules->'include_country' @> '["US"]'::jsonb)
--         OR
--         (t.rules->'include_country' IS NULL AND NOT (t.rules->'exclude_country' @> '["India"]'::jsonb))
--     )
--     AND (
--         (t.rules->'include_os' IS NOT NULL AND t.rules->'include_os' @> '["Android"]'::jsonb)
--         OR
--         (t.rules->'include_os' IS NULL AND NOT (t.rules->'exclude_os' @> '["Android"]'::jsonb))
--     )
--     AND (
--         (t.rules->'include_app' IS NOT NULL AND t.rules->'include_app' @> '["com.duolingo.ludokinggame"]'::jsonb)
--         OR
--         (t.rules->'include_app' IS NULL AND NOT (t.rules->'exclude_app' @> '["com.duolingo.ludokinggame1"]'::jsonb))
--     );