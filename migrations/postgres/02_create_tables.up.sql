
CREATE OR REPLACE FUNCTION create_subscription_trigger() RETURNS trigger
LANGUAGE PLPGSQL
AS
$$
    BEGIN
        INSERT INTO user_subscription_history (
        id ,
        user_subscription_id,
        status,
        start_date,
        end_date,
        created_at,
        price
    )
        SELECT 
        uuid_generate_v4(),
        new.id,
        new.status,
        new.free_trial_start_date ,
        new.free_trial_end_date,
        now(),
        s.price
        FROM subscription_user as su 
        join subscription as s on s.id = su.subscription_id;
    End;
$$;


CREATE TRIGGER subscription_trigger
before INSERT  on subscription_user
FOR EACH ROW EXECUTE FUNCTION create_subscription_trigger();