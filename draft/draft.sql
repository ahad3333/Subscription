select 
    s.title_ru,
    s.title_en,
    s.title_uz,
    s.price,
    s.duration_type,
    s.duration,
    s.is_active,
    s.free_trial
from subscription_user as su 
join subscription as s on id = su.subscription_id
where su.user = 

select 
    s.title_ru,
    s.title_en,
    s.title_uz,
    s.price,
    s.duration_type,
    s.duration,
    s.is_active,
    s.free_trial
from subscription_user as su 
join subscription as s on id = su.subscription_id


select 
        uh.status,
        uh.start_date,
        uh.end_date,
        uh.price,
        s.title_en,
        s.title_ru,
        s.title_uz,
        s.duration
from subscription_user as su
join user_subscription_history as uh on uh.user_subscription_id = su.id
join subscription as s on s.id = su.subscription_id
Order by  uh.status, uh.start_date, uh.end_date, uh.price DESC


DROP TRIGGER subscription_trigger ON subscription_user;


select 
s.price
from subscription as s 
join subscription_user as su on su.subscription_id =s.id

UPDATE 
			subscription_user as s
		SET status  =  case  
              WHEN status = 'FREE' then 'FREE_TRIAL'
			  WHEN status = 'ACTIVE' then 'INACTIVE'
		End
		where id in (
            select user_subscription_id from user_subscription_history 
			Where status ='ACTIVE' or status ='FREE_TRIAL' and 
			now() > end_date)

send_notification boolean
	true
	false
select
status
from subscription_user 
where id = '999a130a-b5f7-4cf4-b9e6-8c4d0c51007b'
