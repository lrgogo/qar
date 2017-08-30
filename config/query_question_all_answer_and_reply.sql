SELECT question.*,answer.*,reply.* FROM reply
RIGHT JOIN answer ON reply.answer_id = answer.id
RIGHT JOIN question ON answer.question_id = question.id
WHERE question_id=200;


select
  t1.q_id as 问题id
, t1.q_user_id as 提问者id
, t1.q_title as 问题内容
, t2.a_id as 回复id
, t2.a_user_id as 回复者id
, t2.a_content as 回复内容
, t3.r_id as 评论id
, t3.r_user_id as 评论者id
, t3.r_content as 评论内容
from
(
select id as q_id ,user_id as q_user_id ,title as q_title from question
) t1 -- 所有的问题列表，用id做唯一性的区分
left outer join
(
select id as a_id ,question_id ,user_id as a_user_id ,content as a_content from answer
) t2 -- 所有的回答列表，用id做唯一性的区分
on t1.q_id = t2.question_id -- 每个question_id对应的回复
left  outer join
(
select id as r_id ,answer_id ,user_id as r_user_id ,content as r_content FROM reply
) t3
on t2.a_id = t3.answer_id -- 每个answer_id对应的评论
WHERE t1.q_id=200;